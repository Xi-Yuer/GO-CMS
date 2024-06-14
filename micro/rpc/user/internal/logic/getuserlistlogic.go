package logic

import (
	"context"
	userModel "micro/model/user"
	"micro/rpc/user/internal/svc"
	"micro/rpc/user/userRPC"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserListLogic) GetUserList(in *userRPC.GetUserListRequest) (*userRPC.GetUserListResponse, error) {
	var userList []userModel.User
	var total int64
	err := l.svcCtx.GormDB.Where(&userModel.User{
		ID:           in.Id,
		Account:      in.Account,
		NickName:     in.Nickname,
		Avatar:       in.Avatar,
		Status:       in.Status,
		DepartmentID: in.Department,
		IsAdmin:      in.IsAdmin,
	}).Offset(int(in.Page)).Limit(int(in.PageSize)).Find(&userList).Count(&total).Error

	if err != nil {
		return nil, err
	}
	userListRPC := make([]*userRPC.GetUserResponse, 0)

	for _, v := range userList {
		userListRPC = append(userListRPC, &userRPC.GetUserResponse{
			Id:       v.ID,
			Account:  v.Account,
			Nickname: v.NickName,
			Avatar:   v.Avatar,
			Status:   v.Status,

			Department: v.DepartmentID,
			IsAdmin:    v.IsAdmin,
			CreateTime: v.CreatedAt.String(),
			UpdateTime: v.UpdatedAt.String(),
		})
	}

	return &userRPC.GetUserListResponse{
		UserList: userListRPC,
		Total:    int32(total),
	}, nil
}
