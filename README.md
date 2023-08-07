# DrinkGPT
English | [简体中文](./README-zh.md)
<p>
	<p align="center">
		<img height=280 src="https://storage.googleapis.com/gopherizeme.appspot.com/gophers/978de269528aebbc31862cb187c41a4e1e40322a.png">
	</p>
	<p align="center">
		<font size=6 face="宋体">Prompt实战：ChatGPT饮品推荐</font>
	<p>
</p>
<p align="center">
<img alt="Go" src="https://img.shields.io/badge/Go-1.18%2B-blue">
<img alt="Mysql" src="https://img.shields.io/badge/Mysql-5.7%2B-brightgreen">
<img alt="Redis" src="https://img.shields.io/badge/Redis-6.2%2B-yellowgreen">
<img alt="go-zero" src="https://img.shields.io/badge/go--zero-1.4.1-orange">
<img alt="OpenAI" src="https://img.shields.io/badge/OpenAI-API-brightgreeni"/>
<img alt="docker" src="https://img.shields.io/badge/docker-2.20.0-yellow"/>
<img alt="license" src="https://img.shields.io/badge/license-MIT-lightgrey"/>
</p>

## 前言
![](https://img.gejiba.com/images/56a9e5f9c65496d7a3d12c1db957a149.png)
之前写了一篇Prompt实战 [《大语言模型Prompt实战-喜茶饮品推荐》](https://mp.weixin.qq.com/s?__biz=MzkxODM4NzM1OQ==&mid=2247485678&idx=1&sn=f6ceb40cb01cc75528fc365babe2b0bf&chksm=c1b3671ef6c4ee089e9d22c7b36dec844b2efaf76a196214ec20e8adc937ec03c97cb458b976&token=1882951763&lang=zh_CN#rd) 反响还不错，但主要是通过ChatGPT和ChatGLM的UI界面进行测试的，通过API交互还没有提到。
所以这里就针对API交互提供一个参考案例供大家参考

## 架构

![](https://img.gejiba.com/images/e2577f4b4314dead58f6edaf50dbd2f8.png)

目前项目关键有三个服务 drinkapi drinkrpc gptapi
1. drinkapi 为新增饮品种类（如喜茶等）、新增饮品和编辑、删除等API请求操作
2. drinkrpc 为drinkapi的CURD实现  
3. gptapi   向用户提供Chat 的api服务器，负责整合Prompt向openai等服务请求  

该三个组件分别在cmd 目录下，采用go-zero框架，通过goctl生成框架代码。详情见deploy/goctl

## 环境依赖

Mysql v5.7+

Redis v6.2.6

etcd v3.5.5

go-zero v1.4.1

Note: 请选择合适的版本进行安装,否则服务启动失败。 go-zero可以参考[官方文档](https://github.com/zeromicro/go-zero)

## 运行
下载依赖（已安装忽略）
```shell
git clone -b v3.5.0 https://github.com/etcd-io/etcd.git
cd etcd
./build.sh
export PATH="$PATH:`pwd`/bin"
etcd --version
```

```shell

#1. 克隆项目
git clone https://github.com/ptonlix/drinkGPT.git

# 2. 启动项目依赖
## 启动Mysql Redis
## 略
## 运行etcd服务(服务发现)
etcd

# 3.修改项目配置文件,运行服务
./cmd/drinks/api/etc/drinks.yaml
./cmd/drinks/rpc/etc/drinks.yaml
./cmd/gpt/api/etc/gpt.yaml
## 修改上述三个配置的Mysql Redis etcd 等地址 
## 还有ChatGPT APIKey BaseUrl 等信息

# 运行drinkrpc 服务
cd ../drinks/rpc
go run drinks.go -f ./etc/drinks.yaml

# 运行drinkapi 服务
cd ../drinks/api
go run drinks.go -f ./etc/drinks.yaml

# 运行gpt服务
cd ../gpt/api
go run gpt.go -f ./etc/gpt.yaml

## 通过API请求，新增饮品种类和饮品信息
## 推荐使用APIFOX请求
```

## API请求参考
[drinks API](./cmd/drinks/api/drinks.md)   
[gpt API](./cmd/gpt/api/gpt.md)


<p align="center">
  <b>SPONSORED BY</b>
</p>
<p align="center">
   <a href="https://www.gogeek.com.cn/" title="gogeek" target="_blank">
      <img height="200px" src="https://img.gejiba.com/images/96b6d150bd758b13d66aec66cb18044e.jpg" title="gogeek">
   </a>
</p>