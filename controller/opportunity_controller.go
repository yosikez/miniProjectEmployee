package controller

import (
	"encoding/json"
	"miniProject/database"
	"miniProject/helper/validation"
	"miniProject/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OpportunityController struct{}

type JsonDataOpportunity struct {
	Code            string     `json:"code" binding:"required"`
	ClientCode      string     `json:"client_code" binding:"required"`
	PicEmail        string     `json:"pic_email" binding:"required,email"`
	OpportunityName string     `json:"opportunity_name" binding:"required"`
	Description     string     `json:"description" binding:"required"`
	SalesEmail      string     `json:"sales_email" binding:"required,email"`
	Status          string     `json:"status" binding:"required"`
	LastModified    string     `json:"last_modified" binding:"required,datetime=2006-01-02 15:04:05"`
	Resources       []Resource `json:"resources" binding:"required,dive"`
}

type Resource struct {
	Qty             int64   `json:"qty" binding:"required"`
	Position        string  `json:"position" binding:"required"`
	Level           string  `json:"level" binding:"required"`
	Ctc             float64 `json:"ctc"`
	ProjectDuration int64   `json:"project_duration" binding:"required"`
}

func NewOpportunityController() *OpportunityController {
	return &OpportunityController{}
}

func (op *OpportunityController) Create(c *gin.Context) {
	var opportunity JsonDataOpportunity

	if err := c.ShouldBind(&opportunity); err != nil {
		errFields := validation.GetErrMess(err)

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "validation error",
			"error":   errFields,
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
