package model

import (
	"database/sql"
	"log"
)

var DB *sql.DB
var ErrNoRows = sql.ErrNoRows

type Rate struct {
	Date     string  `json:"date"`
	Curr     string  `json:"curr"`
	Quantity int     `json:"amount"`
	Rate     float64 `json:"rate"`
}

// save rate to database
func (r *Rate) Save() error {
	_, err := DB.Exec("INSERT INTO currencies (date, curr, quontity, rate) VALUES (?, ?, ?, ?) on conflict (date, curr) do update set rate = ?", r.Date, r.Curr, r.Quantity, r.Rate, r.Rate)
	if err != nil {
		return err
	}
	return nil
}

// find rate in database
func FindRates(date, curr string) ([]Rate, error) {
	var rates = make([]Rate, 0)
	var rows *sql.Rows
	var err error
	if curr == "" {
		rows, err = DB.Query("SELECT date, curr, quontity, rate FROM currencies WHERE date = ?", date)
	} else {
		rows, err = DB.Query("SELECT date, curr, quontity, rate FROM currencies WHERE date = ? AND curr = ?", date, curr)
	}
	if err != nil {
		return rates, err
	}
	defer rows.Close()

	for rows.Next() {
		rate := Rate{}
		err = rows.Scan(&rate.Date, &rate.Curr, &rate.Quantity, &rate.Rate)
		if err != nil {
			log.Println(err)
		}
		rates = append(rates, rate)
	}

	return rates, nil
}

func SetupDB() error {
	// connect to database
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return err
	}
	DB = db

	// migrate database
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS currencies (date text, curr text, quontity integer, rate real, PRIMARY KEY (date, curr))")
	if err != nil {
		return err
	}

	return nil
}
