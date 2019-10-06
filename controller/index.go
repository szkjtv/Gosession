package controller

import (
	"Gosession/session"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	user, _ := session.GetSession(w, r).GetAttr("user")

	t, err := template.ParseFiles("html/index.html")
	checkError(err)

	err = t.Execute(w, user)
	checkError(err)
}
