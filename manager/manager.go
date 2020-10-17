package manager

import (
	context "context"
	"io/ioutil"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	c "github.com/strawhatboy/map_reduce/config"
	pb "github.com/strawhatboy/map_reduce/proto"
	grpc "google.golang.org/grpc"
)

//Manager ...
// the manager
type Manager struct {
	*pb.UnimplementedServer_ServiceServer
	workers map[string] pb.Client_ServiceClient
	currentJob	*Job
	currentJobDone	chan struct{}
	jobs	chan *Job
	logger	*log.Entry
}

func (m *Manager) Init() {
	// listen on 19999
	m.logger = c.GetLogger("manager")
	m.jobs = make(chan *Job, 10)
	m.currentJobDone = make(chan struct{}, 1)
	m.workers = make(map[string]pb.Client_ServiceClient, 100)
}

//MapDone ...
// worker will report when map job is done
func (*Manager) MapDone(ctx context.Context, req *pb.JobDoneRequest) (*pb.CommonResponse, error) {
	log.Info("map job ", req.JobId, "from client: ", req.MapperReducerId, " is done")
	// notify every reducer with the news

	return nil, nil
}

//ReduceDone ...
// worker will report when reduce job is done
func (*Manager) ReduceDone(ctx context.Context, req *pb.JobDoneRequest) (*pb.CommonResponse, error) {
	log.Info("reduce job ", req.JobId, "from client: ", req.MapperReducerId, " is done")
	// the job is done
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

func (m *Manager) GetStatusMessage() string {
	return ""
}

func (m *Manager) QueueJob(j *Job) {
	// put it into the chan
	if j.ID == "" {
		j.ID = uuid.New().String()
	}
	m.logger.Info("queued new job: ", j.ID)
	m.jobs <- j
}

func (m *Manager) pickJob() error {
	for {
		m.logger.Info("going to pick up new job")
		m.currentJob = <- m.jobs
		// allocate workers
		if m.currentJob.MapperCount + m.currentJob.ReducerCount > len(m.workers) {
			// not enough workers
			return &NotEnoughWorkerError{}
		}

		var mapScript string
		f, err := ioutil.ReadFile(m.currentJob.MapperFuncFile)
		if err != nil {
			log.Fatal("failed to load mapper script from ", m.currentJob.MapperFuncFile)
		}
		mapScript = string(f)
		log.Info("loaded mapper script from ", m.currentJob.MapperFuncFile)

		var reduceScript string
		f, err = ioutil.ReadFile(m.currentJob.ReducerFuncFile)
		if err != nil {
			log.Fatal("failed to load reducer script from ", m.currentJob.ReducerFuncFile)
		}
		reduceScript = string(f)
		log.Info("loaded reducer script from ", m.currentJob.ReducerFuncFile)

		// assign 
		var i int64 = 0
		for _, w := range(m.workers) {
			if i < int64(m.currentJob.MapperCount) {
				w.Map(context.Background(), &pb.MapRequest{
					JobId: m.currentJob.ID, 
					AssignedId: i,
					DataProvider: pb.DataProvider(pb.DataProvider_value[m.currentJob.DataProvider]),
					FileFilter: m.currentJob.FilePattern,
					InputFile: m.currentJob.SourceFile,
					IsDirectory: m.currentJob.IsDirectory,
					ReducerCount: int64(m.currentJob.ReducerCount),
					Script: mapScript,
				})
			} else {
				w.Reduce(context.Background(), &pb.ReduceRequest{
					JobId: m.currentJob.ID, 
					AssignedId: i,
					Script: reduceScript,
					OutputPrefix: m.currentJob.ID,
				})
			}

			i++
		}

		// wait for current job done.
		<- m.currentJobDone
		m.logger.Info("current job: ", m.currentJob.ID, " is done.")
	}
}

type NotEnoughWorkerError struct {}
func (ne *NotEnoughWorkerError) Error() string {
	return "Not enough worker here... "
}
