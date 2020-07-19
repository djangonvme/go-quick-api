package service

import (
	"fmt"

	"github.com/jangozw/go-quick-api/model"
	"github.com/jangozw/go-quick-api/param"
	"github.com/jangozw/go-quick-api/pkg/app"
)

const RedisKeyLoginUser = "login_user_token_"

func AppLogout(userId int64) error {
	return app.Redis.Del(loginUserRedisKey(userId))
}

func loginUserRedisKey(userId int64) string {
	return fmt.Sprintf("%s_%d", RedisKeyLoginUser, userId)
}

// 用户列表
func GetUserList(search param.UserListRequest, pager app.Pager) ([]param.UserItem, error) {
	users, err := model.UserList(search, pager)
	if err != nil {
		return nil, err
	}
	data := make([]param.UserItem, len(users))
	for i, u := range users {
		data[i] = param.UserItem{
			Id:     u.ID,
			Mobile: u.Mobile,
			Name:   u.Name,
		}
	}
	return data, nil
}
