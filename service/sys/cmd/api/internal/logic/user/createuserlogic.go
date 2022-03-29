package user

import (
	"back_manage/data/model"
	"back_manage/utils/idutil"
	"back_manage/utils/randomutil"
	"back_manage/utils/security"
	"context"
	"gorm.io/gorm"

	"back_manage/service/sys/cmd/api/internal/svc"
	"back_manage/service/sys/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CreateUser 创建用户信息
func (l *CreateUserLogic) CreateUser(req *types.UserReq) error {
	var (
		user          model.SysUser
		tenantAndUser []model.SysTenantAndUser
		//tenantAndRole []model.SysTenantAndRole
		userId string
		err    error
	)
	userId = idutil.NextId()
	for _, item := range req.TenantList {
		tenantAndUser = append(tenantAndUser, model.SysTenantAndUser{
			ID:       idutil.NextId(),
			TenantId: item.TenantId,
			UserId:   userId,
			Remark:   item.Remark,
		})
	}
	//for _, role := range req.RoleList {
	//
	//}
	l.svcCtx.DbEngin.Transaction(func(tx *gorm.DB) error {
		salt := randomutil.GetRandomString(32)
		_password, _ := security.GenderPassword(req.Password, salt)
		user = model.SysUser{
			ID:             userId,
			UserName:       req.UserName,
			Password:       _password,
			Salt:           salt,
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
		err = tx.Create(&user).Error
		if err != nil {
			return err
		}
		err = tx.Create(&tenantAndUser).Error
		if err != nil {
			return err
		}
		return nil
	})
	return nil
}
