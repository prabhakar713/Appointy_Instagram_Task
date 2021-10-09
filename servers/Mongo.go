package services

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongoDB() {
	Ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(Ctx, options.Client().ApplyURI("mongodb+srv://prabhakar:prabhakar@cluster0.gi0r4.mongodb.net/test"))

	fmt.Println(reflect.TypeOf(Ctx))

	MongoClient = client

	database, err := client.ListDatabaseNames(Ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(database)

	if err != nil {
		log.Fatal(err)
	}
}
