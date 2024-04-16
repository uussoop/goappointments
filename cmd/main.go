package main

import (
	"github.com/uussoop/goappointments"
)

func main() {
	goappointments.InitDBStandlone(goappointments.SQLite)
	// service := &goappointments.Service{
	// 	Name:        "test",
	// 	Description: "test desc",
	// 	Duration:    time.Minute * 15,
	// 	Price:       1.2,
	// }
	// err := service.CreateService()
	// logrus.Warn("error create service: ", err)
	// serviceday := &goappointments.ServiceDayConfig{
	// 	Date:        time.Date(2024, time.April, 16, 14, 0, 0, 0, time.UTC),
	// 	Description: "workingday",
	// 	StartTime:   time.Date(2024, time.April, 16, 14, 0, 0, 0, time.UTC),
	// 	EndTime:     time.Date(2024, time.April, 16, 14, 30, 0, 0, time.UTC),
	// 	ServiceID:   1,
	// }
	// err = serviceday.CreateServiceConfig()
	// logrus.Warn("error create servicedayconfig: ", err)

}
