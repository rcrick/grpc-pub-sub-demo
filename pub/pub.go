package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "example/pb"
)

func main() {
	conn, err := grpc.Dial(":1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	client := pb.NewPubsubServiceClient(conn)

	reply, err := client.Publish(context.Background(), &pb.Msg{Value: "A adas"})
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(reply.GetValue())
}
