package infrastructure

import (
	"context"
	"os"
	"time"

	"github.com/gazelle0130/go-mongo-playground/src/app/interfaces/database"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type KVSHandler struct {
	Client *mongo.Client
}

func (k *KVSHandler) GetCollection(db, col string) *mongo.Collection {
	return k.Client.Database(db).Collection(col)
}

func NewKVSHandler() (database.KVSHandler, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_DB_URI")))
	if err != nil {
		return nil, err
	}
	k := new(KVSHandler)
	k.Client = client
	return k, nil
}
