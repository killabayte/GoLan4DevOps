package main

import "fmt"

func main() {
	customer := getData(1)
	fmt.Println(customer)

	//Get the wrong customer
	customer = getData(3)
	fmt.Println(customer)
}

func getData(customerId int) (customer string) {
	if customerId == 1 {
		return "killabayte Yarick"
	} else if customerId == 2 {
		return "The Mad Hatter"
	} else {
		return "User not found"
	}
}
