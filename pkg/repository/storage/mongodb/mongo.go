package mongodb

import (
	"context"
	"fmt"
	"log"

	"github.com/lordscoba/bible_compass_backend/internal/config"
	"github.com/lordscoba/bible_compass_backend/utility"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ctx         context.Context
	mongoClient *mongo.Client
)

func Connection() (db *mongo.Client) {
	return mongoClient
}

func ConnectToDB() *mongo.Client {
	var err error
	logger := utility.NewLogger()
	uri := config.GetConfig().Mongodb.Url
	mongo_connection := options.Client().ApplyURI(uri)
	mongoClient, err = mongo.Connect(ctx, mongo_connection)
	if err != nil {
		log.Fatal(err)
	}

	// PINGING THE CONNECTION
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	// IF EVERYTHING IS OKAY, THEN CONNECT
	fmt.Println("MONGO CONNECTION ESTABLISHED")
	logger.Info("MONGO CONNECTION ESTABLISHED")

	return mongoClient
}

// getting database collections
func getCollection(collection string) *mongo.Collection {
	databaseName := config.GetConfig().Mongodb.Database
	database := mongoClient.Database(databaseName)
	c := database.Collection(collection)

	return c
}

func MongoPost[T any](collection string, data T) (*mongo.InsertOneResult, error) {
	c := getCollection(collection)

	result, err := c.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func MongoGetOne[T any](collection string, filter map[string]T) (*mongo.SingleResult, error) {
	c := getCollection(collection)
	f := make(bson.M)
	for k, v := range filter {
		f = bson.M{k: v}
	}

	result := c.FindOne(context.TODO(), f)
	if result.Err() != nil {
		return nil, result.Err()
	}

	return result, nil
}

func MongoCount[T any](collection string, filter map[string]T) (int64, error) {
	c := getCollection(collection)
	f := make(bson.M)
	for k, v := range filter {
		f = bson.M{k: v}
	}

	result, err := c.CountDocuments(context.TODO(), f)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func MongoGet[T any](collection string, filter map[string]T) (*mongo.Cursor, error) {
	c := getCollection(collection)

	f := make(bson.M)

	if len(filter) == 1 {
		for k, v := range filter {
			f = bson.M{k: v}
		}
	} else if len(filter) > 1 {
		tf := make([]bson.M, 0, len(filter))
		for k, v := range filter {
			tf = append(tf, bson.M{k: v})
		}

		f = bson.M{"$and": tf}
	}

	cursor, err := c.Find(ctx, f)
	if err != nil {
		return nil, err
	}
	return cursor, nil
}

func MongoUpdate[T any](id string, updateEntries map[string]T, collection string) (*mongo.UpdateResult, error) {
	c := getCollection(collection)
	apiCallCount := "api_call_count"
	user_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if value, ok := updateEntries[apiCallCount]; ok {
		_, err := c.UpdateByID(ctx, user_id, bson.M{"$inc": bson.M{apiCallCount: value}}, options.Update().SetUpsert(true))
		if err != nil {
			return nil, err
		}

		delete(updateEntries, apiCallCount)
		if len(updateEntries) == 0 {
			return nil, nil
		}
	}

	update := make(bson.M, 0)
	for i, j := range updateEntries {
		update[i] = j
	}

	db_data := bson.M{"$set": update}

	result, err := c.UpdateByID(context.TODO(), user_id, db_data)

	if err != nil {
		return nil, err
	}

	return result, nil
}
