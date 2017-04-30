// Package posDateTime returns Time and Dates to the POS application in the prefered format.
package posDateTime

import (
	"time"
)

// PosDate export the current system date.
func PosDate() string {
	s := systemDate()
	return s
}

func systemDate() string {
	layout := "Jan 2, 2006"
	d := time.Now()
	short := d.Format(layout)
	return short
}
