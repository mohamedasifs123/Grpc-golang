
type HelloServer struct {
    pb.GreetServiceServer	
}


pb.RegisterGreetServiceServer(grpcServer, %HelloServer)
log.Fatalf("server start at server")