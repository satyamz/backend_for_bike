package utils

import (
	"gopkg.in/mgo.v2"
)

/*CurrentUserAccessor : Current user accessor data type for storing user key
 */
type CurrentUserAccessor struct {
	key int
}

//NewCurrentUserAccessor : Function to return CurrentUserAccessor instance
func NewCurrentUserAccessor(key int) *CurrentUserAccessor {
	return &CurrentUserAccessor{key}
}
