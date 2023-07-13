package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main(){
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("",)
	r.POST("",)
	r.DELETE("",)
}

func getSession() *mgo.Session{
	s,err := mgo.Dial("mongo://localhost:27107")
	if err != nil{
		panic(err)
	}
	return s
}