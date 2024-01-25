package repository

import (
	"database/sql"
)

type DatabaseRepository interface {
	Connection() *sql.DB
	GetLastStationData(args ...interface{}) (interface{}, error)
	GetBatterySOC(args ...interface{}) (interface{}, error)
}
