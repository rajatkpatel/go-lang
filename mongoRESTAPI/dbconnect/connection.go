package dbconnect

import (
	"log"

	"gopkg.in/mgo.v2"
)

//CreateSession connect to mongoDB and establish a session.
//It return the session if successful otherwise return error.
func CreateSession(hostName string) (*mgo.Session, error) {
	session, err := mgo.Dial(hostName)
	if err != nil {
		log.Printf("Failed to establish connection to mongodb server: %v", err)
		return nil, err
	}
	return session, nil
}
