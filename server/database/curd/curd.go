package curd

import (
	"server/database/mcli"

	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {

	// env file loading
	enverr := godotenv.Load("bkms.env")
	if enverr != nil {
		log.Fatal("curd module Error loading .env file", enverr)
	} else {
		log.Print("curd module Success loading .env file")
	}
}

func FindUser(UserName string) (result map[string]interface{}) {
	var (
		client     = mcli.Getclient()
		collection *mongo.Collection
		err        error
		database   = os.Getenv("BKMS_DTATBASE")
	)

	collection = client.Database(database).Collection("auth")
	filter := bson.M{"username": UserName}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	defer mcli.Disconnect(client)

	return result
}

func FindBookbyName(c *gin.Context) {
	type BookByName struct {
		Name string `json:"name"`
	}
	var (
		client     = mcli.Getclient()
		collection *mongo.Collection
		bkfibyn    *BookByName
		result     []map[string]interface{}
		database   = os.Getenv("BKMS_DTATBASE")
	)

	if err := c.ShouldBindJSON(&bkfibyn); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	collection = client.Database(database).Collection("book")

	filter := bson.M{
		"name": primitive.Regex{
			Pattern: bkfibyn.Name,
			Options: "i",
		},
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	err = cursor.All(context.TODO(), &result)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, result)
}

func FindBookbyAuthor(c *gin.Context) {
	type BookByAuthor struct {
		Author string `json:"author"`
	}
	var (
		client     = mcli.Getclient()
		collection *mongo.Collection
		bkfibyau   *BookByAuthor
		result     []map[string]interface{}
		database   = os.Getenv("BKMS_DTATBASE")
	)

	if err := c.ShouldBindJSON(&bkfibyau); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	collection = client.Database(database).Collection("book")

	filter := bson.M{
		"author": primitive.Regex{
			Pattern: bkfibyau.Author,
			Options: "i",
		},
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	err = cursor.All(context.TODO(), &result)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, result)
}

func FindBookbyISBN(c *gin.Context) {
	type BookByISBN struct {
		ISBN string `json:"isbn"`
	}
	var (
		client     = mcli.Getclient()
		collection *mongo.Collection
		bkfibyisbn *BookByISBN
		result     map[string]interface{}
		database   = os.Getenv("BKMS_DTATBASE")
	)

	if err := c.ShouldBindJSON(&bkfibyisbn); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	collection = client.Database(database).Collection("book")
	if len(bkfibyisbn.ISBN) != 10 && len(bkfibyisbn.ISBN) != 13 {
		c.JSON(http.StatusUnprocessableEntity, "Invalid ISBN Format")
		return
	}
	filter := bson.M{"ISBN": bkfibyisbn.ISBN}

	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusOK, "No result found")
			return
		}
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, result)
}

func AddOneBook(c *gin.Context) {
	type BookInsertStructure struct {
		ISBN   string `json:"isbn"`
		Name   string `json:"name"`
		Author string `json:"author"`
		EP     string `json:"ep"`
		End    bool   `json:"end"`
	}
	type Req struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
	var (
		client     = mcli.Getclient()
		collection *mongo.Collection
		database   = os.Getenv("BKMS_DTATBASE")
		insertBook BookInsertStructure
		sucreq     Req
	)
	if err := c.ShouldBindJSON(&insertBook); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	if len(insertBook.ISBN) != 10 && len(insertBook.ISBN) != 13 {
		sucreq.Success = false
		sucreq.Message = "ISBN format error."
		c.JSON(http.StatusUnprocessableEntity, sucreq)
		return
	}
	collection = client.Database(database).Collection("book")
	docs := bson.D{
		{Key: "ISBN", Value: insertBook.ISBN},
		{Key: "name", Value: insertBook.Name},
		{Key: "author", Value: insertBook.Author},
		{Key: "ep", Value: insertBook.EP},
		{Key: "end", Value: insertBook.End},
	}
	result, err := collection.InsertOne(context.TODO(), docs)
	if err != nil {
		log.Fatal(err)
	}
	sucreq.Success = true
	sucreq.Message = fmt.Sprintf("Insert ID:%v", result.InsertedID)

	c.JSON(http.StatusOK, sucreq)

}

func DeleteOneBook(c *gin.Context) {
	type DeleteBook struct {
		ID primitive.ObjectID `json:"id"`
	}
	type DeleteBookResponse struct {
		Success bool   `json:"success"`
		ID      string `json:"id"`
		Message string `json:"message"`
	}
	var (
		client         = mcli.Getclient()
		collection     *mongo.Collection
		database       = os.Getenv("BKMS_DTATBASE")
		deletebook     DeleteBook
		deleteResponce DeleteBookResponse
	)
	if err := c.ShouldBindJSON(&deletebook); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	collection = client.Database(database).Collection("book")
	filter := bson.D{{Key: "_id", Value: deletebook.ID}}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	if result.DeletedCount == 0 {
		deleteResponce.Success = false
		deleteResponce.ID = deletebook.ID.Hex()
		deleteResponce.Message = "Delete method  can not remove document."
	} else {
		deleteResponce.Success = true
		deleteResponce.ID = deletebook.ID.Hex()
		deleteResponce.Message = fmt.Sprintf("Delete method removed %v document(s)\n", result.DeletedCount)
	}
	c.JSON(http.StatusOK, deleteResponce)
}

func ModifyOneBook(c *gin.Context) {
	type BookModifyStruct struct {
		ID         primitive.ObjectID `json:"id"`
		ModifyItem bson.M             `json:"modifyItem"`
	}
	type BookModifyResponse struct {
		ID      string `json:"id"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
	var (
		client         = mcli.Getclient()
		collection     *mongo.Collection
		database       = os.Getenv("BKMS_DTATBASE")
		modifyBook     BookModifyStruct
		modifyResponse BookModifyResponse
	)
	if err := c.ShouldBindJSON(&modifyBook); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	collection = client.Database(database).Collection("book")
	filter := bson.M{"_id": modifyBook.ID}
	updatedoc := bson.D{
		{"$set", modifyBook.ModifyItem},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, updatedoc)
	if err != nil {
		log.Fatal(err)
	} else {
		modifyResponse.ID = modifyBook.ID.Hex()
		modifyResponse.Success = true
		modifyResponse.Message = fmt.Sprintf("Modify %v item.", result.ModifiedCount)
	}
	c.JSON(http.StatusOK, modifyResponse)
}
