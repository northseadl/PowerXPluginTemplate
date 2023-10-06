package handler

import (
	"net/http"

	"PluginTemplate/internal/logic"
	"PluginTemplate/internal/svc"
	"PluginTemplate/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func EmployeeInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmployeeInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewEmployeeInfoLogic(r.Context(), svcCtx)
		resp, err := l.EmployeeInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
