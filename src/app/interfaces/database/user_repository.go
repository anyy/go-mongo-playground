package database

import (
	"context"
	"time"

	"github.com/gazelle0130/go-mongo-playground/src/app/domain"
	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository struct {
	KVSHandler
}

func (r *UserRepository) Store(u *domain.User) (interface{}, error) {
	col := r.KVSHandler.GetCollection("mongo-playgroud", "user")
	ctx, cf := context.WithTimeout(context.Background(), 5*time.Second)
	defer cf()
	res, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", err
	}
	return res.InsertedID, err
}

func (r *UserRepository) FindALL() ([]*domain.User, error) {
	col := r.KVSHandler.GetCollection("mongo-playgroud", "user")
	res, err := col.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var result []*domain.User
	for res.Next(context.Background()) {
		var u *domain.User
		if err = res.Decode(&u); err != nil {
			return nil, err
		}
		result = append(result, u)
	}
	return result, nil
}
