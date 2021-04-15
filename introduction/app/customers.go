package main

type (
	Customer struct {
		FirstName string
		LastName  string
		FullName  string
	}
)

func GetCustomers() (customers []Customer) {
	killabayte := Customer{FirstName: "killabayte", LastName: "Yarick"}

	customers = append(customers, killabayte,
		Customer{FirstName: "Alice", LastName: "in the wonderland"},
		Customer{FirstName: "March", LastName: "Hare"},
		Customer{FirstName: "Dormose", LastName: "Nobody knows"},
		Customer{FirstName: "The", LastName: "Mad", FullName: "The Mad Hatter"},
	)

	return customers

}
