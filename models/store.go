package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//TODO: write models for store and map
//StoreMap : Structure to create custom store datatype
type StoreMap struct {
	StoreID   bson.ObjectId `bson:"_id, omitempty"`
	StoreName string        `bson:"store_name"`
}
