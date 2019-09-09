package employee

import (
	"gopkg.in/mgo.v2/bson"
)

//Employee is the structure schema for the database
type Employee struct {
	ID          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	FirstName   string        `json:"firstName" bson:"first_name"`
	LastName    string        `json:"lastName" bson:"last_name"`
	Email       string        `json:"email" bson:"email"`
	PhoneNumber int           `json:"phoneNumber" bson:"phone_number"`
}
