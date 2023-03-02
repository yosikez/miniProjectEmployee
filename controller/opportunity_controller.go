package controller

import (
	"encoding/json"
	"miniProject/database"
	"miniProject/helper/validation"
	"miniProject/input"
	"miniProject/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OpportunityController struct{}

func NewOpportunityController() *OpportunityController {
	return &OpportunityController{}
}

func (op *OpportunityController) Create(c *gin.Context) {
	var opportunity input.JsonDataOpportunity

	if err := c.ShouldBind(&opportunity); err != nil {
		errFields := validation.GetErrMess(err)

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

	c.JSON(http.StatusCreated, gin.H{
		"data": opportunity,
	})
}
