package kafka

import (
	"context"
	"encoding/json"
	"log"

	"discount-service/service"

	"github.com/IBM/sarama"
)

type OrderCreatedEvent struct {
	CustomerID   uint    `json:"customer_id"`
	OrderID      uint    `json:"order_id"`
	Amount       float64 `json:"amount"`
	CustomerTier string  `json:"customer_tier"`
}

type KafkaConsumer struct {
	service *service.DiscountService
	config  *sarama.Config
}

func NewKafkaConsumer(service *service.DiscountService, config *sarama.Config) *KafkaConsumer {
	return &KafkaConsumer{service: service, config: config}
}

func (c *KafkaConsumer) ConsumeMessages(brokers []string, topic string, groupID string) {
	consumer, err := sarama.NewConsumerGroup(brokers, groupID, c.config)
	if err != nil {
		log.Fatalf("Kafka consumer oluşturulamadı: %v", err)
	}

	for {
		err := consumer.Consume(context.Background(), []string{topic}, c)
		if err != nil {
			log.Printf("Kafka mesaj tüketimi hatası: %v", err)
		}
	}
}

func (c *KafkaConsumer) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *KafkaConsumer) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *KafkaConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var event OrderCreatedEvent
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			log.Printf("Kafka mesajı işlenemedi: %v", err)
			continue
		}

		discount, err := c.service.ApplyDiscount(event.CustomerID, event.OrderID, event.Amount, event.CustomerTier)
		if err != nil {
			log.Printf("İndirim uygulanamadı: %v", err)
		} else {
			log.Printf("İndirim başarıyla uygulandı: %+v", discount)
		}

		session.MarkMessage(msg, "")
	}
	return nil
}
