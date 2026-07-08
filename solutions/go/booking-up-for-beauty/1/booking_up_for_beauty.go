package booking

import (
    "log"
    "time"
    "strconv"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layout := "1/2/2006 15:04:05"
    t, err := time.Parse(layout, date)
    if err != nil {
    	log.Fatal(err)
    }
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    layout := "January 2, 2006 15:04:05"
    t, err := time.Parse(layout, date)
    if err != nil {
    	log.Fatal(err)
    }
    now := time.Now()
	return now.After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    layout := "Monday, January 2, 2006 15:04:05"
    t, err := time.Parse(layout, date)
    if err != nil {
    	log.Fatal(err)
    }
    hourmin, err := strconv.Atoi(t.Format("1504"))
    if err != nil {
    	log.Fatal(err)
    }
    return hourmin >= 1200 && hourmin <= 1800
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t := Schedule(date)
    return t.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	now := time.Now()
    return time.Date(now.Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
