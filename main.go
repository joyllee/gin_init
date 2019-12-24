package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"local/gin_init/common/config"
	"local/gin_init/middleware/sessionctx"
	"local/gin_init/router"
	_ "local/gin_init/util/logger"
)

func main() {
	// init mongo connection
	//db.InitMongo(&config.ServerConfig.Mongo)
	// init redis connection
	//db.InitRedis(&config.ServerConfig.Redis)
	// init mysql connection
	//err := db.InitMysql(config.ServerConfig.Mysql)
	//if nil != err {
	//	return
	//}

	// get a gin engine
	engine := gin.Default()
	if config.ServerConfig.Mode == "release" {
		// release mode
		gin.SetMode(gin.ReleaseMode)
	}

	engine.Use(sessionctx.NewSessionCtx)

	router.InitRouter(engine)
	port := fmt.Sprintf(":%d", config.ServerConfig.Port)
	logrus.Info("httpclient server listen on port", port)
	engine.Run(port)
}
