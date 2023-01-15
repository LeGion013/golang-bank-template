package api

import "github.com/go-playground/validator/v10"

var validCurrency validator.Func = func(fieledLevel validator.FieldLevel) bool {
	if currency, ok := fieledLevel.Field().Interface().(string); ok {
		// check if currency si supported
		print(currency)
		return true
	}

	return false

}
