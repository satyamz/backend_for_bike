package web

import (
	"github.com/gin-gonic/gin"
	"github.com/satyamz/Bike/utils"
)

//Server : Struct for binding
type Server struct {
	*gin.Engine
}

//NewServer : function to instantiate Server
func NewServer(dba utils.DatabaseAccessor, ca utils.CurrentUserAccessor) *Server {
	router := Server{gin.Default()}
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Ok",
		})
	})
	return &router
}
