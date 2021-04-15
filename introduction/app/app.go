package main

import "fmt"

func main() {
	customer := getData(1)
	fmt.Println(customer)
}

func getData(customerId int) (customer string) {
	var firstName = "killabayte"
	lastName := "Gophers"

	return firstName + " " + lastName
}
