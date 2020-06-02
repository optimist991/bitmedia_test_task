package db

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

type Collection struct {
	db      *Database
	name    string
	Session *mgo.Collection
}

func (c *Collection) Connect() {
	session := *c.db.session.C(c.name)
	c.Session = &session
}
func NewCollectionSession(name string) *Collection {
	var c = Collection{
		db:   NewDBSession(beego.AppConfig.String("DBName")),
		name: name,
	}

	c.Connect()
	return &c
}
func (c *Collection) Close() {
	service.Close(c)
}
