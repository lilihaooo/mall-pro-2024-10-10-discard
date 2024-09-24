package kafka

import (
	"admin-gin/global"
	"encoding/json"
	"github.com/IBM/sarama"
)

type Sender struct {
	Topic string      `json:"topic"`
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func (s Sender) Send() error {
	mByte, err := json.Marshal(s.Value)
	if err != nil {
		return err
	}
	// 创建消息
	msg := sarama.ProducerMessage{
		Topic: s.Topic,
		Key:   sarama.StringEncoder(s.Key),
		Value: sarama.StringEncoder(mByte),
	}
	// 发送消息
	_, _, err = global.KafkaProducer.SendMessage(&msg)
	return err
}
