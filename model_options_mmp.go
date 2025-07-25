/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// MMP Settings
type OptionsMmp struct {
	// Underlying.
	Underlying string `json:"underlying"`
	// Time window (milliseconds), between 1-5000, 0 means disabling MMP.
	Window int32 `json:"window"`
	// Freeze duration (milliseconds), 0 means always frozen, need to call reset API to unfreeze
	FrozenPeriod int32 `json:"frozen_period"`
	// Trading volume upper limit (positive number, up to 2 decimal places).
	QtyLimit string `json:"qty_limit"`
	// Upper limit of net delta value (positive number, up to 2 decimal places).
	DeltaLimit string `json:"delta_limit"`
	// Trigger freeze time (milliseconds), 0 means no freeze is triggered.
	TriggerTimeMs int64 `json:"trigger_time_ms,omitempty"`
	// Unfreeze time (milliseconds). If the freeze duration is not configured, there will be no unfreeze time after the freeze is triggered.
	FrozenUntilMs int64 `json:"frozen_until_ms,omitempty"`
}
