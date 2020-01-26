package main

import (
	"context"
	"fmt"
	"io"
	"log"

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

func request(client pb.NotificationClient, num int32) error {
	req := &pb.NotificationRequest{
		Num: num,
	}
	stream, err := client.Notification(context.Background(), req)
	if err != nil {
		return errors.Wrap(err, "streamエラー")
	}
	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Println("これ：", reply.GetMessage())
	}
	return nil
}

func exec(num int32) error {
	address := getAdress()
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return errors.Wrap(err, "コネクションエラー")
	}
	defer conn.Close()
	client := pb.NewNotificationClient(conn)
	return request(client, num)
}

func main() {
	num := int32(5)
	if err := exec(num); err != nil {
		log.Println(err)
	}
}
