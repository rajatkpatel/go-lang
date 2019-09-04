package employee

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func Route(s *mgo.Session) *mux.Router {
	session = s
	//initialize router
	router := mux.NewRouter()
	//endpoints
	router.HandleFunc("/employees", GetEmployees).Methods("GET")
	// router.HandleFunc("/employees/{id}", GetEmployee).Methods("GET")
	router.HandleFunc("/employees", CreateEmployees).Methods("POST")
	// router.HandleFunc("/employees/{id}", UpdateEmployee).Methods("PUT")
	// router.HandleFunc("/employees/{id}", DeleteEmployee).Methods("DELETE")

	return router
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	s := session.Copy()
	defer s.Close()
	employees, _ := getEmployees(session)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)

}

func CreateEmployees(w http.ResponseWriter, r *http.Request) {
	s := session.Copy()
	defer s.Close()
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	json.NewDecoder(r.Body).Decode(&employee)
	employee, _ = createEmployees(session, employee)
	json.NewEncoder(w).Encode(employee)

}
