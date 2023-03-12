# Student Record REST API

A REST API developed in Go using Fiber which is an Express-inspired framework for Golang.

This API operates on a Student Model/Document and performs all the CRUD operations on it.

The data is stored and retrieved from MongoDB Atlas, a NoSQL Database.

## How an individual entry stored in the database looks like?

```json
    {
        "_id": "628e5ac214322b31dac15601",  // mongoDb objectID
        "name": "John Doe"
        "dob": "1 Jan 1999",
        "percentage": 99.99,
        "address": "8194 Euclid City",
        "description": "Backend Developer",
        "createdAt": "2022-05-25 22:05:14.426684 +0530 IST m=+55.164231301",
    }
```

## Endpoints Description

### Get All Students

This endpoint fethes all the Student documents from the database with their IDS.

```
    URL - *http://localhost:6000/students*
    Method - GET
```

### Get Student By ID

This endpoint fethes a unique Student document from the database with the <User-ID> passed as a request parameter.

```
    URL - *http://localhost:6000/student/<User-ID>*
    Method - GET
```

### Create a new Student

This endpoint creates and publishes a unique Student document to the database.
The attribute "createdAt" is set by the request handler on the time of creation of this new Student.

```
    URL - *http://localhost:6000/student*
    Method - POST
    Request Header - (Content-Type : application/json)
    Request Body -

    {
        "name": "John Doe"
        "dob": "1 Jan 1999",
        "percentage": 99.99,
        "address": "8194 Euclid City",
        "description": "Backend Developer",
        "createdAt": "2022-05-25 22:05:14.426684 +0530 IST m=+55.164231301",
    }

```

### Update Student

This endpoint updates a unique Student document from the database with the <User-ID> passed as a request parameter.

```
    URL - *http://localhost:6000/student/<User-ID>*
    Method - PUT
    Request Header - (Content-Type : application/json)
    Request Body -

    {
        "name": "John Doe"
        "dob": "1 Jan 1999",
        "address": "8194 Euclid City",
        "description": "Backend Developer",
    }
```

### Delete Student

This endpoint deletes a unique Student document from the database with the <User-ID> passed as a request parameter.

```
    URL - *http://localhost:6000/Student/<User-ID>*
    Method - DELETE
```

## Statup Description

To run this project, you must have a MongoDB cluster/database server running and a URI pointing it.

Create a `.env` file in the root directory and paste your uri in it.

```
    MONGOURI=<YOUR MONGODB URI HERE>
```

After doing this, your application would be ready to take off!


## Project Startup

Command to start the server

1. `go run main.go`

## Test Driven Development Description

Command to run all the unit test cases. 
(All the test cases are interlinked and hence some test cases cannot be run independently)

1. `go test -v`

## Hope everything works. Thank you.
