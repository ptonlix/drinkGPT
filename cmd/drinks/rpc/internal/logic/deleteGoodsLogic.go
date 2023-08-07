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

type DeleteGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGoodsLogic {
	return &DeleteGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteGoodsLogic) DeleteGoods(in *drinks.DeleteGoodsReq) (*drinks.DeleteGoodsResp, error) {
	goods, err := l.svcCtx.GoodsModel.FindOne(l.ctx, in.Id)
	if err == model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_NOT_FOUND), "该类型不存在，Id:%d,err:%v", in.Id, err)
	}
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "根据ID查询饮品信息失败，Id:%d,err:%v", in.Id, err)
	}

	err = l.svcCtx.GoodsModel.DeleteSoft(l.ctx, nil, goods)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Delete db goods err:%v, goods:%+v", err, goods)
	}

	return &drinks.DeleteGoodsResp{}, nil
}
