package worker

import (
	"bufio"
	context "context"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"sync"

	"github.com/robertkrimen/otto"
	log "github.com/sirupsen/logrus"
	c "github.com/strawhatboy/map_reduce/config"
	d "github.com/strawhatboy/map_reduce/data"
	model "github.com/strawhatboy/map_reduce/model"
	pb "github.com/strawhatboy/map_reduce/proto"
	grpc "google.golang.org/grpc"
)

//Worker ...
// The worker is the client side in our distributed system
type Worker struct {
	*pb.UnimplementedClient_ServiceServer
	id               string
	currentJobID     string
	mapperReducerID  int64
	currentStatus    pb.ClientStatus //TODO: need mutex
	client           pb.Server_ServiceClient
	vm               *otto.Otto
	mapResults       []*pb.MapPair
	reduceResults    []*pb.MapPair
	mapReceived      *sync.Map
	currentReduceKey string
	mapperCount      int64 // the mappers count, when reducer getting slices, it will use this to make sure all slices got.
	reducerCount     int64
	ip               string
	port             int64
	output           string
}

//Init ...
// init the worker with config after the worker was created
func (w *Worker) Init(config *c.Config) error {
	log.Info("initializing worker")
	w.mapResults = make([]*pb.MapPair, 100)
	var conn, _ = grpc.Dial(config.Server)
	w.client = pb.NewServer_ServiceClient(conn)
	log.Info("client ready")
	w.vm = otto.New()
	w.mapReceived = &sync.Map{}
	var err error = nil
	w.vm.Set(`MR_mapEmit`, func(call otto.FunctionCall) otto.Value {
		key := call.Argument(0).String()
		value := call.Argument(1).String()
		w.mapResults = append(w.mapResults, model.PairCount{First: key, Second: value}.ToPbPair())
		return otto.Value{}
	})
	w.vm.Set(`MR_reduceEmit`, func(call otto.FunctionCall) otto.Value {
		value := call.Argument(0).String()
		w.reduceResults = append(w.reduceResults, model.PairCount{First: w.GetCurrentReduceKey(), Second: value}.ToPbPair())
		return otto.Value{}
	})
	log.Info("vm ready")
	log.Info("initialized")
	return err
}

//Map ...
// to be called by manager to assign map jobs
func (w *Worker) Map(ctx context.Context, req *pb.MapRequest) (*pb.CommonResponse, error) {
	if w.currentStatus != pb.ClientStatus_idle {
		log.Warn("cannot change to a map worker because currently working on: ", w.currentStatus.String())
		return &pb.CommonResponse{Ok: false, Msg: "No. I'm currently working on " + w.currentStatus.String()}, nil
	}
	log.Info("going to work as a map worker")
	var provider d.Provider
	switch req.DataProvider {
	case pb.DataProvider_raw:
		log.Info("using rawdata data provider")
		provider = &d.RawData{FilePath: req.InputFile}
	default:
		break
	}
	w.currentStatus = pb.ClientStatus_working_mapper
	w.mapperReducerID = req.AssignedId
	w.currentJobID = req.JobId
	// launch the go routine to do the map job because it takes time, so we make it asynchronized.
	go func() {
		files := []string{}
		if req.IsDirectory {
			log.Info("need to get all files in ", req.FileFilter)
			// load all files.
			files, _ = w.walkMatch(req.InputFile, req.FileFilter)
		} else {
			log.Info("using single file ", req.InputFile)
			files = append(files, req.InputFile)
		}

		for _, f := range files {
			provider.SetPath(f)
			provider.LoadData()
			w.vm.Run(req.Script)
			d := provider.ReadData()
			for d != nil {
				w.vm.Call(`MR_map`, nil, d)
				d = provider.ReadData()
			}
		}
		log.Info("map done local")
		// almost done, need to partition these results to R (reducer's count) parts
		// emmm how about partLen := math.Ceil(len(mapResults) / float(R))
		// and get the ith slice by mapResults[partLen * i : partLen * (i+1)]
		// good idea

		// call server.MapDone
		r, err := w.client.MapDone(ctx, &pb.JobDoneRequest{JobId: w.currentJobID, MapperReducerId: w.mapperReducerID, ResultPath: fmt.Sprintf("%v:%v", w.ip, w.port)})
		if !r.Ok || err != nil {
			// print error
			log.Error("failed to send result to server: ", err)
			w.currentStatus = pb.ClientStatus_unknown
		}
		log.Info("map done")
		w.currentStatus = pb.ClientStatus_idle
	}()
	return &pb.CommonResponse{Ok: true, Msg: "Ok I'm working on it."}, nil
}

