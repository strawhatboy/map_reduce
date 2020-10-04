
package manager
import (
	context "context"
	pb "github.com/strawhatboy/map_reduce/proto"
)


type Manager struct {
	*pb.UnimplementedServer_ServiceServer
	workers map[string]*pb.Client_ServiceClient
}

func (Manager) MapDone(context.Context, *pb.JobDoneRequest) (*pb.CommonResponse, error) {
}
func (Manager) ReduceDone(context.Context, *pb.JobDoneRequest) (*pb.CommonResponse, error) {
}
func (Manager) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.CommonResponse, error) {
	pb.NewClient_ServiceClient()
	req.ClientId
}
