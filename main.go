// Package "main"
package main


//Importing Essential Libraries
import (
    "fmt"
	"github.com/uzair0007/golangexamples"
)


// Main Function to Check the "golangexamples" Package
func main() {

	// Printing the Original string
	fmt.Println("Original String: ", golangexamples.GoLangString)

	// Converting the String to Byte Slice
	s := golangexamples.String2Slice(golangexamples.GoLangString)
	fmt.Println("Converted Slice: ", s)
	
	// Printing Concatenated String
	fmt.Println("Concatenated String: ", golangexamples.ConcatSlice(s))
	
	// Printing the Encryoted String
	fmt.Println("Encrypted String: ", golangexamples.Encrypt(s, 3))
	
	// Printing the EZ Greetings from imported Package
	fmt.Println("Greetings String: ", golangexamples.EZGreetings("EZ"))
}