package router

import (
	"github.com/gin-gonic/gin"
	"local/gin_init/controller"
)

func InitRouter(engine *gin.Engine) {
	engine.POST("/welcome", controller.HWord)
}
