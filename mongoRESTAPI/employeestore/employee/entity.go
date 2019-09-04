package employee

import (
	"gopkg.in/mgo.v2/bson"
)

type Employee struct {
	Eid         bson.ObjectId `bson: "_id,omitempty"`
	FirstName   string        `bson: "firstName"`
	LastName    string        `bson: "lastName"`
	Email       string        `bson: "email"`
	PhoneNumber int           `bson: "phoneNumber"`
}
