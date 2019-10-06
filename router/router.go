package router

import (
	"Gosession/controller"
	"log"
	"net/http"
)

func Router() {
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/userinfo", controller.Userinfo)
	log.Println("Server is running at http://localhost:8080/. Press Ctrl+C to stop.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
