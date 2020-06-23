package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joyllee/blocks/config"
	"github.com/joyllee/blocks/logger"
	"github.com/joyllee/gin_init/router"
)

func main() {
	logger.InitLogger()
	// get a gin engine
	engine := gin.Default()

	router.InitRouter(engine)

	port := fmt.Sprintf(":%d", config.ServerConfig.Port)
	engine.Run(port)
}
