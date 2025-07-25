/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type BrokerCommission1 struct {
	// Commission Time. (unix timestamp).
	CommissionTime int64 `json:"commission_time,omitempty"`
	// User ID.
	UserId int64 `json:"user_id,omitempty"`
	// Group name.
	GroupName string `json:"group_name,omitempty"`
	// The amount of commission rebates.
	Amount string `json:"amount,omitempty"`
	// Fee.
	Fee string `json:"fee,omitempty"`
	// Fee currency.
	FeeAsset string `json:"fee_asset,omitempty"`
	// The income from rebates, converted to USDT.
	RebateFee string `json:"rebate_fee,omitempty"`
	// Rebate Type: Spot、Futures、Options.、Alpha
	Source string `json:"source,omitempty"`
	// Currency pair.
	CurrencyPair  string                        `json:"currency_pair,omitempty"`
	SubBrokerInfo BrokerCommissionSubBrokerInfo `json:"sub_broker_info,omitempty"`
	// Alpha token address
	AlphaContractAddr string `json:"alpha_contract_addr,omitempty"`
}
