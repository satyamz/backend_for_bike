package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/satyamz/Bike/models"
	"github.com/satyamz/Bike/utils"
	// "gopkg.in/mgo.v2"
	// "net/http"
)

/*AccountController : Structure to bind
database and user
*/
type AccountController struct {
	database    utils.DatabaseAccessor
	currentUser utils.CurrentUserAccessor
}

//NewAccountController : Function to create
func NewAccountController(dba utils.DatabaseAccessor, cua utils.CurrentUserAccessor) *AccountController {
	return &AccountController{
		database:    dba,
		currentUser: cua,
	}

}

//Register : Function to register router for AccountController
func (ac *AccountController) Register(router *gin.Engine) {
	router.POST("/signup", ac.signup)
}

//FindOrCreateUser : While signing up
func (ac *AccountController) signup(c *gin.Context) {
	// user := new(models.User)
	db := ac.database.Get(c)
	userName := c.PostForm("name")
	userEmail := c.PostForm("email")
	userPhoneNumber := c.PostForm("phone")
	userPasswordHash := c.PostForm("password")
	user := models.NewUser(userName, userEmail, userPasswordHash, userPhoneNumber)
	user.Save(db)
	/*	if err != nil {
			c.JSON(http.StatusConflict, gin.H{
				"status":  "Email Exists",
				"message": err,
			})

		}
	*/
}
