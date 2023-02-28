package router

import (
	"miniProject/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine){
	
	employeeController := controller.NewEmployeeController()

	router.GET("/employees", employeeController.FindAll)
	router.GET("/employees/:id", employeeController.FindById)
	router.POST("/employees", employeeController.Create)
	router.PUT("/employees/:id", employeeController.Update)
	router.DELETE("/employees/:id", employeeController.Delete)

}