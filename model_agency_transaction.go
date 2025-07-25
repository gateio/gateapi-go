/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type AgencyTransaction struct {
	// Transaction Time. (unix timestamp).
	TransactionTime int64 `json:"transaction_time,omitempty"`
	// User ID.
	UserId int64 `json:"user_id,omitempty"`
	// Group name.
	GroupName string `json:"group_name,omitempty"`
	// Fee.
	Fee string `json:"fee,omitempty"`
	// Fee currency.
	FeeAsset string `json:"fee_asset,omitempty"`
	// Currency pair.
	CurrencyPair string `json:"currency_pair,omitempty"`
	// Commission Amount.
	Amount string `json:"amount,omitempty"`
	// Commission Asset.
	AmountAsset string `json:"amount_asset,omitempty"`
	// Source. SPOT - SPOT Rebate, FUTURES - Futures Rebate.
	Source string `json:"source,omitempty"`
}
