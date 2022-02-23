package database

import (
	"context"
	"os"

	"github.com/timofeev41/feecc-employees-decoder/internal/structs"
	"github.com/timofeev41/feecc-employees-decoder/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

// Returns MongoDB client
func getMongoClient() *mongo.Client {
	utils.Logger.Infof("Connecting to MongoDB")

	uri := os.Getenv("MONGO_CONNECTION_URL")
	if uri == "" {
		utils.Logger.Errorf("Variable MONGO_CONNECTION_URL not set")
		panic("MONGO_CONNECTION_URL not set")
	}

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	utils.ErrHandler(err)

	utils.Logger.Info("Connected to MongoDB", zap.String("uri", uri))
	return client
}

func getEmployeesCollection(client *mongo.Client) *mongo.Collection {
	database := os.Getenv("MONGO_DATABASE_NAME")

	if database == "" {
		utils.Logger.Errorf("Variable MONGO_DATABASE_NAME not set")
		panic("MONGO_DATABASE_NAME not set")
	}

	return client.Database(database).Collection("Employee-data")
}

type MongoWrapper struct {
	employeesCollection *mongo.Collection
}

func GetWrapper() MongoWrapper {
	return MongoWrapper{getEmployeesCollection(getMongoClient())}
}

func (w MongoWrapper) GetEmployeeByRfid(rfid_card_id string) structs.Employee {
	var result structs.Employee

	err := w.employeesCollection.FindOne(
		context.TODO(),
		bson.D{{"rfid_card_id", rfid_card_id}},
	).Decode(&result)
	utils.ErrHandler(err)

	return result
}
