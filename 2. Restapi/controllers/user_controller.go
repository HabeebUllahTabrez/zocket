// File containing all the handler functions and their logics separated from other code

package controllers

import (
	"context"
	"my-rest-api/configs"
	"my-rest-api/models"
	"my-rest-api/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// variable to the collection
var studentCollection *mongo.Collection = configs.GetCollection(configs.DB, "students")

// special validator variable
var validate = validator.New()

// We are going to validate the request body and check whether the fields/attributes are properly set are not to avoid inconsistency
// We are going test for this in CreateUser and EditUser handlers where we receive json in request body

func GetHome(c *fiber.Ctx) error {
	c.Send([]byte("Welcome to Student Records API!"))
	return nil
}

// function responsible for creating a new user in the database
func CreateStudent(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var student models.Student
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&student); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.StudentResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&student); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.StudentResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	// filling details in the user model
	// the createdAt attribute is set only at the time of user creation
	// it is not manipulated elsewhere (Updating)
	newStudent := models.Student{
		Name:        student.Name,
		DOB:         student.DOB,
		Percentage:  student.Percentage,
		Address:     student.Address,
		Description: student.Description,
		CreatedAt:   time.Now().String(),
	}

	// query to insert a user
	result, err := studentCollection.InsertOne(ctx, newStudent)

	// checking whether an error occured while updating
	// sending an error response to the user if error exists
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.StudentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// sending correct response upon success
	return c.Status(http.StatusCreated).JSON(responses.StudentResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

// function responsible for retrieving a user from the database based on UserID
func GetAStudent(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// extracting userId from params
	userId := c.Params("userId")

	// student model to store fetched data
	var student models.Student

	defer cancel()

	// converting userId from string to ObjectID
	objId, _ := primitive.ObjectIDFromHex(userId)

	// query to fetch an existing users from collection
	err := studentCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&student)

	// checking whether an error occured while fetching
	// sending an error response to the user if error exists
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.StudentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// sending correct response upon success
	return c.Status(http.StatusOK).JSON(responses.StudentResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": student}})
}

// function responsible for editing a user from the database based on UserID
func EditAStudent(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// extracting userId from params
	userId := c.Params("userId")

	// student model to store fetched data
	var student models.Student
	defer cancel()

	// converting userId from string to ObjectID
	objId, _ := primitive.ObjectIDFromHex(userId)

	//validate the request body
	if err := c.BodyParser(&student); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.StudentResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&student); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.StudentResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	// variable which stores the new user attributes after fetching to be updated
	update := bson.M{"name": student.Name, "dob": student.DOB, "percentage": student.Percentage, "address": student.Address, "description": student.Description}

	// query to update a user based on the "_id" value passed
	result, err := studentCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

	// checking whether an error occured while updating
	// sending an error response to the user if error exists
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.StudentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// if updated user count is 0 -> No user updated -> Invalid userId
	// sending error response to the user
	if result.MatchedCount == 0 {
		return c.Status(http.StatusNotFound).JSON(
			responses.StudentResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "User with specified ID not found!"}},
		)
	}

	//get updated user details
	var updatedStudent models.Student

	// After updating the user, fetching back the same user and returning it to the user as a response
	// this code is similar to the fetching a single user code
	if result.MatchedCount == 1 {
		err := studentCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedStudent)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.StudentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	// sending correct response upon success
	return c.Status(http.StatusOK).JSON(responses.StudentResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedStudent}})
}

// function responsible for deleting a user from the database based on UserID
func DeleteAStudent(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// extracting userId from params
	userId := c.Params("userId")
	defer cancel()

	// converting userId from string to ObjectID
	objId, _ := primitive.ObjectIDFromHex(userId)

	// query to delete o user based on the "_id" value passed
	result, err := studentCollection.DeleteOne(ctx, bson.M{"_id": objId})

	// checking whether an error occured while deleting
	// sending an error response to the user if error exists
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.StudentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// if deleted users are less than 1 -> No user deleted -> Invalid userId
	// sending error response to the user
	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			responses.StudentResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "User with specified ID not found!"}},
		)
	}

	// sending correct response upon success
	return c.Status(http.StatusOK).JSON(
		responses.StudentResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "User successfully deleted!"}},
	)
}

// function responsible for retrieving all the user from the database
func GetAllStudents(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// slice to store all the retrieved students
	var students []bson.M
	defer cancel()

	// query to fetch all existing users from collection
	results, err := studentCollection.Find(ctx, bson.M{})

	// checking whether an error occured while fetching
	// sending an error response to the user if error exists
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.StudentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	defer results.Close(ctx)

	// reading from the db in an optimal way
	// fetching an individual user using a curson and appending it to the users slice
	for results.Next(ctx) {
		var singleStudent bson.M

		// sending back error response if error exists
		if err = results.Decode(&singleStudent); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.StudentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		students = append(students, singleStudent)
	}

	// sending correct response upon success
	return c.Status(http.StatusOK).JSON(
		responses.StudentResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": students}},
	)
}
