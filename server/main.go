// gRPC 服务端
// 实现 UserService，监听 50051 端口，处理 GetUser 请求
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/bufferstreamer/grpc-demo/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// 模拟数据库
var mockDB = map[int64]*pb.User{
	1001: {Id: 1001, Name: "Alice", Email: "alice@example.com", Age: 28},
	1002: {Id: 1002, Name: "Bob", Email: "bob@example.com", Age: 34},
	1003: {Id: 1003, Name: "Carol", Email: "carol@example.com", Age: 25},
}

// server 实现 pb.UserServiceServer 接口
type server struct {
	pb.UnimplementedUserServiceServer
}

// GetUser 处理一次查询请求
func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	id := req.GetId()
	log.Printf("[服务端] 收到请求: GetUser(id=%d)", id)

	user, ok := mockDB[id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "user %d not found", id)
	}

	log.Printf("[服务端] 返回响应: %s", user.GetName())
	return &pb.GetUserResponse{User: user}, nil
}

func main() {
	// 1. 监听端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}

	// 2. 创建 gRPC 服务器
	s := grpc.NewServer()

	// 3. 注册我们的服务实现
	pb.RegisterUserServiceServer(s, &server{})

	fmt.Println("[服务端] 已启动，监听 50051 端口...")
	// 4. 启动并阻塞
	if err := s.Serve(lis); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
