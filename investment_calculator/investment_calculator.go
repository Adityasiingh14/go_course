package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	investementAmount,years,expectedReturnRate := 1000.0,10.0,5.5

	// fmt.Print("Investment Amount: ")

	outputText("Inflation Amount :")
	fmt.Scan(&investementAmount)


	futureValue,futureRealValue := calculateValues(investementAmount, expectedReturnRate, years);
	// futureValue := investementAmount * math.Pow(1 + expectedReturnRate/100, years);
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100,years );
	// fmt.Println( "future Value :" ,futureValue);
	// fmt.Printf( "future Value: %v\nfutureRealValue: %v" , futureValue, futureRealValue);

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue )
	formattedRFV := fmt.Sprintf("Future Value: %.1f\n", futureRealValue )

	fmt.Print(formattedFV + formattedRFV)
	// fmt.Printf( "future Value: %.0f\nfutureRealValue: %.0f" , futureValue, futureRealValue);
	// fmt.Println("futureRealValue :" ,futureRealValue)

}


func outputText(text string){
	fmt.Print(text)
}

func calculateValues(investementAmount,years,expectedReturnRate float64) (fv float64, rfv float64){
	fv = investementAmount * math.Pow(1 + expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100,years )
	return
}