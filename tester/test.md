# Test endpoint

### Command
```sh
go test
```

### Endpoint test code
```go
func TestGetProductByIDHandler(t *testing.T) {
	backendApp = backend.App{}
	backendApp.Initialize()
	CheckTablesExists()

	// Create a request with a specific product ID
	req, err := http.NewRequest("GET", "localhost:9003/products/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Create a handler and serve the request to the ResponseRecorder
	backendApp.Router.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Check if the JSON response has the "id" attribute
	var responseMap map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
	if err != nil {
		t.Errorf("Error parsing JSON response: %v", err)
	}

	if _, ok := responseMap["id"]; !ok {
		t.Errorf("JSON response does not contain the 'id' attribute")
	}
}
```