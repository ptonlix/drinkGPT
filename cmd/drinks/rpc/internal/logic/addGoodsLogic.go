package logic

import (
	"context"

	"github.com/ptonlix/drinkGPT/cmd/drinks/model"
	"github.com/ptonlix/drinkGPT/cmd/drinks/rpc/drinks"
	"github.com/ptonlix/drinkGPT/cmd/drinks/rpc/internal/svc"
	"github.com/ptonlix/drinkGPT/common/xerr"
	"github.com/ptonlix/drinkGPT/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGoodsLogic {
	return &AddGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddGoodsLogic) AddGoods(in *drinks.AddGoodsReq) (*drinks.AddGoodsResp, error) {
	gmodel := new(model.Goods)

	gmodel.Name = in.Name
	gmodel.CategoryId = in.CategoryId
	gmodel.Ingredients = in.Ingredients
	gmodel.Tea = in.Tea
	gmodel.CupCapacity = in.CupCapacity

	insertResult, err := l.svcCtx.GoodsModel.Insert(l.ctx, nil, gmodel)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Insert db goods err:%v, goods:%+v", err, gmodel)
	}

	lastId, err := insertResult.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Insert db goods insertResult.LastInsertId err:%v,  goods:%+v", err, gmodel)
	}

	return &drinks.AddGoodsResp{Id: lastId}, nil
}
