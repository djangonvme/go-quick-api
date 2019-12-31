package services

import (
	"github.com/jangozw/gin-api-common/libs"
	"github.com/jangozw/gin-api-common/models"
	"github.com/jangozw/gin-api-common/params"
)

func GetUserList(search params.SearchUserList) (data params.UserList, err error) {
	var users []models.User
	var total int64
	var pageSize int64 = 20
	query := libs.Db.Model(&models.User{})
	if search.Mobile != "" {
		query = query.Where("mobile = ?", search.Mobile)
	}
	if err = query.Count(&total).Error; err != nil {
		return
	}
	limit := (search.Page - 1) * pageSize
	if err = query.Limit(pageSize).Offset(limit).Find(&users).Error; err != nil {
		return
	}
	for _, u := range users {
		data.List = append(data.List, params.UserItem{
			Id:     u.ID,
			Mobile: u.Mobile,
			Name:   u.Name,
		})
	}
	data.PageSize = pageSize
	data.Total = total
	return
}
