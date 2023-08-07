package logic

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/ptonlix/drinkGPT/cmd/drinks/model"
	"github.com/ptonlix/drinkGPT/cmd/drinks/rpc/drinks"
	"github.com/ptonlix/drinkGPT/cmd/drinks/rpc/internal/svc"
	"github.com/ptonlix/drinkGPT/common/xerr"
	"github.com/ptonlix/drinkGPT/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListGoodsByCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetListGoodsByCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListGoodsByCategoryLogic {
	return &GetListGoodsByCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetListGoodsByCategoryLogic) GetListGoodsByCategory(in *drinks.ListGoodsByCategoryReq) (*drinks.ListGoodsByCategoryResp, error) {
	list, err := l.svcCtx.GoodsModel.FindAll(l.ctx, l.svcCtx.GoodsModel.RowBuilder().Where("category_id = ?", in.CategoryId), "")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "err : %v , in : %+v", err, in)
	}
	listResp := make([]*drinks.Goods, 0)

	for _, goods := range list {
		goodsResp := new(drinks.Goods)
		_ = copier.Copy(goodsResp, goods)
		listResp = append(listResp, goodsResp)
	}

	return &drinks.ListGoodsByCategoryResp{GoodsList: listResp}, nil
}
