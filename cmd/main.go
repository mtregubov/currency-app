package main

import (
	"curcli/model"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func Processrow(header *[]string, row string) error {
	columns := strings.Split(row, "|")
	for i := 1; i < len(*header); i++ {
		// separe currency quantity and name
		quantityname := strings.Split((*header)[i], " ")
		// currency name is always second part
		currname := quantityname[1]
		// currency quantity is always first part
		quantity, _ := strconv.Atoi(quantityname[0])
		// date is always first column
		date := columns[0]
		// convert rate to float
		ratevalue, _ := strconv.ParseFloat(columns[i], 64)
		// create rate struct and save to database
		rate := model.Rate{Date: date, Curr: currname, Quantity: quantity, Rate: ratevalue}
		err := rate.Save()
		if err != nil {
			return err
		}
	}

	return nil
}

func Processbody(url string) error {
	// get data from API via http client
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("url status code <> 200")
	}

	// read the body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// split body to slice of strings
	csvrow := strings.Split(string(bodyBytes), "\n")

	// first row is header
	header := csvrow[0]
	columns := strings.Split(header, "|")

	for i := 1; i < len(csvrow)-1; i++ {
		// process each row
		err = Processrow(&columns, csvrow[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	log.Println("Setting up database...")
	err := model.SetupDB()
	if err != nil {
		log.Fatal(err)
	}

	year := "2018"
	if cy := os.Getenv("CURR_YEAR"); cy != "" {
		year = cy
	}
	log.Println("Processing year:", year)

	url := fmt.Sprintf("https://www.cnb.cz/en/financial_markets/foreign_exchange_market/exchange_rate_fixing/year.txt?year=%s", year)
	log.Println("Getting data from URL and saving to database...")

	err = Processbody(url)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Done")
}
