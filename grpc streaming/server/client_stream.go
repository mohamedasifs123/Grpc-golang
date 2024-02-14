package main

import (
	"log"
	pb "github.com/asif"
	"google.golang.com/grpc"


)

func (s *HelloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientServerStreamingServer) error{
	var messages []string
	for{
		req , err := stream.Recv()
		if err==EOF {
			return stream.SendAndClose(&pb.MessageList(Messages ))
		}
		if err!=nil{
			return err
		}

		log.Printf("got request with name %v", req.Name)
		messages = append(messages, "HEllO", req.Name)


	}
}
