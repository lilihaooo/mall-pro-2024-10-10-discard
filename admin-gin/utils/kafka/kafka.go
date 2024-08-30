package kafka

import (
	"admin-gin/global"
	"encoding/json"
	"github.com/IBM/sarama"
)

type Sender struct {
	Topic string  `json:"topic"`
	Key   string  `json:"key"`
	Value Message `json:"value"`
}

type Message struct {
	ID         uint   `json:"id,omitempty"`
	CoverImage string `json:"cover_image,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
}

func NewSender(topic string, key string, value Message) *Sender {
	return &Sender{
		Topic: topic,
		Key:   key,
		Value: value,
	}
}

func (m *Sender) Send() error {
	mByte, err := json.Marshal(m.Value)
	if err != nil {
		return err
	}
	// 创建消息
	msg := sarama.ProducerMessage{
		Topic: m.Topic,
		Key:   sarama.StringEncoder(m.Key),
		Value: sarama.StringEncoder(mByte),
	}
	// 发送消息
	_, _, err = global.KafkaProducer.SendMessage(&msg)
	return err
}
