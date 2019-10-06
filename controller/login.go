package controller

import (
	"Gosession/dao"
	"Gosession/session"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

//登录
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		loginHTML, err := ioutil.ReadFile("html/login.html")
		checkError(err)
		w.Write(loginHTML)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	log.Println("login", username, password)
	if isEmpty(username, password) {
		message(w, r, "字段不能为空")
		return
	}

	user := dao.FindUserByUsernameAndPassword(username, password)
	if user == nil {
		message(w, r, "登录失败！")
		return
	}
	// 登陆成功
	sess := session.GetSession(w, r)
	sess.SetAttr("user", user)
	http.Redirect(w, r, "/", 302)
}

//退出登录
func Logout(w http.ResponseWriter, r *http.Request) {
	sess := session.GetSession(w, r)
	sess.DelAttr("user")
	http.Redirect(w, r, "/", 302)
}

func message(w http.ResponseWriter, r *http.Request, message string) {
	t, err := template.ParseFiles("html/message.html")
	checkError(err)

	err = t.Execute(w, map[string]string{"Message": message})
	checkError(err)
}
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func isEmpty(strs ...string) (isEmpty bool) {
	for _, str := range strs {
		str = strings.TrimSpace(str)
		if str == "" || len(str) == 0 {
			isEmpty = true
			return
		}
	}
	isEmpty = false
	return
}
