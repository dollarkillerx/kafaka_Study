package sarama

import "github.com/Shopify/sarama"

func NewProducer(cfg KafkaConfig) (sarama.SyncProducer, error) {
	kafkaConfig := sarama.NewConfig()
	if cfg.EnableSASL {
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.Mechanism = sarama.SASLTypePlaintext
		kafkaConfig.Net.SASL.User = cfg.User
		kafkaConfig.Net.SASL.Password = cfg.Password
	}

	kafkaConfig.Producer.Retry.Max = 5
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Producer.Partitioner = sarama.NewRandomPartitioner

	sarama.NewAsyncProducer([]string{cfg.BrokerURI}, kafkaConfig)
	return sarama.NewSyncProducer([]string{cfg.BrokerURI}, kafkaConfig)
}

