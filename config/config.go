package config

import (
	"log"
	"os"

	"github.com/IBM/sarama"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ .env dosyası yüklenemedi, varsayılan ayarlar kullanılacak.")
	}
}

func ConnectDB() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {

		dsn = "host=localhost user=postgres password=postgres dbname=discountdb port=5432 sslmode=disable"
	}
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func KafkaConfig() (*sarama.Config, []string) {
	brokers := os.Getenv("KAFKA_BROKERS")
	if brokers == "" {
		brokers = "localhost:9092"
	}
	brokerList := []string{brokers}

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin

	return config, brokerList
}
