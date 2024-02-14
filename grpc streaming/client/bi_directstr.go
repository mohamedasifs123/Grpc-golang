package main

import(
	"log"
	pb "github.com/asif"
	"google.golang.com/grpc"

)


func CallSayHelloBidirectionalStream (client pb.GreetServiceClient,names *pb.NameList) {
	log.Printf("Bidirectional streaming started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err!=nil{
		log.Fatalf("cpuld send names")
		 
	}

	waitc := make(chan struct())
	goo func(){	for{
			message, server:= stream.Recv()
				if err== io.EOF{
					break
			}
			if err!=nil{
				log.Printf("streaming not error")
			}
		log.Println(message)
		}
	close(waitc)	
	}()

    
	for _, name:= range names.Names{
		req:= &pb.HelloRequest{
			Name:name
		}
		if err:= stream.Send(req); err!=nil{
			log.Fatalf("error while sending")
		}
	   time.Sleep(2* time.Second)

	}
    stream.CloseSend()
    <-waitc
    log.Printf("Bidirectional stream ended")

 }












