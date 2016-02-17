package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

//TODO: write models for store and map

//StoreMap : Structure to create custom store datatype
type StoreMap struct {
	StoreID   bson.ObjectId `bson:"_id,omitempty"`
	StoreName string        `bson:"store_name"`
	StoreLong float64       `bson:"store_long"`
	StoreLat  float64       `bson:"store_lat"`

	NumberOfBikesPresent int `bson:"bike_count"`
}

//StoreAddNew : Struct to add new store
type StoreAddNew struct {
	StoreID              bson.ObjectId `bson:"_id,omitempty"`
	StoreName            string        `bson:"store_name"`
	StoreLocation        []float64     `bson:"store_loc"`
	NumberOfBikesPresent int           `bson:"bike_count"`
}

//UserLocation : Struct to store user location
type UserLocation struct {
	UserLong float64 `json:"user_long"`
	UserLat  float64 `json:"user_lat"`
}

//NewStore : Returns StoreMap
func NewStore(StoreName string, userLocation *UserLocation, NoOfBikesPresent int) *StoreMap {
	return &StoreMap{
		StoreName:            StoreName,
		StoreLong:            userLocation.UserLong,
		StoreLat:             userLocation.UserLat,
		NumberOfBikesPresent: NoOfBikesPresent,
	}
}

//Save : Function to add store.
func (sm *StoreMap) Save(db *mgo.Database) error {
	return sm.coll(db).Insert(sm)
}

//SaveStore : Function to store new store
func (sm *StoreAddNew) SaveStore(db *mgo.Database) error {
	return sm.coll(db).Insert(sm)
}

//coll : Returns collection
func (sm *StoreAddNew) coll(db *mgo.Database) *mgo.Collection {
	collection := db.C("store")
	return collection
}

//CheckBikeAvailiblityAtStore :Return the availiblity of bike.
func (sm *StoreMap) CheckBikeAvailiblityAtStore(db *mgo.Database) {

}

//coll : Returns collection.
func (sm *StoreMap) coll(db *mgo.Database) *mgo.Collection {
	collection := db.C("store")
	return collection
}

//FindNearByStore : Function to write query of find_nearby_store
func (sm *StoreMap) FindNearByStore(db *mgo.Database) (StoreMap, error) {
	var result StoreMap
	pipeline := bson.M{
		"store_loc": bson.M{
			"$near": []interface{}{
				sm.StoreLong, sm.StoreLat,
			},
			"$maxDistance": 1,
		},
	}
	// log.Println(pipeline)
	err := sm.coll(db).Find(pipeline).One(&result)
	if err != nil {
		log.Println(err)
	}
	return result, err
}
