package goods

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/ptonlix/drinkGPT/cmd/drinks/api/internal/svc"
	"github.com/ptonlix/drinkGPT/cmd/drinks/api/internal/types"
	"github.com/ptonlix/drinkGPT/cmd/drinks/rpc/drinks"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List() (*types.GetGoodsListResp, error) {
	result, err := l.svcCtx.DrinksRpc.GetListGoods(l.ctx, &drinks.ListGoodsReq{})
	if err != nil {
		return nil, err
	}

	respList := []types.Goods{}
	for _, goods := range result.GoodsList {
		goodsResp := types.Goods{}
		copier.Copy(&goodsResp, goods)
		respList = append(respList, goodsResp)
	}

	return &types.GetGoodsListResp{GoodsList: respList}, nil
}
