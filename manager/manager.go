
package manager
import (
	context "context"
	pb "github.com/strawhatboy/map_reduce/proto"
	grpc "google.golang.org/grpc"
	c "github.com/strawhatboy/map_reduce/config"
	log "github.com/sirupsen/logrus"
)

//Manager ...
// the manager 
type Manager struct {
	*pb.UnimplementedServer_ServiceServer
	workers map[string] pb.Client_ServiceClient
	logger	*log.Entry
}

func (m *Manager) Init() {
	// listen on 19999
	m.logger = c.GetLogger("manager")
	m.workers = make(map[string]pb.Client_ServiceClient, 100)
}

//MapDone ...
// worker will report when map job is done
func (*Manager) MapDone(context.Context, *pb.JobDoneRequest) (*pb.CommonResponse, error) {
	return nil, nil
}

//ReduceDone ...
// worker will report when reduce job is done
func (*Manager) ReduceDone(context.Context, *pb.JobDoneRequest) (*pb.CommonResponse, error) {
	return nil, nil
}

//Register ...
// worker will register itself just after launched
func (m *Manager) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.CommonResponse, error) {
	var conn, err = grpc.Dial(req.ClientEndpoint, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		m.logger.Error("failed to connect to the client: ", req.ClientId, "  ->  ", req.ClientEndpoint)
	} else {
		m.logger.Info("inited connection with client: ", req.ClientId, "  ->  ", req.ClientEndpoint)
	}
	m.workers[req.ClientId] = pb.NewClient_ServiceClient(conn)
	return &pb.CommonResponse{
		Ok: true,
		Msg: "Successfully registered",
	}, nil
}
