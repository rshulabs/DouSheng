// Author: BeYoung
// Date: 2023/2/1 0:42
// Software: GoLand

package services

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/network/models"
	proto "github.com/Go-To-Byte/DouSheng/network/protos"
	"go.uber.org/zap"
)

// 建立 user grpc 连接, 处理用户信息
func getUserInfo(userID int64, toUserId int64) (response models.User, err error) {
	zap.S().Debugf("get UserInfo: %d", userID)
	user := proto.NewUserClient(models.Dials["user"])
	relation := proto.NewRelationClient(models.Dials["relation"])
	userRequest := proto.InfoRequest{UserId: toUserId}
	followListRequest := proto.FollowListRequest{UserId: toUserId}
	followerListRequest := proto.FollowerListRequest{UserId: toUserId}
	followJudgeRequest := proto.FollowJudgeRequest{
		UserId:   userID,
		ToUserId: toUserId,
	}

	// 获取用户信息
	if r, e := user.Info(context.Background(), &userRequest); err != nil {
		zap.S().Errorf("error getting user info: (%v) ==> %v", userID, e)
	} else {
		response.Name = r.User.Name
		response.ID = r.User.Id
	}

	// 获取关注人数
	if r, e := relation.FollowList(context.Background(), &followListRequest); err != nil {
		zap.S().Errorf("error getting relation followList: (%v) ==> %v", userID, e)
	} else {
		response.FollowCount = int64(len(r.UserList))
	}

	// 获取粉丝人数
	if r, e := relation.FollowerList(context.Background(), &followerListRequest); err != nil {
		zap.S().Errorf("error getting relation followerList: (%v) ==> %v", userID, e)
	} else {
		response.FollowerCount = int64(len(r.UserList))
	}

	// 获取是否关注信息
	if r, e := relation.FollowJudge(context.Background(), &followJudgeRequest); err != nil {
		zap.S().Errorf("error getting relation followJudge: (%v) ==> %v", userID, e)
	} else {
		response.IsFollow = r.IsFollow == 1
	}

	return response, err
}
