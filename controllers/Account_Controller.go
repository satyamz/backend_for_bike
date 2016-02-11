package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satyamz/Bike/models"
	"github.com/satyamz/Bike/utils"
	"log"
	"net/http"
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
	router.POST("/login", ac.login)

}

//FindOrCreateUser : While signing up
func (ac *AccountController) signup(c *gin.Context) {
	user := new(models.User)
	// db := ac.database.Get(c)
	db := ac.database.Givedb()
	// userName := c.PostForm("NAME")
	// userEmail := c.PostForm("EMAIL")
	// userPhoneNumber := c.PostForm("MOB_NO")
	// userPasswordHash := c.PostForm("PASSWORD")
	// log.Printf("%s\t%s\t%s\t%s\n", userName, userEmail, userPhoneNumber, userPasswordHash)
	// user := models.NewUser(userName, userEmail, userPasswordHash, userPhoneNumber)
	c.Bind(&user)
	log.Println(user)
	err := user.Save(db)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status":  "Email Exists",
			"message": err,
		})

	}

}

//login : Login utility function
func (ac *AccountController) login(c *gin.Context) {
	db := ac.database.Givedb()
	user := models.User{}
	c.Bind(&user)
	err, Result := user.FindByEmail(user.Email, db)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"flag":    "0",
			"status":  "Email not found",
			"message": "Please check email address",
			"Error":   err,
		})
	} else {
		if user.PasswordHash == Result.PasswordHash {

			c.JSON(http.StatusOK, gin.H{
				"flag":   "1",
				"status": "ok",
			})
		} else {
			fmt.Println("Password error")
			c.JSON(http.StatusOK, gin.H{
				"flag":    "0",
				"status":  "Password is wrong",
				"message": "Please check message",
			})
		}
	}

}
