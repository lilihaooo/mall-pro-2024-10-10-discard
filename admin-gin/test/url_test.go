package test

type Shop struct {
	ID               string `json:"id"`
	SellNum          string `json:"sellNum"`
	Name             string `json:"name"`
	Logo             string `json:"logo"`
	CreaditLogistics string `json:"creadit_logistics"`
	CreaditProduct   string `json:"creadit_product,int32"`
	CreaditService   string `json:"creadit_service"`
}

type GoodsImage struct {
	URL       string `json:"url"`
	ProductID uint   `json:"product_id"`
}

type GoodsReceive struct {
	ID                string       `json:"id"`
	ShopID            string       `json:"shop_id"`
	Shop              Shop         `json:"shop"`
	GoodsImages       []GoodsImage `json:"goods_images"`
	Status            int32        `json:"status"`
	ProductCategoryID uint         `json:"product_category_id"`
	SalesAll          string       `json:"sales_all"`
	Sales24           interface{}  `json:"sales_24"`
	Platform          string       `json:"platform"`
	Title             string       `json:"title"`
	Commission        int32        `json:"commission"`
	BottomPrice       float64      `json:"bottom_price"`
	Url               string       `json:"url"`
	TrueSales         int32        `json:"true_sales"`

	CommissionValue float64 `json:"commission_value"`
	Remaining       int32   `json:"remaining"`
}

//func TestUrl(t *testing.T) {
//	// 将json数据转为结构体
//	var goodsRes GoodsReceive
//	err := json.Unmarshal([]byte(jsonTex), &goodsRes)
//	if err != nil {
//		t.Error(err)
//	}
//}
