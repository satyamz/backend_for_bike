package web

import (
	"github.com/gin-gonic/gin"
	"github.com/satyamz/Bike/controllers"
	"github.com/satyamz/Bike/utils"
)

//Server : Struct for binding

//NewServer : function to instantiate Server
func NewServer(dba utils.DatabaseAccessor, cua utils.CurrentUserAccessor) *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Ok",
		})
	})

	accountController := controllers.NewAccountController(dba, cua)
	accountController.Register(router)

	return router
}
