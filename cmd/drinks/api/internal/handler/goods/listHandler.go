package goods

import (
	"net/http"

	"github.com/ptonlix/drinkGPT/common/result"

	"github.com/ptonlix/drinkGPT/cmd/drinks/api/internal/logic/goods"
	"github.com/ptonlix/drinkGPT/cmd/drinks/api/internal/svc"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := goods.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List()
		result.HttpResult(r, w, resp, err)
	}
}
