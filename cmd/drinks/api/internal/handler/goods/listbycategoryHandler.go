package goods

import (
	"net/http"

	"github.com/ptonlix/drinkGPT/common/result"

	"github.com/ptonlix/drinkGPT/cmd/drinks/api/internal/logic/goods"
	"github.com/ptonlix/drinkGPT/cmd/drinks/api/internal/svc"
	"github.com/ptonlix/drinkGPT/cmd/drinks/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListbycategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetGoodsListByCategoryReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := goods.NewListbycategoryLogic(r.Context(), svcCtx)
		resp, err := l.Listbycategory(&req)
		result.HttpResult(r, w, resp, err)
	}
}
