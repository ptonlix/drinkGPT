package goods

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/ptonlix/drinkGPT/cmd/drinks/api/internal/svc"
	"github.com/ptonlix/drinkGPT/cmd/drinks/api/internal/types"
	"github.com/ptonlix/drinkGPT/cmd/drinks/rpc/drinks"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListbycategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListbycategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListbycategoryLogic {
	return &ListbycategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListbycategoryLogic) Listbycategory(req *types.GetGoodsListByCategoryReq) (resp *types.GetGoodsListByCategoryResp, err error) {
	result, err := l.svcCtx.DrinksRpc.GetListGoodsByCategory(l.ctx, &drinks.ListGoodsByCategoryReq{CategoryId: req.CategoryId})
	if err != nil {
		return nil, err
	}

	respList := []types.Goods{}
	for _, goods := range result.GoodsList {
		goodsResp := types.Goods{}
		copier.Copy(&goodsResp, goods)
		respList = append(respList, goodsResp)
	}

	return &types.GetGoodsListByCategoryResp{GoodsList: respList}, nil
}
