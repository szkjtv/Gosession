package dao

import (
	"Gosession/daomain"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const createUserTable = "CREATE TABLE `user` ( `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, `username` TEXT NOT NULL, `password` TEXT NOT NULL, `email` TEXT );"
const createAdmin = "INSERT INTO user(id, username, password, email) VALUES(1, 'admin', 'admin', 'admin@admin.com');"

func init() {
	initEnv()
	initAdmin()
}

func initEnv() {
	db, err := sql.Open("mysql", "root:aa14514118@tcp(127.0.0.1:3306)/session")
	checkError(err)
	defer db.Close()

	db.Exec(createUserTable)
}

func initAdmin() {
	db, err := sql.Open("mysql", "root:aa1451418@tcp(127.0.0.1:3306)/session")
	checkError(err)
	defer db.Close()
	db.Exec(createAdmin)

}

func getDB() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:aa1451418@tcp(127.0.0.1:3306)/session")
	checkError(err)
	return
}

// FindUserByUsernameAndPassword 通过 username 和 password 查找 User
func FindUserByUsernameAndPassword(username string, password string) (user *daomain.User) {
	sql := "select id, email from user where username=? and password=?"

	db := getDB()
	defer db.Close()

	rows, err := db.Query(sql, username, password)
	checkError(err)
	defer rows.Close()

	if rows.Next() {
		var id int
		var email string
		rows.Scan(&id, &email)

		user = &daomain.User{
			ID:       id,
			Username: username,
			Password: password,
			Email:    email,
		}
	}
	return
}

// AddUser 添加新 User
func AddUser(user *daomain.User) {
	sql := "insert into user(username, password, email) values(?,?,?)"

	db := getDB()
	defer db.Close()

	_, err := db.Exec(sql, user.Username, user.Password, user.Email)
	checkError(err)
}

// UpdateUser 更新 User
func UpdateUser(user *daomain.User) {
	sql := "update user set username=?, password=?, email=? where id=?"

	db := getDB()
	defer db.Close()

	_, err := db.Exec(sql, user.Username, user.Password, user.Email, user.ID)
	checkError(err)
}

// DeleteUser 删除 User
func DeleteUser(id int) {
	sql := "delete from user where id=?"

	db := getDB()
	defer db.Close()

	_, err := db.Exec(sql, id)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
