package category

import (
	"context"

	"github.com/ptonlix/drinkGPT/cmd/drinks/api/internal/svc"
	"github.com/ptonlix/drinkGPT/cmd/drinks/api/internal/types"
	"github.com/ptonlix/drinkGPT/cmd/drinks/rpc/drinks"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateCategoryReq) error {
	_, err := l.svcCtx.DrinksRpc.UpdateCategory(l.ctx, &drinks.UpdateCategoryReq{
		Id:   req.Id,
		Name: req.Name,
		Desc: req.Desc,
	})
	if err != nil {
		return err
	}
	return nil
}
