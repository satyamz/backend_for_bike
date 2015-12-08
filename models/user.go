package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

//User : struct to keep user data
type User struct {
	//Identification Information
	ID            bson.ObjectId `bson:"_id,omitempty"`
	Name          string        `bson:"user_name"`
	Email         string        `bson:"user_email"`
	PhoneNumber   string        `bson:"phone_number"`
	GoogleToken   string        `bson:"google_token"`
	FacebookToken string        `bson:"facebook_token"`
	PasswordHash  string        `bson:"password_hash"`
	//Analytics Information
	SignUpDate   time.Time `bson:"signup_date"`
	LastLoggedIn time.Time `bson:"last_login"`
	LoginCount   uint      `bson:"login_count"`
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
func NewUser(name, email, passwordHash, phoneNumber string) *User {
	return &User{
		ID:           bson.NewObjectId(),
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
		PhoneNumber:  phoneNumber,
		SignUpDate:   time.Now(),
		LastLoggedIn: time.Now(),
		LoginCount:   1,
	}

}

//FindByEmail : To find user by email.
func (u *User) FindByEmail(email string, db *mgo.Database) error {
	return u.Coll(db).Find(bson.M{"email": email}).One(u)
}

//Coll : Returns Collection
func (*User) Coll(db *mgo.Database) *mgo.Collection {
	return db.C("user")
}

//Save : Function to save user in db
func (u *User) Save(db *mgo.Database) error {
	err := u.Coll(db).Insert(u)
	return err
}
