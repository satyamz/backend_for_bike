package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

//TODO: write models for store and map

//StoreMap : Structure to create custom store datatype
type StoreMap struct {
	StoreID              bson.ObjectId `bson:"_id,omitempty"`
	StoreName            string        `bson:"store_name"`
	StoreLocation        []float64     `bson:"store_loc"`
	NumberOfBikesPresent int           `bson:"bike_count"`
}

//UserLocation : Struct to store user location
type UserLocation struct {
	userLocation []float64
}

//NewStore : Returns StoreMap
func NewStore(StoreName string, StoreLocation []float64, NoOfBikesPresent int) *StoreMap {
	return &StoreMap{
		StoreName:            StoreName,
		StoreLocation:        StoreLocation,
		NumberOfBikesPresent: NoOfBikesPresent,
	}
}

//Save : Function to add store.
func (sm *StoreMap) Save(db *mgo.Database) error {
	return sm.coll(db).Insert(sm)
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
func (sm *StoreMap) FindNearByStore(db *mgo.Database) StoreMap {
	var result StoreMap
	pipeline := bson.M{
		"store_loc": bson.M{
			"$near": []interface{}{
				sm.StoreLocation[0], sm.StoreLocation[1],
			},
			"$maxDistance": 1,
		},
	}
	log.Println(pipeline)
	err := sm.coll(db).Find(pipeline).One(&result)
	if err != nil {
		panic(err)
	}
	return result
}
