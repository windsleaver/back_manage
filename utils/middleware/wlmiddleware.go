package middleware

import (
	"back_manage/utils/hsresult"
	token2 "back_manage/utils/token"
	"back_manage/utils/wljwt"
	"context"
	"errors"
	"net/http"
	"strings"
)

/**
 * @Time   : 2022/1/26 23:44
 * @Author : WindsLeaver
 * @File   : wlmiddleware
 * @Version: 1.0.0
 * @Description:
 **/

type WlMiddleware struct {
	UserId string
}

func NewWlAuthMiddleware() *WlMiddleware {
	return &WlMiddleware{}
}

func (m *WlMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if len(authorization) == 0 || authorization == "" {
			w.WriteHeader(http.StatusUnauthorized)
			hsresult.HsResponse(w, nil, errors.New("token无效"))
		} else {
			pointLen := strings.Count(authorization, ".")
			if pointLen < 3 {
				hsresult.HsResponse(w, nil, errors.New("ERR_4011"))
				return
			}
			tLen := strings.Index(authorization, ".")
			typeArg := authorization[0:tLen]
			token := authorization[tLen+1:]
			jwtType := wljwt.ConvertTokenArg(typeArg)
			if jwtType == 0 {
				hsresult.HsResponse(w, nil, errors.New("ERR_4012"))
				return
			}
			jwtInfo, err := wljwt.ParseToken(jwtType, token)
			if err != nil {
				hsresult.HsResponse(w, nil, errors.New("ERR_4013"))
				return
			}
			err = token2.ParseMark(jwtInfo.(*wljwt.SysManageUserClaims).MarkInfo, typeArg)
			if err != nil {
				hsresult.HsResponse(w, nil, errors.New("ERR_4012"))
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
