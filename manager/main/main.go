package main

import (
	"fmt"
	"net"

	c "github.com/strawhatboy/map_reduce/config"
	"github.com/strawhatboy/map_reduce/manager"
	"google.golang.org/grpc"
	pb "github.com/strawhatboy/map_reduce/proto"
)

func main() {
	mainLogger := c.GetLogger("main")
	var lis net.Listener
	var err error
	retryTimes := 9999
	port := 19999
	for ;retryTimes > 0; {
		lis, err = net.Listen("tcp4", fmt.Sprintf("127.0.0.1:%d", port))
		if err == nil {
			mainLogger.Info("listening at ", port)
			break
		}
		retryTimes--
		port++
		mainLogger.Fatal(fmt.Sprintf("failed to listen: %v, retrying with port %v", err, port))
	}
	// var opts []grpc.ServerOption
	grpcServer := grpc.NewServer()
	s := &manager.Manager{}
	scheduler := &manager.Scheduler{}
	s.Init()
	scheduler.Init(s)
	pb.RegisterServer_ServiceServer(grpcServer, s)
	err = grpcServer.Serve(lis)
	if err != nil {
		mainLogger.Fatal(fmt.Sprintf("failed to serve.. %v", err))
	} else {
		mainLogger.Info("serving")
	}
}
