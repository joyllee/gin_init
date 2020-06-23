package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joyllee/gin_init/controller"
)

func InitRouter(engine *gin.Engine) {
	engine.POST("/welcome", controller.HWord)
	engine.GET("/ws", controller.KafkaToWs)
}
