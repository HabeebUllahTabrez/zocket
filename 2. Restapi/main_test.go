package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"my-rest-api/controllers"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// This file consists of a series of tests in which every end point of the api is checked with various test cases
// a user is created, retrieved, edited and deleted in the end of the sequence

// global variable to store the objectId when a new user is created
var objId string

func TestGetAllStudents(t *testing.T) {
	tests := []struct {
		description  string // description of the test case
		method       string
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		{
			description:  "get HTTP status 200",
			method:       "GET",
			route:        "/students",
			expectedCode: 200,
		},
		{
			description:  "get HTTP status 404, when route does not exists",
			method:       "GET",
			route:        "/not-found",
			expectedCode: 404,
		},
	}

	app := fiber.New()

	app.Get("/students", controllers.GetAllStudents)

	for _, test := range tests {
		// Create a new http request with the route from the test case
		req := httptest.NewRequest(test.method, test.route, nil)

		// Perform the request plain with the app,
		// the second argument is a request latency
		// (set to -1 for no latency)
		resp, _ := app.Test(req)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestCreateStudent(t *testing.T) {
	tests := []struct {
		description  string // description of the test case
		method       string
		route        string // route path to test
		jsonStr      []byte
		expectedCode int // expected HTTP status code
	}{
		{
			description:  "get HTTP status 201",
			method:       "POST",
			route:        "/student",
			jsonStr:      []byte(`{"name":"Spiderman","dob":"69 Dec 2002","percentage": 99.99,"address":"8194 NowayhomeCity","description":"Go Developer"}`),
			expectedCode: 201,
		},
		{
			description:  "get HTTP status 400, when invalid parameters given",
			method:       "POST",
			route:        "/student",
			jsonStr:      []byte(`{"name":"Spiderman","dob":"69 Dec 2002","percentage": "99.99","address":"8194 NowayhomeCity","description":"Go Developer"}`),
			expectedCode: 400,
		},
	}

	app := fiber.New()
	app.Post("/student", controllers.CreateStudent)

	for i, test := range tests {
		req := httptest.NewRequest(test.method, test.route, bytes.NewBuffer(test.jsonStr))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		if i == 0 {
			body, _ := ioutil.ReadAll(resp.Body)
			var result map[string]interface{}
			json.Unmarshal([]byte(body), &result)
			objId = fmt.Sprintf("%v", result["data"].(map[string]interface{})["data"].(map[string]interface{})["InsertedID"])
			fmt.Println("The object is created with an object id of", objId)
		}

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestGetStudent(t *testing.T) {
	tests := []struct {
		description  string // description of the test case
		method       string
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		{
			description:  "get HTTP status 200",
			method:       "GET",
			route:        "/student/",
			expectedCode: 200,
		},
		{
			description:  "get HTTP status 500, when invalid object id specified",
			method:       "GET",
			route:        "/student/ksdflj45ljk",
			expectedCode: 500,
		},
	}

	app := fiber.New()
	app.Get("/student/:userId", controllers.GetAStudent)

	for i, test := range tests {
		var completeRoute string
		if i == 0 {
			completeRoute = test.route + objId
		} else {
			completeRoute = test.route
		}
		// fmt.Println(completeRoute)
		req := httptest.NewRequest(test.method, completeRoute, nil)
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestEditStudent(t *testing.T) {
	tests := []struct {
		description  string // description of the test case
		method       string
		route        string // route path to test
		jsonStr      []byte
		expectedCode int // expected HTTP status code
	}{
		{
			description:  "get HTTP status 201",
			method:       "PUT",
			route:        "/student/",
			jsonStr:      []byte(`{"name":"Spiderman XD","dob":"61 Dec 2002","percentage": 99.88,"address":"8194 NowayhomeCity","description":"Go Developer"}`),
			expectedCode: 200,
		},
		{
			description:  "get HTTP status 400, when invalid parameters given",
			method:       "PUT",
			route:        "/student/",
			jsonStr:      []byte(`{"name":"Spiderman XD","dob":"69 Dec 2002","percentage": 99.99,"addddress":"8194 NowayhomeCity","description":"Go Developer"}`),
			expectedCode: 400,
		},
		{
			description:  "get HTTP status 400, when invalid userId given",
			method:       "PUT",
			route:        "/student/3bfdjn3f",
			jsonStr:      []byte(`{"name":"Spiderman XD","dob":"69 Dec 2002","percentage": 99.99,"address":"8194 NowayhomeCity","description":"Go Developer"}`),
			expectedCode: 404,
		},
	}

	app := fiber.New()
	app.Put("/student/:userId", controllers.EditAStudent)

	for i, test := range tests {
		var completeRoute string
		if i == 2 {
			completeRoute = test.route
		} else {
			completeRoute = test.route + objId
		}

		req := httptest.NewRequest(test.method, completeRoute, bytes.NewBuffer(test.jsonStr))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestDeleteStudent(t *testing.T) {
	tests := []struct {
		description  string // description of the test case
		method       string
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		{
			description:  "get HTTP status 200",
			method:       "DELETE",
			route:        "/student/",
			expectedCode: 200,
		},
		{
			description:  "get HTTP status 404, when invalid object id specified",
			method:       "DELETE",
			route:        "/student/ksdflj45ljk",
			expectedCode: 404,
		},
	}

	app := fiber.New()
	app.Delete("/student/:userId", controllers.DeleteAStudent)

	for i, test := range tests {
		var completeRoute string
		if i == 0 {
			completeRoute = test.route + objId
		} else {
			completeRoute = test.route
		}
		// fmt.Println(completeRoute)
		req := httptest.NewRequest(test.method, completeRoute, nil)

		resp, _ := app.Test(req)

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
