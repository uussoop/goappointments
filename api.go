package goappointments

import "time"

func CreateAppointment(userId, serviceid uint, starttime, endtime time.Time, status StatusType) error {

	appo := &Appointment{
		UserID:    userId,
		StartTime: starttime,
		EndTime:   endtime,
		ServiceID: serviceid,
		Status:    status,
	}

	return appo.createAppointment()
}
func GetAppointmentByUserId(userid uint) ([]Appointment, error) {
	return getAppointmentsByUserID(userid)

}
func CreateService(userid uint, name, description string, duration time.Duration, price int64) error {
	s := &Service{
		UserID:      userid,
		Name:        name,
		Description: description,
		Duration:    duration,
		Price:       price,
	}
	return s.createService()
}
func GetServiceByUserId(userid uint) ([]Service, error) {
	return getServiceByUserId(userid)

}
func CreateServiceConfig(date, starttime, endtime time.Time, description string, serviceid uint) error {
	serviceday := &ServiceDayConfig{
		Date:        date,
		Description: "workingday",
		StartTime:   starttime,
		EndTime:     endtime,
		ServiceID:   serviceid,
	}
	return serviceday.createServiceConfig()
}
func GetServiceConfigByServiceId(userid uint) ([]ServiceDayConfig, error) {
	return getServiceConfigByServiceId(userid)

}
