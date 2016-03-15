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
	router.POST("/start_ride", bc.StartRide)
	router.POST("/stop_ride", bc.StopRide)
	router.POST("/finish_ride", bc.RideEnd)

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
			"status":  "No nearby store available",
			"message": "Please try from other location",
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
				fmt.Println(err)
			} else {
				c.JSON(200, gin.H{
					//TODO: Send details about SM who is delivering bike.
					"status":  "ok",
					"message": "Ride confirmed",
				})
			}
		}
	}
}

//StartRide : Function to start ride
// TODO: Update database, Think more on maintaining RideID
func (bc *BookingController) StartRide(c *gin.Context) {
	db := bc.database.Givedb()
	StartRideInstance := new(models.Ride)
	c.Bind(&StartRideInstance)
	RideInstanceOnStart := models.NewStartRide(StartRideInstance)
	err := RideInstanceOnStart.StartRide(db)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"Error": "Error"})
	} else {
		//Implement RabbitMQ API for sending notification to user.
		c.JSON(200, gin.H{"status": "Ride Started"}) // Response to SM.
	}
}

//StopRide : Function to stop ride from user -> Server.
func (bc *BookingController) StopRide(c *gin.Context) {
	db := bc.database.Givedb()
	StopRideInstance := new(models.Ride)
	c.Bind(&StopRideInstance)
	EndStoreInstance := models.NewStore(" ", StopRideInstance.EndUserLocation, 1)
	EndStoreResult, err := EndStoreInstance.FindNearByStore(db)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  "No nearby store available",
			"message": "Please try from other location",
		})
	} else {
		//TODO: Send request to EndStoreResult store manager may be GCM.
		StopRideInstance.StoreManagerEndID = EndStoreResult.StoreID.Hex()
		StopRideInstanceUpdateTime := models.NewStopRide(StopRideInstance)
		fmt.Println(*StopRideInstanceUpdateTime)
		err := StopRideInstanceUpdateTime.StopRide(db)
		if err != nil {
			c.JSON(200, gin.H{
				"status": "Error",
			})
		} else {
			c.JSON(200, gin.H{
				"status":  "Ride stop request confirmed",
				"message": "Please wait till store manager comes",
			})
		}
	}
}

//ConfirmEndRide : Function to confirm end of ride form SM.
//Send notification to the user "Bike will be picked up by so and so SM.".
func (bc *BookingController) ConfirmEndRide(c *gin.Context) {
	ConfirmEndRideInstance := new(models.Ride)
	c.Bind(&ConfirmEndRideInstance)

}

//RideEnd : Function to finish ride.
//Main function of this utility is to take ride end data and build
func (bc *BookingController) RideEnd(c *gin.Context) {
	db := bc.database.Givedb()
	RideEndInstance := new(models.Ride)
	c.Bind(&RideEndInstance)
	RideFinish := models.NewEndRide(RideEndInstance)
	err := RideFinish.RideEnd(db) //Update fields
	if err != nil {
		c.JSON(200, gin.H{
			"Status": "Error occured",
		})
	} else {
		c.JSON(200, gin.H{
			"Status": "Ride end",
		})
	}
}
