package controllers

import (
	"net/http"

	"github.com/gazelle0130/go-mongo-playground/src/app/domain"
	"github.com/gazelle0130/go-mongo-playground/src/app/interfaces/database"
	"github.com/gazelle0130/go-mongo-playground/src/app/interfaces/helper"
	"github.com/gazelle0130/go-mongo-playground/src/app/usecase"
	"github.com/go-chi/render"
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
	user := &domain.User{}
	if err := render.Bind(r, user); err != nil {
		render.Render(w, r, helper.ErrInvalidRequest(err))
		return
	}

	_, err := c.Interactor.Add(user)
	if err != nil {
		render.Render(w, r, helper.ErrInvalidRequest(err))
		return
	}
	render.Status(r, http.StatusCreated)
	render.Render(w, r, user)
}

func (c *UserController) Index(w http.ResponseWriter, r *htttp.Request) {
	res, err := c.Interactor.Get()
	if err != nil {
		render.Render(w, r, helper.ErrInvalidRequest(err))
		return
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, res)
}
