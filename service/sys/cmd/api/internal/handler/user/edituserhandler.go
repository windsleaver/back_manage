package user

import (
	"back_manage/utils/hsresult"
	"net/http"

	"back_manage/service/sys/cmd/api/internal/logic/user"
	"back_manage/service/sys/cmd/api/internal/svc"
	"back_manage/service/sys/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func EditUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewEditUserLogic(r.Context(), svcCtx)
		err := l.EditUser(&req)
		if err != nil {
			hsresult.HsResponse(w, nil, err)
		} else {
			hsresult.HsResponse(w, nil, err)
		}
	}
}
