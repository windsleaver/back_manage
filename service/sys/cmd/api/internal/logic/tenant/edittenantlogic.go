package tenant

import (
	"back_manage/data/model"
	"context"

	"back_manage/service/sys/cmd/api/internal/svc"
	"back_manage/service/sys/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditTenantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditTenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditTenantLogic {
	return &EditTenantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// EditTenant 编辑租户信息
func (l *EditTenantLogic) EditTenant(req *types.TeanantEditReq) error {
	err := l.svcCtx.DbEngin.Model(&model.SysTenant{}).Where("id = ?", req.Id).Updates(&model.SysTenant{
		TenantName:  req.TenantName,
		Name:        req.Name,
		Description: req.Description,
		IsActive:    req.IsActive,
		Remark:      req.Remark,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
