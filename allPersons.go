package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func allPersons(c *gin.Context) {
	data := getPersons()

	res := gin.H{
		"nurses": data,
	}
	//c.Writer.Header().Set("Content-Type", "application/json")

	c.JSON(http.StatusBadRequest, res)
	return

}

func getPersons() []Persons {
	Data := []Persons{}

	SQL := `SELECT email,first_name,last_name,date_of_birth,sex FROM "persons"`

	rows, err := DB.Query(SQL)

	if err != nil {
		log.Println("Failed to execute query: ", err)
		return Data
	}
	defer rows.Close()
	person := Persons{}

	for rows.Next() {
		rows.Scan(&person.Email, &person.FirstName, &person.LastName, &person.DateOfBirth, &person.Sex)

		Data = append(Data, person)

	}

	fmt.Println(Data)
	return Data

}
