package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"pv-monitor-telegram-bot/pkg/data"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

// GetLastStationData() gets the last Station data from the database
func (m *PostgresDBRepo) GetLastStationData(args ...interface{}) (interface{}, error) {
	var stationData = data.Station{
		LastUpdateTimeRead:   true,
		TotalProductionRead:  true,
		FeedInRead:           true,
		BatteryChargeRead:    true,
		SelfUsedRead:         true,
		TotalConsumptionRead: true,
		PowerPurchasedRead:   true,
		BatteryDischargeRead: true,
		ProductionRead:       true,
		BatterySOCRead:       true,
	}

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `select last_update_ts, total_production, feed_in, battery_charge, self_used, total_consumption, power_purchased, battery_discharge, production, battery_soc
	           from "Station".Station
		      order by last_update_ts desc
			  limit 1`

	rows, err := m.DB.QueryContext(ctx, stmt)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error reading stationData: %s", rows.Err()))
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&stationData.LastUpdateTime,
			&stationData.TotalProduction,
			&stationData.FeedIn,
			&stationData.BatteryCharge,
			&stationData.SelfUsed,
			&stationData.TotalConsumption,
			&stationData.PowerPurchased,
			&stationData.BatteryDischarge,
			&stationData.Production,
			&stationData.BatterySOC,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}
	}

	return stationData, nil
}

// GetBatterySOC() (float32, error) gets the BatterySOC from the database
func (m *PostgresDBRepo) GetBatterySOC(args ...interface{}) (interface{}, error) {
	var batterySOC float32

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `select battery_soc from "Station".Station order by last_update_ts desc limit 1`

	rows, err := m.DB.QueryContext(ctx, stmt)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error reading stationData: %s", rows.Err()))
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&batterySOC,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}
	}

	return batterySOC, nil
}
