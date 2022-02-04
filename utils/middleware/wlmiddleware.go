package middleware

import (
	"back_manage/utils/hsresult"
	"context"
	"errors"
	"net/http"
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
			//如果没有传递token获取token为空直接返回http状态码401
			w.WriteHeader(http.StatusUnauthorized)
			hsresult.HsResponse(w, nil, errors.New("200"))
		} else {
			ctx := r.Context()
			ctx = context.WithValue(ctx, "Authorization", authorization)
			next(w, r.WithContext(ctx))
		}

	}
}
