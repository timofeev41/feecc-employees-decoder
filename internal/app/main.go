package app

import (
	"fmt"

	"github.com/timofeev41/feecc-employees-decoder/internal/database"
	"github.com/timofeev41/feecc-employees-decoder/internal/structs"
	"go.mongodb.org/mongo-driver/bson"
)

var wrapper database.MongoWrapper = database.GetWrapper()


func Run() {
	result := wrapper.GetEmployeeByRfid("id")

	fmt.Println(result)

	var employee = structs.Employee{}
	bsonBytes, _ := bson.Marshal(result)
	bson.Unmarshal(bsonBytes, &employee)

	fmt.Println(employee)
}
