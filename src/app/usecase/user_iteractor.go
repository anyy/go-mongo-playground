package usecase

import (
	"github.com/gazelle0130/go-mongo-playground/src/app/domain"
	"github.com/go-chi/render"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (i *UserInteractor) Add(u *domain.User) (interface{}, error) {
	return i.UserRepository.Store(u)
}

func (i *UserInteractor) Get() ([]render.Renderer, error) {
	res, err := i.UserRepository.FindALL()
	if err != nil {
		return nil, err
	}
	var ren []render.Renderer
	for _, v := range res {
		ren = append(ren, v)
	}
	return ren, nil
}
