/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type TransferOrderStatus struct {
	// Order id.
	TxId string `json:"tx_id,omitempty"`
	// Transfer status, PENDING - in process, SUCCESS - successful transfer, FAIL - failed transfer, PARTIAL_SUCCESS - Partially successful (this status will appear when transferring between sub-subs)
	Status string `json:"status,omitempty"`
}
