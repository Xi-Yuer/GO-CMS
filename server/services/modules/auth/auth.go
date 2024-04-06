package authServiceModules

import (
	"errors"
	"github.com/Xi-Yuer/cms/dto"
	pagesResponsiesModules "github.com/Xi-Yuer/cms/dto/modules/pages"
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
	// 验证密码
	if err := utils.Bcrypt.VerifyPassword(params.Password, user.Password); err != nil {
		return nil, errors.New("密码错误")
	}
	// 查找用户角色ID
	rolesID := repositories.UsersAndRolesRepositorys.FindUserRolesID(user.ID)
	// 生成token
	jwtPayload := &dto.JWTPayload{
		ID:           user.ID,
		NickName:     user.Nickname,
		RoleID:       rolesID,
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
			RolesID:      rolesID,
			DepartmentID: user.DepartmentID,
			Status:       user.Status,
			Account:      user.Account,
			CreateTime:   user.CreateTime,
			UpdateTime:   user.UpdateTime,
		},
	}, nil
}

// CreateUserRoleRecord 给用户分配角色信息
func (a *authService) CreateUserRoleRecord(id string, params *dto.AuthorizationManagementParams) error {
	// 检查角色是否存在
	if err := repositories.RoleRepositorysModules.CheckRolesExistence(params.RoleID); err != nil {
		return err
	}
	// 插入数据
	return repositories.UsersAndRolesRepositorys.CreateRecords(id, params.RoleID)
}

// CreateRolePermissionsRecord 给角色分配权限
func (a *authService) CreateRolePermissionsRecord(params *dto.CreateRolePermissionRecordParams) error {
	// 检查角色是否存在
	if err := repositories.RoleRepositorysModules.CheckRolesExistence([]string{params.RoleID}); err != nil {
		return err
	}
	// 检查页面是否存在
	if err := repositories.PageRepositorysModules.CheckPagesExistence(params.PageID); err != nil {
		return err
	}
	// 插入数据
	return repositories.RolesAndPagesRepository.CreateRecord(params)
}

func (a *authService) GetUserMenus(id string) ([]*pagesResponsiesModules.SinglePageResponse, error) {
	// 查找用户角色ID
	rolesID := repositories.UsersAndRolesRepositorys.FindUserRolesID(id)
	// 查找用户页面权限
	var pagesID []string
	for _, roleID := range rolesID {
		page, err := repositories.RolesAndPagesRepository.GetRecordsByRoleID(roleID)
		if err != nil {
			return nil, err
		}
		pagesID = append(pagesID, page...)
	}
	// pagesID 去重
	pagesID = utils.Unique(pagesID)
	// 获取页面详情
	var pagesDetail []*dto.SinglePageResponse
	for _, pageID := range pagesID {
		pageDetail, err := repositories.PageRepositorysModules.FindPageByID(pageID)
		if err != nil {
			return nil, err
		}
		pagesDetail = append(pagesDetail, pageDetail)
	}
	menu := utils.BuildPages(pagesDetail)
	return menu, nil
}
