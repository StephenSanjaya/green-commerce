package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Orders *mongo.Collection
}

var instance *Database
var once sync.Once

func GetMongoConnection() *Database {
	once.Do(func() {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("failed to load env")
			return
		}

		mongo_uri := os.Getenv("MONGO_URI")
		mongo_name := os.Getenv("MONGO_NAME")
		mongo_coll := os.Getenv("MONGO_COLLECTION")

		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongo_uri))
		if err != nil {
			log.Fatal("failed to connect mongo db")
			return
		}

		orderColl := client.Database(mongo_name).Collection(mongo_coll)
		instance = &Database{Orders: orderColl}
	})

	fmt.Println("DB CONNECTED")

	return instance
}
