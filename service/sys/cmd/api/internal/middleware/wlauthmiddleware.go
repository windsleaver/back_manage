package middleware

import (
	"back_manage/utils/hsresult"
	"back_manage/utils/wljwt"
	"context"
	"errors"
	"net/http"
	"strings"
)

type WlAuthMiddleware struct {
}

func NewWlAuthMiddleware() *WlAuthMiddleware {
	return &WlAuthMiddleware{}
}

func (m *WlAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if len(authorization) == 0 || authorization == "" {
			w.WriteHeader(http.StatusUnauthorized)
			hsresult.HsResponse(w, nil, errors.New("token无效"))
		} else {
			pointLen := strings.Count(authorization, ".")
			if pointLen < 3 {
				hsresult.HsResponse(w, nil, errors.New("ERR_401"))
				return
			}
			tLen := strings.Index(authorization, ".")
			typeArg := authorization[0:tLen]
			token := authorization[tLen+1:]
			jwtType := wljwt.ConvertTokenArg(typeArg)
			if jwtType == 0 {
				hsresult.HsResponse(w, nil, errors.New("ERR_401"))
				return
			}
			jwtInfo, err := wljwt.ParseToken(jwtType, token)
			if err != nil {
				hsresult.HsResponse(w, nil, errors.New("ERR_401"))
				return
			}
			switch jwtType {
			case wljwt.SYSMANAGE:
				wljwt.SmAuth = jwtInfo.(*wljwt.SysManageUserClaims)
				wljwt.JType = wljwt.SYSMANAGE
			case wljwt.MERCHANT:
				wljwt.MeAuth = jwtInfo.(*wljwt.MerchantCustomClaims)
				wljwt.JType = wljwt.MERCHANT
			case wljwt.RIDER:
				wljwt.RdAuth = jwtInfo.(*wljwt.RiderCustomClaims)
				wljwt.JType = wljwt.RIDER
			case wljwt.MAINUSERAPP:
				wljwt.MuaAuth = jwtInfo.(*wljwt.MainUserAppCustomClaims)
				wljwt.JType = wljwt.MAINUSERAPP
			case wljwt.MAINUSERMINIPROGRAM:
				wljwt.MumAuth = jwtInfo.(*wljwt.MainUserMiniProgramCustomClaims)
				wljwt.JType = wljwt.MAINUSERMINIPROGRAM
			}
			ctx := r.Context()
			ctx = context.WithValue(ctx, "Authorization", jwtInfo)
			next(w, r.WithContext(ctx))
		}
	}
}
