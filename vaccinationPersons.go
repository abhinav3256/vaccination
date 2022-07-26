package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func vaccinationPersons(c *gin.Context) {
	data := getvaccinationPersons()

	res := gin.H{
		"nurses": data,
	}
	//c.Writer.Header().Set("Content-Type", "application/json")

	c.JSON(http.StatusBadRequest, res)
	return

}

func getvaccinationPersons() []Persons {
	Data := []Persons{}

	SQL := "SELECT persons.email,persons.first_name,persons.last_name,persons.date_of_birth,persons.sex from persons, vaccinations WHERE persons.email=vaccinations.recipient"

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
