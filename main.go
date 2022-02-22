package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type Employee struct {
	ID         string `bson:"_id", json:"_id"`
	RfidCardId string `bson:"rfid_card_id", json:"rfid_card_id"`
	Name       string `bson:"name", json:"name"`
	Position   string `bson:"position", json:"position"`
}

func main() {
	result := getEmployeeByRfid("smth")

	fmt.Println(result)

	var employee = Employee{}
	bsonBytes, _ := bson.Marshal(result)
	bson.Unmarshal(bsonBytes, &employee)

	fmt.Println(employee)
}
