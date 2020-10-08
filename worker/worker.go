package worker

import (
	context "context"
	"github.com/robertkrimen/otto"
	c "github.com/strawhatboy/map_reduce/config"
	d "github.com/strawhatboy/map_reduce/data"
	pb "github.com/strawhatboy/map_reduce/proto"
	grpc "google.golang.org/grpc"
	"math"
	"sync"
)

//Worker ...
// The worker is the client side in our distributed system
type Worker struct {
	*pb.UnimplementedClient_ServiceServer
	id               string
	currentJobID	 string
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
}

//Init ...
// init the worker with config after the worker was created
func (w *Worker) Init(config *c.Config) error {
	w.mapResults = make([]*pb.MapPair, 100)
	var conn, _ = grpc.Dial(config.Server)
	w.client = pb.NewServer_ServiceClient(conn)
	w.vm = otto.New()
	w.mapReceived = &sync.Map{}
	var err error = nil
	w.vm.Set(`MR_mapEmit`, func(call otto.FunctionCall) otto.Value {
		key := call.Argument(0).String()
		var value int64
		value, err = call.Argument(1).ToInteger()
		w.mapResults = append(w.mapResults, &pb.MapPair{First: key, Second: value})
		return otto.Value{}
	})
	w.vm.Set(`MR_reduceEmit`, func(call otto.FunctionCall) otto.Value {
		var value int64
		value, err = call.Argument(0).ToInteger()
		w.reduceResults = append(w.reduceResults, &pb.MapPair{First: w.GetCurrentReduceKey(), Second: value})
		return otto.Value{}
	})
	return err
}

//Map ...
// to be called by manager to assign map jobs
func (w *Worker) Map(ctx context.Context, req *pb.MapRequest) (*pb.CommonResponse, error) {
	if w.currentStatus != pb.ClientStatus_idle {
		return &pb.CommonResponse{Ok: false, Msg: "No. I'm currently working on " + w.currentStatus.String()}, nil
	}
	var provider d.Provider
	switch req.DataProvider {
	case pb.DataProvider_raw:
		provider = &d.RawData{FilePath: req.InputFile}
	default:
		break
	}
	w.currentStatus = pb.ClientStatus_working_mapper
	w.mapperReducerID = req.AssignedId
	w.currentJobID = req.JobId
	// launch the go routine to do the map job because it takes time, so we make it asynchronized.
	go func() {
		provider.LoadData()
		w.vm.Run(req.Script)
		d := provider.ReadData()
		for d != nil {
			w.vm.Call(`MR_map`, nil, d)
			d = provider.ReadData()
		}
		// almost done, need to partition these results to R (reducer's count) parts
		// emmm how about partLen := math.Ceil(len(mapResults) / float(R))
		// and get the ith slice by mapResults[partLen * i : partLen * (i+1)]
		// good idea

		// call server.MapDone
		r, err := w.client.MapDone(ctx, &pb.JobDoneRequest{JobId: w.currentJobID, MapperReducerId: w.mapperReducerID, ResultPath: ""})
		if !r.Ok || err != nil {
			// print error
			w.currentStatus = pb.ClientStatus_unknown
		}
		w.currentStatus = pb.ClientStatus_idle
	}()
	return &pb.CommonResponse{Ok: true, Msg: "Ok I'm working on it."}, nil
}

//MapDone ...
// to be called by manager to notify that one of the map job was done. now the reducer can start to fetch the result
func (w *Worker) MapDone(ctx context.Context, req *pb.JobDoneRequest) (*pb.CommonResponse, error) {
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
			}

			w.mapResults = append(w.mapResults, res.Pairs...)
			w.mapReceived.Store(req.MapperReducerId, true)
			allRecieved := true
			for i := int64(0); i < w.mapperCount; i++ {
				_, ok := w.mapReceived.Load(i)
				if !ok {
					allRecieved = false
					break
				}
			}
			if allRecieved {
				// do the reduce job and send back the reduce done request.
			}
		}()
	}

	return nil, nil
}

//Reduce ...
// to be called by manager to assign reduce job
func (w *Worker) Reduce(ctx context.Context, req *pb.ReduceRequest) (*pb.CommonResponse, error) {
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
