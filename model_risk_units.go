/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type RiskUnits struct {
	// Risk unit flag
	Symbol string `json:"symbol,omitempty"`
	// Spot hedging utilization
	SpotInUse string `json:"spot_in_use,omitempty"`
	// Maintenance margin for risk unit
	MaintainMargin string `json:"maintain_margin,omitempty"`
	// Initial margin for risk unit
	InitialMargin string `json:"initial_margin,omitempty"`
	// Total Delta of risk unit
	Delta string `json:"delta,omitempty"`
	// Total Gamma of risk unit
	Gamma string `json:"gamma,omitempty"`
	// Total Theta of risk unit
	Theta string `json:"theta,omitempty"`
	// Total Vega of risk unit
	Vega string `json:"vega,omitempty"`
}