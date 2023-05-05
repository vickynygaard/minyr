package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// konstant variabel som holder filnavnet til inputfilen som programmet leser data fra og
// outputfil hvor de konverterte dataene blir skrevet inn i en ny fil.
const (
	inputFile  = "kjevik-temp-celsius-20220318-20230318.csv"
	outputFile = "kjevik-temp-fahr-20220318-20230318.csv"
)

func main() {

	fmt.Println("Type 'convert' to convert temperature from Celsius to Fahrenheit and save the results in a new file")
	fmt.Println("Type 'average' to calculate the average temperature")
	fmt.Println("Type 'q' to exit the program")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := strings.TrimSpace(strings.ToLower(scanner.Text()))

		switch input {
		case "convert":
			fmt.Println("Convert")
			if err := handleConvertOption(); err != nil {
				log.Fatal(err)
			}
		case "average":
			fmt.Println("Average")
			if err := handleAverageOption(); err != nil {
				log.Fatal(err)
			}
		case "exit", "q":
			fmt.Println("Exiting program...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func handleConvertOption() error {

	//se om output filen allerede eksisterer
	if _, err := os.Stat(outputFile); err == nil {

		// Om filen eksisterer får brukeren et valg om å generere på nytt
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("The file already exists, do you wish to regenerate it? (y/n):", outputFile)
		confirm, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		confirm = strings.TrimSpace(strings.ToLower(confirm))

		switch confirm {
		case "y", "yes", "ja":
			// Genererer outputfil på nytt
			if err := generateOutputFile(); err != nil {
				return err
			}
			fmt.Println("Fil generert")
		default:
			// Do not generate output file again
			fmt.Println("Avslutter progammet")
		}
	} else {
		// Output file does not exist, generate it
		if err := generateOutputFile(); err != nil {
			return err
		}
		fmt.Println("Fil generert")
	}
	return nil
}

//evt annen kode å bruke til samme func^^
/* input, err := os.Open(inputFile)
if err != nil {
	return fmt.Errorf("failed to open input file: %v", err)
}
defer input.Close()

//lager output fil
output, err := os.Create(outputFile)
if err != nil {
	return fmt.Errorf("failed to create output file: %v", err)
}
defer output.Close()
*/

func handleAverageOption() error {

}
