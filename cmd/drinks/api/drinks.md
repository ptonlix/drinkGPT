### 1. "获取饮品种类列表"

1. route definition

- Url: /v1/drinks/category/list
- Method: GET
- Request: `-`
- Response: `GetCategoryListResp`

2. request definition



3. response definition



```golang
type GetCategoryListResp struct {
	CategoryList []Category `json:"categoryList"`
}
```

### 2. "新增饮品种类"

1. route definition

- Url: /v1/drinks/category/add
- Method: POST
- Request: `AddCategoryReq`
- Response: `AddCategoryResp`

2. request definition



```golang
type AddCategoryReq struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}
```


3. response definition



```golang
type AddCategoryResp struct {
	Id int64 `json:"id"`
}
```

### 3. "删除饮品种类"

1. route definition

- Url: /v1/drinks/category/delete
- Method: POST
- Request: `DeleteCategoryReq`
- Response: `-`

2. request definition



```golang
type DeleteCategoryReq struct {
	Id int64 `json:"id"`
}
```


3. response definition


### 4. "更新饮品种类"

1. route definition

- Url: /v1/drinks/category/update
- Method: POST
- Request: `UpdateCategoryReq`
- Response: `-`

2. request definition



```golang
type UpdateCategoryReq struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}
```


3. response definition


### 5. "获取饮品实例列表"

1. route definition

- Url: /v1/drinks/goods/list
- Method: GET
- Request: `-`
- Response: `GetGoodsListResp`

2. request definition



3. response definition



```golang
type GetGoodsListResp struct {
	GoodsList []Goods `json:"list"`
}
```

### 6. "新增饮品实例"

1. route definition

- Url: /v1/drinks/goods/add
- Method: POST
- Request: `AddGoodsReq`
- Response: `AddGoodsResp`

2. request definition



```golang
type AddGoodsReq struct {
	CategoryId int64 `json:"categoryId"`
	Name string `json:"name"`
	Ingredients string `json:"ingredients"`
	Tea string `json:"tea"`
	CupCapacity string `json:"cup_capacity"`
}
```


3. response definition



```golang
type AddGoodsResp struct {
	Id int64 `json:"id"`
}
```

### 7. "删除饮品实例"

1. route definition

- Url: /v1/drinks/goods/delete
- Method: POST
- Request: `DeleteGoodsReq`
- Response: `-`

2. request definition



```golang
type DeleteGoodsReq struct {
	Id int64 `json:"id"`
}
```


3. response definition


### 8. "更新饮品实例"

1. route definition

- Url: /v1/drinks/goods/update
- Method: POST
- Request: `UpdateGoodsReq`
- Response: `-`

2. request definition



```golang
type UpdateGoodsReq struct {
	Id int64 `json:"id"`
	CategoryId int64 `json:"categoryId"`
	Name string `json:"name"`
	Ingredients string `json:"ingredients"`
	Tea string `json:"tea"`
	CupCapacity string `json:"cup_capacity"`
}
```


3. response definition


