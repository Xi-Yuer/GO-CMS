package authServiceModules

import (
	"errors"
	"github.com/Xi-Yuer/cms/dto"
	usersResponsiesModules "github.com/Xi-Yuer/cms/dto/modules/users"
	repositories "github.com/Xi-Yuer/cms/repositories/modules"
	userServiceModules "github.com/Xi-Yuer/cms/services/modules/users"
	"github.com/Xi-Yuer/cms/utils"
)

var AuthService = &authService{}

type authService struct {
}

// Login 登录
func (a *authService) Login(params *dto.LoginRequestParams) (*dto.LoginResponse, error) {
	user, exist := userServiceModules.UserService.FindUserByAccount(params.Account)
	if !exist {
		return nil, errors.New("账号不存在")
	}
	if user.Status == "0" {
		return nil, errors.New("账号被禁用")
	}
	// 验证密码
	if err := utils.Bcrypt.VerifyPassword(params.Password, user.Password); err != nil {
		return nil, errors.New("密码错误")
	}
	// 查找用户角色ID
	rolesID := repositories.UsersAndRolesRepositorys.FindUserRolesID(user.ID)
	// 查找用户的接口信息
	var interfaceDic []string
	for _, s := range rolesID {
		interfaceID, err := repositories.RolesAndInterfacesRepository.GetRecordByRoleID(s)
		if err != nil {
			continue
		}
		for _, s2 := range interfaceID {
			inter, e := repositories.InterfaceRepository.GetInterfaceByID(s2)
			if !e {
				continue
			}
			interfaceDic = append(interfaceDic, inter.InterfaceDic)
		}
	}
	// 接口权限去重
	interfaceDic = utils.Unique(interfaceDic)
	// 生成token
	jwtPayload := &dto.JWTPayload{
		ID:           user.ID,
		Account:      user.Account,
		IsAdmin:      user.IsAdmin,
		RoleID:       rolesID,
		InterfaceDic: interfaceDic,
		DepartmentID: user.DepartmentID,
	}
	tokenUsingHs256, err := utils.Jsonwebtoken.GenerateTokenUsingHs256(jwtPayload)
	if err != nil {
		return nil, err
	}
	return &dto.LoginResponse{
		Token: tokenUsingHs256,
		User: &usersResponsiesModules.SingleUserResponse{
			ID:           user.ID,
			Nickname:     user.Nickname,
			Avatar:       user.Avatar,
			IsAdmin:      user.IsAdmin,
			RolesID:      rolesID,
			InterfaceDic: interfaceDic,
			DepartmentID: user.DepartmentID,
			Status:       user.Status,
			Account:      user.Account,
			CreateTime:   user.CreateTime,
			UpdateTime:   user.UpdateTime,
		},
	}, nil
}
