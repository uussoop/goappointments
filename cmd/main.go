package main

import (
	"time"

	"github.com/uussoop/goappointments"
)

func main() {
	goappointments.InitDBStandlone(goappointments.SQLite)
	service := &goappointments.Service{
		Name:        "test",
		Description: "test desc",
		Duration:    time.Minute * 15,
		Price:       1.2,
	}
	service.CreateService()

}
