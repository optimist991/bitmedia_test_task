package models

import (
	"bitmedia_test_task/models/db"

	"gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"
)

type Users struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Email     string        `json:"email" bson:"email"`
	LastName  string        `json:"last_name" bson:"last_name" validate:"min=3, max=40, regexp=^[a-zA-Z]*$"`
	Country   string        `json:"country" bson:"country" validate:"min=3, max=40, regexp=^[a-zA-Z]*$"`
	City      string        `json:"city" bson:"city" validate:"min=3, max=40, regexp=^[a-zA-Z]*$"`
	Gender    string        `json:"gender" bson:"gender" validate:"min=3, max=40, regexp=^[a-zA-Z]*$"`
	BirthDate string        `json:"birth_date" bson:"birth_date"`
}

func NewUsersCollection() *db.Collection {
	return db.NewCollectionSession("users")
}

func CreateUser(user Users) (Users, error) {
	var err error

	// Get users collection connection
	c := NewUsersCollection()

	// set default mongodb ID
	user.Id = bson.NewObjectId()
	if err = c.Session.EnsureIndex(mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		return Users{}, err
	}

	// Insert user to mongodb
	err = c.Session.Insert(&user)
	if err != nil {
		return user, err
	}

	defer c.Close()
	return user, err
}

func (user Users) UpdateUser(postParam Users) (Users, error) {
	var (
		err error
	)
	// Get user collection connection
	c := NewUsersCollection()
	defer c.Close()
	// update user

	if err = c.Session.EnsureIndex(mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		return Users{}, err
	}

	err = c.Session.Update(bson.M{
		"_id": user.Id,
	}, bson.M{
		"$set": bson.M{
			"email":      postParam.Email,
			"last_name":  postParam.LastName,
			"country":    postParam.Country,
			"city":       postParam.City,
			"gender":     postParam.Gender,
			"birth_date": postParam.BirthDate,
		},
	})
	if err != nil {
		return user, err
	}
	return user, err
}

func FindAllUsers(limit, skip int) ([]Users, error) {
	var (
		err   error
		users []Users
	)
	// Get user collection connection
	c := NewUsersCollection()
	defer c.Close()
	// get users
	err = c.Session.Find(nil).Sort("_id").Limit(limit).Skip(skip).All(&users)
	if err != nil {
		return users, err
	}
	return users, err
}

func FindUser(id bson.ObjectId) (Users, error) {
	var (
		err  error
		user Users
	)
	// Get post collection connection
	c := NewUsersCollection()
	defer c.Close()
	// get post
	err = c.Session.FindId(id).One(&user)
	if err != nil {
		return user, err
	}
	return user, err
}

func DeleteUser(userId bson.ObjectId) (err error) {

	// Get user collection connection
	c := NewUsersCollection()
	defer c.Close()
	// remove post
	err = c.Session.Remove(bson.M{"_id": userId})
	if err != nil {
		return err

	}
	return nil

}
