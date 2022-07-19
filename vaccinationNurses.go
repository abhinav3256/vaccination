package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func vaccinationNurses(c *gin.Context) {
	data := getvaccinationNurses()

	res := gin.H{
		"nurses": data,
	}
	//c.Writer.Header().Set("Content-Type", "application/json")

	c.JSON(http.StatusBadRequest, res)
	return

}

func getvaccinationNurses() []Nurse {
	Data := []Nurse{}

	SQL := "select nurses.email from nurses, vaccinations where nurses.email=vaccinations.recipient"

	rows, err := DB.Query(SQL)

	if err != nil {
		log.Println("Failed to execute query: ", err)
		return Data
	}
	defer rows.Close()
	nurse := Nurse{}

	for rows.Next() {
		rows.Scan(&nurse.Email)

		Data = append(Data, nurse)

	}

	fmt.Println(Data)
	return Data

}
