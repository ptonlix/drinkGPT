### 1. "获取推荐饮品"

1. route definition

- Url: /v1/gpt/drinks/chat
- Method: POST
- Request: `DrinkGptReq`
- Response: `DrinkGptResp`

2. request definition



```golang
type DrinkGptReq struct {
	CategoryId int64 `json:"categoryId"`
	Phrases []string `json:"prases"`
}
```


3. response definition



```golang
type DrinkGptResp struct {
	Drink Goods `json:"drink"`
}

type Goods struct {
	Id int64 `json:"id"`
	CategoryId int64 `json:"categoryId"`
	Name string `json:"name"`
	Ingredients string `json:"ingredients"`
	Tea string `json:"tea"`
	CupCapacity string `json:"cup_capacity"`
	Reason string `json:"reason"`
}
```

