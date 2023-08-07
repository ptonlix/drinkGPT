package category

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

func (l *ListLogic) List() (*types.GetCategoryListResp, error) {
	result, err := l.svcCtx.DrinksRpc.GetListCategory(l.ctx, &drinks.ListCategoryReq{})
	if err != nil {
		return nil, err
	}
	respList := []types.Category{}
	for _, category := range result.CategoryList {
		categoryResp := types.Category{}
		copier.Copy(&categoryResp, category)
		respList = append(respList, categoryResp)
	}
	return &types.GetCategoryListResp{CategoryList: respList}, nil
}
