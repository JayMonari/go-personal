package util

type Currency uint

//go:generate go run github.com/dmarkham/enumer -type=Currency -trimprefix=Currency -json -sql
const (
	CurrencyUnknown Currency = iota
	CurrencyUSD
	CurrencyEUR
	CurrencyCAD
)
