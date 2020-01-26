package main

import (
	"context"
	"fmt"
	"log"
	"io"

	"github.com/pkg/errors"

	pb "grpc-sample/pb/notification"

	"google.golang.org/grpc"
)

func getAdress() string {
	const (
		host = "localhost"
		port = "50051"
	)
	return fmt.Sprintf("%s:%s", host, port)
}

func exec(num int32) error {
	address := getAdress()
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return errors.Wrap(err, "コネクションエラー")
	}
	defer conn.Close()
	client := pb.NewNotificationClient(conn)

	req := &pb.NotificationRequest{
		Num: num,
	}
	stream, err := client.Notification(context.Background(), req)
	if err != nil {
		log.Fatalf("えらー：%v", err)
	}
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		log.Println("これ：", feature.GetMessage())
	}

	return nil
}

func main() {
	exec(5)
}
