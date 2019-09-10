package employee

import (
	"errors"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//DbName is the database name in the mongoDB.
var DbName = "employeeDB"

//CollectionName is the collection name in the databse DbName.
var CollectionName = "employees"

func getEmployees(session *mgo.Session) ([]Employee, error) {
	var employees []Employee
	db := session.DB(DbName).C(CollectionName)
	err := db.Find(bson.M{}).All(&employees)
	if err != nil {
		log.Println("Failed to fetch employees: ", err)
		return nil, err
	}
	return employees, err
}

func createEmployees(session *mgo.Session, employee Employee) (Employee, error) {
	db := session.DB(DbName).C(CollectionName)
	err := db.Insert(employee)
	if err != nil {
		log.Println("Failed to add employee: ", err)
		return Employee{}, err
	}
	return employee, err
}

func getEmployee(session *mgo.Session, id string) (Employee, error) {
	var employee Employee
	db := session.DB(DbName).C(CollectionName)
	err := db.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&employee)
	if err != nil {
		log.Printf("Failed to fetch employee having id %v : %v", id, err)
		return Employee{}, err
	}
	return employee, err
}

func replaceEmployee(session *mgo.Session, id string, employee Employee) (Employee, error) {
	db := session.DB(DbName).C(CollectionName)
	err := db.Update(bson.M{"_id": bson.ObjectIdHex(id)}, employee)
	if err != nil {
		log.Printf("Failed to update employee having id %v : %v", id, err)
		return Employee{}, err
	}
	replacedEmployee, err := getEmployee(session, id)
	if err != nil {
		return Employee{}, err
	}
	return replacedEmployee, err
}

func updateEmployee(session *mgo.Session, id string, employee map[string]interface{}) (Employee, error) {
	db := session.DB(DbName).C(CollectionName)
	previousEmployee, err := getEmployee(session, id)
	if err != nil {
		return Employee{}, err
	}
	err = valueChecker(&previousEmployee, employee)
	if err != nil {
		return Employee{}, err
	}
	err = db.Update(bson.M{"_id": bson.ObjectIdHex(id)}, previousEmployee)
	if err != nil {
		log.Printf("Failed to update employee having id %v : %v", id, err)
		return Employee{}, err
	}
	updatedEmployee, err := getEmployee(session, id)
	if err != nil {
		return Employee{}, err
	}
	return updatedEmployee, err
}

func deleteEmployee(session *mgo.Session, id string) error {
	db := session.DB(DbName).C(CollectionName)
	err := db.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		log.Printf("Failed to delete employee having id %v : %v", id, err)
	}
	return err
}

func valueChecker(previousEmployee *Employee, employee map[string]interface{}) error {
	if _, ok := employee["firstName"]; ok {
		previousEmployee.FirstName = employee["firstName"].(string)
	}
	if _, ok := employee["lastName"]; ok {
		previousEmployee.LastName = employee["lastName"].(string)
	}
	if _, ok := employee["email"]; ok {
		previousEmployee.Email = employee["email"].(string)
	}
	if _, ok := employee["phoneNumber"]; ok {
		previousEmployee.PhoneNumber = int(employee["phoneNumber"].(float64))
	}

	if len(employee) == 0 {
		log.Println("No new Data to Update")
		return errors.New("No new Data to Update")
	}
	return nil
}
