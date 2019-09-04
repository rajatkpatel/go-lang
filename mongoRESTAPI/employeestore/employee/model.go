package employee

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const DBNAME = "employeeDB"
const COLLECTIONNAME = "employees"

func getEmployees(session *mgo.Session) ([]Employee, error) {
	var employees []Employee
	db := session.DB(DBNAME).C(COLLECTIONNAME)
	err := db.Find(bson.M{}).All(&employees)
	if err != nil {
		log.Println("Failed to fetch employees: ", err)
		return nil, err
	}
	return employees, err
}

func createEmployees(session *mgo.Session, employee Employee) (Employee, error) {
	db := session.DB(DBNAME).C(COLLECTIONNAME)
	err := db.Insert(employee)
	if err != nil {
		log.Println("Failed to add employee: ", err)
		return Employee{}, err
	}
	return employee, err
}
