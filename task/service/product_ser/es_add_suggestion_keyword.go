package product_ser

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/xuri/excelize/v2"
	"log"
	"math"
	"strconv"
	"strings"
	"task/global"
	"task/models/esmodel"
	"task/models/gormmodel"
	"time"
)

type SuggestionService struct{}

func (s *SuggestionService) EsAddSuggestionKeyword() {
	start := time.Now() // 记录开始时间

	addCategory()
	addTxt()

	elapsed := time.Since(start) // 计算耗时
	fmt.Printf("执行时间: %s\n", elapsed)
}

func addCategory() {
	// 插入分类名的词组
	var categoryName []string
	err := global.DB.Model(gormmodel.Category{}).Select("name").Scan(&categoryName).Error
	if err != nil {
		global.Logrus.Error(err.Error())
	}
	var keyword []string
	for _, one := range categoryName {
		if strings.Contains(one, "/") {
			parts := strings.Split(one, "/")
			keyword = append(keyword, parts...)
		} else {
			keyword = append(keyword, one)
		}
	}
	// 将品牌信息加入到建议词汇中
	var brandName []string
	err = global.DB.Model(gormmodel.Brand{}).Debug().Pluck("name", &brandName).Error
	if err != nil {
		global.Logrus.Error(err.Error())
	}
	keyword = append(keyword, brandName...)

	// 写入es中, 将关键词作为id避免重复插入
	var esM esmodel.Suggestion
	for _, k := range keyword {
		_, err = global.ESClient.Index().
			Index(esM.Index()).
			Id(k).
			BodyJson(map[string]string{"keyword": k}).
			Do(context.Background())
		if err != nil {
			global.Logrus.Error(err)
		}
	}
	fmt.Println("添加完成")
}

func addTxt() {
	f, err := excelize.OpenFile("./resource/keyword.xlsx")
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
		return
	}
	defer func() {
		// 关闭文件
		if err = f.Close(); err != nil {
			log.Fatalf("无法关闭文件: %v", err)
		}
	}()

	// 获取所有工作表的名字
	sheetList := f.GetSheetList()
	// 读取第一个工作表的数据
	sheetName := sheetList[0]
	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Fatalf("无法获取行数据: %v", err)
		return
	}

	// 每次读取的行数
	batchSize := 1000
	totalRows := len(rows)
	totalBatches := int(math.Ceil(float64(totalRows) / float64(batchSize)))

	// 按批次读取数据并将其存储到数组中
	for batch := 0; batch < totalBatches; batch++ {
		start := batch * batchSize
		end := start + batchSize
		// 确保最后一批数据的行数不超过总行数
		if end > totalRows {
			end = totalRows
		}
		// 将数据加入到桶中
		bulkRequest := global.ESClient.Bulk()
		for i := start; i < end; i++ {
			req := elastic.NewBulkIndexRequest().
				Index("suggestion_index").
				Id(rows[i][0]). // 使用关键词作为文档的唯一 ID
				Doc(map[string]interface{}{
					"keyword": rows[i][0],
				})
			bulkRequest = bulkRequest.Add(req)
		}
		// 执行批量请求
		_, err = bulkRequest.Do(context.Background())
		if err != nil {
			log.Fatalf("批量请求失败: %v", err)
		}
		log.Println("批量" + strconv.Itoa(batch) + "插入完成")

	}

}
