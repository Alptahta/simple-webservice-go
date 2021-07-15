package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Alptahta/simple-webservice-go/internal"
)

type UserService interface {
	Create(name string) error
	Find(id uint) (internal.User, error)
}

type UserHandler struct {
	service UserService
}

func (u UserHandler) Register(r *http.ServeMux) {
	r.HandleFunc("/users", u.create)
}

type CreateUserRequest struct {
	Name string `json:"name"`
}

func (u UserHandler) create(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		return
	}
	defer r.Body.Close()

	err := u.service.Create(req.Name)
	if err != nil {
		log.Println(err)
		return
	}

	renderResponse(w, struct{}{}, http.StatusOK)
}

type FindUserRequest struct {
	ID uint `json:"id"`
}

type UserResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (u UserHandler) find(w http.ResponseWriter, r *http.Request) {
	var req FindUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		return
	}
	defer r.Body.Close()

	user, err := u.service.Find(req.ID)
	if err != nil {
		log.Println(err)
		return
	}

	renderResponse(w, UserResponse{ID: user.ID, Name: user.Name}, http.StatusOK)
}

func renderResponse(w http.ResponseWriter, res interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")

	content, err := json.Marshal(res)
	if err != nil {
		// error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)

	if _, err = w.Write(content); err != nil {
		//  error
	}
}
