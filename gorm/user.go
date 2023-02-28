package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var db *gorm.DB
var connectFailed string

// func InitialMigration() {
// 	db, err := gorm.Open("mysql", "testdb")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		panic("Failed to connect to database")
// 	}
// 	fmt.Println("Success")
// 	defer db.Close()
// 	db.AutoMigrate(&User{})
// }

func AllUsers(w http.ResponseWriter, r *http.Request) {
	dsn := "root:Anhminh96@@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	connectFailed = "Could not connect the data base"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Aaaa")
		panic(connectFailed)
	}
	// defer db.
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New User Endpoint Hit")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User Endpoint Hit")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update User Endpoint Hit")
}
