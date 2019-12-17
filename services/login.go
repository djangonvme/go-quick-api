package services

import (
	"errors"
	"fmt"
	"gin-api-common/configs"
	"gin-api-common/consts"
	"gin-api-common/databases"
	"gin-api-common/models"
	"gin-api-common/utils"
	"github.com/jinzhu/gorm"
	"time"
)

//try login a user and return the token to client, client need to store the receive token and put it in http header before api request
func AppLogin(account string, pwd string) (jwtToken string, err error) {
	var user models.User
	if err = databases.Db.Model(&user).Where("mobile=?", account).First(&user).Error; err != nil {
		return
	}
	if utils.Sha256(pwd) != user.Password {
		return "", errors.New("invalid account or pwd")
	}
	var token string
	if token, err = refreshUserToken(user); err != nil {
		return
	}
	//generate a jwt token for client
	return utils.GenerateJwtToken(utils.AppPayload{UserId: user.ID, UserToken: token})
}

func AppLogout(userId int64) error {
	return databases.DelKey(loginUserRedisKey(userId))
}

//Verify token from request headers every time
//1, jwtToken is valid (not expired and the sign hash is right)
//2, get user's token from jwtToken and check the user's token not expired and equal to the token in redis
func VerifyAppToken(jwtToken string) (jwt *utils.JwtCustomClaims, err error) {
	//1, verify jwt sign and expires
	if jwt, err = utils.ParseJwtToken(jwtToken); err != nil {
		return
	}
	//2,verify redis token
	var redisToken string
	if redisToken, err = redisGetLoginUser(jwt.UserId); err != nil {
		return
	}
	//ok
	if len(redisToken) > 0 && redisToken == jwt.UserToken {
		return jwt, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

//重新生成用户的token
func refreshUserToken(user models.User) (token string, err error) {
	token = utils.RandToken()
	var userToken models.UserToken
	if err = databases.Db.Model(&models.UserToken{}).Where("user_id=?", user.ID).First(&userToken).Error; err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	expSeconds := configs.GetTokenExpireSeconds()
	if userToken.ID == 0 {
		t := models.UserToken{
			UserID:    user.ID,
			Token:     token,
			ExpiredAt: time.Now().Unix() + expSeconds,
		}
		err = databases.Db.Create(&t).Error
	} else {
		err = databases.Db.Model(&models.UserToken{}).Where("id=?", userToken.ID).Update(map[string]interface{}{
			"expired_at": time.Now().Unix() + expSeconds,
			"token":      token,
		}).Error
	}
	if err != nil {
		return "", errors.New("create user token data failed")
	}
	return token, redisSetLoginUser(user.ID, token, expSeconds)
}

//set user's token in expires
func redisSetLoginUser(userId int64, token string, exp int64) error {
	if err := databases.SetKey(loginUserRedisKey(userId), token, exp); err != nil {
		return errors.New(fmt.Sprintf("redis set login user failed:%d, %s", userId, err.Error()))
	}
	return nil
}

//
func redisGetLoginUser(userId int64) (string, error) {
	return databases.GetKey(loginUserRedisKey(userId))
}
func loginUserRedisKey(userId int64) string {
	return fmt.Sprintf("%s_%d", consts.RedisKeyLoginUser, userId)
}
