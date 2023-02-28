package main

import (
	"fmt"
	"log"
	"miniProject/database"
	"miniProject/helper/validation"
	"miniProject/router"

	"github.com/gin-gonic/gin"
)

func main() {
	err := database.Connect()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	router.RegisterRouter(r)
	
	validation.RegisterCustomValidator()

	err = r.Run(":8000")
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}

	fmt.Println("Server started on port 8080")
}