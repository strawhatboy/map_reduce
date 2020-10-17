
package main

import (
	"flag"
	"fmt"
	"net"
	grpc "google.golang.org/grpc"
	pb "github.com/strawhatboy/map_reduce/proto"
	c "github.com/strawhatboy/map_reduce/config"
	worker "github.com/strawhatboy/map_reduce/worker"
)

func main() {
	mainLogger := c.GetLogger("main")
	flag.Parse()
	config, err := c.InitConfig()
	if err != nil {
		mainLogger.Error("Failed to init config: ", err)
		return
	}
	mainLogger.Info("worker starting with config: ", config)
	s := &worker.Worker{}
	var lis net.Listener
	retryTimes := 9999
	for ;retryTimes > 0; {
		lis, err = net.Listen("tcp4", fmt.Sprintf("127.0.0.1:%d", config.Port))
		if err == nil {
			mainLogger.Info("listening at ", config.Port)
			s.SetIpAndPort("127.0.0.1", int64(config.Port))
			break
		}
		retryTimes--
		config.Port++
		mainLogger.Fatal(fmt.Sprintf("failed to listen: %v, retrying with port %v", err, config.Port))
	}
	// var opts []grpc.ServerOption
	s.Init(config)
	grpcServer := grpc.NewServer()
	pb.RegisterClient_ServiceServer(grpcServer, s)
	grpcServer.Serve(lis)
}
