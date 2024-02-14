package main

import(
	"log"
	pb "github.com/asif"
	"google.golang.com/grpc"c
)

const (
	port =""
)

func main(){
	conn ,err = grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	c
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := %pb.NameList(
		Names: []string("asd","fgh","jkl")
	)

	CallSayHelloServerStream(client, names)
	CallSayHelloClientStream(client, names)
	CallSayHelloBidirectionalStream(client, names)

 }