package curd

import (
	"server/database/mcli"

	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type JSONResponse struct {
	StatusCode int                    `json:"status"`
	Headers    map[string]string      `json:"headers"`
	Body       map[string]interface{} `json:"body"`
}

func FindoneUser(c *gin.Context) {
	type UserJSON struct {
		Database string `json:"database"`
		Table    string `json:"table"`
		Usern    string `json:"username"`
	}
	var ujs *UserJSON
	var (
		client     = mcli.Getclient()
		collection *mongo.Collection
		err        error
		result     map[string]interface{}
	)
	if err := c.ShouldBindJSON(&ujs); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}

	collection = client.Database(ujs.Database).Collection(ujs.Table)
	filter := bson.M{"username": ujs.Usern}

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	defer mcli.Disconnect(client)

	// fmt.Println(result)
	c.JSON(http.StatusAccepted, result)

}

func FindUser(UserName string) (result map[string]interface{}) {
	var (
		client     = mcli.Getclient()
		collection *mongo.Collection
		err        error
	)

	collection = client.Database("testDatabase").Collection("auth")
	filter := bson.M{"username": UserName}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	defer mcli.Disconnect(client)

	return result
}
