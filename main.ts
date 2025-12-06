/**
 * @author Miles Aube
 * @version 1.0.0
 * @date 2025-12-06
 * @fileoverview This program asks the user for two ASCII values (between 32 and 126)
 */

// get user input
let startValue: number = Number(prompt("Please enter a number larger than 32, and less than 126:"));
let endValue: number = Number(prompt("Please enter a number larger than " + startValue + " and less than 126:"));

// check if input is correct 
while (endValue <= startValue || startValue < 32 || endValue > 126) {
  endValue = Number(prompt("Please enter a number larger than " + startValue + " and less than 126:"));
}

// set variables
let counter: number = 0;
let output: string = "";

// loop through ASCII values
for (counter = startValue; counter <= endValue; counter = counter + 1) {
  output = output + counter + " = " + String.fromCharCode(counter) + "\n";
}

// display results
console.log(output);