package logic

import (
	"context"
	"micro/rpc/user/userRPC"

	"micro/api/user/internal/svc"
	"micro/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.GetUserListRequest) (resp *types.GetUserListResponse, err error) {
	list, _ := l.svcCtx.UserService.GetUserList(l.ctx, &userRPC.GetUserListRequest{
		Id:         req.ID,
		Account:    req.Account,
		Nickname:   req.Nickname,
		Avatar:     req.Avatar,
		Status:     req.Status,
		Department: req.DepartmentID,
		IsAdmin:    req.IsAdmin,
		Page:       req.Offset,
		PageSize:   req.Limit,
	})
	userList := make([]types.UserResponse, 0)

	for _, user := range list.UserList {
		userList = append(userList, types.UserResponse{
			ID:           user.Id,
			Account:      user.Account,
			Nickname:     user.Nickname,
			Avatar:       user.Avatar,
			Status:       user.Status,
			DepartmentID: user.Department,
			IsAdmin:      user.IsAdmin,
			CreateTime:   user.CreateTime,
			UpdateTime:   user.UpdateTime,
		})
	}

	return &types.GetUserListResponse{
		Code: 200,
		Msg:  "成功",
		Data: types.UserListResponse{
			Total: int(list.Total),
			List:  userList,
		},
	}, nil
}
