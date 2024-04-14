package goappointments

import (
	"time"

	"gorm.io/gorm"
)

// Appointment struct represents a booking appointment
type Appointment struct {
	gorm.Model
	UserID    uint
	ServiceID uint
	Service   Service
	StartTime time.Time
	EndTime   time.Time
	Status    string // e.g. "Pending", "Confirmed", "Cancelled"
}

// Service struct represents a service offered by the booking system
type Service struct {
	gorm.Model
	Name               string
	Description        string
	Duration           time.Duration
	Price              float64
	Appointments       []Appointment
	ServiceDayConfigId []ServiceDayConfig
}

// Service struct represents a service offered by the booking system
type ServiceDayConfig struct {
	gorm.Model
	Date        time.Time
	Description string
	StartTime   time.Time
	EndTime     time.Time
	ServiceID   uint
}

// Location struct represents a location where services are offered
type Location struct {
	gorm.Model
	Name      string
	Address   string
	Latitude  float64
	Longitude float64
	Services  []Service
}
