package utils

import (
	"fmt"
	"log"
	"time"
)

func FormatDate(location, date string) string {
	loc, err := time.LoadLocation(location)
	if err != nil {
		log.Printf("LOG [error]: %v", err)
	}
	t, err := time.ParseInLocation(time.RFC3339, date, loc)
	if err != nil {
		log.Printf("LOG [error]: %v", err)
	}
	return fmt.Sprintf("%v", t.Format(time.RFC822))
}
