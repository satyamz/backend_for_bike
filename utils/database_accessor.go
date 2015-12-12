package utils

import (
	// "github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	// "net/http"
)

//DatabaseAccessor : struct for binding mgo session, db url, db name and key.
type DatabaseAccessor struct {
	*mgo.Session
	url  string
	name string
	key  string
}

/*NewDatabaseAccessor : Method to return DatabaseAccessor
  instance and to initialize databse
*/
func NewDatabaseAccessor(url, name, key string) *DatabaseAccessor {
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}

	//TODO: Implement Method to add unique index on user_email and user_phone.
	//REVIEW: Using manually db.user.createIndex({"user_email" : 1 , "user_phone" : 1},{ unique : true})

	/*
		index := mgo.Index{
			Key:    []string{"user_email", "user_phone"},
			Unique: true,
		}
		 session.DB(name).C("user").EnsureIndex(index)
	*/
	return &DatabaseAccessor{session, url, name, key}
}

//Givedb : Function to return db
func (d *DatabaseAccessor) Givedb() *mgo.Database {
	return d.Session.DB("App")
}

/*
//Get : To keep track of db and request
func (d *DatabaseAccessor) Get(c *gin.Context) *mgo.Database {
	if returnValue, _ := c.Get(d.key); returnValue != nil {
		return returnValue.(*mgo.Database)
	}
	return nil
}
*/
