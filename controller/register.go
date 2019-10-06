package controller

import (
	"Gosession/dao"
	"Gosession/daomain"
	"io/ioutil"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		registerHTML, err := ioutil.ReadFile("html/register.html")
		checkError(err)
		w.Write(registerHTML)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	password2 := r.FormValue("password2")
	email := r.FormValue("email")

	if isEmpty(username, password, password2, email) {
		message(w, r, "字段不能为空")
		return
	}

	if password != password2 {
		message(w, r, "两次密码不相符")
		return
	}

	user := &daomain.User{
		Username: username,
		Password: password,
		Email:    email,
	}
	dao.AddUser(user)
	message(w, r, "注册成功！")
}
