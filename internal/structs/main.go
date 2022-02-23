package structs

type Employee struct {
	ID         string `bson:"_id", json:"_id"`
	RfidCardId string `bson:"rfid_card_id", json:"rfid_card_id"`
	Name       string `bson:"name", json:"name"`
	Position   string `bson:"position", json:"position"`
}