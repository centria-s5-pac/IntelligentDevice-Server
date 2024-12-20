package sensor_test

import (
	"helios/internal/api/handlers/sensor"
	service "helios/internal/api/service/data"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPutInvalidRequestBody(t *testing.T) {

	req, err := http.NewRequest("PUT", "/sensor", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Body = io.NopCloser(strings.NewReader(`Plain text, not JSON`))
	rr := httptest.NewRecorder()
	sensor.PutHandler(rr, req, log.Default(), &service.MockDataServiceSuccessful{})

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	expected := `{"error": "Invalid request data. Please check your input."}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestPutHandlerError(t *testing.T) {

	req, err := http.NewRequest("PUT", "/sensor", strings.NewReader(`{"id":1,"type":2,"value":1,"timestamp":"2021-01-01T00:00:00Z"}`))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	sensor.PutHandler(rr, req, log.Default(), &service.MockDataServiceError{})

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	expected := `{"error": "Error updating data."}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestPutHandlerSuccess(t *testing.T) {

	req, err := http.NewRequest("PUT", "/sensor", strings.NewReader(`{"id":1,"type":2,"value":1.1,"timestamp":"2021-01-01T00:00:00Z"}`))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	sensor.PutHandler(rr, req, log.Default(), &service.MockDataServiceSuccessful{})

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"id":1,"type":2,"value":1.1,"timestamp":"2021-01-01T00:00:00Z"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
