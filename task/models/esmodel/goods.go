package esmodel

import (
	"context"
	"github.com/olivere/elastic/v7"
	"task/global"
)

type Goods struct {
	ID              uint    `json:"id"`
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	CategoryID      uint    `json:"category_id"`
	CategoryName    string  `json:"category_name"`
	CoverImage      string  `json:"cover_image"`
	ShopName        string  `json:"shop_name"`
	Price           float64 `json:"price"`
	Status          int     `json:"status"`
	CouponID        uint    `json:"coupon_id"`
	CouponBeginTime string  `json:"coupon_begin_time"`
	CouponEndTime   string  `json:"coupon_end_time"`
	CouponAmount    float64 `json:"coupon_amount"`
	CouponTotal     int64   `json:"coupon_total"`
	CouponCover     int64   `json:"coupon_cover"`
	SalesAll        int32   `json:"sales_all"`
	SalesMonth      int32   `json:"sales_month"`
	SalesDay        int32   `json:"sales_day"`
	Sales2Hour      int32   `json:"sales_2_hour"`
	PlatformName    string  `json:"platform_name"`
	PlatformID      uint    `json:"platform_id"`
	MediaUID        string  `json:"media_uid"`
	MediaName       string  `json:"media_name"`
	CommissionRate  int32   `json:"commission_rate"`
	CommissionValue float64 `json:"commission_value"`
	PostCouponPrice float64 `json:"post_coupon_price"`
	Inventory       int32   `json:"inventory"`
	ExperienceScore float64 `json:"experience_score"`
	CreatedAt       string  `json:"created_at"`
	PushHot         int     `json:"push_hot"`
	DataFrom        int     `json:"data_from"`
	BrandName       string  `json:"brand_name"`
	BrandID         uint    `json:"brand_id"`
	Tags            []uint  `json:"tags"`
}

func (g *Goods) Index() string {
	return "goods_index"
}

// 返回的最大结果, 默认10000
func (g *Goods) Mapping() map[string]interface{} {
	type mi = map[string]interface{}
	var mapping = mi{
		"settings": mi{
			"number_of_shards":   3,
			"number_of_replicas": 2,
			"index": mi{
				"max_result_window": 1000,
			},
			"analysis": mi{
				"analyzer": mi{
					"text_analyzer": mi{
						"tokenizer": "ik_smart",
						"filter":    []string{"pinyin"},
					},
				},
				"filter": mi{
					"pinyin": mi{
						"type":                         "pinyin",
						"keep_full_pinyin":             false,
						"keep_joined_full_pinyin":      true,
						"keep_original":                true,
						"limit_first_letter_length":    16,
						"remove_duplicated_term":       true,
						"none_chinese_pinyin_tokenize": false,
					},
				},
			},
		},
		"mappings": mi{
			"properties": mi{
				"id": mi{
					"type": "keyword",
				},
				"title": mi{
					"type":    "text",
					"copy_to": "all",
				},
				"description": mi{
					"type":    "text",
					"copy_to": "all",
				},
				"category_id": mi{
					"type": "integer",
				},
				"data_from": mi{
					"type": "integer",
				},
				"category_name": mi{
					"type":  "keyword",
					"index": false,
				},
				"brand_name": mi{
					"type":    "text",
					"copy_to": "all",
				},
				"brand_id": mi{
					"type": "integer",
				},
				"cover_image": mi{
					"type":  "keyword",
					"index": false,
				},
				"shop_name": mi{
					"type":    "text",
					"copy_to": "all",
				},
				"price": mi{
					"type": "double",
				},
				"status": mi{
					"type": "integer",
				},
				"coupon_id": mi{
					"type":  "integer",
					"index": false,
				},
				"coupon_amount": mi{
					"type": "double",
				},
				"coupon_total": mi{
					"type": "integer",
				},
				"coupon_cover": mi{
					"type": "integer",
				},

				"coupon_begin_time": mi{
					"type":   "date",
					"format": "yyyy-MM-dd HH:mm:ss",
				},
				"coupon_end_time": mi{
					"type":   "date",
					"format": "yyyy-MM-dd HH:mm:ss",
				},
				"sales_all": mi{
					"type": "integer",
				},
				"sales_month": mi{
					"type": "integer",
				},
				"sales_day": mi{
					"type": "integer",
				},
				"sales_2_hour": mi{
					"type": "integer",
				},
				"platform_name": mi{
					"type":  "keyword",
					"index": false,
				},
				"platform_id": mi{
					"type": "integer",
				},
				"media_uid": mi{
					"type":  "keyword",
					"index": false,
				},
				"media_name": mi{
					"type":  "keyword",
					"index": false,
				},
				"commission_rate": mi{
					"type": "integer",
				},
				"commission_value": mi{
					"type": "double",
				},
				"post_coupon_price": mi{
					"type": "double",
				},
				"inventory": mi{
					"type":  "integer",
					"index": false,
				},
				"created_at": mi{
					"type":   "date",
					"format": "yyyy-MM-dd HH:mm:ss",
				},
				"push_hot": mi{
					"type": "integer",
				},
				"experience_score": mi{
					"type": "double",
				},
				"goods_data_from": mi{
					"type":  "integer",
					"index": false,
				},
				"tags": mi{
					"type": "integer",
				},
				"all": mi{
					"type":            "text",
					"analyzer":        "text_analyzer",
					"search_analyzer": "ik_smart", // 搜索的时候不能使用自定义 text_analyzer分析器
				},
			},
		},
	}
	return mapping
}

