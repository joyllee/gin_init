package middleware

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/cskr/pubsub"
	"github.com/joyllee/blocks"
	"github.com/joyllee/blocks/config"
	"github.com/joyllee/blocks/kafka"
	"github.com/panjf2000/ants/v2"
)

var PubSub *pubsub.PubSub
var pool *ants.Pool

//var gPool *ants.PoolWithFunc

func init() {
	PubSub = pubsub.New(512)
	//gPool, _ = ants.NewPoolWithFunc(512, HandleFaceData, ants.WithExpiryDuration(10*time.Second))
	var err error
	if pool, err = ants.NewPool(128); err != nil {
		panic(err)
	}
}

func InitKafka(ctx *blocks.HCtx) {
	kafka.InitDefaultConsumerGroup(config.ServerConfig.Kafka)
	defer kafka.ConsumerGroup().Close()

	go func() {
		for err := range kafka.ConsumerGroup().Errors() {
			ctx.Warn(err)
		}
	}()

	for {
		err := kafka.ConsumerGroup().Consume(context.Background(), []string{"test-face"}, func(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
			for msg := range claim.Messages() {
				PubSub.Pub(msg, msg.Topic)
				sess.MarkMessage(msg, "")
			}
			return nil
		})
		if err != nil {
			ctx.Warn(err)
		}
	}
}
