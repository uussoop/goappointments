package goappointments

import (
	"time"
)

// CreateAppointment creates a new appointment
func (r *Appointment) CreateAppointment(userId uint, serviceId uint, startTime time.Time, endTime time.Time, status string) error {
	appointment := Appointment{
		UserID:    userId,
		ServiceID: serviceId,
		StartTime: startTime,
		EndTime:   endTime,
		Status:    status,
	}
	return DB.Create(&appointment).Error
}

// GetAppointmentsByUserID retrieves all appointments for a given user ID
func (r *Appointment) GetAppointmentsByUserID(userId uint) ([]Appointment, error) {
	var appointments []Appointment
	return appointments, DB.Where("user_id = ?", userId).Find(&appointments).Error
}

// UpdateAppointment updates an existing appointment
func (r *Appointment) UpdateAppointment(id uint, startTime time.Time, endTime time.Time, status string) error {
	return DB.Save(r).Error
}

// DeleteAppointment deletes an appointment by its ID
func (r *Appointment) DeleteAppointment() error {
	return DB.Delete(r).Error
}

// CreateService creates a new service
func (r *Service) CreateService() error {

	return DB.Create(r).Error
}

// GetAllServices retrieves all available services
func (r *Service) GetAllServices() ([]Service, error) {
	var services []Service
	return services, DB.Find(&services).Error
}

// UpdateService updates an existing service
func (r *Service) UpdateService() error {
	return DB.Save(r).Error
}

// DeleteService deletes a service by its ID
func (r *Service) DeleteService() error {
	return DB.Delete(r).Error
}

// CreateService creates a new service
func (r *ServiceDayConfig) CreateServiceConfig(Date, StartTime, EndTime time.Time, serviceId uint, Description string) error {
	serviceConfig := ServiceDayConfig{
		Date:        Date,
		StartTime:   StartTime,
		EndTime:     EndTime,
		ServiceID:   serviceId,
		Description: Description,
	}
	return DB.Create(&serviceConfig).Error
}

// GetAllServices retrieves all available services
func (r *ServiceDayConfig) GetAllServicesConfig() ([]ServiceDayConfig, error) {
	var servicesconfig []ServiceDayConfig
	return servicesconfig, DB.Find(&servicesconfig).Error
}

// UpdateService updates an existing service
func (r *ServiceDayConfig) UpdateServiceConfig() error {

	return DB.Save(r).Error
}

// DeleteService deletes a service by its ID
func (r *ServiceDayConfig) DeleteServiceConfig() error {
	return DB.Delete(r).Error
}
