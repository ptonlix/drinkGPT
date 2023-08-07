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

type GetListGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetListGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListGoodsLogic {
	return &GetListGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetListGoodsLogic) GetListGoods(in *drinks.ListGoodsReq) (*drinks.ListGoodsResp, error) {
	list, err := l.svcCtx.GoodsModel.FindAll(l.ctx, l.svcCtx.GoodsModel.RowBuilder(), "")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "err : %v , in : %+v", err, in)
	}

	listResp := make([]*drinks.Goods, 0)

	for _, goods := range list {
		goodsResp := new(drinks.Goods)
		_ = copier.Copy(goodsResp, goods)
		listResp = append(listResp, goodsResp)
	}

	return &drinks.ListGoodsResp{GoodsList: listResp}, nil
}
