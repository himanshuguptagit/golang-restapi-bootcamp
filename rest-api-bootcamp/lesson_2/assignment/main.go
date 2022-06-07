package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func main() {
	PORT := 8081

	// Create an empty user list for storing users data.
	userList := make([]User, 0)

	//Initialize DB
	db := Db{
		Users: userList,
	}

	//Initialize controller
	userController := UserController{
		COUNT: 0,
		DB:    db,
	}

	fmt.Println("newn enwlanfdkslamfklas")
	// Initializing mux router and routes
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", userController.CreateUser).Methods("POST")
	router.HandleFunc("/users", userController.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{user_id}", userController.GetUserById).Methods("GET")
	router.HandleFunc("/users/{user_id}", userController.UpdateUserById).Methods("PUT")
	router.HandleFunc("/users/{user_id}", userController.DeleteUserById).Methods("DELETE")

	//start the http server
	fmt.Println("Starting http server on port: " + strconv.Itoa(PORT))
	err := http.ListenAndServe(":"+strconv.Itoa(PORT), router)
	if err != nil {
		fmt.Println("error in starting server")
		return
	}

}
