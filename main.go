/**
 * @author Miles Aube
 * @version 1.0.0
 * @date 2025-12-06
 * @fileoverview This program asks the user for two ASCII values (between 32 and 126)
 */

package main

import (
	"fmt"
)

func main() {

	// get user input
	var startValue int
	var endValue int

	fmt.Print("Please enter a number larger than 32, and less than 126: ")
	fmt.Scan(&startValue)

	fmt.Print("Please enter a number larger than ", startValue, " and less than 126: ")
	fmt.Scan(&endValue)

	// check if input is correct
	for endValue <= startValue || startValue < 32 || endValue > 126 {
		fmt.Print("Please enter a number larger than ", startValue, " and less than 126: ")
		fmt.Scan(&endValue)
	}

	// set variables
	counter := 0
	output := ""

	// loop through ASCII values
	for counter = startValue; counter <= endValue; counter = counter + 1 {
		output = output + fmt.Sprint(counter, " = ", string(rune(counter)), "\n")
	}

	// display results
	fmt.Println(output)
}