package main

import "fmt"

type CardNumber string
type CardScheme uint8

const (
	Visa CardScheme = iota
	MasterCard
	UnionPay
	Verve_Card
)

func (ct CardScheme) String() string {
	switch ct {
	case Visa:
		return "Visa"
	case MasterCard:
		return "MasterCard"
	case UnionPay:
		return "UnionPay"
	case Verve_Card:
		return "Verve_Card"
	default:
		return ""
	}
}

type CreditCardInfo struct {
	CardScheme
	CardNumber
}

type CheckNumber int
type Cash struct{}

func (c Cash) String() string {
	return fmt.Sprintf("Paid in cash")
}

func (cn CheckNumber) String() string {
	return fmt.Sprintf("Paid By Check: %d\n", cn)
}

func (cci CreditCardInfo) String() string {
	return fmt.Sprintf("Paid with %s %s", cci.CardScheme, cci.CardNumber)
}

type PaymentMethod interface {
	Cash | CheckNumber | CreditCardInfo
}

//////////////////////////////////////////
type Currency uint8

const (
	EUR Currency = iota
	YEN
	USD
)

type PaymentAmount float32

type Payment[pm PaymentMethod] struct {
	PaymentAmount
	Currency
	Method pm
}

func printPayment[T PaymentMethod](pm T) {
	fmt.Println(pm)
}

func main() {
	fmt.Println(1^0, 0^1, 1^1, 0^0)
	var cNo CheckNumber = 123
	printPayment(cNo)
	cash := Cash{}
	cci := CreditCardInfo{CardScheme: Visa, CardNumber: "38290029 02190"}
	printPayment(cash)
	printPayment(cci)
	pCash := Payment[Cash]{PaymentAmount: 12342.3, Currency: EUR}
	pCN := Payment[CheckNumber]{PaymentAmount: 12342.3, Currency: EUR}
	pCCI := Payment[CreditCardInfo]{PaymentAmount: 12342.3, Currency: EUR}
	fmt.Printf("%#v\n", pCash)
	fmt.Printf("%#v\n", pCN)
	fmt.Printf("%#v\n", pCCI)
}
