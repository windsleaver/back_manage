package tenant

import (
	"back_manage/data/model"
	"back_manage/service/sys/cmd/api/internal/svc"
	"back_manage/service/sys/cmd/api/internal/types"
	"back_manage/utils/idutil"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTenantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTenantLogic {
	return &CreateTenantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CreateTenant 创建租户信息
func (l *CreateTenantLogic) CreateTenant(req *types.TeanantReq) error {
	var (
		err error
		num int64
	)
	err = l.svcCtx.DbEngin.Where("tenant_name = ?", req.TenantName).Count(&num).Error
	if err != nil {
		return err
	}
	if num > 0 {
		return errors.New("当前租户信息已存在")
	}
	err = l.svcCtx.DbEngin.Create(&model.SysTenant{
		ID:             idutil.NextId(),
		ParentTenantId: req.ParentTenantId,
		TenantName:     req.TenantName,
		Name:           req.Name,
		Description:    req.Description,
		IsActive:       req.IsActive,
		Remark:         req.Remark,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
