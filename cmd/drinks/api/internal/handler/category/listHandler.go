package category

import (
	"net/http"

	"github.com/ptonlix/drinkGPT/common/result"

	"github.com/ptonlix/drinkGPT/cmd/drinks/api/internal/logic/category"
	"github.com/ptonlix/drinkGPT/cmd/drinks/api/internal/svc"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := category.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List()
		result.HttpResult(r, w, resp, err)
	}
}
