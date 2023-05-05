package yr

import (
	"bufio"
	"os"
	"testing"
)

// tester om antall linjer i filen er 16756
func TestLineCount(t *testing.T) {

	//åpner filen som skal leses
	file, err := os.Open("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		t.Fatalf("failed to open file %v", err)
	}
	defer file.Close()

	//teller antall linjer i filen
	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	//sjekker at linjeantallet samsvarer med det vi vil ha
	expectedCount := 16756
	if lineCount != expectedCount {
		t.Errorf("Unexpected number of lines. Got %v, expected %v", lineCount, expectedCount)
	}

}

func TestCelsiusToFahrenheitLine(t *testing.T) {

	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: "Kjevik;SN39040;18.03.2022 01:50;6", want: "Kjevik;SN39040;18.03.2022 01:50;42.8"},
		{input: "Kjevik;SN39040;07.03.2023 18:20;0", want: "Kjevik;SN39040;07.03.2023 18:20;32.0"},
		{input: "Kjevik;SN39040;08.03.2023 02:20;-11", want: "Kjevik;SN39040;08.03.2023 02:20;12.2"},
	}

	for _, tc := range tests {
		got, _ := CelsiusToFahrenheitLine(tc.input)
		if !(tc.want == got) {
			t.Errorf("expected %s, got: %s", tc.want, got)
		}
	}

}

func TestLastLineOfFile(t *testing.T) {

	type test struct {
		input string
		want  string
	}
	tests := []test{

		{input: "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;",
			want: "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Victoria Nygaard"},
	}

	for _, tc := range tests {
		got := LastLineOfFile(tc.input)
		if !(tc.want == got) {
			t.Errorf("expected %s, got: %s", tc.want, got)
		}
	}
}

/* func TestAverageTemp(t *testing.T) {

	file, err := os.Open("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	//leser gjennomsnittstemperaturen
	avg, err := AverageTemperature(file, "C")
	if err != nil {
		t.Fatalf("failed to calculate average temperature: %s", err)
	}

	//sjekker om verdien er 8.56
	expected := 8.56
	if avg != expected {
		t.Errorf("Avg %v, want %v", avg, expected)
	}

}
*/
