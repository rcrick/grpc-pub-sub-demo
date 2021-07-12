package main

import (
	"context"
	pb "example/pb"
	"log"
	"net"
	"strings"
	"time"

	"github.com/docker/docker/pkg/pubsub"
	"google.golang.org/grpc"
)

type PubSubService struct {
	pub *pubsub.Publisher
}

func NewPubSubService() *PubSubService {
	return &PubSubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

func (p *PubSubService) Publish(ctx context.Context, arg *pb.Msg) (*pb.Msg, error) {
	p.pub.Publish(arg.GetValue())
	return &pb.Msg{Value: "val"}, nil
}

func (p *PubSubService) Subscribe(arg *pb.Msg, stream pb.PubsubService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})
	for c := range ch {
		if err := stream.Send(&pb.Msg{Value: c.(string)}); err != nil {
			return err
		}
	}
	return nil
}
func main() {
	server := grpc.NewServer()
	pb.RegisterPubsubServiceServer(server, NewPubSubService())
	lis, err := net.Listen("tcp", ":1234")
	if nil != err {
		log.Fatal(err)
	}

	server.Serve(lis)
}
