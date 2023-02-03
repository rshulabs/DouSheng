// Author: BeYoung
// Date: 2023/2/1 15:18
// Software: GoLand

package services

import (
	"fmt"
	"github.com/Go-To-Byte/DouSheng/network/milddlewares"
	"github.com/Go-To-Byte/DouSheng/network/models"
	proto "github.com/Go-To-Byte/DouSheng/network/protos"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func Follow(ctx *gin.Context) {
	zap.S().Debugf("Follow")
	c := proto.NewRelationClient(models.Dials["relation"])

	// JWT Authorization
	var err error
	jwt := milddlewares.NewJWT()
	token := &models.TokenClaims{}
	if token, err = jwt.ParseToken(ctx.Query("token")); err != nil {
		zap.S().Panicf("Invalid token value (token: %v): %v", token, err)
		ctx.JSON(http.StatusForbidden, models.FollowResponse{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Invalid token value: %v", token),
		})
		ctx.Abort()
		return
	}

	// Parse to_user_id to int64
	var toUserID int64
	if toUserID, err = strconv.ParseInt(ctx.Query("to_user_id"), 10, 64); err != nil {
		zap.S().Panicf("Parse to_user_id value failed(id: %v): %v", ctx.Query("to_user_id"), err)
		ctx.JSON(http.StatusBadRequest, models.FollowResponse{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Parse to_user_id value failed: %v", ctx.Query("to_user_id")),
		})
		ctx.Abort()
		return
	}

	// Parse ActionType
	var actionType int64
	if actionType, err = strconv.ParseInt(ctx.Query("ActionType"), 10, 32); err != nil {
		zap.S().Panicf("Parse ActionType value failed(id: %v): %v", ctx.Query("to_user_id"), err)
		ctx.JSON(http.StatusBadRequest, models.FollowResponse{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Parse ActionType value failed: %v", ctx.Query("to_user_id")),
		})
		ctx.Abort()
		return
	}

	// 发出请求 && 处理响应
	request := proto.FollowRequest{
		UserId:     token.ID,
		ToUserId:   toUserID,
		ActionType: int32(actionType),
	}
	if response, err := c.Follow(ctx, &request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.FollowResponse{
			StatusCode: 1,
			StatusMsg:  response.StatusMsg,
		})
		ctx.Abort()
	} else {
		ctx.JSON(http.StatusOK, models.FollowResponse{
			StatusCode: 0,
			StatusMsg:  response.StatusMsg,
		})
	}

}

func FollowList(ctx *gin.Context) {
	zap.S().Debugf("FollowList")
	c := proto.NewRelationClient(models.Dials["relation"])

	// JWT Authorization
	var err error
	jwt := milddlewares.NewJWT()
	token := &models.TokenClaims{}
	if token, err = jwt.ParseToken(ctx.Query("token")); err != nil {
		zap.S().Panicf("Invalid token value(token: %v): %v", token, err)
		ctx.JSON(http.StatusForbidden, models.FollowListResponse{
			StatusCode: strconv.Itoa(1),
			StatusMsg:  fmt.Sprintf("Invalid token value: %v", token),
		})
		ctx.Abort()
		return
	}

	// 发出请求, 获取关注列表 && 处理响应
	request := proto.FollowListRequest{UserId: token.ID}
	if response, err := c.FollowList(ctx, &request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.FollowListResponse{
			StatusCode: string(response.StatusCode),
			StatusMsg:  response.StatusMsg,
		})
	} else {
		// 获得关注用户的id列表后，处理响应数据
		followListResponse := models.FollowListResponse{
			StatusCode: string(response.StatusCode),
			StatusMsg:  response.StatusMsg,
		}
		zap.S().Debugf("relation follow list: len(user_list) = %d", len(response.UserList))

		// 依据 user_id 提取 user_info
		for i := 0; i < len(response.UserList); i++ {
			user, err := getUserInfo(request.UserId, response.UserList[i])
			if err != nil {
				continue
			}
			followListResponse.UserList = append(followListResponse.UserList, user)
		}
		ctx.JSON(http.StatusOK, followListResponse)
	}
}

func FollowerList(ctx *gin.Context) {
	zap.S().Debugf("FollowerList")
	c := proto.NewRelationClient(models.Dials["relation"])

	// JWT Authorization
	var err error
	jwt := milddlewares.NewJWT()
	token := &models.TokenClaims{}
	if token, err = jwt.ParseToken(ctx.Query("token")); err != nil {
		zap.S().Panicf("Invalid token value (token: %v): %v", token, err)
		ctx.JSON(http.StatusForbidden, models.FollowerListResponse{
			StatusCode: strconv.Itoa(1),
			StatusMsg:  fmt.Sprintf("Invalid token value: %v", token),
		})
		ctx.Abort()
		return
	}

	// 发出请求, 获取粉丝列表 && 处理响应
	request := proto.FollowerListRequest{UserId: token.ID}
	if response, err := c.FollowerList(ctx, &request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.FollowerListResponse{
			StatusCode: string(response.StatusCode),
			StatusMsg:  response.StatusMsg,
		})
	} else {
		// 获得关注用户的id列表后，处理响应数据
		followerListResponse := models.FollowerListResponse{
			StatusCode: string(response.StatusCode),
			StatusMsg:  response.StatusMsg,
		}
		zap.S().Debugf("relation follower list: len(user_list) = %d", len(response.UserList))

		// 依据 user_id 提取 user_info
		for i := 0; i < len(response.UserList); i++ {
			user, err := getUserInfo(request.UserId, response.UserList[i])
			if err != nil {
				continue
			}
			followerListResponse.UserList = append(followerListResponse.UserList, user)
		}
		ctx.JSON(http.StatusOK, followerListResponse)
	}
}

func FriendList(ctx *gin.Context) {
	zap.S().Debugf("FollowerList")
	c := proto.NewRelationClient(models.Dials["relation"])

	// JWT Authorization
	var err error
	jwt := milddlewares.NewJWT()
	token := &models.TokenClaims{}
	if token, err = jwt.ParseToken(ctx.Query("token")); err != nil {
		zap.S().Panicf("Invalid token value (token: %v): %v", token, err)
		ctx.JSON(http.StatusForbidden, models.FriendListResponse{
			StatusCode: strconv.Itoa(1),
			StatusMsg:  fmt.Sprintf("Invalid token value: %v", token),
		})
		ctx.Abort()
		return
	}

	// 发出请求, 获取好友列表 && 处理响应
	request := proto.FriendListRequest{UserId: token.ID}
	if response, err := c.FriendList(ctx, &request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.FriendListResponse{
			StatusCode: string(response.StatusCode),
			StatusMsg:  response.StatusMsg,
		})
	} else {
		// 获得好友的id列表后，处理响应数据
		friendListResponse := models.FriendListResponse{
			StatusCode: string(response.StatusCode),
			StatusMsg:  response.StatusMsg,
		}
		zap.S().Debugf("relation follower list: len(user_list) = %d", len(response.UserList))

		// 依据 user_id 提取 user_info
		for i := 0; i < len(response.UserList); i++ {
			user, err := getUserInfo(request.UserId, response.UserList[i])
			if err != nil {
				continue
			}
			friendListResponse.UserList = append(friendListResponse.UserList, user)
		}
		ctx.JSON(http.StatusOK, friendListResponse)
	}
}
