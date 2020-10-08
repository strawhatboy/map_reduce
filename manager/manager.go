
package manager
import (
	context "context"
	pb "github.com/strawhatboy/map_reduce/proto"
	grpc "google.golang.org/grpc"
)

//Manager ...
// the manager 
type Manager struct {
	*pb.UnimplementedServer_ServiceServer
	workers map[string] pb.Client_ServiceClient
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
	var conn, _ = grpc.Dial(req.ClientId)
	m.workers[req.ClientId] = pb.NewClient_ServiceClient(conn)
	return &pb.CommonResponse{
		Ok: true,
		Msg: "Successfully registered",
	}, nil
}
