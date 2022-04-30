package mcli

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {

	// env file loading
	enverr := godotenv.Load("bkms.env")
	if enverr != nil {
		log.Fatal("mcli module Error loading .env file", enverr)
	} else {
		log.Print("mcli module Success loading .env file")
	}

}

func initEngine() {
	var err error
	db_server := os.Getenv("MONGODB_SERVER")

	clientOptions := options.Client().ApplyURI(db_server)

	// connect to MongoDB
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// connection check
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
}
func Getclient() *mongo.Client {
	if client == nil {
		initEngine()
	}
	return client
}

func Disconnect(client *mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	log.Print("Mongodb Disconnect")
}
