package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"

	"github.com/parth/mongo-golang/controllers"
)

func main(){
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id",uc.GetUser)
	r.POST("/user",uc.CreateUser)
	r.DELETE("/user/:id",uc.DeleteUser)
	http.ListenAndServe("localhost:8080",r)
}

func getSession() *mgo.Session{
	s,err := mgo.Dial("YOURMONGOURL")
	if err != nil{
		panic(err)
	}
	return s
}