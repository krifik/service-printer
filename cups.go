// package main

// import (
// 	"fmt"
// 	"log"

// 	// "github.com/michaelbironneau/go-cups/cups"
// )

// func main() {
// 	// Initialize the CUPS client
// 	client, err := cups.NewDefaultClient()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Get the default printer
// 	printer, err := client.Default()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Print a file
// 	filePath := "/path/to/file.txt" // Replace with the path of the file you want to print
// 	jobID, err := printer.PrintFile(filePath, "Print Job", nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("Print job submitted. Job ID: %d\n", jobID)

// 	// Close the CUPS client
// 	client.Close()
// }
