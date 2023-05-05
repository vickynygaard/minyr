package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/vickynygaard/funtemps/conv"
)

func main() {
	var input string
	lines := Opnelese()
	fmt.Print("velg convert eller average")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = scanner.Text()
		if input == "q" || input == "exit" {
			fmt.Println("exit")
			os.Exit(0)

		} else if input == "convert" {
			//fmt.Println("Konverterer alle målingene gitt i grader Celsius til grader Fahrenheit.")
			fmt.Print("vil du ha en ny fil? y for ja or n for nei ---> ")

			var uinput string
			fmt.Scan(&uinput)

			if uinput == "j" {
				newfile, err := os.Create("kjevik-temp-fahr-20220318-20230318.csv")
				if err != nil {
					log.Println(err)
					continue

				}

				defer newfile.Close()
				writer := bufio.NewWriter(newfile)
				defer writer.Flush()

				for i := 0; i <= 16755; i++ {
					line := lines[i]
					fields := strings.Split(line, ";")

					if len(fields) == 2 {

						writer.WriteString("Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Victoria Nygaard")
						writer.Flush()
						continue

					} else if len(fields) == 4 && fields[3] == "Lufttemperatur" {

						writer.WriteString("Navn;Stasjon;Tid(norsk normaltid);Lufttemperatur\n")

						continue

					} else {

						temp, err := strconv.ParseFloat(fields[len(fields)-1], 64)
						if err != nil {

							log.Println(err)
							continue
						}

						fahrenheit := conv.CelsiusToFahrenheit(temp)
						ls := fmt.Sprintf("%2.1f\n", fahrenheit)

						lastIndex := strings.LastIndex(line, ";")
						if lastIndex != -1 {
							line = line[:lastIndex]
							line += ";"

							//lastline(lines)
							ls2 := fmt.Sprint(line, ls)
							writer.WriteString(ls2)
							writer.Flush()

						}

					}
				}
				fmt.Println("ferdig Konvertere alle lingene fra grader Celsius til grader Fahrenheit.velg average eller exit:")
			} else if uinput == "n" {

				fmt.Print("ok,skal ikke gjøre noe")
			} else {
				fmt.Print("bare j eller n")

			}

		} else if input == "average" {
			fmt.Print("c or f")
			var uinput string
			fmt.Scan(&uinput)

			if uinput == "c" {
				sum := 0.0

				for i := 1; i <= 16754; i++ {
					line := lines[i]
					fields := strings.Split(line, ";")
					temp, err := strconv.ParseFloat(fields[len(fields)-1], 64)

					if err != nil {
						log.Println(err)
						continue
					}

					sum += temp
				}
				fmt.Printf("gjennomsnittstemperatur (C) er : %0.2f", sum/16754)
			} else if uinput == "f" {
				sum2 := 0.0
				for i := 1; i <= 16754; i++ {
					line := lines[i]
					fields := strings.Split(line, ";")

					temp, err := strconv.ParseFloat(fields[len(fields)-1], 64)
					if err != nil {

						log.Println(err)
						continue

					}

					fahrenheit := conv.CelsiusToFahrenheit(temp)

					sum2 += fahrenheit

				}
				fmt.Printf("gjennomsnittstemperatur (F) er : %0.2f", sum2/16754)
			}
		} else {
			fmt.Println("Venligst velg convert, average eller exit:")
		}
	}
}

func Opnelese() []string {
	var lines []string
	fill, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Println(err)
	}

	defer fill.Close()
	scanner := bufio.NewScanner(fill)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	lines[16755] = "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av khalel alsamael"
	return lines
}

/*package main

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

	fmt.Print("Type 'convert' to convert temperature from Celsius to Fahrenheit and save the results in a new file")
	fmt.Print("Type 'average' to calculate the average temperature")
	fmt.Print("Type 'q' to exit the program")
	fmt.Print(">>>")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		//fjerner alle tomme spaces før og etter input i tillegg til å gjøre inputen om til små bokstaver
		input := strings.TrimSpace(strings.ToLower(scanner.Text()))

		switch input {
		case "convert":
			fmt.Print("Convert")
			if err := handleConvertOption(); err != nil {
				log.Fatal(err)
			}
		case "average":
			fmt.Print("Average")
			if err := handleAverageOption(); err != nil {
				log.Fatal(err)
			}
		case "exit", "q":
			fmt.Print("Exiting program...")
			os.Exit(0)
		default:
			fmt.Print("Invalid choice. Please try again.")
		}
	}
}

func handleConvertOption() error {

	// funksjon som åpner fil, leser linjer, gjør endringer og lagrer nye linjer i en ny fil

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
		case "y", "yes":
			// genererer outputfil på nytt
			if err := generateOutputFile(); err != nil {
				return err
			}
			fmt.Print("Fil generert")
		default:
			fmt.Print("Avslutter progammet")
		}
	} else {
		// Output filen finnes ikke fra før og en ny vil bli generert
		if err := generateOutputFile(); err != nil {
			return err
		}
		fmt.Printf("%s generated\n", inputFile)
	}
	return nil
}

func handleAverageOption() error {

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
