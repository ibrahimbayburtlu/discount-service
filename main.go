package main

import (
	"discount-service/api"
	"discount-service/config"
	"discount-service/kafka"
	"discount-service/models"
	"discount-service/repository"
	"discount-service/service"

	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config.LoadEnv()

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Veritabanına bağlanılamadı: %v", err)
	}

	db.AutoMigrate(&models.Discount{})

	repo := repository.NewDiscountRepository(db)
	discountService := service.NewDiscountService(repo)
	apiHandler := api.NewDiscountAPI(discountService)

	r := gin.Default()
	r.GET("/discounts/:customerID", apiHandler.GetCustomerDiscounts)

	go func() {
		if err := r.Run(":8083"); err != nil {
			log.Fatalf("API başlatılamadı: %v", err)
		}
	}()

	kafkaConfig, brokers := config.KafkaConfig()
	consumer := kafka.NewKafkaConsumer(discountService, kafkaConfig)
	go consumer.ConsumeMessages(brokers, "order.created", "discount-service-group")

	select {}
}
