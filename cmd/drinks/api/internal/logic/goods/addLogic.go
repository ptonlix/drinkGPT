package goods

import (
	"context"

	"github.com/ptonlix/drinkGPT/cmd/drinks/api/internal/svc"
	"github.com/ptonlix/drinkGPT/cmd/drinks/api/internal/types"
	"github.com/ptonlix/drinkGPT/cmd/drinks/rpc/drinks"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.AddGoodsReq) (*types.AddGoodsResp, error) {
	result, err := l.svcCtx.DrinksRpc.AddGoods(l.ctx, &drinks.AddGoodsReq{
		Name:        req.Name,
		CategoryId:  req.CategoryId,
		Ingredients: req.Ingredients,
		Tea:         req.Tea,
		CupCapacity: req.CupCapacity,
	})
	if err != nil {
		return nil, err
	}
	return &types.AddGoodsResp{Id: result.Id}, nil
}
