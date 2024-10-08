package process

import (
	"admin-gin/global"
	"admin-gin/model/product"
	"sync"
	"time"
)

var (
	ToMysqlJobCh = make(chan product.Goods, 100000)
	once         sync.Once
)

func toMysqlJob() {
	var batch []product.Goods
	var batchSize = 1000
	for {
		select {
		case goods, _ := <-ToMysqlJobCh:
			batch = append(batch, goods)
			if len(batch) == batchSize {
				SaveGoods(batch)
				batch = nil
			}
		case <-time.After(time.Second * 20):
			if len(batch) > 0 {
				SaveGoods(batch)
				batch = nil
			}

		}

	}
}

func SaveGoods(gs []product.Goods) {
	err := global.XTK_DB.Exec("SET sql_log_bin = 0;").Save(&gs).Error // 禁用二进制日志
	if err != nil {
		global.GVA_LOG.Error(err.Error())
	}
}

func init() {
	once.Do(func() {
		go toMysqlJob()
		go toMysqlJob()
	})
}
