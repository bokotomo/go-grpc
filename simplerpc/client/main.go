package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"

	pb "grpc-sample/simplerpc/helloworld"

	"google.golang.org/grpc"
)

func getAdress() string {
	const (
		host = "localhost"
		port = "50051"
	)
	return fmt.Sprintf("%s:%s", host, port)
}

func exec(message string) error {
	address := getAdress()
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return errors.Wrap(err, "did not connect")
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	reply, err := client.SayHello(ctx, &pb.HelloRequest{
		Name: message,
	})
	if err != nil {
		return errors.Wrap(err, "could not greet")
	}
	log.Printf("サーバからの受け取り\n %s", reply.GetMessage())
	return nil
}

func main() {
	message := "メッセージ内容"
	exec(message)
}
