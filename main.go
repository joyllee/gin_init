package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joyllee/blocks/logger"
	"local/gin_init/router"
)

func main() {
	logger.InitLogger()
	// get a gin engine
	engine := gin.Default()

	router.InitRouter(engine)

	port := fmt.Sprintf(":%d", logger.ServerConfig.Port)
	logger.Info("httpclient server listen on port", port)
	engine.Run(port)
}
