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

type UpdateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(in *drinks.UpdateCategoryReq) (*drinks.UpdateCategoryResp, error) {
	category, err := l.svcCtx.CategoryModel.FindOne(l.ctx, in.Id)
	if err == model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_NOT_FOUND), "该类型不存在，Id:%d,err:%v", in.Id, err)
	}
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "根据ID查询类别信息失败，Id:%d,err:%v", in.Id, err)
	}
	category.Name = in.Name
	category.Desc = in.Desc

	err = l.svcCtx.CategoryModel.UpdateWithVersion(l.ctx, nil, category)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Update db category err:%v,user:%+v", err, category)
	}

	return &drinks.UpdateCategoryResp{}, nil
}
