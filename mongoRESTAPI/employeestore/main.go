package main

import (
	"log"
	"net/http"

	"github.com/go-lang/mongoRESTAPI/dbconnect"
	"github.com/go-lang/mongoRESTAPI/employeestore/employee"
)

//Host is the Host URL of mongoDB.
var Host = "mongodb://127.0.0.1:27017"

//Port is the exposed Port for REST API.
var Port = ":5000"

var listenAndServe = http.ListenAndServe

func main() {
	log.Println("Connecting with Mongo DB...")
	session, err := dbconnect.CreateSession(Host)
	if err != nil {
		log.Println("Session not established, Exiting from Main...")
		return
	}
	log.Println("Session established")

	defer session.Close()

	router := employee.Route(session)
	log.Fatal(listenAndServe(Port, router))
}

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
