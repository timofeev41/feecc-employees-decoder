package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)


type Employee struct {
	RfidCardId string `json:"rfid_card_id"`
	Name string `json:"name"`
	Position string `json:"position"`
}

func main() {
	client := getMongoClient()

	employeesCollection := getEmployeesCollection(client)

	var result bson.M
	err := employeesCollection.FindOne(context.TODO(), bson.D{{"rfid_card_id", "000915495313945449"}}).Decode(&result)
	errHandler(err)

	fmt.Print(result)
}