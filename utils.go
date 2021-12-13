package main

import (
	"log"
	"time"
)

func DateParsing(layoutFormat string, date string) (time.Time, error) {
	t, err := time.Parse(layoutFormat, date)
	if err != nil {
		log.Fatal(err)
	}

	return t, err
}
