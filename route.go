package main

import (
	"code-math/component/appctx"

	"github.com/gin-gonic/gin"
)

func setupRoute(ctx appctx.AppContext, v1 *gin.RouterGroup) {

	courses := v1.Group("/courses")
	courses.POST("/", func( c *gin.Context){
		//hàm trả về http reponse
		//c: đối tưởng kiểm soát hết toàn bộ việc lấy request data và response data
		//chỉ chạy 1 lần đầu
		c.JSON(200, gin.H{"message": "Hello, world!"})	
		
	})
}