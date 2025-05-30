package timer

import (
	"fmt"
	"time"
)

const (
	targetEpoch = 2147483647 // 2038-01-19 03:14:07 UTC
)

// Status holds the formatted time information.
type Status struct {
	UTCString    string
	Epoch        int64
	BinaryChunks []string
	Years        int
	Months       int
	Days         int
	Hours        int
	Minutes      int
	Seconds      int
}

// GetStatus calculates the current time details and returns them in a Status struct.
func GetStatus() (*Status, error) {
	now := time.Now().UTC()
	epoch := now.Unix()

	// Format UTC time
	utcString := now.Format("2006-01-02 15:04:05")

	// Binary representation, padded to 32 bits
	binaryString := fmt.Sprintf("%032b", epoch)

	// Split binary into 8-bit chunks
	var binaryChunks []string
	for i := 0; i < 32; i += 8 {
		binaryChunks = append(binaryChunks, binaryString[i:i+8])
	}

	// Time left until 2038 problem
	remainingSeconds := float64(targetEpoch - epoch)

	// Time units in seconds
	secondsInMinute := 60.0
	secondsInHour := 60.0 * secondsInMinute
	secondsInDay := 24.0 * secondsInHour
	secondsInYear := 365.25 * secondsInDay // Account for leap years
	secondsInMonth := secondsInYear / 12.0

	// Calculate remaining time components
	years := int(remainingSeconds / secondsInYear)
	remainingSeconds -= float64(years) * secondsInYear

	months := int(remainingSeconds / secondsInMonth)
	remainingSeconds -= float64(months) * secondsInMonth

	days := int(remainingSeconds / secondsInDay)
	remainingSeconds -= float64(days) * secondsInDay

	hours := int(remainingSeconds / secondsInHour)
	remainingSeconds -= float64(hours) * secondsInHour

	minutes := int(remainingSeconds / secondsInMinute)
	seconds := int(remainingSeconds - float64(minutes)*secondsInMinute)

	return &Status{
		UTCString:    utcString,
		Epoch:        epoch,
		BinaryChunks: binaryChunks,
		Years:        years,
		Months:       months,
		Days:         days,
		Hours:        hours,
		Minutes:      minutes,
		Seconds:      seconds,
	}, nil
}
