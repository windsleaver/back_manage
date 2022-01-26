package logic

import (
	"back_manage/service/sys/cmd/api/internal/svc"
	"back_manage/service/sys/cmd/api/internal/types"
	"back_manage/utils/wljwt"
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/tal-tech/go-zero/core/logx"
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

func (l *LoginLogic) Login(req types.LoginReq) (resp *types.LoginResp, err error) {
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	loginIn := wljwt.RiderCustomClaims{
		TenantId:       "12",
		RiderId:        "321",
		UnionId:        "321",
		OpenId:         "321",
		SchoolId:       "321",
		AppleId:        "321",
		StandardClaims: jwt.StandardClaims{},
	}
	token, errs := wljwt.CreateJwtToken(wljwt.RIDER, loginIn)
	if errs != nil {
		return nil, errors.New("创建Token错误")
	}
	fmt.Println(token, accessExpire)
	return
}
