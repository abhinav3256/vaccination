package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Nurse struct {
	Email       string `json:"email" binding:"required"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	DateOfBirth string `json:"date_of_birth" binding:"required"`
	Sex         string `json:"sex" binding:"required"`
	City        string `json:"city" binding:"required"`
}

type Persons struct {
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	Sex         string `json:"sex"`
}

func main() {
	CreateDBConnection()
	fmt.Println("test")
	r := gin.Default()
	setupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setupRoutes(r *gin.Engine) {

	r.GET("all_nurses", allNurses)
	r.GET("all_persons", allPersons)
	r.GET("vaccination_nurses", vaccinationNurses)
	r.GET("vaccination_person", vaccinationPersons)

	r.POST("nurse/add", addNurseHandler)

}
