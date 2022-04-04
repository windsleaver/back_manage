package login

import (
	"back_manage/data/model"
	"back_manage/service/sys/cmd/api/internal/svc"
	"back_manage/service/sys/cmd/api/internal/types"
	"back_manage/utils/security"
	"back_manage/utils/token"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Login 登陆
func (l *LoginLogic) Login(req types.LoginReq) (resp *types.LoginResp, err error) {
	var (
		user          model.SysUser
		tenantAndUser []model.SysTenantAndUser
		tenantNumber  int
		userToken     string
	)
	if req.UserName == "" || req.Password == "" {
		return nil, errors.New("用户名或者密码不能为空！")
	}
	err = l.svcCtx.DbEngin.Where("user_name = ?", req.UserName).First(&user).Error
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.DbEngin.Where("user_id = ?", user.ID).Find(&tenantAndUser).Error
	if err != nil {
		return nil, err
	}
	for _, item := range tenantAndUser {
		var tenant model.SysTenant
		err = l.svcCtx.DbEngin.Where("id = ?", item.TenantId).First(&tenant).Error
		if err != nil {
			return nil, err
		}
		if !tenant.IsActive {
			tenantNumber++
		}
	}
	if tenantNumber == len(tenantAndUser) {
		return nil, errors.New("您所在的组织已被禁用，登陆失败")
	}
	if user.IsLock {
		//用户被锁定
		return nil, errors.New("您的账号已被锁定，登陆失败")
	}
	if !user.IsActive {
		//用户被禁用
		return nil, errors.New("您的账号已被禁用，登陆失败")
	}
	//判断密码
	if !security.VerifyPassword(req.Password, user.Salt, user.Password) {
		//if user.LoginNumberErr >= 5 {
		//	//每日密码输入错误5次则锁定账户，不允许登陆
		//	m := map[string]interface{}{
		//		"is_lock":          true,
		//		"login_number_err": user.LoginNumberErr + 1,
		//	}
		//	err = s.userRepo.Update(ctx, user.ID, m)
		//	if err != nil {
		//		return hsresp.HsRes(hscode.INTERNALSERVER, nil, nil)
		//	}
		//} else {
		//	//记录密码输入错误次数
		//	m := map[string]interface{}{
		//		"login_number_err": user.LoginNumberErr + 1,
		//	}
		//	err = s.userRepo.Update(ctx, user.ID, m)
		//	if err != nil {
		//		return hsresp.HsRes(hscode.INTERNALSERVER, nil, nil)
		//	}
		//}
		return nil, errors.New("用户名或密码错误")
	}
	//成功登陆
	userToken, err = token.GenerateUserBackToken(user.ID)
	if err != nil {
		return nil, errors.New("其它错误")
	}
	return &types.LoginResp{
		UserName:  user.UserName,
		UserToken: userToken,
	}, nil
}
