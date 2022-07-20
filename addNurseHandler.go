package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addNurseHandler(c *gin.Context) {
	reqBody := Nurse{}
	err := c.Bind(&reqBody)

	if err != nil {
		res := gin.H{

			"Error": err.Error(),
		}

		c.JSON(http.StatusBadRequest, res)
		return
	}
	if !isCityExist(reqBody.City) {
		res := gin.H{

			"Error": "City Does not Exist",
		}

		c.JSON(http.StatusBadRequest, res)
		return
	}

	result, _ := insertNurse(reqBody)

	if result == false {
		res := gin.H{

			"Error": "Something Went Wrong",
		}

		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := gin.H{
		"message": "Successfully Inserted",
		"status":  200,
	}
	c.JSON(http.StatusCreated, res)
}

func insertNurse(reqbody Nurse) (bool, string) {
	var result = true
	var err_responce = ""

	sqlStatement := `
INSERT INTO persons(first_name, last_name,email, date_of_birth, sex,city)
VALUES ($1, $2, $3, $4,$5,$6)`
	_, err2 := DB.Exec(sqlStatement, reqbody.FirstName, reqbody.LastName, reqbody.Email, reqbody.DateOfBirth, reqbody.Sex, reqbody.City)
	fmt.Println(err2)
	if err2 != nil {

		err_responce = "Something went wrong"
		result = false
		return result, err_responce
	}

	sqlStatement2 := `
INSERT INTO nurses(email) VALUES ($1)`
	_, err3 := DB.Exec(sqlStatement2, reqbody.Email)
	fmt.Println(err3)
	//fmt.Println(user)
	if err3 != nil {

		err_responce = "Something went wrong"
		result = false
		return result, err_responce
	}
	return result, err_responce
}
