package repository

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/qubic-pool/model"
	"gitlab.com/qubic-pool/pkg/db"
)

func GetUserByUsername(username string) (*model.TUser, error) {
	var user model.TUser
	err := db.Instance.Model(&model.TUser{}).Where("username=?", username).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByID(uid int64) (*model.TUser, error) {
	var user model.TUser
	err := db.Instance.Model(&model.TUser{}).Where("id=?", uid).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByMinerID(minerId string) (*model.TUser, error) {
	var user model.TUser
	err := db.Instance.Model(&model.TUser{}).Where("miner_id=?", minerId).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(username, encodePwd, minerId, phone, email string) error {
	user := model.TUser{
		MinerID:  minerId,
		Username: username,
		Password: encodePwd,
		Phone:    phone,
		Email:    email,
	}
	return db.Instance.Model(&model.TUser{}).Create(&user).Error
}

func UpdateUserInfo(username, minerId, phone, email string) error {
	var update = map[string]any{
		"miner_id": minerId,
		"phone":    phone,
		"email":    email,
	}
	return db.Instance.Model(&model.TUser{}).Where("username=?", username).Update(update).Error
}

func UpdateUserToken(uid int64, token string) error {
	var update = map[string]any{
		"token": token,
	}
	return db.Instance.Model(&model.TUser{}).Where("id=?", uid).Update(update).Error

}
