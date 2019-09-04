package dbConnect

import (
	"log"

	"gopkg.in/mgo.v2"
)

func CreateSession(hostName string) (*mgo.Session, error) {
	session, err := mgo.Dial(hostName)
	if err != nil {
		log.Printf("Failed to establish connection to mongodb server: %v", err)
		return nil, err
	}
	return session, nil
}
