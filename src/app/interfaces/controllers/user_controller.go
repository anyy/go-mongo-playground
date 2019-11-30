package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gazelle0130/go-mongo-playground/src/app/domain"
	"github.com/gazelle0130/go-mongo-playground/src/app/interfaces/database"
	"github.com/gazelle0130/go-mongo-playground/src/app/usecase"
	"github.com/go-chi/chi"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(k database.KVSHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				KVSHandler: k,
			},
		},
	}
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user domain.User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = json.Unmarshal(body, &user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = c.Interactor.Add(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (c *UserController) Index(w http.ResponseWriter, r *http.Request) {
	res, err := c.Interactor.Get()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	err := c.Interactor.DeleteByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