// IndexExists 索引是否存在
func (g *Goods) IndexExists() bool {
	exists, err := global.ESClient.
		IndexExists(g.Index()).
		Do(context.Background())
	if err != nil {
		global.Logrus.Error(err)
		return exists
	}
	return exists
}

// CreateIndex 创建索引
func (g *Goods) CreateIndex() error {
	if g.IndexExists() {
		// 有索引
		err := g.RemoveIndex()
		if err != nil {
			return err
		}
	}
	// 创建索引
	createIndex, err := global.ESClient.
		CreateIndex(g.Index()).
		BodyJson(g.Mapping()).
		Do(context.Background())
	if err != nil {
		global.Logrus.Error(err)
		return err
	}
	if !createIndex.Acknowledged {
		global.Logrus.Error("创建失败")
		return err
	}
	global.Logrus.Infof("索引 %s 创建成功", g.Index())
	return nil
}

// RemoveIndex 删除索引
func (g *Goods) RemoveIndex() error {
	global.Logrus.Info("索引存在，删除索引")
	// 删除索引
	indexDelete, err := global.ESClient.DeleteIndex(g.Index()).Do(context.Background())
	if err != nil {
		global.Logrus.Error(err)
		return err
	}
	if !indexDelete.Acknowledged {
		global.Logrus.Error("删除索引失败")
		return err
	}
	global.Logrus.Info("索引删除成功")
	return nil
}

// ISExistData 是否存在该文章
func (g *Goods) ISExistData() bool {
	res, err := global.ESClient.
		Search(g.Index()).
		Query(elastic.NewTermQuery("keyword", g.Title)).
		Size(1).
		Do(context.Background())
	if err != nil {
		global.Logrus.Error(err)
		return false
	}
	if res.Hits.TotalHits.Value > 0 {
		return true
	}
	return false
}

//func (p Product) EsGetArticleList(pageReq *req.PaginationReq, key string) (list []Product, count int64, err error) {
//offset := req.GetOffset(pageReq)
//if offset < 0 {
//	offset = 0
//}
//boolQuery := elastic.NewBoolQuery()
//if key != "" {
//	boolQuery.Must(
//		elastic.NewMatchQuery("title", key),
//	)
//}
//searchResult, err := global.ESClient.Search().
//	Index(g.Index()).
//	Query(boolQuery).                    // 查询条件
//	From(offset).Size(pageReq.PageSize). // 分页
//	Do(context.Background())
//if err != nil {
//	return nil, 0, err
//}
//count = searchResult.Hits.TotalHits.Value
//if count == 0 {
//	return []Product{}, count, nil
//}
//for _, hit := range searchResult.Hits.Hits {
//	err = json.Unmarshal(hit.Source, &p)
//	if err != nil {
//		return nil, 0, err
//	}
//	g.ID = hit.Id
//	list = append(list, p)
//}
//return list, count, nil

//}
