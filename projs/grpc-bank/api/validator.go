package api

import (
	"example.xyz/bank/internal/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	currency, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	if _, err := util.CurrencyString(currency); err != nil {
		return false
	}
	return true
}
