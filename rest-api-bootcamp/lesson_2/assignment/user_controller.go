package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UserController struct {
	COUNT int // count number of users for generating id
	DB    Db  //in memory db
}

// Following functions are bind to UserController Struct -> its like declaring functions in a class
// Also since all these functions will be added to mux route so they will have request and response as params

// This is a format to bind function to a struct
func (controller *UserController) CreateUser(rw http.ResponseWriter, rq *http.Request) {

	// request body contains the user data in JSON format for which the user needs to be created. This request comes from View(frontend)
	reqBody, err := ioutil.ReadAll(rq.Body)
	if err != nil {
		//error
		fmt.Println("error in reading")
		return
	}

	// Convert request body from bytes to json and store in user object
	var userData User
	err = json.Unmarshal(reqBody, &userData)
	if err != nil {
		fmt.Println("error in unmarshaling")

		return
	}

	// assign some id to user
	controller.COUNT = controller.COUNT + 1
	userData.Id = controller.COUNT

	// Add userData to DB
	controller.DB.Users = append(controller.DB.Users, userData)

	// Sending user information in response
	userDataJsonBytes, err := json.Marshal(userData)
	if err != nil {
		fmt.Println("error in marshaling")

		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(userDataJsonBytes)

}

func (controller *UserController) GetAllUsers(rw http.ResponseWriter, rq *http.Request) {
	usersDataJsonBytes, err := json.Marshal(controller.DB.Users)
	if err != nil {
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(usersDataJsonBytes)
}

func (controller *UserController) GetUserById(rw http.ResponseWriter, rq *http.Request) {

}

func (controller *UserController) UpdateUserById(rw http.ResponseWriter, rq *http.Request) {

}

func (controller *UserController) DeleteUserById(rw http.ResponseWriter, rq *http.Request) {

}
