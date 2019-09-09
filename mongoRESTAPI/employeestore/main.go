package main

import (
	"log"
	"net/http"

	"github.com/go-lang/mongoRESTAPI/dbconnect"
	"github.com/go-lang/mongoRESTAPI/employeestore/employee"
)

//HOST is the host URL of mongoDB.
const HOST = "mongodb://127.0.0.1:27017"

//PORT is the exposed port for REST API.
const PORT = ":5000"

func main() {
	log.Println("Connecting with Mongo DB...")
	session, err := dbconnect.CreateSession(HOST)
	if err != nil {
		log.Println("Session not established, Exiting from Main...")
		return
	}
	log.Println("Session established")

	defer session.Close()

	router := employee.Route(session)
	log.Fatal(http.ListenAndServe(PORT, router))
}

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
