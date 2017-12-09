package data

import "gopkg.in/mgo.v2"

type Context struct {
	Session *mgo.Session
}

func (c *Context)Close() {
	c.Session.Close()
}

func (c *Context)DBCollection(name string) *mgo.Collection {
	return c.Session.DB(DBName).C(CName)
}

func NewContext () *Context{
	session := GetSession()
	c := &Context{
		Session: session,
	}
	return c
}