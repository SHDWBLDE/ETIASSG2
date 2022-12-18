package models

import (
	"fmt"
	"html"
	"strings"

	"ride-sharing/internal/token"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
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
	FirstName            string `gorm:"size:128;not null;unique" json:"first_name"`
	LastName             string `gorm:"size:128;not null;unique" json:"last_name"`
	MobileNumber         string `gorm:"size:20;not null;unique" json:"mobile_number"`
	Email                string `gorm:"not null;unique" json:"email"`
	IdentificationNumber string `gorm:"not null;unique" json:"identification_number"`
	LicenceNumber        string `gorm:"not null;unique" json:"licence_number"`
	Password             string `gorm:"size:255;not null" json:"password"`
}

var DB = OpenDb()

func OpenDb() *gorm.DB {
	DB, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&Passenger{})
	DB.AutoMigrate(&Driver{})
	return DB
}

func (u *Passenger) SaveUser() (*Passenger, error) {
	var err error
	fmt.Println(u)
	err = DB.Create(&u).Error
	if err != nil {
		return &Passenger{}, err
	}
	return u, nil
}

func (u *Passenger) BeforeSave(*gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	return nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func CheckLogin(username string, password string) (string, error) {
	var err error
	u := Passenger{}

	err = DB.Model(Passenger{}).Where("email = ?", u.Email).Take(&u).Error

	if err != nil {
		return "", err
	}
	err = VerifyPassword(password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	token, err := token.GenerateToken(u.ID)
	if err != nil {
		return "", err
	}

	return token, nil

}
