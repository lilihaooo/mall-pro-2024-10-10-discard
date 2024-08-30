package initialize

import (
	"admin-gin/global"
	"github.com/IBM/sarama"
)

func InitKafkaProducer() {
	// 初始化Kafka生产者
	brokers := global.GVA_CONFIG.Kafka.Url()
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal // 要求leader 收到消息即可
	config.Producer.Retry.Max = 5

	//config.Producer.Return.Successes = true // 设置为 true 时，生产者在调用 SendMessage 方法时会阻塞，直到收到 Kafka broker 的确认（即消息被成功写入 Kafka）。
	//config.Producer.RequiredAcks = sarama.WaitForAll
	/*
		sarama.WaitForAll 表示生产者会等待所有的副本都确认收到消息后才返回。还有其他可选值：
		NoResponse (0): 生产者不会等待任何确认，消息可能会丢失。
		WaitForLocal (1): 生产者只等待 leader 副本确认，性能较高但可靠性较低。  默认false
	*/
	producer, err := sarama.NewSyncProducer([]string{brokers}, nil)
	if err != nil {
		panic(err)
	}
	global.KafkaProducer = producer
}
