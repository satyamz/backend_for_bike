package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/satyamz/Bike/models"
	"github.com/satyamz/Bike/utils"
	"log"
	"strconv"
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
	router.GET("/find_store", sm.FindBikesInStore)
	router.POST("/add_store", sm.AddStore)

}

//FindBikesInStore : Find store near to user location and send the sttud
func (sm *StoreMapController) FindBikesInStore(c *gin.Context) {
	userLocationLat := c.Query("user_lat")
	userLocationLong := c.Query("user_long")
	storeLocationLat := c.Query("store_lat")
	storeLocationLong := c.Query("store_long")
	ulat, _ := strconv.ParseFloat(userLocationLat, 32)
	ulong, _ := strconv.ParseFloat(userLocationLong, 32)
	slat, _ := strconv.ParseFloat(storeLocationLat, 32)
	slong, _ := strconv.ParseFloat(storeLocationLong, 32)

	userLocation := []float64{ulat, ulong}
	storeLocation := []float64{slat, slong}
	log.Println(userLocation)
	log.Println(storeLocation)
}

//AddStore : add new store
func (sm *StoreMapController) AddStore(c *gin.Context) {
	db := sm.database.Givedb()
	storeLocation := make([]float64, 2)
	storeName := c.PostForm("store_name")

	storeLocation[0], _ = strconv.ParseFloat(c.PostForm("store_lat"), 32)
	storeLocation[1], _ = strconv.ParseFloat(c.PostForm("store_long"), 32)
	NumberOfBikesPresent, _ := strconv.Atoi(c.PostForm("bikes_present"))
	Store := models.NewStore(storeName, storeLocation, NumberOfBikesPresent)
	err := Store.Save(db)
	if err != nil {
		panic(err)
	}

}
