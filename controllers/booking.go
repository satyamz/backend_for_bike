//Package controllers :
//0. Ride Confirm
//1. Write Controller /ConfirmBooking for S.M to confirm booking from his side. (Send status : ok)
//2. Start Ride Handler (On S.M Side): It'll add start time, start reading, License photo, Meter photo.
//3. Stop Ride Handler(On User side) : Update stop time. Send request to nearby store. With userID, UserLocation
//4. End Ride Confirm (On S.M side) : To confirm pickup. Send request to user and tell that S.M is arriving.
//5. RideEnd (On S.M Side) : Send UserID, Meter reading. -> Calculate time -> Calculate the bill Send Bill to the user and S.M
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

/*RideConfirm : When user Asks for delivery RideConfirm fuction will take inputs
1. UserID
2. User location

find nearby store and send request to respective store.
Right now we're extracting only one nearby. This function fails if there's no
bike available in nearby store. This can be adjusted later.

Variable ResultStore (type Store) will give nearby store info.
Update the db and add storeID, UserID, UserLocation, Request Time.
*/
func (bc *BookingController) RideConfirm(c *gin.Context) {
	db := bc.database.Givedb()
	ConfirmRideInstance := new(models.Ride)
	c.Bind(&ConfirmRideInstance)

	storeInstance := models.NewStore(" ", ConfirmRideInstance.StartUserLocation, 1)
	Result, err := storeInstance.FindNearByStore(db)

	if err != nil {
		c.JSON(200, gin.H{
			"status": "No nearby store available",
		})
	} else {
		if Result.NumberOfBikesPresent == 0 {
			c.JSON(200, gin.H{
				"status":  "Not found",
				"message": "No bikes are available in your nearby store",
			})
		} else {
			//Send notification to the store

			utils.StartRabbitMq(Result, ConfirmRideInstance)
			ConfirmRideInstance.StoreManagerID = Result.StoreID.Hex()
			RideInstance := models.NewRide(ConfirmRideInstance)
			fmt.Println(*RideInstance)
			err := RideInstance.ConfirmRide(db)
			if err != nil {
				c.JSON(500, gin.H{
					"status": "Error while booking ride",
				})
				// panic(err)
				fmt.Println(err)

			} else {
				c.JSON(200, gin.H{
					"status":  "ok",
					"message": "Ride confirmed",
				})
			}
		}
	}

	// c.Bind(&RideInstance)
	// c.Bind(&ConfirmRideInstance)
	// fmt.Println(RideInstance)
	// c.JSON(200, ConfirmRideInstance)

}
