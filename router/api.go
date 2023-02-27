package router

import (
	"miniProject/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine){
	
	employeeController := controller.NewEmployeeController()

	router.GET("/employees", employeeController.FindAll)
	router.GET("/employee/:id", employeeController.FindById)
	router.POST("/employee", employeeController.Create)
	router.PUT("/employee/:id", employeeController.Update)
	router.DELETE("/employee/:id", employeeController.Delete)

}