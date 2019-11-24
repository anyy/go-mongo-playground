package usecase

import "github.com/gazelle0130/go-mongo-playground/src/app/domain"

type UserRepository interface {
	Store(*domain.User) (interface{}, error)
	FindALL() ([]domain.User, error)
}
