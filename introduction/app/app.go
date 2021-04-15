package main

import (
	"fmt"
)

func main() {
	customers := getData()
	fmt.Println(customers)
	fmt.Println(len(customers))
}

func getData() (customers []string) {
	customers = []string{"killabayte", "The Mad Hatter"}
	customers = append(customers, "Alice")
	customers = append(customers, "March Hare")
	customers = append(customers, "Dormouse")

	for _, customer := range customers {
		fmt.Println(customer)
	}

	return customers
}
