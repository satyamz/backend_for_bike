package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/satyamz/Bike/models"
	"github.com/satyamz/Bike/utils"
)

//StoreMapController : controller structure
type StoreMapController struct {
	database    utils.DatabaseAccessor
	currentUser utils.CurrentUserAccessor
}

//NewStoreMapController : Function to return new map controller instance
func NewStoreMapController(dba utils.DatabaseAccessor, cua utils.CurrentUserAccessor) *StoreMapController {
	return &StoreMapController{
		database:    dba,
		currentUser: cua,
	}
}

//Register : Function to register router for Map_controller
func (sm *StoreMapController) Register(router *gin.Engine) {
	router.POST("/find_store", sm.FindNearStore)

}

//FindNearStore : Find store near to user location and send the sttud
func (sm *StoreMapController) FindNearStore() {

}
