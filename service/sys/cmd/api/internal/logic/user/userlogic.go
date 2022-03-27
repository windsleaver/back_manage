package user

import (
	"back_manage/data/model"
	"back_manage/service/sys/cmd/api/internal/svc"
	"back_manage/service/sys/cmd/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserLogic {
	return UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req types.UserReq) error {
	var (
		err error
	)
	user := model.SysUser{
		ID:             "1111",
		UserName:       req.UserName,
		Password:       req.Password,
		Salt:           "",
		Name:           req.Name,
		HeadImg:        req.HeadImg,
		Phone:          req.Phone,
		Email:          req.Email,
		IsBindWechat:   false,
		UnionId:        "",
		OpenId:         "",
		IsActive:       false,
		IsAdmin:        false,
		LastLoginTime:  0,
		LoginNumberErr: 0,
		IsLock:         false,
	}
	err = l.svcCtx.DbEngin.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
