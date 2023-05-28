package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/fabiovalinhos/grpc-go/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {

	log.Println("doMax was invoked")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while opening stream: %v\n")
	}

	waitc := make(chan struct{})

	go func() {
		numbers := []int32{4, 9, 7, 45, 3, 32, 77}

		for _, number := range numbers {
			log.Printf("Sending number: %d\n", number)

			stream.Send(&pb.MaxRequest{
				Number: number,
			})

			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Problem while reading stream: %v\n", err)
				break
			}

			log.Printf("Received a new maximum: %d\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
