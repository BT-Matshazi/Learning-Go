package main

import (
	"fmt"
)

// Assignment
// Complete the placeOrder function.

// It returns a bool indicating whether the order was successful (true is a success) and 
// a float64 representing the user's balance after the order. The placeOrder function 
// should always return the account balance regardless of if it was adjusted or not.

// The amountInStock and calcPrice functions can be used to look up the current stock and price of an item.

// If the quantity is greater than the amount in stock, the order should be rejected.
// If the user doesn't have enough money in their account, the order should be rejected.
// Otherwise, the order should be accepted and you should return the new balance.

func main(){

	orderStatus, balance := placeOrder("2", 25, 459.00)

	fmt.Printf("The order was: %t, the balance is %.2f", orderStatus, balance)
}

func placeOrder(productID string, quantity int, accountBalance float64) (bool, float64) {

	isInStock := amountInStock(productID) >= quantity

	if (!isInStock) { 
		return false, accountBalance
	}

	purchasePrice := calcPrice(productID, quantity)

	if(purchasePrice <= accountBalance){
		return true, accountBalance - purchasePrice
	}

	return false, accountBalance
}

func placeOrderImproved(productID string, quantity int, accountBalance float64) (bool, float64) {
    if amountInStock(productID) < quantity {
        return false, accountBalance
    }

    purchasePrice := calcPrice(productID, quantity)
    if purchasePrice > accountBalance {
        return false, accountBalance
    }

    return true, accountBalance - purchasePrice
}

// Don't touch below this line

func calcPrice(productID string, quantity int) float64 {
	return priceList(productID) * float64(quantity)
}

func priceList(productID string) float64 {
	if productID == "1" {
		return 1.50
	} else if productID == "2" {
		return 2.25
	} else if productID == "3" {
		return 3.00
	} else if productID == "4" {
		return 1.00
	} else if productID == "5" {
		return 2.50
	} else if productID == "6" {
		return 8.99
	} else if productID == "7" {
		return 22.50
	} else if productID == "8" {
		return 50.00
	} else if productID == "9" {
		return 999.99
	} else {
		return 0.00
	}
}

func amountInStock(productID string) int {
	if productID == "1" {
		return 11
	} else if productID == "2" {
		return 25
	} else if productID == "3" {
		return 4
	} else if productID == "4" {
		return 6
	} else if productID == "5" {
		return 50
	} else if productID == "6" {
		return 2
	} else if productID == "7" {
		return 0
	} else if productID == "8" {
		return 99
	} else if productID == "9" {
		return 1
	} else {
		return 0
	}
}
