package svc

import (
	"github.com/ptonlix/drinkGPT/cmd/drinks/rpc/drinksClient"
	"github.com/ptonlix/drinkGPT/cmd/gpt/api/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	DrinksRpc drinksClient.Drinks
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		DrinksRpc: drinksClient.NewDrinks(zrpc.MustNewClient(c.DrinksRpcConf)),
	}
}
