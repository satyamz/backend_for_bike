package models

import (
	"gopkg.in/mgo.v2"
	// "encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

//Ride : Structure to bind ride Information together
type Ride struct {
	RideID                 bson.ObjectId `bson:"_id,omitempty"`
	UserID                 string        `bson:"user_id" json:"user_id"`
	StoreManagerID         string        `bson:"sm_id" json:"sm_id"`
	ConfirmRideTime        time.Time     `bson:"ride_confirm_time" json:"ride_confirm_time"`
	StartUserLocation      []float64     `bson:"start_user_loc" json:"start_loc"`
	RideStartTime          time.Time     `bson:"ride_start_time" json:"start_time"`
	RideStartReading       float64       `bson:"start_reading" json:"start_reading"`
	EndUserLocation        []float64     `bson:"end_user_loc" json:"end_loc"`
	RideEndTime            time.Time     `bson:"ride_end_time" json:"end_time"`
	RideEndReading         float64       `bson:"end_reading" json:"end_reading"`
	TotalDistanceTravelled float64       `bson:"total_distance" json:"total_distance"`
	TotalTimeTravelled     time.Time     `json:"total_time" json:"total_time"`
	TotalBill              float64       `json:"total_bill" json:"total_bill"`
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
	return &Ride{}
}

//ConfirmRide : Function to confirm ride.
func (ride *Ride) ConfirmRide(db *mgo.Database) error {
	fmt.Println(ride)
	err := ride.coll(db).Insert(ride)
	return err
}

//coll : Function to create collection.
func (ride *Ride) coll(db *mgo.Database) *mgo.Collection {
	return db.C("ride")
}
