package database

import "go.mongodb.org/mongo-driver/mongo"

type KVSHandler interface {
	GetCollection(string, string) *mongo.Collection
}
