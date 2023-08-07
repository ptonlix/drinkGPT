package category

import (
	"context"

	"github.com/ptonlix/drinkGPT/cmd/drinks/api/internal/svc"
	"github.com/ptonlix/drinkGPT/cmd/drinks/api/internal/types"
	"github.com/ptonlix/drinkGPT/cmd/drinks/rpc/drinks"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.AddCategoryReq) (*types.AddCategoryResp, error) {
	result, err := l.svcCtx.DrinksRpc.AddCategory(l.ctx, &drinks.AddCategoryReq{
		Name: req.Name,
		Desc: req.Desc,
	})
	if err != nil {
		return nil, err
	}

	return &types.AddCategoryResp{Id: result.Id}, nil
}
