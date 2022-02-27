package main

import "fmt"

type CheckNumber int
type CardNumber string
type CardType uint8

const (
	Visa CardType = iota
	MasterCard
)

type CreditCardInfo struct {
	CardType
	CardNumber
}

type PaymentMethod interface {
	isPaymentMethod()
	// | Cash
	// | CheckNumber
	// | CreditCardInfo
}

type Cash struct{}

func (c Cash) isPaymentMethod() {}

func (c CheckNumber) isPaymentMethod() {}

func (c CreditCardInfo) isPaymentMethod() {}

type PaymentAmount float32
type Currency uint8

const (
	EUR Currency = iota
	YEN
	USD
)

type Payment struct {
	PaymentAmount
	Currency
	PaymentMethod
}

func main() {
	p1 := Payment{
		PaymentAmount: 1234.2,
		Currency:      EUR,
		PaymentMethod: CheckNumber(123421),
	}
	p2 := Payment{
		PaymentAmount: 12000,
		Currency:      YEN,
	}
	fmt.Println("payment:", p1, p2)
}
