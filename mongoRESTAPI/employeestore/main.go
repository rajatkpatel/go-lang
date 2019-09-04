package main

import (
	"log"
	"net/http"

	"github.com/go-lang/mongoRESTAPI/dbConnect"
	"github.com/go-lang/mongoRESTAPI/employeestore/employee"
)

const HOST = "mongodb://127.0.0.1:27017"
const PORT = ":5000"

func main() {
	log.SetPrefix("LOG: ")
	log.Println("Connecting with Mongo DB...")
	session, err := dbConnect.CreateSession(HOST)
	if err != nil {
		log.Fatal("Session not established, Exiting from Main...")
	} else {
		log.Println("Session established")
	}

	defer session.Close()

	router := employee.Route(session)
	log.Fatal(http.ListenAndServe(PORT, router))
}
