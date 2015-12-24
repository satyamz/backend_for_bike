package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/satyamz/Bike/models"
	"github.com/satyamz/Bike/utils"
	// "gopkg.in/mgo.v2"
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
	// user := new(models.User)
	// db := ac.database.Get(c)
	db := ac.database.Givedb()
	userName := c.PostForm("NAME")
	userEmail := c.PostForm("EMAIL")
	userPhoneNumber := c.PostForm("MOB_NO")
	userPasswordHash := c.PostForm("PASSWORD")
	log.Printf("%s\t%s\t%s\t%s\n", userName, userEmail, userPhoneNumber, userPasswordHash)
	user := models.NewUser(userName, userEmail, userPasswordHash, userPhoneNumber)
	err := user.Save(db)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status":  "Email Exists",
			"message": err,
		})

	}

}

//login : Login utility
func (ac *AccountController) login(c *gin.Context) {
	db := ac.database.Givedb()
	userEmail := c.PostForm("user_email")
	userPassword := c.PostForm("user_password")
	user := models.User{Email: userEmail, PasswordHash: userPassword}
	err := user.FindByEmail(userEmail, userPassword, db)
	c.JSON(http.StatusConflict, gin.H{
		"status":  "ok",
		"message": err})
	/*
		    if err != nil {
				c.JSON(http.StatusConflict, gin.H{
					"status":  "Email Exists",
					"message": err,
				})

			} else {
				c.JSON(http.StatusNotFound, gin.H{
					"status":  "Email not exists",
					"message": "Please sign up",
				})
			}
	*/
}
