package main

import (
	"net/http"
	"ride-sharing/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Passenger struct {
	ID           uint   `gorm:"primaryKey;not null;unique;autoIncrement" json:"id"`
	FirstName    string `gorm:"size:128;not null;unique" json:"first_name"`
	LastName     string `gorm:"size:128;not null;unique" json:"last_name"`
	MobileNumber string `gorm:"size:20;not null;unique" json:"mobile_number"`
	Email        string `gorm:"not null;unique" json:"email"`
	Password     string `gorm:"size:255;not null" json:"password"`
}

type Driver struct {
	gorm.Model
	FirstName            string
	LastName             string
	MobileNumber         string
	Email                string
	IdentificationNumber string
	LicenceNumber        string
}

type Config struct {
	port string
}

type application struct {
	config *Config
}

func main() {

	config := &Config{port: "8080"}
	app := &application{config: config}

	r := gin.New()

	register := r.Group("/api/register")

	// route 127.0.0.1:8080/api/register/passenger
	register.POST("/passenger", controllers.RegisterPassenger)
	// route 127.0.0.1:8080/api/register/driver
	register.POST("/driver", controllers.RegisterDriver)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":" + app.config.port)
}
