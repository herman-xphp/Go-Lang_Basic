package main

import "fmt"

func main() {
	type noKTP string
	type Married bool

	var noKtpEko noKTP = "12121211211221"
	var marriedStatus Married = true

	fmt.Println(noKtpEko)
	fmt.Println(marriedStatus)
}
