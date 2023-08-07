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

type AddCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCategoryLogic {
	return &AddCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCategoryLogic) AddCategory(in *drinks.AddCategoryReq) (*drinks.AddCategoryResp, error) {

	cmodel := new(model.Category)

	cmodel.Name = in.Name
	cmodel.Desc = in.Desc

	insertResult, err := l.svcCtx.CategoryModel.Insert(l.ctx, nil, cmodel)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Insert db category err:%v, category:%+v", err, cmodel)
	}

	lastId, err := insertResult.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Insert db category insertResult.LastInsertId err:%v,  category:%+v", err, cmodel)
	}

	return &drinks.AddCategoryResp{Id: lastId}, nil
}
