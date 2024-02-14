package main 

import(
	"log"
	pb "github.com/asif"
	"google.golang.com/grpc"c
)

func CallSayHelloServerStream (client pb.GreetServiceClient,names *pb.NameList) {
	log.Printf("streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err!=nil{
		log.Fatalf("cpuld send names")
		 
	}
	for{
		message, server:= stream.Recv()
			if err== io.EOF{
				break
		}
		if err!=nil{
			log.Printf("streaming not error")
		}
	log.Println(message)
	}
	log.Printf("streaming finished")
}