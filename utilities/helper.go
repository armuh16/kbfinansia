package utilities

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}

func TimeParse(layout, value string) time.Time {
	// Parse the string into a time.Time value
	parsedTime, err := time.Parse(layout, value)
	if err != nil {
		fmt.Println("Error parsing the date:", err)
	}

	return parsedTime
}
