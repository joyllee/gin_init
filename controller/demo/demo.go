package demo

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func HWord(c *gin.Context)  {
	logrus.Info("hello word")
	c.String(http.StatusOK,"hello word")
}