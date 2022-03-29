package tenant

import (
	"back_manage/utils/hsresult"
	"net/http"

	"back_manage/service/sys/cmd/api/internal/logic/tenant"
	"back_manage/service/sys/cmd/api/internal/svc"
	"back_manage/service/sys/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func EditTenantHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TeanantReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := tenant.NewEditTenantLogic(r.Context(), svcCtx)
		err := l.EditTenant(&req)
		if err != nil {
			hsresult.HsResponse(w, nil, err)
		} else {
			hsresult.HsResponse(w, nil, err)
		}
	}
}
