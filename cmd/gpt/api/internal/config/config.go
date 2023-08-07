package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type GptConfig struct {
	APIKey  string
	BaseUrl string
}

type Config struct {
	rest.RestConf

	DrinksRpcConf zrpc.RpcClientConf

	GptConf GptConfig
}
