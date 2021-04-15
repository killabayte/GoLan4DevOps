package main

import "fmt"

func main() {
	customers := getData()
	fmt.Println(customers)
	fmt.Println(len(customers))
}

func getData() (customers [2]string) {
	customer := "killabayte Yarick"
	customers[0] = customer
	customers[1] = "The Mad Hatter"

	return customers
}
