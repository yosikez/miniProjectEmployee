package router

import (
	"miniProject/controller"
	"miniProject/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine){
	
	employeeController := controller.NewEmployeeController()
	opportunityController := controller.NewOpportunityController()

	router.GET("/employees", employeeController.FindAll)
	router.GET("/employees/:id", employeeController.FindById)
	router.POST("/employees", middleware.JsonValidEmp(), employeeController.Create)
	router.PUT("/employees/:id", middleware.JsonValidEmp(), employeeController.Update)
	router.DELETE("/employees/:id", employeeController.Delete)

	router.POST("/opportunities", middleware.JsonValidOpportunity(), opportunityController.Create)
}