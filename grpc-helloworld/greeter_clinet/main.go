package main

import (
	"log"
	"os"
	"google.golang.org/grpc"
	pb "grpc-helloworld/helloworld"
	"golang.org/x/net/context"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := "张川林" +
		"" +
			""
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
/*
客户端的思路也很清晰，建立一个rpc客户端连接，将这个连接用pb.NewGreeterClient和协议绑定，返回一个client对象，用这个对象就可以调用远程的函数了。
 */