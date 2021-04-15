package main

import "fmt"

func main() {
	var (
		userName string
	)

	fmt.Scan(&userName)
	fmt.Printf("Wellcome to the GoLang for DevOps cource: %s\n", userName)
}
