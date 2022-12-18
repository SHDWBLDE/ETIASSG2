package controllers

import (
	"fmt"
	"net/http"
	"ride-sharing/models"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PassengerInput struct {
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	MobileNumber string `json:"mobile_number" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Password     string `json:"password" binding:"required"`
}

type DriverInput struct {
	FirstName            string `gorm:"size:128;not null;unique" json:"first_name"`
	LastName             string `gorm:"size:128;not null;unique" json:"last_name"`
	MobileNumber         string `gorm:"size:20;not null;unique" json:"mobile_number"`
	Email                string `gorm:"not null;unique" json:"email"`
	IdentificationNumber string `gorm:"not null;unique" json:"identification_number"`
	LicenceNumber        string `gorm:"not null;unique" json:"licence_number"`
	Password             string `json:"password" binding:"required"`
}

type PassengerLoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {

	var input PassengerLoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.Passenger{}

	user.Email = input.Email
	user.Password = input.Password

	token, err := models.CheckLogin(user.Email, user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func RegisterPassenger(c *gin.Context) {
	var input PassengerInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.Passenger{}

	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.MobileNumber = input.MobileNumber
	user.Email = input.Email
	user.Password = input.Password

	user.SaveUser()

	fmt.Println(user.ID)

	c.JSON(http.StatusOK, gin.H{"data": "registration successful!"})
}

func RegisterDriver(c *gin.Context) {
	var input DriverInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.Driver{}

	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.MobileNumber = input.MobileNumber
	user.Email = input.Email
	user.Password = input.Password
	user.LicenceNumber = input.LicenceNumber
	user.IdentificationNumber = input.IdentificationNumber

	// TODO: save the driver registration

	fmt.Println(user.ID)

	c.JSON(http.StatusOK, gin.H{"data": "registration successful!"})
}
