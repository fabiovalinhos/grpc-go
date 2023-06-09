package main

import (
	"context"
	"log"

	pb "github.com/fabiovalinhos/grpc-go/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("---createBlog was involked---")

	blog := &pb.Blog{
		AuthorId: "fabiovalinhos",
		Title:    "My thirdh blog",
		Content:  "Content of first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpect error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)
	return res.Id
}
