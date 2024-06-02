package client

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

type SocialNetworkProvider interface {
	SendSocial(Social string)
}

type SocialProvider struct {
	conn *kafka.Conn
}

func NewSocialProvider() SocialNetworkProvider {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "kafka:9092", "social", 0)
	if err != nil {
		return nil
	}
	return &SocialProvider{
		conn: conn,
	}
}

func (s *SocialProvider) SendSocial(Social string) {
	go s.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
}
