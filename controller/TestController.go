package controller

import (
	"server.tpl/service"

	"github.com/gin-gonic/gin"
)



// 创意素材审核状态查询
func HelloWorld(c *gin.Context) {
	response := &Response{}
	hash, err := service.GenerateJWT()
	if err != nil {
		response.Code = REQUEST_ERROR
		OutPut(c, response)
		return
	}
	response.Data = hash
	OutPut(c, response)
}
