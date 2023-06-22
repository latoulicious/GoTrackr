package main

import "fmt"

func main() {
	var firstName string = "John"

	lastName := "Doe"

	var first, second, third string = "A", "B", "C"

	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	var decimalNumber = 2.62

	var exist bool = true

	var message = `Nama saya "John Doe".
	Salam Kenal
	Mari Belajar "Golang".`

	const namaDepan string = "John"
	const namaBelakang = "Wick"


	fmt.Printf("Hello %s %s!\n", firstName, lastName)
	fmt.Printf("Hello %s %s %s!\n", first, second, third)
	fmt.Printf("Positive Number: %d\n", positiveNumber)
	fmt.Printf("negative Number: %d\n", negativeNumber)
	fmt.Printf("Decimal Number: %f\n", decimalNumber)
	fmt.Printf("Decimal Number: %.3f\n", decimalNumber)
	fmt.Printf("Exist: %t\n", exist)
	fmt.Printf("Message: %s\n", message)
	fmt.Print("Hello " , namaDepan , "!\n", "Nice to meet you " , namaBelakang , "!\n")
}