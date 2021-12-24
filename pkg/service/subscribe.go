package service

import (
	"context"

	pb "github.com/tkeel-io/core-broker/api/dapr"
	"github.com/tkeel-io/core-broker/pkg/types"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SubscribeService struct {
	pb.UnimplementedSubscribeServer
}

func NewSubscribeService() *SubscribeService {
	return &SubscribeService{}
}

func (s *SubscribeService) GetSubscribe(ctx context.Context, req *emptypb.Empty) (*pb.ListTopicSubscriptionsResponse, error) {
	resp := &pb.ListTopicSubscriptionsResponse{}
	resp.Subscriptions = append(resp.Subscriptions, &pb.TopicSubscription{
		PubsubName: types.PubsubName,
		Topic:      types.Topic,
		Metadata:   map[string]string{},
		Route:      "/v1/topic",
	})

	return resp, nil
}
