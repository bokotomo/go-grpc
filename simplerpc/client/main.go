package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"

	pb "grpc-sample/pb/calc"

	"google.golang.org/grpc"
)

func getAdress() string {
	const (
		host = "localhost"
		port = "50051"
	)
	return fmt.Sprintf("%s:%s", host, port)
}

func exec(a, b int32) error {
	address := getAdress()
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return errors.Wrap(err, "did not connect")
	}
	defer conn.Close()
	client := pb.NewCalcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	reply, err := client.Sum(ctx, &pb.SumRequest{
		A: a,
		B: b,
	})
	if err != nil {
		return errors.Wrap(err, "could not greet")
	}
	log.Printf("サーバからの受け取り\n %s", reply.GetMessage())
	return nil
}

func main() {
	exec(300, 500)
}
