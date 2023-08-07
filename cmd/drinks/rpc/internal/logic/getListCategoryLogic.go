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

type GetListCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetListCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListCategoryLogic {
	return &GetListCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetListCategoryLogic) GetListCategory(in *drinks.ListCategoryReq) (*drinks.ListCategoryResp, error) {
	list, err := l.svcCtx.CategoryModel.FindAll(l.ctx, l.svcCtx.CategoryModel.RowBuilder(), "")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "err : %v , in : %+v", err, in)
	}
	listResp := make([]*drinks.Category, 0)

	for _, category := range list {
		categoryResp := new(drinks.Category)
		_ = copier.Copy(categoryResp, category)
		listResp = append(listResp, categoryResp)
	}
	l.Error(123, len(listResp))
	return &drinks.ListCategoryResp{CategoryList: listResp}, nil
}
