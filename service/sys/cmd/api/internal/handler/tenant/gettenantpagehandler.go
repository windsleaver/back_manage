package tenant

import (
	"back_manage/utils/hsresult"
	"net/http"

	"back_manage/service/sys/cmd/api/internal/logic/tenant"
	"back_manage/service/sys/cmd/api/internal/svc"
	"back_manage/service/sys/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetTenantPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TenantPageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := tenant.NewGetTenantPageLogic(r.Context(), svcCtx)
		resp, err := l.GetTenantPage(&req)
		if err != nil {
			hsresult.HsResponse(w, resp, err)
		} else {
			hsresult.HsResponse(w, resp, err)
		}
	}
}
