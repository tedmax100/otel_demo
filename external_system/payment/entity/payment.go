package entity

import "time"

type Payment struct {
	PaymentId   uint64
	PayerName   string
	OrderId     uint64
	Amount      uint
	PaymentTime time.Time
}

func (p *Payment) Pay() bool {
	return true
}
