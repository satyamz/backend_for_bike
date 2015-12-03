package utils

import (
	"gopkg.in/mgo.v2"
)

//DatabaseAccessor : struct for binding mgo session, db url, db name and key.
type DatabaseAccessor struct {
	*mgo.Session
	url  string
	name string
	key  int
}

/*NewDatabaseAccessor : Method to return DatabaseAccessor
  instance and to initialize databse
*/
func NewDatabaseAccessor(url, name string, key int) *DatabaseAccessor {
	session, _ := mgo.Dial(url)
	session.DB(name).C("").EnsureIndex()
	return &DatabaseAccessor{session, url, name, key}
}
