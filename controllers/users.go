package controllers

import (
	"bitmedia_test_task/models"
	"encoding/json"
	"errors"
	"regexp"

	"github.com/astaxie/beego/logs"
	"gopkg.in/validator.v2"

	"gopkg.in/mgo.v2/bson"
)

// Operations about Users
type UsersController struct {
	BaseController
}

// URLMapping ...
func (c *UsersController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.Get)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

func validEmail(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\." +
		"[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(email)
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UsersController) Post() {
	var user models.Users
	var err error
	if err = json.Unmarshal(u.Ctx.Input.RequestBody, &user); err != nil {
		logs.Error(err)
		u.Response(400, nil, err)
	}

	if !validEmail(user.Email) {
		logs.Error(err)
		u.Response(400, nil, errors.New("not valid Email"))
	}

	if err = validator.Validate(user); err != nil {
		logs.Error(err)
		u.Response(400, nil, err)
	}
	if user, err = models.CreateUser(user); err != nil {
		logs.Error(err)
		u.Response(500, nil, err)
	}
	u.Response(201, user, nil)
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UsersController) GetAll() {
	var users []models.Users
	var err error
	var limit, skip int
	if limit, err = u.GetInt("limit"); err != nil {
		err = errors.New("incorrect param limit")
		logs.Error(err)
		u.Response(400, nil, err)
	}

	if skip, err = u.GetInt("skip"); err != nil {
		err = errors.New("incorrect param skip")
		logs.Error(err)
		u.Response(400, nil, err)
	}

	if users, err = models.FindAllUsers(limit, skip); err != nil {
		logs.Error(err)
		u.Response(400, nil, err)
	}
	u.Response(200, users, nil)

}

// @Title Get
// @Description get user by id
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:id [get]
func (u *UsersController) Get() {
	userId := u.GetString(":id")
	var err error
	if userId == "" {
		err = errors.New("User id is empty")
		logs.Error(err)
		u.Response(403, nil, err)
	}

	var user models.Users
	if user, err = models.FindUser(bson.ObjectIdHex(userId)); err != nil {
		logs.Error(err)
		u.Response(400, nil, err)
	}
	u.Response(200, user, nil)
}

// @Title Update
// @Description update the user
// @Param	id		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.Users	true		"body for user content"
// @Success 200 {object} models.Users
// @Failure 403 :id is not
// @router /:id [put]
func (u *UsersController) Put() {
	userId := u.GetString(":id")
	var err error
	if userId == "" {
		err = errors.New("user id is empty")
		logs.Error(err)
		u.Response(403, nil, err)
	}
	var userParams models.Users
	if err = json.Unmarshal(u.Ctx.Input.RequestBody, &userParams); err != nil {
		logs.Error(err)
		u.Response(400, nil, err)
	}
	var userObject models.Users
	if userObject, err = models.FindUser(bson.ObjectIdHex(userId)); err != nil {
		logs.Error(err)
		u.Response(400, nil, err)
	}
	if !validEmail(userParams.Email) {
		err = errors.New("not valid Email")
		logs.Error(err)
		u.Response(400, nil, err)
	}
	if err = validator.Validate(userParams); err != nil {
		logs.Error(err)
		u.Response(400, nil, err)
	}
	if userObject, err = userObject.UpdateUser(userParams); err != nil {
		logs.Error(err)
		u.Response(500, nil, err)
	}
	u.Response(200, userObject, nil)
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:id [delete]
func (u *UsersController) Delete() {
	userId := u.GetString(":id")
	var err error
	if err = models.DeleteUser(bson.ObjectIdHex(userId)); err != nil {
		logs.Error(err)
		u.Response(500, nil, err)
	}
	u.Response(200, "ok", nil)
}
