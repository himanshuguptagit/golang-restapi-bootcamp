package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type UserController struct {
	COUNT int
	USERS []User // USERS is list of type user which acts as a kind of in memory db to store users.
}

// Following functions are bind to UserController Struct -> its like declaring functions in a class
// Also since all these functions will be added to mux route so they will have request and response as params

// This is a format to bind function to a struct
func (controller *UserController) CreateUser(rw http.ResponseWriter, rq *http.Request) {

	// request body contains the user data in JSON format for which the user needs to be created. This request comes from View(frontend)

	reqBody, err := ioutil.ReadAll(rq.Body)
	if err != nil {
		//error
		return
	}

	// Convert request body from bytes to json and store in user object
	var userData User
	err := json.Unmarshal(reqBody, &userData)
	if err != nil {
		return
	}

	// assign some id to user
	controller.COUNT = controller.COUNT + 1
	userData.Id = controller.COUNT

	// Add userData to DB
	controller.USERS = append(controller.USERS, userData)

	// Sending user information in response
	userDataJsonBytes, err := json.Marshal(userData)
	if err != nil {
		return
	}
	rw.WriteHeader(200)
	rw.Write(userDataJsonBytes)

}

func (controller *UserController) GetAllUsers(rw http.ResponseWriter, rq *http.Request) {

}

func (controller *UserController) GetUserById(rw http.ResponseWriter, rq *http.Request) {

}

func (controller *UserController) UpdateUserById(rw http.ResponseWriter, rq *http.Request) {

}

func (controller *UserController) DeleteUserById(rw http.ResponseWriter, rq *http.Request) {

}

