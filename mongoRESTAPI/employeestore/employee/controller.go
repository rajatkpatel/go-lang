package employee

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

//Route is the router function for employee.
func Route(s *mgo.Session) *mux.Router {
	session = s
	//initialize router
	router := mux.NewRouter()
	//endpoints
	router.HandleFunc("/employees", GetEmployees).Methods("GET")
	router.HandleFunc("/employees", CreateEmployees).Methods("POST")
	router.HandleFunc("/employee/{id}", GetEmployee).Methods("GET")
	router.HandleFunc("/employee/{id}", ReplaceEmployee).Methods("PUT")
	router.HandleFunc("/employee/{id}", UpdateEmployee).Methods("PATCH")
	router.HandleFunc("/employee/{id}", DeleteEmployee).Methods("DELETE")

	return router
}

//GetEmployees displays the list of employees details.
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	s := session.Copy()
	defer s.Close()
	employees, err := getEmployees(session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

//CreateEmployees create a employee.
func CreateEmployees(w http.ResponseWriter, r *http.Request) {
	s := session.Copy()
	defer s.Close()
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	employee, err = createEmployees(session, employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(employee)
}

//GetEmployee displays the employee details of the queried employee id.
func GetEmployee(w http.ResponseWriter, r *http.Request) {
	s := session.Copy()
	defer s.Close()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	employee, err := getEmployee(session, id)
	if err != nil && err == mgo.ErrNotFound {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(employee)
}

//ReplaceEmployee replace the employee details whose id is provided with new details.
func ReplaceEmployee(w http.ResponseWriter, r *http.Request) {
	s := session.Copy()
	defer s.Close()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	var employee Employee
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	employee, err = replaceEmployee(session, id, employee)
	if err != nil && err == mgo.ErrNotFound {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(employee)
}

//UpdateEmployee update the provided employee details whose id is provided with new details.
//Details which are not provided keep intact.
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	s := session.Copy()
	defer s.Close()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	var employee map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedEmployee, err := updateEmployee(session, id, employee)
	if err != nil && err == mgo.ErrNotFound {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedEmployee)
}

//DeleteEmployee remove the employee from the database whose id is provided.
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	s := session.Copy()
	defer s.Close()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	err := deleteEmployee(session, id)
	if err != nil && err == mgo.ErrNotFound {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(id)
}
