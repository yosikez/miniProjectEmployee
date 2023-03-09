package controller

import (
	"context"
	"encoding/json"
	"miniProject/config"
	"miniProject/database"
	"miniProject/input"
	"miniProject/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	cusMessage "github.com/yosikez/custom-error-message"
)

type OpportunityController struct {
	rmq    *config.RabbitMQConnection
	rmqCfg *config.RabbitMQ
}

func NewOpportunityController(rqConnection *config.RabbitMQConnection, rqConfig *config.RabbitMQ) *OpportunityController {
	return &OpportunityController{
		rmq:    rqConnection,
		rmqCfg: rqConfig,
	}
}

func (op *OpportunityController) Create(c *gin.Context) {
	var opportunity input.JsonDataOpportunity

	if err := c.ShouldBind(&opportunity); err != nil {
		errFields := cusMessage.GetErrMess(err, opportunity, input.Resource{})

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "validation error",
			"errors":  errFields,
		})
		return
	}

	jsonData, err := json.Marshal(&opportunity)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed marshal to json",
			"error":   err.Error(),
		})
		return
	}

	modelData := model.Opportunity{Data: jsonData}

	if err := database.DB.Create(&modelData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create opportunity",
			"error":   err.Error(),
		})
		return
	}

	q, err := op.rmq.Channel.QueueDeclare("opportunity_create_queue", false, false, false, false, nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to declare queue rabbitmq",
			"error":   err.Error(),
		})
		return
	}

	err = op.rmq.Channel.QueueBind("opportunity_create_queue", "opportunity_create_queue", op.rmqCfg.ExchangeName, false, nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to bind a queue",
			"error":   err.Error(),
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to marshal json for message rabbitmq",
			"error":   err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = op.rmq.Channel.PublishWithContext(ctx, op.rmqCfg.ExchangeName, q.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonData,
	},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to publish message to rabbitmq",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": modelData,
	})
}
