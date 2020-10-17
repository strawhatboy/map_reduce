
package main

import (
	"flag"
	"fmt"
	"net"
	grpc "google.golang.org/grpc"
	pb "github.com/strawhatboy/map_reduce/proto"
	log "github.com/sirupsen/logrus"
	c "github.com/strawhatboy/map_reduce/config"
	worker "github.com/strawhatboy/map_reduce/worker"
)

func main() {
	flag.Parse()
	config, err := c.InitConfig()
	if err != nil {
		log.Error("Failed to init config: ", err)
		return
	}
	log.Info("Worker starting with config: ", config)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", config.Port))
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	s := &worker.Worker{}
	s.Init(config)
	pb.RegisterClient_ServiceServer(grpcServer, s)
	grpcServer.Serve(lis)
}
