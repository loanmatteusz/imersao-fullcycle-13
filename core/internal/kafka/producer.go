package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

type Producer struct {
	ConfigMap *ckafka.ConfigMap
}

func NewKafkaProducer(configMap *ckafka.ConfigMap) *Producer {
	return &Producer{
		ConfigMap: configMap,
	}
}

func (producer *Producer) Publish(message interface{}, key []byte, topic string) error {
	prod, err := ckafka.NewProducer(producer.ConfigMap)
	if err != nil {
		return err
	}
	msg := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{
			Topic:     &topic,
			Partition: ckafka.PartitionAny,
		},
		Key:   key,
		Value: message.([]byte),
	}

	err = prod.Produce(msg, nil)
	if err != nil {
		return err
	}

	return nil
}
