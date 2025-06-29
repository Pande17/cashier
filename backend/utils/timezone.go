package utils

import (
	"log"
	"time"
)

// SetTimezone sets the time zone for the application
func SetTimezone() (*time.Location, error) {
	// Load the timezone (Asia/Jakarta corresponds to GMT+8)
	loc, err := time.LoadLocation("Asia/Singapore")
	if err != nil {
		log.Printf("Error loading timezone: %s\n", err.Error())
		return nil, err
	}
	return loc, nil
}

// GetCurrentTimeInTimezone returns the current time in the specified timezone
func GetCurrentTimeInTimezone(loc *time.Location) time.Time {
	return time.Now().In(loc)
}

// FormatDate formats the date to "13-Jun-2025" format
func FormatDate(t time.Time) string {
	return t.Format("02-Jan-2006")
}
