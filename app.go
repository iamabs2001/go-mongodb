package main

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://username:password@atlas-cluster/hero?authSource=admin&retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	User := client.Database("hero").Collection("users")
	var result bson.M
	err = User.FindOne(context.TODO(), bson.D{{"username", "hero"}}).Decode(&result)
	jsonData, err := json.MarshalIndent(result, "", "    ")
	fmt.Printf("%s\n", jsonData)
}
