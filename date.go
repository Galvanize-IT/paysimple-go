package paysimple

import "time"

const paysimpleDateFormat = `2006-01-02`

// Date discards time of day information
type Date struct {
	time.Time
}

func (date Date) String() string {
	return date.Format(paysimpleDateFormat)
}

// NewDate creates a new Date
func NewDate(year, month, day int) Date {
	return Date{
		Time: time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local),
	}
}
