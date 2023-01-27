// Author: BeYoung
// Date: 2023/1/26 2:47
// Software: GoLand

package service

import (
	"github.com/Go-To-Byte/DouSheng/dal/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func register(context *gin.Context) {
	zap.S().Debugf("register")
	user := model.UserInfo{
		ID:       Node.Generate().Int64(),
		Username: context.Query("username"),
		Passwd:   context.Query("password"),
	}

	// TODO: md5.Sum(password)
	passwd := make([]byte, 128)
	copy(passwd, user.Passwd)

	result := FindByName(user)
	if result == nil || len(result) > 0 {
		context.JSON(http.StatusForbidden, gin.H{
			"user_id":     0,
			"status_code": http.StatusForbidden,
			"status_msg":  "failed",
			"token":       "",
		})
		return
	}

	zap.S().Debugf("add user: %+v", user)
	Add(user)
	context.JSON(http.StatusOK, gin.H{
		"user_id":     user.ID,
		"status_code": http.StatusOK,
		"status_msg":  "success",
		"token":       "token",
	})
}
