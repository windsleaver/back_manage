package user

import (
	"context"

	"back_manage/service/sys/cmd/api/internal/svc"
	"back_manage/service/sys/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditUserLogic {
	return &EditUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditUserLogic) EditUser(req *types.UserReq) error {
	// todo: add your logic here and delete this line

	return nil
}
