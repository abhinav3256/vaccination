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

	SQL := "SELECT persons.email from persons, vaccinations WHERE persons.email=vaccinations.recipient"

	rows, err := DB.Query(SQL)

	if err != nil {
		log.Println("Failed to execute query: ", err)
		return Data
	}
	defer rows.Close()
	person := Persons{}

	for rows.Next() {
		rows.Scan(&person.Email)

		Data = append(Data, person)

	}

	fmt.Println(Data)
	return Data

}
