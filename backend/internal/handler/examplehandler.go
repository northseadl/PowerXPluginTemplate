package handler

import (
	"PluginTemplate/internal/logic"
	"PluginTemplate/internal/svc"
	"PluginTemplate/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func ExampleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExampleRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewExampleLogic(r.Context(), svcCtx)
		resp, err := l.Example(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
