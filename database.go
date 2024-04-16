package goappointments

import (
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)
var tables = []interface{}{&Appointment{}, &Service{}, &ServiceDayConfig{}}

type DatabaseType int

const (
	// SQLite database
	SQLite DatabaseType = iota
	// MySQL database
	MySQL
	// PostgreSQL database
	PostgreSQL
)

func InitSqlDB() {
	err := InitDBStandlone(SQLite)
	if err != nil {
		logrus.Error("Error: ", err)
		panic(err)
	}

}

func InitDBStandlone(dbtype DatabaseType) (err error) {

	gCnf := &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent),
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	logrus.Info("Initializing sqlite database...")

	err = os.Mkdir("database", 0777)
	if err != nil {
		logrus.Info(err)

	}

	DB, err = openDB("./database/testsqlite.db", gCnf)

	if err != nil {
		logrus.Info("failed to connect database: ", err)
	}

	err = DB.AutoMigrate(

		//set up structs to get in
		tables...,
	)
	if err != nil {

		logrus.Info("error migration ", err)
		return
	}

	logrus.Info("SQL Database has been initialized")
	return

}
func openDB(filepath string, gCnf *gorm.Config) (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open(filepath), gCnf)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func initDBApplication(db *gorm.DB) (err error) {
	DB = db

	err = DB.AutoMigrate(
		//set up structs to get in
		tables...,
	)
	if err != nil {

		logrus.Warn("error migration ", err)
	}

	logrus.Info("Appointments SQL Database has been initialized")

	return

}
