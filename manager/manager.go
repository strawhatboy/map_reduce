package manager

import (
	context "context"
	"io/ioutil"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	c "github.com/strawhatboy/map_reduce/config"
	pb "github.com/strawhatboy/map_reduce/proto"
	grpc "google.golang.org/grpc"
	"github.com/strawhatboy/map_reduce/util"
)

//Manager ...
// the manager
type Manager struct {
	*pb.UnimplementedServer_ServiceServer
	workers        map[string]pb.Client_ServiceClient
	mappers        []string
	reducers       []string
	mapDone		map[string]bool
	reduceDone		map[string]bool
	currentJob     *Job
	currentJobDone chan struct{}
	jobs           chan *Job
	logger         *log.Entry
}

func (m *Manager) Init() {
	// listen on 19999
	m.logger = c.GetLogger("manager")
	m.jobs = make(chan *Job, 10)
	m.currentJobDone = make(chan struct{}, 1)
	m.mappers = make([]string, 0)
	m.reducers = make([]string, 0)
	m.mapDone = make(map[string]bool)
	m.reduceDone = make(map[string]bool, 0)
	m.workers = make(map[string]pb.Client_ServiceClient, 100)

	// start to pick jobs.
	go m.pickJob()
}

//MapDone ...
// worker will report when map job is done
func (m *Manager) MapDone(ctx context.Context, req *pb.JobDoneRequest) (*pb.CommonResponse, error) {
	log.Info("map job ", req.JobId, "from client: ", req.MapperReducerId, " is done")
	m.mapDone[m.mappers[req.MapperReducerId]] = true
	// notify every reducer with the news

	for _, cid := range m.mappers {
		res, err := m.workers[cid].MapDone(ctx, req)
		if err != nil || !res.Ok {
			log.Error("job: ", req.JobId, ", failed to notify reducer: ", cid)
		} else {
			log.Info("job: ", req.JobId, ", notified reducer: ", cid)
		}
	}
	return nil, nil
}

//ReduceDone ...
// worker will report when reduce job is done
func (m *Manager) ReduceDone(ctx context.Context, req *pb.JobDoneRequest) (*pb.CommonResponse, error) {
	log.Info("reduce job ", req.JobId, "from client: ", req.MapperReducerId, " is done")
	m.reduceDone[m.reducers[req.MapperReducerId]] = true
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
		Ok:  true,
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
		m.currentJob = <-m.jobs
		m.logger.Info("new job got: ", m.currentJob)
		// allocate workers
		if m.currentJob.MapperCount+m.currentJob.ReducerCount > len(m.workers) {
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

		// get files:
		files := []string{}
		for _, f := range m.currentJob.SourceFile {
			if m.currentJob.IsDirectory {
				m.logger.Info("need to get all files in ", f)
				// load all files.
				files, _ = util.WalkMatch(f, m.currentJob.FilePattern)
			} else {
				m.logger.Info("adding single file ", f)
				files = append(files, f)
			}
		}

		filesCountForEachMapper := int64(len(files) / m.currentJob.MapperCount + 1)

		// assign
		var i int64 = 0
		for cid, w := range m.workers {
			if i < int64(m.currentJob.MapperCount) {
				res, err := w.Map(context.Background(), &pb.MapRequest{
					JobId:        m.currentJob.ID,
					AssignedId:   i,
					DataProvider: pb.DataProvider(pb.DataProvider_value[m.currentJob.DataProvider]),
					FileFilter:   m.currentJob.FilePattern,
					InputFiles:   files[filesCountForEachMapper * i : filesCountForEachMapper * (i+1)],
					IsDirectory:  m.currentJob.IsDirectory,
					ReducerCount: int64(m.currentJob.ReducerCount),
					Script:       mapScript,
				})
				if err != nil || !res.Ok {
					log.Error("failed to assign map task ", i, " to worker: ", cid)
					continue
				}
				log.Info("assigned map task: ", i, " to worker: ", cid)
				m.mappers = append(m.mappers, cid)
			} else {
				theID := i - int64(m.currentJob.MapperCount)
				res, err := w.Reduce(context.Background(), &pb.ReduceRequest{
					JobId:        m.currentJob.ID,
					AssignedId:   theID,
					Script:       reduceScript,
					OutputPrefix: m.currentJob.ID,
				})
				if err != nil || !res.Ok {
					log.Error("failed to assign reduce task ", theID, " to worker: ", cid)
					continue
				}
				log.Info("assigned reduce task: ", theID, " to worker: ", cid)
				m.reducers = append(m.reducers, cid)
			}

			i++
		}

		// wait for current job done.
		<-m.currentJobDone
		m.logger.Info("current job: ", m.currentJob.ID, " is done.")
	}
}

type NotEnoughWorkerError struct{}

func (ne *NotEnoughWorkerError) Error() string {
	return "Not enough worker here... "
}
