package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/astaxie/beego/logs"
)

type UsersAll struct {
	Objects []Users `json:"objects"`
}

type Users struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Email     string             `json:"email" bson:"email"`
	LastName  string             `json:"last_name" bson:"last_name"`
	Country   string             `json:"country" bson:"country"`
	City      string             `json:"city" bson:"city"`
	Gender    string             `json:"gender" bson:"gender"`
	BirthDate string             `json:"birth_date" bson:"birth_date"`
}

func main() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		logs.Error(err)
		return
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		logs.Error(err)
	}
	fmt.Println("Connected to MongoDB!")
	collection := client.Database("bitmedia_test_task").Collection("users")

	// read our opened xmlFile as a byte array.
	var byteValue []byte
	if byteValue, err = ioutil.ReadFile("users.json"); err != nil {
		logs.Error(err)
	}

	// we initialize our Users array
	var users UsersAll

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'objects' which we defined above
	if err = json.Unmarshal(byteValue, &users); err != nil {
		logs.Error(err)
	}

	// we iterate through every user within our objects array and
	for _, user := range users.Objects {
		user.Id = primitive.NewObjectID()
		insertResult, err := collection.InsertOne(context.TODO(), user)
		if err != nil {
			logs.Error(err)
		}
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	}
	err = client.Disconnect(context.TODO())

	if err != nil {
		logs.Error(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
