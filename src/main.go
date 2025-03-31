package main

import (
	"fmt"
	"time"
)

const (
	targetEpoch = 2147483647 // 2038-01-19 03:14:07 UTC
)

func main() {
	for {
		printCurrentStatus()
		time.Sleep(1 * time.Second)
	}
}

func printCurrentStatus() {
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
	remaining := float64(targetEpoch - epoch)

	// Time units in seconds
	secondsInMinute := 60.0
	secondsInHour := 60.0 * secondsInMinute
	secondsInDay := 24.0 * secondsInHour
	secondsInYear := 365.25 * secondsInDay // Account for leap years
	secondsInMonth := secondsInYear / 12.0

	// Calculate remaining time components
	years := int(remaining / secondsInYear)
	remaining = remaining - float64(years)*secondsInYear

	months := int(remaining / secondsInMonth)
	remaining = remaining - float64(months)*secondsInMonth

	days := int(remaining / secondsInDay)
	remaining = remaining - float64(days)*secondsInDay

	hours := int(remaining / secondsInHour)
	remaining = remaining - float64(hours)*secondsInHour

	minutes := int(remaining / secondsInMinute)
	seconds := int(remaining - float64(minutes)*secondsInMinute)

	// Clear screen (works on Unix)
	fmt.Print("\033[H\033[2J")

	// Output
	fmt.Println("==== unixy2k CLI Countdown ====")
	fmt.Printf("UTC Date      : %s\n", utcString)
	fmt.Printf("Epoch Time    : %d\n", epoch)
	fmt.Printf("Epoch Binary  : %s\n", binaryChunks)
	fmt.Println()
	fmt.Printf("Remaining Time until 2038-01-19 03:14:07 UTC:\n")
	fmt.Printf("%d years, %d months, %d days, %d hours, %d minutes, %d seconds\n",
		years, months, days, hours, minutes, seconds)
}
