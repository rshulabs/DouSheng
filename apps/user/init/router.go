// Author: BeYoung
// Date: 2023/1/29 14:16
// Software: GoLand

package init

import (
	"github.com/Go-To-Byte/DouSheng/apps/user/model"
	"github.com/Go-To-Byte/DouSheng/apps/user/service"
	"github.com/gin-gonic/gin"
)

func initRouter() {
	model.Router = gin.Default()
	user := model.Router.Group("/douyin/user")
	{
		user.GET("/")
		user.POST("/login")
		user.POST("/register", service.Register)
	}
}
