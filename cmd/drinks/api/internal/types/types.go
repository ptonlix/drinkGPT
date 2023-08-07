// Code generated by goctl. DO NOT EDIT.
package types

type Category struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type GetCategoryListResp struct {
	CategoryList []Category `json:"categoryList"`
}

type AddCategoryReq struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type AddCategoryResp struct {
	Id int64 `json:"id"`
}

type UpdateCategoryReq struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type DeleteCategoryReq struct {
	Id int64 `json:"id"`
}

type Goods struct {
	Id          int64  `json:"id"`
	CategoryId  int64  `json:"categoryId"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
	Tea         string `json:"tea"`
	CupCapacity string `json:"cup_capacity"`
}

type GetGoodsListResp struct {
	GoodsList []Goods `json:"list"`
}

type GetGoodsListByCategoryReq struct {
	CategoryId int64 `form:"categoryId"`
}

type GetGoodsListByCategoryResp struct {
	GoodsList []Goods `json:"list"`
}

type AddGoodsReq struct {
	CategoryId  int64  `json:"categoryId"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
	Tea         string `json:"tea"`
	CupCapacity string `json:"cup_capacity"`
}

type AddGoodsResp struct {
	Id int64 `json:"id"`
}

type UpdateGoodsReq struct {
	Id          int64  `json:"id"`
	CategoryId  int64  `json:"categoryId"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
	Tea         string `json:"tea"`
	CupCapacity string `json:"cup_capacity"`
}

type DeleteGoodsReq struct {
	Id int64 `json:"id"`
}