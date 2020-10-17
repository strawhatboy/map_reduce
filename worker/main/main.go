
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
	var portFromCmd int
	flag.IntVar(&portFromCmd, "p", 20000, "the port")
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
	var port int
	if portFromCmd != 20000 {
		port = portFromCmd
	} else {
		port = config.Port
	}
	retryTimes := 9999
	for ;retryTimes > 0; {
		lis, err = net.Listen("tcp4", fmt.Sprintf("127.0.0.1:%d", port))
		if err == nil {
			mainLogger.Info("listening at ", port)
			s.SetIpAndPort("127.0.0.1", int64(port))
			break
		}
		retryTimes--
		port++
		mainLogger.Fatal(fmt.Sprintf("failed to listen: %v, retrying with port %v", err, port))
	}
	// var opts []grpc.ServerOption
	s.Init(config)
	grpcServer := grpc.NewServer()
	pb.RegisterClient_ServiceServer(grpcServer, s)
	grpcServer.Serve(lis)
}
