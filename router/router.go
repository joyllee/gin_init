package router

import (
	"github.com/gin-gonic/gin"
	"local/gin_init/controller/demo"
)

func InitRouter(engine *gin.Engine) {
	engine.GET("/welcome", demo.HWord)
}
