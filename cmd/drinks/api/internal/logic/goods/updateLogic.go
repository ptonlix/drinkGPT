package goods

import (
	"context"

	"github.com/ptonlix/drinkGPT/cmd/drinks/api/internal/svc"
	"github.com/ptonlix/drinkGPT/cmd/drinks/api/internal/types"
	"github.com/ptonlix/drinkGPT/cmd/drinks/rpc/drinks"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateGoodsReq) error {
	_, err := l.svcCtx.DrinksRpc.UpdateGoods(l.ctx, &drinks.UpdateGoodsReq{
		Id:          req.Id,
		Name:        req.Name,
		CategoryId:  req.CategoryId,
		Ingredients: req.Ingredients,
		Tea:         req.Tea,
		CupCapacity: req.CupCapacity,
	})
	if err != nil {
		return err
	}

	return nil
}
