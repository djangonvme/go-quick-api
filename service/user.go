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
func GetUserList(search param.UserListRequest, pager *app.Pager) (data []param.UserItem, err error) {
	var users []model.User
	query := app.Db.Model(&model.User{})
	if search.Mobile != "" {
		query = query.Where("mobile = ?", search.Mobile)
	}
	if err = query.Count(&pager.Total).Error; err != nil {
		return
	}
	if err = query.Limit(pager.Limit()).Offset(pager.Offset()).Find(&users).Error; err != nil {
		return
	}
	for _, u := range users {
		data = append(data, param.UserItem{
			Id:     u.ID,
			Mobile: u.Mobile,
			Name:   u.Name,
		})
	}
	// data.SetPager(search.Page, search.PageSize, total)
	return
}

