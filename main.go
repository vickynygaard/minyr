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
	fmt.Print("Choose convert, average or 'q' for exit")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = scanner.Text()
		if input == "q" || input == "exit" {
			fmt.Println("exit")
			os.Exit(0)
		} else if input == "convert" {

			fmt.Print("Do you want to generate a new file? y for yes or n for no")
			var uinput string
			fmt.Scan(&uinput)
			if uinput == "y" {
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
				fmt.Println("Done with conversions. Choose average or exit:")
			} else if uinput == "n" {
				fmt.Print("...")
			} else {
				fmt.Print("type in (y/n)")
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
				fmt.Printf("Average temp (C) is : %0.2f", sum/16754)
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
				fmt.Printf("Average temp (F) is : %0.2f", sum2/16754)
			}
		} else {
			fmt.Println("Please choose convert, average or exit:")
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
	lines[16755] = "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Victoria Nygaard"
	return lines
}
