package category

import (
	"net/http"

	"github.com/ptonlix/drinkGPT/common/result"

	"github.com/ptonlix/drinkGPT/cmd/drinks/api/internal/logic/category"
	"github.com/ptonlix/drinkGPT/cmd/drinks/api/internal/svc"
	"github.com/ptonlix/drinkGPT/cmd/drinks/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteCategoryReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := category.NewDeleteLogic(r.Context(), svcCtx)
		err := l.Delete(&req)
		result.HttpResult(r, w, nil, err)
	}
}
