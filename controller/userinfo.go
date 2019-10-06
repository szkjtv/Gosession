package controller

import (
	"Gosession/dao"
	"Gosession/daomain"
	"Gosession/session"
	"html/template"
	"log"
	"net/http"
)

//用户信息
func Userinfo(w http.ResponseWriter, r *http.Request) {
	sess := session.GetSession(w, r)
	user, exist := sess.GetAttr("user")
	if !exist {
		http.Redirect(w, r, "/", 302)
		return
	}

	if r.Method == "GET" {
		t, err := template.ParseFiles("html/userinfo.html")
		checkError(err)
		err = t.Execute(w, user)
		checkError(err)
		return
	}

	// POST 更新用户信息
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	if isEmpty(username, password, email) {
		message(w, r, "字段不能为空")
		return
	}

	switch user := user.(type) {
	case *daomain.User:
		user.Username = username
		user.Password = password
		user.Email = email
		dao.UpdateUser(user)
	default:
		log.Println(":userinfo:user.(type)", user)
	}
	http.Redirect(w, r, "/userinfo", 302)
}
