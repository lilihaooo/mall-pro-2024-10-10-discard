package esmodel

import (
	"context"
	"task/global"
)

type Suggestion struct {
	Keyword []string
}

func (s *Suggestion) Index() string {
	return "suggestion_index"
}

func (s *Suggestion) Mapping() map[string]interface{} {
	type mi = map[string]interface{}
	mapping := mi{
		"settings": mi{
			"analysis": mi{
				"analyzer": mi{
					"completion_analyzer": mi{
						"tokenizer": "keyword",
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
				"keyword": mi{
					"type":            "completion",
					"analyzer":        "completion_analyzer",
					"search_analyzer": "keyword",
				},
			},
		},
	}
	return mapping
}

// IndexExists 索引是否存在
func (s *Suggestion) IndexExists() bool {
	exists, err := global.ESClient.
		IndexExists(s.Index()).
		Do(context.Background())
	if err != nil {
		global.Logrus.Error(err)
		return exists
	}
	return exists
}

// CreateIndex 创建索引
func (s *Suggestion) CreateIndex() error {
	if s.IndexExists() {
		// 有索引
		err := s.RemoveIndex()
		if err != nil {
			return err
		}
	}
	// 创建索引
	createIndex, err := global.ESClient.
		CreateIndex(s.Index()).
		BodyJson(s.Mapping()).
		Do(context.Background())
	if err != nil {
		global.Logrus.Error(err)
		return err
	}
	if !createIndex.Acknowledged {
		global.Logrus.Error("创建失败")
		return err
	}
	global.Logrus.Infof("索引 %s 创建成功", s.Index())
	return nil
}

// RemoveIndex 删除索引
func (s *Suggestion) RemoveIndex() error {
	global.Logrus.Info("索引存在，删除索引")
	// 删除索引
	indexDelete, err := global.ESClient.DeleteIndex(s.Index()).Do(context.Background())
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
