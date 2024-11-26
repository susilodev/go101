package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	investmentAmount, expectedReturnRate, years := 1000.0, 5.5, 10.0

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("expectedReturnRate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("years: ")
	fmt.Scan(&years)

	_, futureRealValue := calcFutureValue(investmentAmount, expectedReturnRate, years)

	fmt.Printf("the future value is %.1f \n", futureRealValue)
}

func calcFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv := fv / math.Pow(1+inflationRate/100, years)

	return fv, frv
}
