package main

import (
	"fmt"
	"time"
)


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
		fmt.Println("Посчитать сумму? y/n")
		var question string
		fmt.Scan(&question)
		if question != "y"{
			return
		}
		sum:= SumTransaction(transactions)
		for  range 5{
			time.Sleep(1 * time.Second)
			fmt.Print(" _ ")
		}
		fmt.Println("Ваша сумма =",sum)
	}

	func ScanTransaction()float64{
		var transaction int
		fmt.Scan(&transaction)
		return float64(transaction)
	}

	func SumTransaction(transaction []float64)float64{
		var sum float64

		for _,elem := range transaction{
			sum += elem
		}

		return sum
	} 
