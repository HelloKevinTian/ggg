package dao

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// TestMongo ...
func TestMongo() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.1.73:27000"))
	if err != nil {
		fmt.Println("mongo.NewClient error")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Ping error", err)
	} else {
		fmt.Println("Ping OK")
	}

	collection := client.Database("bi_plat").Collection("googleaccounts")
	var result struct {
		ID          string    `bson:"_id"`
		AccountID   string    `bson:"account_id"`
		AccountName string    `bson:"account_name"`
		V           int       `bson:"__v"`
		UpdateTime  time.Time `bson:"update_time"`
		CreateTime  time.Time `bson:"create_time"`
	}
	filter := bson.M{"account_id": "487-641-6360"}
	defer cancel()
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		// fmt.Println("FindOne error", err)
		log.Fatal(err)
	}
	fmt.Println(result, result.AccountID, result.AccountName, result.CreateTime)
	fmt.Println(result.CreateTime.Format("2006-01-02 15:04:05"))
}
