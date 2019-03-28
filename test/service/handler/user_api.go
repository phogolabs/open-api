//
// Swagger Petstore
// 
// This is a sample Petstore server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/). 
// 
// API version: 1.0.0
// Contact: apiteam@swagger.io
//
// Code generated by OpenAPI Generator (https://http://github.com/openAPITools/openapi-generator)

// Package handler contains all API models and handlers
package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/phogolabs/rest"
)

// UserAPI handler
type UserAPI struct{
 // TODO: add your dependencies here
}

// Mount mounts the UserAPI handler
func (h *UserAPI) Mount(r chi.Router) {
  r.Group(func(r chi.Router) {
     r.Post("/user", h.createUser)
     r.Post("/user/createWithArray", h.createUsersWithArrayInput)
     r.Post("/user/createWithList", h.createUsersWithListInput)
     r.Delete("/user/{username}", h.deleteUser)
     r.Get("/user/{username}", h.getUserByName)
     r.Get("/user/login", h.loginUser)
     r.Get("/user/logout", h.logoutUser)
     r.Put("/user/{username}", h.userUsernamePut)
  })
}

// CreateUser operation: Create user
// This can only be done by the logged in user.
// @param User: Created user object
func (h *UserAPI) createUser(w http.ResponseWriter, r *http.Request) {
  input := &CreateUserInput{}

  if err := rest.Bind(w, r, input); err != nil {
    rest.Error(w, r, err)
    return
  }

  //TODO: implement the operation createUser
}


// CreateUsersWithArrayInput operation: Creates list of users with given input array
// @param User: List of user object
func (h *UserAPI) createUsersWithArrayInput(w http.ResponseWriter, r *http.Request) {
  input := &CreateUsersWithArrayInputInput{}

  if err := rest.Bind(w, r, input); err != nil {
    rest.Error(w, r, err)
    return
  }

  //TODO: implement the operation createUsersWithArrayInput
}


// CreateUsersWithListInput operation: Creates list of users with given input array
// @param User: List of user object
func (h *UserAPI) createUsersWithListInput(w http.ResponseWriter, r *http.Request) {
  input := &CreateUsersWithListInputInput{}

  if err := rest.Bind(w, r, input); err != nil {
    rest.Error(w, r, err)
    return
  }

  //TODO: implement the operation createUsersWithListInput
}


// DeleteUser operation: Delete user
// This can only be done by the logged in user.
// @param Username: The name that needs to be deleted (default: null)
func (h *UserAPI) deleteUser(w http.ResponseWriter, r *http.Request) {
  input := &DeleteUserInput{}

  if err := rest.Bind(w, r, input); err != nil {
    rest.Error(w, r, err)
    return
  }

  //TODO: implement the operation deleteUser
}


// GetUserByName operation: Get user by user name
// @param Username: The name that needs to be fetched. Use user1 for testing. (default: null)
// @return User
func (h *UserAPI) getUserByName(w http.ResponseWriter, r *http.Request) {
  input := &GetUserByNameInput{}

  if err := rest.Bind(w, r, input); err != nil {
    rest.Error(w, r, err)
    return
  }

  //TODO: implement the operation getUserByName

  output := &GetUserByNameOutput{}
  rest.Render(w, r, output)
}


// LoginUser operation: Logs user into the system
// @param Username: The user name for login (default: null)
// @param Password: The password for login in clear text (default: null)
// @return string
func (h *UserAPI) loginUser(w http.ResponseWriter, r *http.Request) {
  input := &LoginUserInput{}

  if err := rest.Bind(w, r, input); err != nil {
    rest.Error(w, r, err)
    return
  }

  //TODO: implement the operation loginUser

  output := &LoginUserOutput{}
  rest.Render(w, r, output)
}


// LogoutUser operation: Logs out current logged in user session
func (h *UserAPI) logoutUser(w http.ResponseWriter, r *http.Request) {
  //TODO: implement the operation logoutUser
}


// UserUsernamePut operation: Updated user
// This can only be done by the logged in user.
// @param Username: name that need to be updated (default: null)
// @param User: Updated user object
func (h *UserAPI) userUsernamePut(w http.ResponseWriter, r *http.Request) {
  input := &UserUsernamePutInput{}

  if err := rest.Bind(w, r, input); err != nil {
    rest.Error(w, r, err)
    return
  }

  //TODO: implement the operation userUsernamePut
}

