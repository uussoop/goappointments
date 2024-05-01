package goappointments

import (
	"time"

	"gorm.io/gorm"
)

type StatusType int16

func (s *StatusType) String() string {
	switch *s {
	case 0:
		return "Pending"
	case 1:
		return "Confirmed"
	case 2:
		return "Cancelled"
	case 3:
		return "Done"
	default:
		return "Pending"
	}
}

const (
	Cancelled = iota
	Pending
	Confirmed
	Done
)

// Appointment struct represents a booking appointment
type Appointment struct {
	gorm.Model
	UserID    uint
	ServiceID uint
	Service   Service
	StartTime time.Time
	EndTime   time.Time
	Status    StatusType // e.g. "Pending", "Confirmed", "Cancelled"
}

// Service struct represents a service offered by the booking system
type Service struct {
	gorm.Model
	Name               string `gorm:"unique"`
	UserID             uint
	Description        string
	Duration           time.Duration
	Price              int64
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
