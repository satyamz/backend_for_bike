package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satyamz/Bike/models"
	"github.com/satyamz/Bike/utils"
)

//BookingController : Structure for booking controller
type BookingController struct {
	database    utils.DatabaseAccessor
	currentUser utils.CurrentUserAccessor
}

//NewBookingController : Returns BookingController instance
func NewBookingController(dba utils.DatabaseAccessor, cua utils.CurrentUserAccessor) *BookingController {
	return &BookingController{
		database:    dba,
		currentUser: cua,
	}
}

//Register : Function to register router.
func (bc *BookingController) Register(router *gin.Engine) {
	router.POST("/confirm_ride", bc.RideConfirm)
}

//RideConfirm : Function to confirm the ride(S.M will confirm ride)
func (bc *BookingController) RideConfirm(c *gin.Context) {
	db := bc.database.Givedb()
	ConfirmRideInstance := new(models.Ride)
	c.Bind(&ConfirmRideInstance)
	RideInstance := models.NewRide(ConfirmRideInstance)
	fmt.Println(*RideInstance)
	// c.Bind(&RideInstance)
	err := RideInstance.ConfirmRide(db)
	if err != nil {
		c.JSON(500, gin.H{
			"status": "Failure",
		})
		// panic(err)
		fmt.Println(err)
	} else {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "confirmed",
		})
	}
	// c.Bind(&ConfirmRideInstance)
	fmt.Println(RideInstance)
	// c.JSON(200, ConfirmRideInstance)

}
