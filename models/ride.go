package models

import (
	"gopkg.in/mgo.v2"
	// "encoding/json"
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Ride : Structure to bind ride Information together
type Ride struct {
	//Ride metadata fields
	RideID         bson.ObjectId `bson:"_id,omitempty"`
	UserID         string        `bson:"user_id" json:"user_id"`
	StoreManagerID string        `bson:"sm_id" json:"sm_id"`

	//Ride Start fields
	ConfirmRideTime       time.Time `bson:"ride_confirm_time" json:"ride_confirm_time"`
	StartUserLocation     []float64 `bson:"start_user_loc" json:"start_loc"`
	RideStartTime         time.Time `bson:"ride_start_time" json:"start_time"`
	RideStartTimeOnServer time.Time `bson:"ride_start_time_server" json:"start_time_server"`
	RideStartReadingImage string    `bson:"start_meter_image" json:"start_meter_image"`
	UserLicenseImage      string    `bson:"user_license_image" json:"user_license_image"`
	RideStartReading      float64   `bson:"start_reading" json:"start_reading"`

	//Ride End fields
	EndUserLocation       []float64 `bson:"end_user_loc" json:"end_loc"`
	RideEndTime           time.Time `bson:"ride_end_time" json:"end_time"`
	UserRideEndTimeServer time.Time `bson:"ride_end_time_server" json:"end_time_server"`
	RideEndReading        float64   `bson:"end_reading" json:"end_reading"`
	RideEndReadingImage   string    `bson:"end_meter_image" json:"end_meter_image"`

	//Ride Calculations
	TotalDistanceTravelled float64   `bson:"total_distance" json:"total_distance"`
	TotalTimeTravelled     time.Time `json:"total_time" json:"total_time"`
	TotalBill              float64   `json:"total_bill" json:"total_bill"`
}

//NewRide : Function to return new ride instance
func NewRide(r *Ride) *Ride {
	return &Ride{
		RideID:          bson.NewObjectId(),
		UserID:          r.UserID,
		StoreManagerID:  r.StoreManagerID,
		ConfirmRideTime: time.Now(),
	}
}

//NewStartRide : Function to instantiate StartRide instance.
func NewStartRide(r *Ride) *Ride {
	return &Ride{
		StoreManagerID:        r.StoreManagerID,
		UserID:                r.UserID,
		StartUserLocation:     r.StartUserLocation,
		RideStartTime:         r.RideStartTime,
		RideStartReading:      r.RideStartReading,
		RideStartReadingImage: r.RideStartReadingImage,
		UserLicenseImage:      r.UserLicenseImage,
		RideStartTimeOnServer: time.Now(),
	}
}

//NewStopRide : Function to update time on stop ride.
//Update time when request hits server.
func NewStopRide(r *Ride) *Ride {
	return &Ride{
		UserID:                r.UserID,
		EndUserLocation:       r.EndUserLocation,
		RideEndTime:           r.RideEndTime,
		UserRideEndTimeServer: time.Now(),
	}
}

//NewConfirmEndRide : Function to update
func NewConfirmEndRide(r *Ride) *Ride {
	return &Ride{}
}

//ConfirmRide : Function to confirm ride.
func (ride *Ride) ConfirmRide(db *mgo.Database) error {
	fmt.Println(ride)
	err := ride.coll(db).Insert(ride)
	return err
}

//StartRide : Function to update ride instance
func (ride *Ride) StartRide(db *mgo.Database) error {
	UserIDQuery := bson.M{"user_id": ride.UserID}
	UpdateQuery := bson.M{"$set": bson.M{"sm_id": ride.StoreManagerID, "start_user_loc": ride.StartUserLocation, "ride_start_time": ride.RideStartTime, "start_reading": ride.RideStartReading, "start_meter_image": ride.RideStartReadingImage, "user_license_image": ride.UserLicenseImage, "ride_start_time_server": ride.RideStartTimeOnServer}}
	err := ride.coll(db).Update(UserIDQuery, UpdateQuery)
	return err
}

//coll : Function to create collection.
func (ride *Ride) coll(db *mgo.Database) *mgo.Collection {
	return db.C("ride")
}
