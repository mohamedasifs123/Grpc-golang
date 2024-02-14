package main

import (
	"log"
	pb "github.com/asif"
	"google.golang.com/grpc"


)

func SayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NameList){
	stream,err := client.SayHelloClientStreaming(context.Background())
	
	if err!=nil{
		log.Printf(" name error")
	}

	for _, name:= range names.Names{
		req:= &pb.HelloRequest{
			Name:name
		}
		if err:= stream.Send(res); err!=nil{
			log.Fatalf("error while sending")
		}
		log.Printf("Sent the request with name:%s",name)
	    time.Sleep(2* time.Second)

	}
	res,err := stream.SendAndClose()
	log.Printf("client streaam finished")
	if err!=nil{
		log.Fatalf(" receiving error")
	}
	log.Printf("%v",res.Messages)

}














