package goappointments

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (a *Appointment) BeforeCreate(tx *gorm.DB) (err error) {
	service, err := getServiceById(a.ServiceID)
	if err != nil {
		return fmt.Errorf("service doesnt exist")
	}
	if (a.EndTime.Sub(a.StartTime)) != service.Duration {

		return fmt.Errorf("appointment time frame is less than or bigger than its service ")
	}
	var serviceConfig ServiceDayConfig
	err = tx.Where("service_id = ? AND start_time < ? AND end_time > ?", a.ServiceID, a.StartTime, a.EndTime).First(&serviceConfig).Error
	if err != nil {
		return fmt.Errorf("cannot make this appointment  ", err)

	}

	// Check for overlapping appointments
	var overlappingAppointments int64
	err = tx.Model(&Appointment{}).
		Where("user_id = ? AND status > 0 AND ((? >= start_time AND ? <= end_time) OR (? >= start_time AND ? <= end_time))", a.UserID, a.StartTime, a.EndTime, a.StartTime, a.EndTime).
		Count(&overlappingAppointments).Error
	if err != nil {
		return err
	}

	if overlappingAppointments > 0 {
		return fmt.Errorf("reserving appointment failed. overlapping appointment found!")
	}

	return nil
}

func (u *Appointment) AfterCreate(tx *gorm.DB) (err error) {

	return
}
func (sd *ServiceDayConfig) BeforeCreate(tx *gorm.DB) (err error) {
	logrus.Info("before service day config create")
	// Check for overlapping ServiceDayConfig records
	var overlappingConfigs []ServiceDayConfig
	err = tx.Where("date BETWEEN ? AND ?", sd.Date, sd.Date).
		Where("(start_time BETWEEN ? AND ?) OR (end_time BETWEEN ? AND ?)", sd.StartTime, sd.EndTime, sd.StartTime, sd.EndTime).
		Find(&overlappingConfigs).Error

	if err != nil {
		return err
	}
	if len(overlappingConfigs) > 0 {
		logrus.Errorf("overlapping ServiceDayConfig found for service ID %d on date %s", sd.ServiceID, sd.Date.Format("2006-01-02"))

		return fmt.Errorf("overlapping ServiceDayConfig found for service ID %d on date %s", sd.ServiceID, sd.Date.Format("2006-01-02"))
	}

	return nil
}
