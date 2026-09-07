package database

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	Client          *mongo.Client
	UserCollection  *mongo.Collection
	HobbyCollection *mongo.Collection
	Ctx             = context.Background()
)

type config struct {
	uri             string
	database        string
	userCollection  string
	hobbyCollection string
}

func configFromEnv() config {
	return config{
		uri:             envOrDefault("MONGODB_URI", "mongodb://localhost:27017"),
		database:        envOrDefault("MONGODB_DATABASE", "go-mongodb"),
		userCollection:  envOrDefault("MONGODB_USERS_COLLECTION", "users"),
		hobbyCollection: envOrDefault("MONGODB_HOBBIES_COLLECTION", "hobbies"),
	}
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

func Setup() error {
	cfg := configFromEnv()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.uri))
	if err != nil {
		return err
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		_ = client.Disconnect(ctx)
		return err
	}

	Client = client
	db := client.Database(cfg.database)
	UserCollection = db.Collection(cfg.userCollection)
	HobbyCollection = db.Collection(cfg.hobbyCollection)
	return nil
}

func Disconnect(ctx context.Context) error {
	if Client == nil {
		return nil
	}

	return Client.Disconnect(ctx)
}
