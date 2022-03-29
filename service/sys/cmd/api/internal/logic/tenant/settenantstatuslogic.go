package tenant

import (
	"back_manage/data/model"
	"context"

	"back_manage/service/sys/cmd/api/internal/svc"
	"back_manage/service/sys/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetTenantStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetTenantStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetTenantStatusLogic {
	return &SetTenantStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// SetTenantStatus 修改租户状态
func (l *SetTenantStatusLogic) SetTenantStatus(req *types.TeanantStatusReq) error {
	var (
		tenant model.SysTenant
		err    error
	)
	err = l.svcCtx.DbEngin.Where("id = ?", req.Id).First(&tenant).Error
	if err != nil {
		return err
	}
	err = l.svcCtx.DbEngin.Model(&model.SysTenant{}).Where("id = ?", req.Id).Update("is_active", !tenant.IsActive).Error
	if err != nil {
		return err
	}
	return nil
}
