package main

func CallSayHello(client pb.GreetServiceClien){
	ctx , cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx , &pb.NoParam())
	if err!= nil{
		log.Fatalf("did not greet")
	}

	log.Printf("%s", res.Message)
	
}