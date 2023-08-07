package category

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/ptonlix/drinkGPT/cmd/drinks/rpc/drinks"
	"github.com/ptonlix/drinkGPT/cmd/gpt/api/internal/prompt"
	"github.com/ptonlix/drinkGPT/cmd/gpt/api/internal/svc"
	"github.com/ptonlix/drinkGPT/cmd/gpt/api/internal/types"
	"github.com/ptonlix/drinkGPT/common/util"
	"github.com/sashabaranov/go-openai"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatLogic {
	return &ChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatLogic) Chat(req *types.DrinkGptReq) (*types.DrinkGptResp, error) {
	// 根据饮品分类获取饮品种类列表
	result, err := l.svcCtx.DrinksRpc.GetListGoodsByCategory(l.ctx, &drinks.ListGoodsByCategoryReq{CategoryId: req.CategoryId})
	if err != nil {
		return nil, err
	}

	goodsList := []types.Goods{}
	for _, goods := range result.GoodsList {
		goodsResp := types.Goods{}
		copier.Copy(&goodsResp, goods)
		goodsList = append(goodsList, goodsResp)
	}
	goodsListStr, err := json.Marshal(goodsList)
	if err != nil {
		return nil, err
	}
	userInput := strings.Join(req.Phrases, "、")
	// 生成Prompt
	templatestr, err := util.GetTemplateByString(prompt.DrinkChatTemplate,
		map[string]interface{}{
			"drinkclass": template.HTML(string(goodsListStr)),
			"userinput":  userInput})
	if err != nil {
		return nil, err
	}
	fmt.Println(templatestr)
	// 调用 ChatGPT API
	cc := openai.DefaultConfig(l.svcCtx.Config.GptConf.APIKey)
	cc.BaseURL = l.svcCtx.Config.GptConf.BaseUrl
	client := openai.NewClientWithConfig(cc)
	gptresp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: templatestr,
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	resp := &types.Goods{}
	err = json.Unmarshal([]byte(gptresp.Choices[0].Message.Content), resp)
	if err != nil {
		l.Errorf("ChatGPT Return result error: %+v", gptresp)
		return nil, err
	}
	return &types.DrinkGptResp{Drink: *resp}, nil
}
