package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/pkg/errors"

	pb "grpc-sample/pb/chat"

	"google.golang.org/grpc"
)

func getAdress() string {
	const (
		host = "localhost"
		port = "50051"
	)
	return fmt.Sprintf("%s:%s", host, port)
}

func request(client pb.ChatClient) error {
	stream, err := client.Chat(context.Background())
	if err != nil {
		return err
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("エラー: %v", err)
			}
			log.Printf("サーバから：%s", in.Message)
		}
	}()

	if err := stream.Send(&pb.ChatRequest{
		Message: "こんにちは",
	}); err != nil {
		return err
	}

	stream.CloseSend()
	<-waitc
	return nil
}

func exec() error {
	address := getAdress()
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return errors.Wrap(err, "コネクションエラー")
	}
	defer conn.Close()
	client := pb.NewChatClient(conn)
	return request(client)
}

func main() {
	if err := exec(); err != nil {
		log.Println(err)
	}
}
