package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "github.com/jinzhu/gorm"
)

type User struct {
	Name string `json:"name"`
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "AAA")
}
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	// myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{name}/{email}", NewUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
func handleDatabase() {
	db, err := sql.Open("mysql", "root:Anhminh96@@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}

	// insert, err := db.Query("INSERT INTO users VALUES('ELLIOT')")

	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()

	// fmt.Println("succesfully inserted")

	defer db.Close()
	fmt.Println("Successfully connected")
	results, err := db.Query("SELECT name FROM users")

	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var user User

		err = results.Scan(&user.Name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.Name)
	}
}

func main() {
	// handleRequests()
	dsn := "root:Anhminh96@@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	connectFailed = "Could not connect the data base"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("error", err)
		panic(connectFailed)
	}

	err = Migrate(db)
	log.Printf("Migration did run successfully")
	if err != nil {
		fmt.Println("error", err)
		panic(connectFailed)
	}

}
