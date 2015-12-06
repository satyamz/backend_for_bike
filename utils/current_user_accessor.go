package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/satyamz/Bike/models"
	// "gopkg.in/mgo.v2"
)

/*CurrentUserAccessor : Current user accessor data type for storing user key
 */
type CurrentUserAccessor struct {
	key string
}

//NewCurrentUserAccessor : Function to return CurrentUserAccessor instance
func NewCurrentUserAccessor(key string) *CurrentUserAccessor {
	return &CurrentUserAccessor{key}
}

//Get :function to return user instance associated with
func (cua *CurrentUserAccessor) Get(c *gin.Context) *models.User {
	if returnValue, _ := c.Get(cua.key); returnValue != "" {
		return returnValue.(*models.User)
	}
	return nil
}

//Set : function to set the user and key together
func (cua *CurrentUserAccessor) Set(c *gin.Context, user *models.User) {
	c.Set(cua.key, user)
}
