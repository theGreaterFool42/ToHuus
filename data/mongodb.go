package data

import (
	"gopkg.in/mgo.v2"
	"log"
)

const (
	DBName = "toHuusdb"
	CName = "users"
)

var session *mgo.Session

func createDBSession(){
	var err error
	session, err = mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
}

func GetSession() *mgo.Session{
	if session == nil {
		createDBSession()
	}
	return session
}

func Init() {
	createDBSession()
}
