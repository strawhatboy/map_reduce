package worker

import (
	context "context"

	"github.com/robertkrimen/otto"
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
	currentStatus    pb.StatusResponse_ClientStatus		//TODO: need mutex
	client           pb.Server_ServiceClient
	vm               *otto.Otto
	mapResults       []model.Pair
	reduceResults    []model.Pair
	currentReduceKey string
}

func (w *Worker) Init(config *c.Config) error {
	w.mapResults = make([]model.Pair, 100)
	var conn, _ = grpc.Dial(config.Server)
	w.client = pb.NewServer_ServiceClient(conn)
	w.vm = otto.New()
	var err error = nil
	w.vm.Set(`MR_mapEmit`, func(call otto.FunctionCall) otto.Value {
		key := call.Argument(0).String()
		var value int64
		value, err = call.Argument(1).ToInteger()
		w.mapResults = append(w.mapResults, model.Pair{key, value})
		return otto.Value{}
	})
	w.vm.Set(`MR_reduceEmit`, func(call otto.FunctionCall) otto.Value {
		var value int64
		value, err = call.Argument(0).ToInteger()
		w.reduceResults = append(w.reduceResults, model.Pair{w.GetCurrentReduceKey(), value})
		return otto.Value{}
	})
	return err
}

func (w *Worker) Map(ctx context.Context, req *pb.MapReduceRequest) (*pb.CommonResponse, error) {
	var provider d.Provider
	switch req.DataProvider {
	case pb.MapReduceRequest_raw:
		provider = &d.RawData{FilePath: req.InputFile}
	default:
		break
	}
	w.currentStatus = pb.StatusResponse_working_mapper
	// launch the go routine to do the map job
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
	}()
	return &pb.CommonResponse{Ok: true, Msg: "Ok I'm working on it."}, nil
}

func (w *Worker) Reduce(context.Context, *pb.MapReduceRequest) (*pb.CommonResponse, error) {
	return nil, nil
}
func (w *Worker) Status(context.Context, *pb.Empty) (*pb.StatusResponse, error) {
	return &pb.StatusResponse{Status: w.currentStatus}, nil
}

func (w *Worker) GetCurrentReduceKey() string {
	return w.currentReduceKey
}
