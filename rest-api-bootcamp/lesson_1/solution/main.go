package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func homePage(rw http.ResponseWriter, rq *http.Request) {
	fmt.Fprintf(rw, "Hello World")
}

func usersPage(rw http.ResponseWriter, rq *http.Request) {
	fmt.Fprintf(rw, "<html><b>users list</b></html>")
}

func userPage(rw http.ResponseWriter, rq *http.Request) {
	userId := mux.Vars(rq)["user_id"]
	fmt.Fprintf(rw, "hello user: "+userId)
}

func main() {
	PORT := 8081

	// Initializing mux router and routes
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/users", usersPage).Methods("GET")
	router.HandleFunc("/users/{user_id}", userPage).Methods("GET")

	//start the http server
	fmt.Println("Starting http server on port: " + strconv.Itoa(PORT))
	err := http.ListenAndServe(":"+strconv.Itoa(PORT), router)
	if err != nil {
		fmt.Println("error in starting server")
		return
	}

}
