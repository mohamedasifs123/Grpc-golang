package main

import{
	"log"
	pb "github.com/asif"
	"google.golang.com/grpc"
}

func (s *HelloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResp , error)  {
	return &pb.HelloResp{
		message : "Hello", 
	},nil

}