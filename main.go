package main

import (
	"fmt"
	"log"
	"miniProject/database"
	"miniProject/helper/validation"
	"miniProject/rabbitmq"
	"miniProject/router"

	"github.com/gin-gonic/gin"
)

func main() {
	err := database.Connect()
	if err != nil {
		panic(err)
	}

	rmqCfg, rmq, err := rabbitmq.NewRabbitMQ()
	
	if err != nil {
		log.Fatalf("failed to connect to rabbitmq : %v", err)
	}

	defer rmq.Connection.Close()
	defer rmq.Channel.Close()

	err = rmq.Channel.ExchangeDeclare(
		rmqCfg.ExchangeName,
		rmqCfg.ExchangeKind,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("failed to declare exchange : %v", err)
	}

	r := gin.Default()
	
	router.RegisterRouter(r, rmq, rmqCfg)
	
	validation.RegisterCustomValidator()

	err = r.Run(":8000")
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}

	fmt.Println("Server started on port 8000")
}