//MapDone ...
// to be called by manager to notify that one of the map job was done. now the reducer can start to fetch the result
func (w *Worker) MapDone(ctx context.Context, req *pb.JobDoneRequest) (*pb.CommonResponse, error) {

	if w.currentStatus != pb.ClientStatus_working_reducer {
		log.Warn("cannot work as a reducer to receive map result because currently on: ", w.currentStatus.String())
		return &pb.CommonResponse{Ok: false, Msg: "No. I'm currently working on " + w.currentStatus.String()}, nil
	}
	log.Info("going to receive map results from map worker: ", req.MapperReducerId)

	_, ok := w.mapReceived.Load(req.MapperReducerId)
	if !ok {
		go func() {
			// ask for
			conn, _ := grpc.Dial(req.ResultPath)
			client := pb.NewClient_ServiceClient(conn)
			res, err := client.GetReduceSlice(ctx, &pb.ReduceSliceRequest{ReduceId: w.mapperReducerID})
			// put res
			if err != nil {
				// boom ?
				log.Error("failed to init connection with map worker: ", req.MapperReducerId, " ", req.ResultPath)
			}

			log.Info("connection to map worker: ", req.MapperReducerId, " ", req.ResultPath, " established")

			w.mapResults = append(w.mapResults, res.Pairs...)
			w.mapReceived.Store(req.MapperReducerId, true)
			allReceived := true
			for i := int64(0); i < w.mapperCount; i++ {
				_, ok := w.mapReceived.Load(i)
				if !ok {
					allReceived = false
					break
				}
			}
			if allReceived {
				log.Info("all slices received")
				// do the reduce job and send back the reduce done request.
				sort.Slice(w.mapResults, func(i int, j int) bool {
					return w.mapResults[i].First < w.mapResults[j].First
				})

				_len := len(w.mapResults)
				if _len > 0 {
					w.currentReduceKey = w.mapResults[0].First
					var count int64
					count = 0
					//TODO: problem here.
					for i, x := range w.mapResults {
						if x.First != w.currentReduceKey || i == _len-1 {
							w.reduceResults = append(w.reduceResults, model.PairCount{First: w.currentReduceKey, Second: strconv.FormatInt(count, 10)}.ToPbPair())
							w.currentReduceKey = x.First
							count = 0
						} else {
							mp := model.PairCount{}
							mp.FromPbPair(x)
							c, _ := strconv.Atoi(mp.Second)
							count = count + int64(c)
						}
					}
				}
				log.Info("reduce results generated")

				// write reduce results to file
				outputfile, err := os.OpenFile(w.output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				dw := bufio.NewWriter(outputfile)
				for _, rr := range(w.reduceResults) {
					_rr := model.PairCount{}
					_rr.FromPbPair(rr)
					dw.WriteString(fmt.Sprintf("%v,%v", _rr.First, _rr.Second))
				}
				dw.Flush()
				outputfile.Close()
				log.Info("reduce results wrote to file")
				// send reduce done
				r, err := w.client.ReduceDone(ctx, &pb.JobDoneRequest{JobId: w.currentJobID})
				if !r.Ok || err != nil {
					// print error
					log.Error("reduce error: ", err)
					w.currentStatus = pb.ClientStatus_unknown
				}
				log.Info("reduce done")
				w.currentStatus = pb.ClientStatus_idle
			}
		}()
	}

	return &pb.CommonResponse{Ok: true, Msg: "Ok, I'm on it"}, nil
}

//Reduce ...
// to be called by manager to assign reduce job
func (w *Worker) Reduce(ctx context.Context, req *pb.ReduceRequest) (*pb.CommonResponse, error) {
	if w.currentStatus != pb.ClientStatus_idle {
		return &pb.CommonResponse{Ok: false, Msg: "No. I'm currently working on " + w.currentStatus.String()}, nil
	}
	w.currentStatus = pb.ClientStatus_working_reducer
	return nil, nil
}

//Status ...
// to be called by manager to check the current status of this worker
func (w *Worker) Status(context.Context, *pb.Empty) (*pb.StatusResponse, error) {
	return &pb.StatusResponse{Status: w.currentStatus}, nil
}

//GetReduceSlice ...
// to be called by reduce worker to get the slice for reducing
func (w *Worker) GetReduceSlice(ctx context.Context, req *pb.ReduceSliceRequest) (*pb.ReduceSliceResponse, error) {
	partLen := int64(math.Ceil(float64(len(w.mapResults)) / float64(w.reducerCount)))
	start := partLen * req.ReduceId
	end := partLen * (req.ReduceId + 1)
	_len := int64(len(w.mapResults))
	if end > _len {
		end = _len
	}
	return &pb.ReduceSliceResponse{Pairs: w.mapResults[start:end]}, nil
}

//Reset ...
// to be called by manager, when a new mapreduce job comes.
func (w *Worker) Reset(ctx context.Context, req *pb.ResetRequest) (*pb.CommonResponse, error) {
	if w.currentStatus == pb.ClientStatus_idle {
		w.mapperCount = req.MapperCount
		w.reducerCount = req.ReducerCount
		w.mapReceived.Range(func(key interface{}, value interface{}) bool {
			w.mapReceived.Delete(key)
			return true
		})
		return &pb.CommonResponse{Ok: true, Msg: "Ok I'm reset."}, nil
	}
	return &pb.CommonResponse{Ok: false, Msg: "No. I'm currently working on " + w.currentStatus.String()}, nil
}

//GetCurrentReduceKey ...
// to be used when reducing.
func (w *Worker) GetCurrentReduceKey() string {
	return w.currentReduceKey
}

func (w *Worker) walkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, filepath.Join(root, path))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}
