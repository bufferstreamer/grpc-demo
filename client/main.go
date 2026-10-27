// gRPC 客户端
// 连接服务端，像调用本地函数一样发起远程调用
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/bufferstreamer/grpc-demo/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1. 建立到服务端的连接
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()

	// 2. 创建客户端桩（stub）
	client := pb.NewUserServiceClient(conn)

	// 3. 设置超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 4. 像调用本地函数一样发起远程调用
	req := &pb.GetUserRequest{Id: 1001}
	fmt.Printf("[客户端] 发起调用: GetUser(id=%d)\n", req.GetId())

	resp, err := client.GetUser(ctx, req)
	if err != nil {
		log.Fatalf("调用失败: %v", err)
	}

	// 5. 拿到结构化响应
	user := resp.GetUser()
	fmt.Println("[客户端] 收到响应:")
	fmt.Printf("        id    = %d\n", user.GetId())
	fmt.Printf("        name  = %s\n", user.GetName())
	fmt.Printf("        email = %s\n", user.GetEmail())
	fmt.Printf("        age   = %d\n", user.GetAge())
}
