package product

import (
	"admin-gin/global"
	"admin-gin/model/product/request"
	"context"
	"github.com/olivere/elastic/v7"
)

type SearchService struct{}

// 获取搜索词条
func (*SearchService) SearchSuggestionKeywordList(req request.Suggestion) (collect []string) {
	collect = make([]string, 0)
	keyword := req.Keyword
	// 构建 Suggest 查询
	suggestName := "keyword_suggest"
	suggestQuery := elastic.NewCompletionSuggester(suggestName).
		Prefix(keyword).
		Field("keyword").
		Size(10)

	// 执行查询
	searchResult, err := global.ESClient.Search().
		Index("suggestion_index"). // 设置索引名称
		Suggester(suggestQuery).   // 添加 suggest 查询
		Do(context.Background())   // 执行查询
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil
	}

	// 处理查询结果
	if suggestResult, found := searchResult.Suggest[suggestName]; found {
		for _, entry := range suggestResult {
			for _, option := range entry.Options {
				collect = append(collect, option.Text)
			}
		}
	} else {
		global.GVA_LOG.Info("无推荐")
		return
	}
	return
}
