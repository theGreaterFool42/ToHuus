package main

import (
	//"fmt"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"html/template"
	"log"
	"net/http"
	"ToHuus/data"
)

type Person struct {
	Name  string
	Phone string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", someFunc)
	http.ListenAndServe(":8080", nil)
	data.GetSession()
}

func someFunc(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Println(err)
	}
}
