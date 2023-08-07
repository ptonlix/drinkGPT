package category

import (
	"net/http"

	"github.com/ptonlix/drinkGPT/common/result"

	"github.com/ptonlix/drinkGPT/cmd/gpt/api/internal/logic/category"
	"github.com/ptonlix/drinkGPT/cmd/gpt/api/internal/svc"
	"github.com/ptonlix/drinkGPT/cmd/gpt/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DrinkGptReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := category.NewChatLogic(r.Context(), svcCtx)
		resp, err := l.Chat(&req)
		result.HttpResult(r, w, resp, err)
	}
}
