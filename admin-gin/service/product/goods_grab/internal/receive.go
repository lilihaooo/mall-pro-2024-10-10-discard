package internal

import (
	"admin-gin/model/product"
	"strconv"
)

var (
	DOU_YIN   = 1
	KUAI_SHOU = 2
	WEI_XIN   = 3
)

type Shop struct {
	ID               interface{} `json:"id"`
	Name             interface{} `json:"name"`
	Logo             interface{} `json:"logo"`
	CreaditLogistics interface{} `json:"creadit_logistics"`
	CreaditProduct   interface{} `json:"creadit_product,int32"`
	CreaditService   interface{} `json:"creadit_service"`
}

type GoodsImage struct {
	URL interface{} `json:"url"`
}

type GoodsReceive struct {
	ID                interface{}  `json:"id"`
	ShopID            interface{}  `json:"shop_id"`
	Shop              Shop         `json:"shop"`
	GoodsImages       []GoodsImage `json:"productImages"`
	Status            interface{}  `json:"status"`
	ProductCategoryID interface{}  `json:"product_category_id"`
	SalesAll          interface{}  `json:"sales_all"`
	Sales24           interface{}  `json:"sales_24"`
	Platform          interface{}  `json:"platform"`
	Title             interface{}  `json:"title"`
	Commission        interface{}  `json:"commission"`
	BottomPrice       interface{}  `json:"bottom_price"`
	Url               interface{}  `json:"url"`
	Remaining         interface{}  `json:"remaining"`
}

func (p *GoodsReceive) GoodsReceive2model() (goods product.Goods) {
	// 商品赋值
	goods.ShopID = uint(I2Int(p.ShopID))
	goods.ID = uint(I2Int(p.ID))
	goods.Status = I2Int(p.Status)
	goods.SalesAll = int32(I2Int(p.SalesAll))
	goods.Title = I2Str(p.Title)
	goods.CommissionRate = int32(I2Int(p.Commission))
	goods.Price = I2Float64(p.BottomPrice)
	goods.Inventory = int32(I2Int(p.Remaining))
	goods.MediaUID = nil
	goods.PlatformID = I2UintByPlatform(p.Platform)
	// 店铺信息赋值
	goods.Shop.ID = uint(I2Int(p.Shop.ID))
	goods.Shop.Name = I2Str(p.Shop.Name)
	goods.Shop.Logo = I2Str(p.Shop.Logo)
	goods.Shop.LogisticsScore = I2Float64(p.Shop.CreaditLogistics)
	goods.Shop.ProductScore = I2Float64(p.Shop.CreaditProduct)
	goods.Shop.ServiceScore = I2Float64(p.Shop.CreaditService)
	// 图片赋值
	var imgs []product.Image
	for _, one := range p.GoodsImages {
		img := product.Image{
			Url: I2Str(one.URL),
		}
		imgs = append(imgs, img)
	}
	goods.Images = imgs

	return
}

func I2UintByPlatform(i interface{}) uint {
	id := 0
	switch v := i.(type) {
	case string:
		switch v {
		case "douyin":
			id = DOU_YIN
		case "kuaishou":
			id = KUAI_SHOU
		case "weixin":
			id = WEI_XIN
		}
	}
	return uint(id)
}

func I2Int(i interface{}) int {
	switch v := i.(type) {
	case int:
		return v
	case string:
		num, err := strconv.Atoi(v)
		if err != nil {
			return 0
		}
		return num
	case float64:
		num := int(v)
		return num
	default:
		return 0
	}
}

func I2Float64(i interface{}) float64 {
	switch v := i.(type) {
	case float64:
		return v
	case string:
		num, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0
		}
		return num
	default:
		return 0
	}
}

func I2Str(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	default:
		return ""
	}
}
