package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type PublicHolidays struct {
	Date      string `json:"date"`
	LocalName string `json:"localName"`
	Name      string `json:"name"`
}

func main() {
	url := "https://date.nager.at/Api/v1/Get/UA/2020"

	spaceClient := http.Client{
		//--------------------
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var public []PublicHolidays
	jsonErr := json.Unmarshal(body, &public)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	// formatting an actual date
	t := time.Now()
	tShort := t.Format("2006-01-02")

	// check the nearest date
	for i := 0; i <= len(public)-1; i++ {
		if tShort == public[i].Date {

			// this block will execute
			// when condition_1 is true
			fmt.Println("It’s a holiday today: ", public[i].Date)
			fmt.Println(public[i].LocalName)
			fmt.Println(public[i].Name)
			publicItemDateParsed, _ := time.Parse("2006-01-02", public[i].Date)
			month := publicItemDateParsed.Month()
			weekday := publicItemDateParsed.Weekday()
			day := publicItemDateParsed.Day()
			fmt.Printf("And the day is: %+v\n", weekday)

			// using Switch for our weekend length evaluation
			switch weekday {
			case time.Friday:
				fmt.Printf("So we have long weekend, that will last 3 days: %v %v - %v %v ", month, day, month, day+2)
			case time.Saturday:
				fmt.Printf("So we have long weekend, that will last 3 days: %v %v - %v %v ", month, day, month, day+2)
			case time.Sunday:
				fmt.Printf("So we have long weekend, that will last 2 days: %v %v - %v %v ", month, day, month, day+1)
			default:

			}
			break

		} else if tShort < public[i].Date {

			// this block will execute
			// when condition2 is true
			fmt.Println("It's not a holiday today. Next closest holiday is:", public[i].Date)
			fmt.Println(public[i].LocalName)
			fmt.Println(public[i].Name)
			publicItemDateParsed, _ := time.Parse("2006-01-02", public[i].Date)
			month := publicItemDateParsed.Month()
			weekday := publicItemDateParsed.Weekday()
			day := publicItemDateParsed.Day()
			fmt.Printf("And the day is: %+v\n", weekday)

			// using Switch for our weekend length evaluation
			switch weekday {
			case time.Friday:
				fmt.Printf("So we have long weekend, that will last 3 days: %v %v - %v %v ", month, day, month, day+2)
			case time.Saturday:
				fmt.Printf("So we have long weekend, that will last 3 days: %v %v - %v %v ", month, day, month, day+2)
			case time.Sunday:
				fmt.Printf("So we have long weekend, that will last 2 days: %v %v - %v %v ", month, day, month, day+1)
			default:

			}
			break
		}

	}

}
