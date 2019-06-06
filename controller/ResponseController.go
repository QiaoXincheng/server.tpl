package controller

import (
	"server.tpl/base"

	"github.com/gin-gonic/gin"

	"strconv"
)

const (
	REQUEST_SUCCESS       = 0       //请求成功
	REQUEST_ERROR         = 1		//请求失败
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func OutPut(c *gin.Context, response *Response) {
	if response.Code > 0 {
		codeStr := strconv.Itoa(response.Code)
		response.Message = base.Tr("zh-CN", codeStr) + response.Message
	}
	c.JSON(200, response)
}

func getHeader(ctx *gin.Context, key string) string {
	return ctx.Request.Header.Get(key)
}
