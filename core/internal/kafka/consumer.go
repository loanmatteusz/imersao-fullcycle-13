package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

type Consumer struct {
	ConfigMap *ckafka.ConfigMap
	Topics    []string
}

func NewConsumer(configMap *ckafka.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
	}
}

func (consumer *Consumer) Consume(msgChannel chan *ckafka.Message) error {
	cons, err := ckafka.NewConsumer(consumer.ConfigMap)
	if err != nil {
		panic(err)
	}
	err = cons.SubscribeTopics(consumer.Topics, nil)
	if err != nil {
		return err
	}
	for {
		msg, err := cons.ReadMessage(-1)
		if err != nil {
			msgChannel <- msg
		}
	}
}
