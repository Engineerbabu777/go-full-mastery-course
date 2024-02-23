package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()

	fmt.Println(currentTime.Format("2006-01-02 15:04:05 Monday"))

	// CREATE A NEW DATE!
	newDate := time.Date(2020, 12, 25, 10, 45, 0, 0, time.UTC)
	fmt.Println(newDate.Format("2006-01-02"))
}
