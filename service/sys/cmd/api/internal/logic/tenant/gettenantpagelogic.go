package tenant

import (
	"back_manage/data/model"
	"back_manage/utils/paging"
	"context"

	"back_manage/service/sys/cmd/api/internal/svc"
	"back_manage/service/sys/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTenantPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTenantPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTenantPageLogic {
	return &GetTenantPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetTenantPage 获取租户分页列表
func (l *GetTenantPageLogic) GetTenantPage(req *types.TenantPageReq) (resp *types.TenantPageResp, err error) {
	var (
		tenant []model.SysTenant
		list   []types.TenantResp
		count  int64
	)
	page := paging.PageHandle(req.Page, req.PageSize, req.Filter)
	query := l.svcCtx.DbEngin.Model(&model.SysTenant{})
	if len(req.Filter) > 0 {
		query = query.Where("tenant_name LIKE ? OR name LIKE ? ", page.Filter, page.Filter)
	}
	err = query.Count(&count).Error
	if err != nil {
		return nil, err
	}
	err = query.Offset(page.Page).Limit(page.PageSize).Order("created_at DESC").Find(&tenant).Error
	if err != nil {
		return nil, err
	}
	for _, item := range tenant {
		list = append(list, types.TenantResp{
			Id:             item.ID,
			ParentTenantId: item.ParentTenantId,
			TenantName:     item.TenantName,
			Name:           item.Name,
			Description:    item.Description,
			IsActive:       item.IsActive,
			Remark:         item.Remark,
			CreatedAt:      item.CreatedAt,
		})
	}
	resp = &types.TenantPageResp{
		Total: count,
		Items: list,
	}
	return resp, nil
}
