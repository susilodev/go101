package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const inflationText = "inflation.txt"

func getInflationRate() (float64, error) {
	rawFile, err := os.ReadFile(inflationText)
	if err != nil {
		return 0.0, errors.New("Can't os.ReadFile")
	}

	inflationString := strings.TrimSpace(string(rawFile))
	inflationRate, err := strconv.ParseFloat(inflationString, 64)
	if err != nil {
		fmt.Println("can't parse float64", err)
		return 0.0, errors.New("Cant parse float64")
	}

	return inflationRate, nil
}

func main() {
	investmentAmount, expectedReturnRate, years := 1000.0, 5.5, 10.0

	fmt.Println("Welcome to the Salaf Investment")

	fmt.Print("full fill your Investment Amount: ")
	fmt.Scan(&investmentAmount)

	if investmentAmount < 0.0 {
		fmt.Println("are you broke?")
		investmentAmount = 0.0
	}

	fmt.Print("expectedReturnRate: ")
	fmt.Scan(&expectedReturnRate)

	if expectedReturnRate < 0.0 {
		fmt.Println("wrong input, try again latter")
		expectedReturnRate = 0.0
	}

	fmt.Print("years: ")
	fmt.Scan(&years)

	if years < 0.0 {
		fmt.Println("wrong years input")
		years = 0.0
	}

	futureRealValue, err := calcFutureValue(investmentAmount, expectedReturnRate, years)
	if err != nil {
		fmt.Println("Error calculating future value:", err)
		return

	}

	fmt.Printf("the future value is %.1f", futureRealValue)
}

func calcFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, error) {
	inflationRate, err := getInflationRate()
	if err != nil {
		return 0.0, err
	}

	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv := fv / math.Pow(1+inflationRate/100, years)

	return frv, nil
}
