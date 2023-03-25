package backend_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"example.com/backend"
)

var a backend.App

const tableProductCreationQuery = `CREATE TABLE IF NOT EXISTS products 
(
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	productCode VARCHAR(25) NOT NULL,
	name VARCHAR(256) NOT NULL,
	inventory INTEGER NOT NULL,
	price INTEGER NOT NULL,
	status VARCHAR(64) NOT NULL
)`

func TestMain(m *testing.M) {
	a = backend.App{}
	a.Initialize()
	ensureTableExists()
	code := m.Run()

	clearProductTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableProductCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearProductTable() {
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("DELETE FROM sqlite_sequence WHERE name = 'products'")
}

func TestGetNonExistentProduct(t *testing.T) {
	clearProductTable()

	req, _ := http.NewRequest("GET", "/product/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusInternalServerError, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["error"] != "sql: no rows in result set" {
		t.Errorf("Expected the 'error' key of the response to be set to 'sql: no rows in result set'")
	}
}

func TestCreateProduct(t *testing.T) {
	clearProductTable()

	payload := []byte(`{"productCode":"TEST12345", "name":"ProductTest", "inventory":1, "price":1, "status": "testing"}`)

	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["productCode"] != "TEST12345" {
		t.Errorf("Expected the 'productCode' key of the response to be set to 'TEST12345'. Got %v", m["productCode"])
	}

	if m["name"] != "ProductTest" {
		t.Errorf("Expected the 'name' key of the response to be set to 'ProductTest'. Got %v", m["name"])
	}

	if m["inventory"] != 1.0 {
		t.Errorf("Expected the 'inventory' key of the response to be set to '1'. Got %v", m["inventory"])
	}

	if m["price"] != 1.0 {
		t.Errorf("Expected the 'price' key of the response to be set to '1'. Got %v", m["price"])
	}

	if m["status"] != "testing" {
		t.Errorf("Expected the 'status' key of the response to be set to 'testing'. Got %v", m["status"])
	}

	if m["id"] != 1.0 {
		t.Errorf("Expected the 'id' key of the response to be set to '1'. Got %v", m["id"])
	}
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}

}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()

	a.Router.ServeHTTP(rr, req)

	return rr
}
