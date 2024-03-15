package apiv1

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/qubic-pool/model"
	"gitlab.com/qubic-pool/param"
	"gitlab.com/qubic-pool/repository"
	"gitlab.com/qubic-pool/service"
	"golang.org/x/xerrors"
)

func getLoginUser(c *gin.Context) (*model.TUser, error) {
	value, ok := c.Get("loginUser")
	if !ok {
		return nil, xerrors.Errorf("not login")
	}
	user, ok := value.(model.TUser)
	if ok {
		return &user, nil
	}
	return nil, xerrors.Errorf("unexpect parse user err")
}

func UserRegister(c *gin.Context) (any, error) {
	var req param.UserRegister
	if err := c.ShouldBind(&req); err != nil {
		return nil, err
	}
	if err := service.UserRegister(req.Username, req.Password, req.MinerId, req.Phone, req.Email); err != nil {
		return nil, err
	}
	user, err := repository.GetUserByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	return map[string]int64{"uid": user.ID}, nil
}

func UserLogin(c *gin.Context) (any, error) {
	var req param.UserAuthorize
	if err := c.ShouldBind(&req); err != nil {
		return nil, err
	}
	user, err := repository.GetUserByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, xerrors.Errorf("invalid user")
	}
	if user.MinerID == "" {
		return nil, xerrors.Errorf("Authorize failed! please bind your qubic miner id")
	}
	if !service.CheckPassword(req.Password, user.Password) {
		return nil, xerrors.Errorf("Authorize failed: invalid user!")
	}
	token, err := service.UpdateUserToken(user.ID)
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"token": token,
	}, nil
}

func UserUpdateInfo(c *gin.Context) (any, error) {
	var req param.UserUpdate
	if err := c.ShouldBind(&req); err != nil {
		return nil, err
	}
	return nil, repository.UpdateUserInfo(req.Username, req.MinerId, req.Phone, req.Email)
}

func GetUserInfo(c *gin.Context) (any, error) {
	userInfo, err := getLoginUser(c)
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"id":         userInfo.ID,
		"minerId":    userInfo.MinerID,
		"username":   userInfo.Username,
		"phone":      userInfo.Phone,
		"email":      userInfo.Email,
		"token":      userInfo.Token,
		"createdAt":  userInfo.CreatedAt.Unix(),
		"createdAt2": userInfo.CreatedAt.String(),
		"updatedAt":  userInfo.UpdatedAt.Unix(),
		"updatedAt2": userInfo.UpdatedAt.String(),
	}, nil

}
