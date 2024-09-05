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
	ID uint `json:"id,omitempty"`
	// 封面
	CoverImage string `json:"cover_image,omitempty"`
	// 基本信息
	Description     string  `json:"description,omitempty"`
	Price           float64 `json:"price,omitempty"`
	CommissionRate  int32   `json:"commission_rate,omitempty"`
	CommissionValue float64 `json:"commission_value,omitempty"`
	Tags            []uint  `json:"tags,omitempty"`
	BrandName       string  `json:"brand_name,omitempty"`
	// 优惠券信息
	CouponID        uint    `json:"coupon_id,omitempty"`
	CouponAmount    float64 `json:"coupon_amount,omitempty"`
	CouponBeginTime string  `json:"coupon_begin_time,omitempty"`
	CouponEndTime   string  `json:"coupon_end_time,omitempty"`
	CouponTotal     int     `json:"coupon_total,omitempty"`
	CouponCover     int     `json:"coupon_cover,omitempty"`
	PostCouponPrice float64 `json:"post_coupon_price,omitempty"`
	// 添加时间
	CreatedAt string `json:"created_at,omitempty"`
}

func NewSender(topic string, key string, value Message) *Sender {
	return &Sender{
		Topic: topic,
		Key:   key,
		Value: value,
	}
}

func (s *Sender) Send() error {
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
