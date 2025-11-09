package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const BaseURL = "http://localhost:8083/api/v1"

// HTTPResponse wraps the response from HTTP requests
type HTTPResponse struct {
	StatusCode int
	Body       map[string]interface{}
	RawBody    string
}

// GET makes a GET request to the API
func GET(endpoint string, token ...string) *HTTPResponse {
	return makeRequest("GET", endpoint, nil, token...)
}

// POST makes a POST request to the API
func POST(endpoint string, body interface{}, token ...string) *HTTPResponse {
	return makeRequest("POST", endpoint, body, token...)
}

// PUT makes a PUT request to the API
func PUT(endpoint string, body interface{}, token ...string) *HTTPResponse {
	return makeRequest("PUT", endpoint, body, token...)
}

// DELETE makes a DELETE request to the API
func DELETE(endpoint string, token ...string) *HTTPResponse {
	return makeRequest("DELETE", endpoint, nil, token...)
}

// makeRequest is the core function that makes HTTP requests
func makeRequest(method, endpoint string, body interface{}, token ...string) *HTTPResponse {
	var reqBody io.Reader

	if body != nil {
		jsonData, _ := json.Marshal(body)
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, _ := http.NewRequest(method, BaseURL+endpoint, reqBody)
	req.Header.Set("Content-Type", "application/json")

	if len(token) > 0 && token[0] != "" {
		req.Header.Set("Authorization", "Bearer "+token[0])
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic("HTTP request failed: " + err.Error())
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	var bodyMap map[string]interface{}
	json.Unmarshal(bodyBytes, &bodyMap)

	return &HTTPResponse{
		StatusCode: resp.StatusCode,
		Body:       bodyMap,
		RawBody:    string(bodyBytes),
	}
}

// GenerateTestJWT creates a valid JWT token for testing
func GenerateTestJWT(userID uint, username, email string) string {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"email":    email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("test-secret-key-for-testing-only"))

	return tokenString
}
