package authServiceModules

import (
	"errors"
	"github.com/Xi-Yuer/cms/dto"
	userServiceModules "github.com/Xi-Yuer/cms/services/modules/users"
	"github.com/Xi-Yuer/cms/utils"
)

var AuthService = &authService{}

type authService struct {
}

func (a *authService) Login(params *dto.LoginRequestParams) (error, string) {
	user, exist := userServiceModules.UserService.FindUserByAccount(params.Account)
	if !exist {
		return errors.New("账号不存在"), ""
	}
	// 验证密码
	err := utils.Bcrypt.VerifyPassword(params.Password, user.Password)
	if err != nil {
		return errors.New("密码错误"), ""
	}

	// TODO 查找用户页面和按钮权限并赋值
	// 生成token
	jwtPayload := &dto.JWTPayload{
		ID:               user.ID,
		NickName:         user.Nickname,
		Role:             "",
		PagePermission:   nil,
		ButtonPermission: nil,
	}
	tokenUsingHs256, err := utils.Jsonwebtoken.GenerateTokenUsingHs256(jwtPayload)
	if err != nil {
		return err, ""
	}
	return nil, tokenUsingHs256
}
