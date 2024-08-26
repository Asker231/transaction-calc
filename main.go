package main

import "fmt"


	func main() {
		var transactions []float64 = []float64{}
		for{
			fmt.Print("Your transaction:")
			transaction := ScanTransaction()
			transactions = append(transactions, transaction)
			if transaction == 0{
				fmt.Println(transactions)
				break
			}
		}
	}

	func ScanTransaction()float64{
		var transaction int
		fmt.Scan(&transaction)
		return float64(transaction)
	}
