package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/joyllee/blocks"
	"github.com/joyllee/blocks/http_pool"
	"github.com/valyala/fasthttp"
)

func HWord(c *gin.Context) {
	ctx := blocks.NewHTTPContext()
	http_pool.InitClient(3)
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	req.Header.SetMethod(fasthttp.MethodPost)
	req.SetRequestURI("http://172.16.9.201:62100/api/v1/account/login")
	req.SetBodyString(`{
    "user_name":"admin",
    "password":"123456"
}`)

	err := http_pool.Client().Do(req, res)
	if err != nil {
		ctx.Error(err)
		c.JSON(200, gin.H{"message": "error"})
	}
	ctx.Info(string(res.Body()))
}

