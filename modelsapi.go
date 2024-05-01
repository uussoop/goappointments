package goappointments

import (
	"time"
)

// CreateAppointment creates a new appointment
func (r *Appointment) createAppointment() error {

	return DB.Create(&r).Error
}

// GetAppointmentsByUserID retrieves all appointments for a given user ID
func getAppointmentsByUserID(userId uint) ([]Appointment, error) {
	var appointments []Appointment
	return appointments, DB.Where("user_id = ?", userId).Find(&appointments).Error
}

// UpdateAppointment updates an existing appointment
func (r *Appointment) updateAppointment(id uint, startTime time.Time, endTime time.Time, status string) error {
	return DB.Save(r).Error
}

// DeleteAppointment deletes an appointment by its ID
func (r *Appointment) deleteAppointment() error {
	return DB.Delete(r).Error
}

// CreateService creates a new service
func (r *Service) createService() error {

	return DB.Create(r).Error
}

// GetAllServices retrieves all available services
func (r *Service) getAllServices() ([]Service, error) {
	var services []Service
	return services, DB.Find(&services).Error
}

func getServiceByUserId(userId uint) ([]Service, error) {
	var services []Service
	return services, DB.Where("user_id = ?", userId).Find(&services).Error
}
func getServiceById(Id uint) (Service, error) {
	var services Service
	return services, DB.Where("id = ?", Id).First(&services).Error
}

// UpdateService updates an existing service
func (r *Service) updateService() error {
	return DB.Save(r).Error
}

// DeleteService deletes a service by its ID
func (r *Service) deleteService() error {
	return DB.Delete(r).Error
}

// CreateService creates a new service
func (r *ServiceDayConfig) createServiceConfig() error {

	return DB.Create(&r).Error
}

// GetAllServices retrieves all available services
func (r *ServiceDayConfig) getAllServicesConfig() ([]ServiceDayConfig, error) {
	var servicesconfig []ServiceDayConfig
	return servicesconfig, DB.Find(&servicesconfig).Error
}

// UpdateService updates an existing service
func (r *ServiceDayConfig) updateServiceConfig() error {

	return DB.Save(r).Error
}
func getServiceConfigByServiceId(userId uint) ([]ServiceDayConfig, error) {
	var servicesConfig []ServiceDayConfig
	return servicesConfig, DB.Where("service_id = ?", userId).Find(&servicesConfig).Error
}

// DeleteService deletes a service by its ID
func (r *ServiceDayConfig) deleteServiceConfig() error {
	return DB.Delete(r).Error
}
