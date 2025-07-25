/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type CreateCollateralOrder struct {
	// Collateral amount.
	CollateralAmount string `json:"collateral_amount"`
	// Collateral.
	CollateralCurrency string `json:"collateral_currency"`
	// Borrowing amount.
	BorrowAmount string `json:"borrow_amount"`
	// Borrowed currency.
	BorrowCurrency string `json:"borrow_currency"`
}
