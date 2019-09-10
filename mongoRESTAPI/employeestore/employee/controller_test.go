package employee

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-lang/mongoRESTAPI/dbconnect"
)

var hostName = "mongodb://127.0.0.1:27017"

func init() {
	testSession, err := dbconnect.CreateSession(hostName)
	if err != nil {
		log.Println("Session not established: ", err)
		return
	}
	Route(testSession)
}

func TestRoute(t *testing.T) {
	testSession, err := dbconnect.CreateSession(hostName)
	if err != nil {
		t.Error("Session not established: ", err)
		return
	}
	Route(testSession)
}

func TestGetEmployees(t *testing.T) {
	req, err := http.NewRequest("GET", "/employees", nil)
	if err != nil {
		t.Error("Error occured while testing: ", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetEmployees)
	handler.ServeHTTP(rr, req)
	assert.Equalf(t, http.StatusOK, rr.Code, "Status code expected: %v got %v", http.StatusOK, rr.Code)
}

var createEmployeesInput = []struct {
	inputBody      []byte
	expectedStatus int
}{
	{[]byte(`{"firstName": "valid", "lastName": "test","email": "test@a.com","phoneNumber": 123}`), 200},
	{[]byte(`{"firstName": "invalid", "lastName": "test","email": "test@a.com","phoneNumber": "invalid"}`), 400},
}

func TestCreateEmployees(t *testing.T) {
	for _, item := range createEmployeesInput {
		req, err := http.NewRequest("POST", "/employees", bytes.NewReader(item.inputBody))
		if err != nil {
			t.Error("Error occured while testing: ", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateEmployees)
		handler.ServeHTTP(rr, req)
		assert.Equalf(t, item.expectedStatus, rr.Code, "Status code expected: %v got %v", item.expectedStatus, rr.Code)
	}
}

// type getEmployeeInput struct {
// 	inputURL       string
// 	expectedStatus int
// }

// func TestGetEmployee(t *testing.T) {

// 	input := []getEmployeeInput{{"/employee/5d774bdd8de65cb8e1507513", 200}}

// 	for _, item := range input {
// 		req, err := http.NewRequest("GET", item.inputURL, nil)
// 		if err != nil {
// 			t.Error("Error occured while testing: ", err)
// 		}
// 		rr := httptest.NewRecorder()
// 		handler := http.HandlerFunc(GetEmployee)
// 		handler.ServeHTTP(rr, req)
// 		assert.Equalf(t, item.expectedStatus, rr.Code, "Status code expected: %v got %v", item.expectedStatus, rr.Code)
// 	}
// }

// mux := http.NewServeMux()
// 	mux.HandleFunc("/", GetEmployee)
// 	srv := httptest.NewServer(mux)
// 	client := &http.Client{
// 		Timeout: 20 * time.Second,
// 	}
// 	r, _ := http.NewRequest("GET", srv.URL+"/employee/5d774bdd8de65cb8e1507513", nil)
// 	res, _ := client.Do(r)
// 	fmt.Println(res.StatusCode)
