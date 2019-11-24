package usecase

import (
	"github.com/gazelle0130/go-mongo-playground/src/app/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (i *UserInteractor) Add(u *domain.User) (interface{}, error) {
	return i.UserRepository.Store(u)
}

func (i *UserInteractor) Get() ([]domain.User, error) {
	return i.UserRepository.FindALL()
}
