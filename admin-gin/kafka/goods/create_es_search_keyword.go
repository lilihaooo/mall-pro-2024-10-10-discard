package goods

import "admin-gin/kafka"

func NewCreateEsSearchKeywordSender(value string) *kafka.Sender {
	return &kafka.Sender{
		Topic: "search-keyword",
		Key:   "create_es_search_keyword",
		Value: value,
	}
}
