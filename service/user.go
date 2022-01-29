package service

import (
	"context"
	"fmt"

	"gitlab.com/task-dispatcher/model"
	"gitlab.com/task-dispatcher/pkg/app"
	"gitlab.com/task-dispatcher/types"
)

const RedisKeyLoginUser = "login_user_token_"

func AppLogout(userId int64) error {
	ctx := context.Background()
	_, err := app.Redis().Del(ctx, loginUserRedisKey(userId)).Result()
	if err != nil {
		return err
	}
	return nil
}

func loginUserRedisKey(userId int64) string {
	return fmt.Sprintf("%s_%d", RedisKeyLoginUser, userId)
}

// 用户列表
func GetUserList(search types.UserListRequest, pager app.Pager) ([]types.UserItem, error) {
	users, err := model.UserList(search, pager)
	if err != nil {
		return nil, err
	}
	data := make([]types.UserItem, len(users))
	for i, u := range users {
		data[i] = types.UserItem{
			Id:     u.ID,
			Mobile: u.Mobile,
			Name:   u.Name,
		}
	}
	return data, nil
}
