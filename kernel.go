package main

import (
	// "github.com/gin-gonic/gin"
	// "github.com/satyamz/Bike/models"
	"github.com/satyamz/Bike/utils"
	"github.com/satyamz/Bike/web"
)

func main() {
	dbURL := "127.0.0.1"
	Key := "user1"
	dbAccessor := utils.NewDatabaseAccessor(dbURL, "App", Key)
	cuAccessor := utils.NewCurrentUserAccessor(Key)
	router := web.NewServer(*dbAccessor, *cuAccessor)
	router.Run(":8080")
}
