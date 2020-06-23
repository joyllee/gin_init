package controller

import (
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"github.com/joyllee/blocks"
	"github.com/joyllee/blocks/utils"
	"github.com/joyllee/blocks/ws"
	"github.com/joyllee/gin_init/middleware"
	"net/http"
)

func KafkaToWs(c *gin.Context) {
	ctx := blocks.NewHTTPContext()
	ctx.Request = c.Request
	ctx.ResponseWriter = c.Writer

	//允许跨域
	ws.SetCheckOrigin(func(r *http.Request) bool {
		return true
	})
	wsIns, err := ws.NewWS(ctx, nil)
	if err != nil {
		return
	}
	defer wsIns.Close()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				ctx.Warn(err, string(utils.GetStack()))
			}
		}()
		for {
			select {
			case <-ctx.Ctx.Done():
				return
			default:
				_, _, err := wsIns.ReadMessage()
				if wsIns.IsWebSocketCloseError(err) {
					ctx.Cancel()
					return
				}
			}
		}
	}()

	sub := middleware.PubSub.Sub("test")
	defer middleware.PubSub.Unsub(sub, "test")

	for {
		select {
		case <-ctx.Ctx.Done():
			return
		case msg:=<-sub:
			kafkaMsg := msg.(*sarama.ConsumerMessage)
			ctx.Info(kafkaMsg.Topic)
			ctx.Info(string(kafkaMsg.Value))
			if err := wsIns.WriteTextMessage(kafkaMsg.Value);err != nil {
				ctx.Warn(err)
				return
			}
		}
	}
}
