package controllers

import (
	"fmt"
)

func ValidateTypeProperty(property string) error {
	if property == "DEPOSIT" || property == "WITHDRAWN" || property == "TRADE" {
		return nil
	}
	return fmt.Errorf("Invalid property: %s", property)
}
func ValidateReceivedCurrency(property string) error {
	if property != "" {
		return nil
	}
	return fmt.Errorf("Invalid property: %s", property)
}
