package main

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/uussoop/goappointments"
)

func main() {
	goappointments.InitDBStandlone(goappointments.SQLite)
	err := goappointments.CreateService(1, "phone", "by phone", 30*time.Minute, 250000)
	logrus.Error(err)
	err = goappointments.CreateServiceConfig(time.Date(2024, time.April, 0, 8, 30, 0, 0, time.UTC), time.Date(2024, time.April, 16, 8, 30, 0, 0, time.UTC), time.Date(2024, time.April, 16, 20, 30, 0, 0, time.UTC), "time", 2)
	logrus.Error(err)

	err = goappointments.CreateAppointment(2, 2, time.Date(2024, time.April, 16, 14, 0, 0, 0,
		time.UTC), time.Date(2024, time.April, 16, 14, 30, 0, 0, time.UTC), goappointments.Pending)
	logrus.Error(err)

}
