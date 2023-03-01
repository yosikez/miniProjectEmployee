package router

import (
	"miniProject/controller"
	"miniProject/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine){
	
	employeeController := controller.NewEmployeeController()

	router.GET("/employees", employeeController.FindAll)
	router.GET("/employees/:id", employeeController.FindById)
	router.POST("/employees", middleware.JsonValid(), employeeController.Create)
	router.PUT("/employees/:id", middleware.JsonValid(), employeeController.Update)
	router.DELETE("/employees/:id", employeeController.Delete)

}