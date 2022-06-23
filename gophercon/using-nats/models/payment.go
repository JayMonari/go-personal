package models

import (
	"fmt"
	"math/rand"
)

const maxRandomAmount = 500
const currency = "GBP"

type Payment struct {
	id       int
	amount   int
	currency string
	sender   string
	receiver string
}

func (p Payment) String() string {
	return fmt.Sprintf("[%d]:%s sent a payment of %d %s to %s", p.id, p.sender,
		p.amount, p.currency, p.receiver)
}

func GetRandomPayment(id int) string {
	return Payment{
		id:       id,
		amount:   rand.Intn(maxRandomAmount),
		currency: currency,
		sender:   fmt.Sprintf("SENDER%d", rand.Intn(maxRandomAmount)),
		receiver: fmt.Sprintf("RECEIVER%d", rand.Intn(maxRandomAmount)),
	}.String()
}
