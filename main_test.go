package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

const serverBase = "http://localhost:8080"

func TestSimpleHelloCase(t *testing.T) {
	response, err := http.Get(serverBase + "/hello")
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "Hello from HTTP Mock Service", string(body))
}

func TestSimpleHelloCaseHttpMethodGetOnly(t *testing.T) {
	response, err := http.Get(serverBase + "/hello_get_only")
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "Hello GET only from HTTP Mock Service", string(body))
}

func TestHelloJson(t *testing.T) {
	response, err := http.Get(serverBase + "/hello_json")
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, response.Header.Get("Content-Type"), "application/json")
	assert.Equal(t, response.Header.Get("Cache-Control"), "no-cache")
	assert.Equal(t, response.Header.Get("X-Powered-By"), "http-mock-service")
	assert.Equal(t, `{
    "product_id": "12345",
    "name": "Wireless Bluetooth Headphones",
    "brand": "TechSound",
    "price": 79.99,
}
`, string(body))
}

func TestHelloXml(t *testing.T) {
	response, err := http.Get(serverBase + "/hello_xml")
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, response.Header.Get("Content-Type"), "application/xml")
	assert.Equal(t, response.Header.Get("Cache-Control"), "no-cache")
	assert.Equal(t, response.Header.Get("X-Powered-By"), "http-mock-service")
	assert.Equal(t, `<product>
  <product_id>12345</product_id>
  <name>Wireless Bluetooth Headphones</name>
  <brand>TechSound</brand>
  <price>79.99</price>
</product>
`, string(body))
}

func TestHelloOptions(t *testing.T) {
	request, err := http.NewRequest("OPTIONS", serverBase+"/hello_options", nil)
	if err != nil {
		t.Error(err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 204, response.StatusCode)
	assert.Equal(t, response.Header.Get("Allow"), "OPTIONS, GET, HEAD, POST")
	assert.Equal(t, response.Header.Get("Cache-Control"), "no-cache")
	assert.Equal(t, response.Header.Get("X-Powered-By"), "http-mock-service")
	assert.Empty(t, string(body))
}

func TestNonExistingEndpoint(t *testing.T) {
	response, err := http.Get(serverBase + "/hello_non_existing_endpoint")
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 501, response.StatusCode)
	assert.Equal(t, "No handler found for request.", string(body))
}

func TestResponseToAnyHttpMethodIfHttpMethodNotDefined(t *testing.T) {
	request, err := http.NewRequest("DELETE", serverBase+"/hello_json", nil)
	if err != nil {
		t.Error(err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, `{
    "product_id": "12345",
    "name": "Wireless Bluetooth Headphones",
    "brand": "TechSound",
    "price": 79.99,
}
`, string(body))
}

func TestDoesNotRespondToDifferentHttpMethodThenDefined(t *testing.T) {
	request, err := http.NewRequest("DELETE", serverBase+"/hello_get_only", nil)
	if err != nil {
		t.Error(err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 501, response.StatusCode)
	assert.Equal(t, "No handler found for request.", string(body))
}

func setup() {
	createServer()
	go startServer()
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
}
