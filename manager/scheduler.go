package manager

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	c "github.com/strawhatboy/map_reduce/config"
)

//Scheduler serve as a http server
//to control the manager?, like starting a new job?
type Scheduler struct {
	logger *log.Entry
}

func (s *Scheduler) Init(m *Manager) {
	s.logger = c.GetLogger("scheduler")
	s.logger.Info("initializing")
	r := gin.Default()
	r.GET("/status", func(c *gin.Context) {
		// return current system status
		c.String(http.StatusOK, m.GetStatusMessage())
	})
	r.POST("/job/new", func(c *gin.Context) {
		// init a new job with:
		//	mapper count
		//	reducer count
		//	mapper function file
		//	reduce function file
		//	source file
		//	type: raw, audio, csv...
		//	file_split_strategy
		job := &Job{}
		if err := c.ShouldBindJSON(job); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 9999,
				"msg": fmt.Sprintf("bad request: %v", err),
			})
		} else {
			// create the job
			m.QueueJob(job)
		}
	})
	s.logger.Info("web api registered")
	go r.Run("127.0.0.1:18080")
}
