package main

import (
	"fmt"
	"net/http"
)

func homePage(rw http.ResponseWriter, rq *http.Request) {
	// Bind this function to mux route
	fmt.Fprintf(rw, "Hello world")
}

func usersPage(rw http.ResponseWriter, rq *http.Request) {
	// Complete and Bind this function to mux route

}

func userPage(rw http.ResponseWriter, rq *http.Request) {
	// Complete and Bind this function to mux route

}

func main() {
	PORT := 8081

	// Initializing mux router and add routes

	//start the http server on PORT

}
