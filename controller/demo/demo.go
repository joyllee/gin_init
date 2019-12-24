package demo

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"local/gin_init/middleware/api"
	"local/gin_init/model/dto"
	"net/http"
)

func HWord(c *gin.Context) {
	logrus.Info("hello word")
	user := &dto.User{
		UserName: "admin",
		PassWord: "123456",
	}
	call := api.CallClient{
		Host: "http://127.0.0.1:62005",
		Uri:  "/api/v1/account/login",
	}
	auth := &dto.AutoGenerated{}

	err := call.PostCall(user, auth)
	if err != nil {
		logrus.Warn(err)
		return
	}

	c.JSON(http.StatusOK, auth)
}
