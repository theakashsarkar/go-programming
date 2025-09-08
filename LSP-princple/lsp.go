package main

import "fmt"

type Payment interface {
	Pay(amount float64) string
}

type CreditCardPayment struct{}

func (c CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f with Credit Card", amount)
}

func processPayment(p Payment, amount float64) {
	fmt.Println(p.Pay(amount))
}

func main() {
	cc := CreditCardPayment{}
	processPayment(cc, 100)
}
