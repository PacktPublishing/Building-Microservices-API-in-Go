package dto

import (
	"github.com/ashishjuyal/banking-lib/errs"
	"strings"
)

// New account request details
// swagger:model newAccountRequest
type NewAccountRequest struct {
	// Customer Id
	CustomerId string `json:"customer_id"`
	// Type of Account: saving or checking
	AccountType string `json:"account_type"`
	// Account opening amount
	Amount float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("To open a new account you need to deposit atleast 5000.00")
	}
	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidationError("Account type should be checking or saving")
	}
	return nil
}
