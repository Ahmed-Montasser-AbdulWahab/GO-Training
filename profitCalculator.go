package main

import (
	"fmt"
	"os"
)

const fileName string = "records.csv"

func main() {

	for {
		var choice int8

		fmt.Println("Do you want to calculate profit? (1 for Yes, 0 for No)")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			var taxRate, revenue, expense float64

			taxRate = takeInput("Please Enter Tax Rate (Ex : 0.5)")

			revenue = takeInput("Please Enter Revenue")

			expense = takeInput("Please Enter Expense")

			if !(isAboveZero(taxRate) && isAboveZero(revenue) && isAboveZero(expense)) {
				fmt.Println("All inputs should be above 0. Please try again.")
				continue
				//panic("All inputs should be above 0")
				//break
			}

			profitBeforeTax, profitAfterTax, ratio := calculateValues(revenue, expense, taxRate)

			fmt.Printf("Profit Before Tax: %4.2f\n", profitBeforeTax)
			fmt.Printf("Profit After Tax: %4.2f\n", profitAfterTax)
			fmt.Printf("Ratio: %4.2f\n", ratio)

			writeToFile(revenue, expense, taxRate, profitBeforeTax, profitAfterTax, ratio)

		case 0:
			fmt.Println("Exiting the program.")
			return
		}
	}
}

func takeInput(prompt string) float64 {

	var input float64

	fmt.Println(prompt)

	fmt.Scan(&input)

	return input
}

func isAboveZero(input float64) bool {
	return input > 0
}

func calculateValues(revenue, expense, taxRate float64) (pBT, pAT, ratio float64) {
	pBT = revenue - expense

	pAT = pBT * (1 - taxRate)

	ratio = pBT / pAT

	return
}

// func writeToFile(revenue, expense, taxRate, pBT, pAT, ratio float64) {
// 	if !fileExists(fileName) {
// 		os.WriteFile(fileName, []byte("revenue,expense,taxRate,pBT,pAT,ratio\n"), 0644)
// 	}
// 	record := fmt.Sprint(revenue, ",", expense, ",", taxRate, ",", pBT, ",", pAT, ",", ratio, "\n")
// 	os.WriteFile(fileName, []byte(record), 0644)

// }

func writeToFile(revenue, expense, taxRate, pBT, pAT, ratio float64) {
	if !fileExists(fileName) {
		os.WriteFile(fileName, []byte("revenue,expense,taxRate,pBT,pAT,ratio\n"), 0644)
	}
	record := fmt.Sprint(revenue, ",", expense, ",", taxRate, ",", pBT, ",", pAT, ",", ratio, "\n")

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(record)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
