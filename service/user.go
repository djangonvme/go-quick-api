package service

import (
	"fmt"
	"gitlab.com/qubic-pool/config"
	"gitlab.com/qubic-pool/model"
	"gitlab.com/qubic-pool/pkg/jwt"
	"gitlab.com/qubic-pool/pkg/util"
	"gitlab.com/qubic-pool/repository"
	"golang.org/x/xerrors"
	"strconv"
)

const (
	PasswordMixed = "BFC2EC17-923C-2337-08D3-C2B4D1361581"
)

func UserRegister(username, password, minerId, phone, email string) error {
	exists, err := repository.GetUserByUsername(username)
	if err != nil {
		return err
	}
	if exists != nil {
		return xerrors.Errorf("user already registered by username: %s", username)
	}
	password = encodePassword(password)
	if err = repository.CreateUser(username, password, minerId, phone, email); err != nil {
		return err
	}
	return nil
}

func CheckPassword(pwd, encoded string) bool {
	return encodePassword(pwd) == encoded
}

func UpdateUserToken(uid int64) (string, error) {
	token, err := jwt.GenerateToken(jwt.DefaultSecret, fmt.Sprintf("%d", uid), config.Instance.Jwt.Expire)
	if err != nil {
		return "", err
	}
	if err = repository.UpdateUserToken(uid, token); err != nil {
		return "", err
	}
	return token, nil
}

func CheckLoginUserByToken(token string) (*model.TUser, error) {
	data, err := jwt.ParseToken(jwt.DefaultSecret, token)
	if err != nil {
		return nil, err
	}
	uid, _ := strconv.ParseInt(data.Id, 10, 64)
	if uid == 0 {
		return nil, xerrors.Errorf("invalid uid: %d", uid)
	}
	user, err := repository.GetUserByID(uid)
	if err != nil {
		return nil, err
	}
	if user == nil || user.MinerID == "" || user.Username == "" {
		return nil, xerrors.Errorf("invalid user profile")
	}
	return user, nil
}

func encodePassword(pwd string) string {
	return util.MD5String(fmt.Sprintf("%s%s", pwd, PasswordMixed))
}
