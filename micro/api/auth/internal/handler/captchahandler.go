package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"micro/api/auth/internal/logic"
	"micro/api/auth/internal/svc"
	"micro/api/auth/internal/types"
)

func CaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmptyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.Captcha(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
