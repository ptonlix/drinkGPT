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

type UpdateGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGoodsLogic {
	return &UpdateGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGoodsLogic) UpdateGoods(in *drinks.UpdateGoodsReq) (*drinks.UpdateGoodsResp, error) {
	goods, err := l.svcCtx.GoodsModel.FindOne(l.ctx, in.Id)
	if err == model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_NOT_FOUND), "该类型不存在，Id:%d,err:%v", in.Id, err)
	}
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "根据ID查询饮品信息失败，Id:%d,err:%v", in.Id, err)
	}
	goods.Name = in.Name
	goods.Ingredients = in.Ingredients
	goods.Tea = in.Tea
	goods.CupCapacity = in.CupCapacity
	goods.CategoryId = in.CategoryId

	err = l.svcCtx.GoodsModel.UpdateWithVersion(l.ctx, nil, goods)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Update db goods err:%v,user:%+v", err, goods)
	}

	return &drinks.UpdateGoodsResp{}, nil
}
