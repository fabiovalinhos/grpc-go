package main

import (
	"context"
	"io"
	"log"

	pb "github.com/fabiovalinhos/grpc-go/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := pb.GreetRequest{
		FirstName: "Valinhos",
	}

	stream, err := c.GreetManyTimes(context.Background(), &req)

	if err != nil {
		log.Fatalln("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error white reading the stream: %v\n", err)
		}

		log.Printf("GreetingManyTimes: %s\n", msg.Result)
	}
}
