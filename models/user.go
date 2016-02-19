package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "log"
	"time"
)

//User : struct to keep user data
type User struct {
	//Identification Information
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `bson:"user_name" json:"user_name"`
	Email       string        `bson:"user_email" json:"user_email"`
	PhoneNumber string        `bson:"phone_number" json:"phone_number"`
	// GoogleToken   string        `bson:"google_token" `
	// FacebookToken string        `bson:"facebook_token"`
	PasswordHash string `bson:"password_hash" json:"password_hash"`
	//Analytics Information
	SignUpDate   time.Time `bson:"signup_date" json:"signup_date"`
	LastLoggedIn time.Time `bson:"last_login" json:"last_login"`
	LoginCount   uint      `bson:"login_count" json:"login_count"`
}

//UserLogin : Struct for user login
type UserLogin struct {
	Email    string `bson:"user_email" json:"user_email"`
	Password string `bson:"password_hash" json:"password_hash"`
}

//REVIEW: #Issue 1 on Bitbucket in this module
//TODO: Implement features for Google and Facebook login api's
/*
//NewUserGoogle : Function to create new user using Google Token.
func NewUserGoogle(name, email, googleToken string, phoneNumber int64) *User {
	return &User{
		ID:           bson.NewObjectId(),
		Name:         name,
		Email:        email,
		GoogleToken:  googleToken,
		PhoneNumber:  phoneNumber,
		SignUpDate:   time.Now(),
		LastLoggedIn: time.Now(),
		LoginCount:   1,
	}

}

//NewUserFacebook : Function to create new user using Facebook Token.
func NewUserFacebook(name, email, facebookToken string, phoneNumber int64) *User {
	return &User{
		ID:            bson.NewObjectId(),
		Name:          Name,
		Email:         Email,
		FacebookToken: facebookToken,
		PhoneNumber:   phoneNumber,
		SignUpDate:    time.Now(),
		LastLoggedIn:  time.Now(),
		LoginCount:    1,
	}

}
*/

//NewUser : Function to create new user for normal signup
func NewUser(u *User) *User {
	return &User{
		ID:           bson.NewObjectId(),
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		PhoneNumber:  u.PhoneNumber,
		SignUpDate:   time.Now(),
		LastLoggedIn: time.Now(),
		LoginCount:   1,
	}

}

//FindByEmail : To find user by email.
func (u *User) FindByEmail(email string, db *mgo.Database) (error, *User) {
	Result := new(User)
	err := u.coll(db).Find(bson.M{"user_email": email}).One(&Result)
	// log.Println("-->")
	return err, Result

}

//Coll : Returns Collection
func (*User) coll(db *mgo.Database) *mgo.Collection {
	collection := db.C("user")
	return collection
}

//Save : Function to save user in db
func (u *User) Save(db *mgo.Database) error {
	err := u.coll(db).Insert(u)
	return err
}
