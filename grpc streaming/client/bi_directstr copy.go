package main

import(

)

func (s *HelloServer)CallSayHelloBidrectionalStreaming (req *pb.NameList, stream pb.GreetService_SayHelloBidrectionalStreamingServer) error {
	for{
		req, err:= stream.Recv()
			if err== io.EOF{
				return nil
		}
		if err!=nil{
			return err
		}
		log.Printf("got request with name:%v",req.Name)
		res:=  &pb.HelloResp{
			Message:= "Hello"+req.Name,
		}
		if err:= stream.Send(res); err!=nil{
			return err
		}
		
}
