package main

import (
	"net"
	"log"
	"google.golang.org/grpc/reflection"
	"golang.org/x/net/context"
	pb "grpc-helloworld/helloworld"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)
//实现协议接口
// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	say:=in.Name+"yes"
	return &pb.HelloReply{Message: "Hello " + say}, nil
}
/*
然后需要在接口函数里面实现我们具体的业务逻辑，这里仅仅把请求里面的内容读出来，再写回到响应里面

你还可以为这个类增加其他的函数，比如初始化之类的，根据你具体的业务需求就好
 */
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
/*
服务器端主要逻辑就是实现之前协议中的SayHello方法，这里是将字符串Hello和参数拼接在一起返回。

协议生成的go文件给了一个RegisterGreeterServer方法，我们用这个方法绑定实现函数的结构体和server。
 */