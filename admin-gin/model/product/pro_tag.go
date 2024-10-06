package product

type Tag struct {
	ID    uint    `gorm:"column:id;primaryKey;autoIncrement;type:bigint;comment:'标签ID'" json:"id"`
	Title string  `gorm:"column:title;type:varchar(30);not null;comment:'标签标题'" json:"title"`
	Group string  `gorm:"column:group;type:varchar(20);not null;comment:'标签分组'" json:"group"`
	Goods []Goods `gorm:"many2many:goods_tag;foreignKey:ID;joinForeignKey:TagID;References:ID;joinReferences:GoodsID;comment:'标签关联的商品'" json:"goods"`
}

//var jsonData = `
//[
//  { "id": 36, "title": "精选推荐" },
//  { "id": 1, "title": "低价好卖" },
//  { "id": 3, "title": "买家好评" },
//  { "id": 5, "title": "精推素材" },
//  { "id": 6, "title": "朋友圈素材" },
//  { "id": 7, "title": "短视频" },
//  { "id": 8, "title": "运费险" },
//  { "id": 9, "title": "花呗免息" },
//  { "id": 10, "title": "破损保障" },
//  { "id": 12, "title": "西藏" },
//  { "id": 11, "title": "新疆" },
//  { "id": 13, "title": "内蒙" },
//  { "id": 16, "title": "甘肃" },
//  { "id": 17, "title": "宁夏" },
//  { "id": 14, "title": "海南" },
//  { "id": 18, "title": "官方旗舰店" },
//  { "id": 19, "title": "阿里健康大药房" },
//  { "id": 20, "title": "天猫国际" },
//  { "id": 21, "title": "进口超市" },
//  { "id": 22, "title": "淘宝冠店" },
//  { "id": 23, "title": "淘宝企业店铺" },
//  { "id": 24, "title": "定向计划" },
//  { "id": 25, "title": "聚划算" },
//  { "id": 26, "title": "拍多件" },
//  { "id": 27, "title": "前N件" },
//  { "id": 28, "title": "额外满减" },
//  { "id": 29, "title": "买赠活动" },
//  { "id": 30, "title": "首单礼金" },
//  { "id": 31, "title": "实时榜单" },
//  { "id": 32, "title": "30天榜单" },
//  { "id": 33, "title": "30天新品" },
//  { "id": 34, "title": "30天低价" },
//  { "id": 35, "title": "30天高佣金" }
//]
//`
