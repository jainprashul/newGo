package server

import (
	"xpJain.co/bookserver/auth"
	"xpJain.co/bookserver/server/middleware"
)

func AuthRouteInit() {

	usrRoute := New_GormRouteHandler(auth.UserDB)

	usrRoute.router.Use(middleware.Authentication)

	usrRoute.InitService()
	

	Router.HandleFunc("/login", auth.Login).Methods("POST")
	Router.HandleFunc("/register", auth.Register).Methods("POST")
	Router.HandleFunc("/logout", auth.Logout).Methods("GET")

}