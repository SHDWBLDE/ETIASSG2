package models

import (
	"time"

	"gorm.io/gorm"
)

type Rides struct {
	gorm.Model
	DriverName    string    `gorm:"not null;unique" json:"driver_name"`
	MobileNumber  string    `gorm:"size:20;not null;unique" json:"mobile_number"`
	Status        string    `gorm:"not null" json:"status"`
	StartTripTime time.Time `gorm:"not null" json:"start_trip_time"`
	EndTripTime   time.Time `gorm:"not null" json:"end_trip_time"`
	Details       string    `gorm:"not null" json:"details"`
}
