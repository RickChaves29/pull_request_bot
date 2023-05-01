package utils

import (
	"fmt"
	"log"
	"time"
)

func FormatDate(date string) string {
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		log.Printf("LOG [error]: %v", err)
	}
	return fmt.Sprintf("%v", t.Format(time.RFC822))
}
