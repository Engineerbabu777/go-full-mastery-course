
package main

import (
    "fmt"
    "time"
)

func main() {
    currentTime := time.Now() // WILL GIVE US THE CURRENT TIME!
    fmt.Println("Current time:", currentTime)

	// THIS WILL FORMAT THE TIME!
    formattedTime := currentTime.Format("2006-01-02 11:04:05")
    fmt.Println("Formatted time:", formattedTime)

	// INCREASING THE TIME!
	increasedTime := currentTime.Add(10 * time.Minute) 
	// ADDING 10 MINUTES TO THE CURRENT TIME!
	fmt.Println(increasedTime.Format("2006-01-02 11:04:05"))

	// SLEEP METHOD!
	time.Sleep(5 * time.Second) // AT THIS POINT THE EXECUTION OF PROGRAMS STOPS FOR 5 seconds!
	fmt.Println("Woke up!")

}