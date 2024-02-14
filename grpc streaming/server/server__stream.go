package main 

import(
	"log"
	pb "github.com/asif"
	"google.golang.com/grpc"
)

func (s *HelloServer)CallSayHelloServerStreaming (req *pb.NameList, stream pb.GreetService_SayHelloClientServerStreamingServer) error {
	llog.Printf("got request with names : %s",req.Names)
	for _, names := range req.Names{
		res := &pb.HelloResp{
			Message:= "Hello"+names,
		}
		if err:= stream.Send(res); err!=nil{
			return err
		}
	    time.Sleep(2*time.Second)
	}
    return nil

}

