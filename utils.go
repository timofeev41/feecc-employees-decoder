package main

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func errHandler(err error) {
	if err != nil {
		sugar.Errorf("Error %v", err)
		panic(err)
	}
}

// Returns MongoDB client
func getMongoClient() *mongo.Client  {
	sugar.Infof("Connecting to MongoDB")

	uri := os.Getenv("MONGO_CONNECTION_URL")
	if uri == "" {
		sugar.Errorf("Variable MONGO_CONNECTION_URL not set")
		panic("MONGO_CONNECTION_URL not set")
	}

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	errHandler(err)

	sugar.Info("Connected to MongoDB", zap.String("uri", uri))
	return client
}

func getEmployeesCollection(client *mongo.Client) *mongo.Collection{
	database := os.Getenv("MONGO_DATABASE_NAME")
	if database == "" {
		sugar.Errorf("Variable MONGO_DATABASE_NAME not set")
		panic("MONGO_DATABASE_NAME not set")
	}

	return client.Database(database).Collection("Employee-data")
}

// Returns zap sugared logger to handle log messages
func getLogger() *zap.Logger {
	logger, _ := zap.NewProduction()

	defer logger.Sync()
	return logger
}

var sugar *zap.SugaredLogger = getLogger().Sugar()
