/*
 * Gate API v4
 *
 * Welcome to Gate API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type LedgerRecord struct {
	// Record ID
	Id string `json:"id,omitempty"`
	// Hash record of the withdrawal
	Txid string `json:"txid,omitempty"`
	// User-defined order number when withdrawing. Default is empty. When not empty, the specified user-defined order number record will be queried
	WithdrawOrderId string `json:"withdraw_order_id,omitempty"`
	// Operation time
	Timestamp string `json:"timestamp,omitempty"`
	// Currency amount
	Amount string `json:"amount"`
	// Currency name
	Currency string `json:"currency"`
	// Withdrawal address. Required for withdrawals
	Address string `json:"address,omitempty"`
	// Additional remarks with regards to the withdrawal
	Memo string `json:"memo,omitempty"`
	// The withdrawal record id starts with w, such as: w1879219868. When withdraw_id is not empty, the value querys this withdrawal record and no longer querys according to time
	WithdrawId string `json:"withdraw_id,omitempty"`
	// The currency type of withdrawal record is empty by default. It supports users to query the withdrawal records in the main and innovation areas on demand. Value range: SPOT, PILOT  SPOT: Main Zone  PILOT: Innovation Zone
	AssetClass string `json:"asset_class,omitempty"`
	// Record status.  - DONE: done - CANCEL: cancelled - REQUEST: requesting - MANUAL: pending manual approval - BCODE: GateCode operation - EXTPEND: pending confirm after sending - FAIL: pending confirm when fail - INVALID: invalid order - VERIFY: verifying - PROCES: processing - PEND: pending - DMOVE: required manual approval - REVIEW: Under review
	Status string `json:"status,omitempty"`
	// Name of the chain used in withdrawals
	Chain string `json:"chain"`
}
