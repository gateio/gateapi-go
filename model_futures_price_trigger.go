/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type FuturesPriceTrigger struct {
	// Trigger Policy   - 0: Price trigger, that is, when the price meets the conditions  - 1: Price spread trigger, i.e. the last price specified in `price_type` minus the second-last price difference At present, only 0 is the latest transaction price
	StrategyType int32 `json:"strategy_type,omitempty"`
	// Price type. 0 - latest deal price, 1 - mark price, 2 - index price.
	PriceType int32 `json:"price_type,omitempty"`
	// Value of price on price triggered, or price gap on price gap triggered.
	Price string `json:"price,omitempty"`
	// Price Condition Type  - 1: Indicates that the price calculated based on `strategy_type` and `price_type` is greater than or equal to `Trigger.Price` Trigger, while Trigger.Price must > last_price - based on `strategy_type` and `price_type` is less than or equal to `Trigger.Price` Trigger, and Trigger.Price must < last_price
	Rule int32 `json:"rule,omitempty"`
	// How long (in seconds) to wait for the condition to be triggered before cancelling the order.
	Expiration int32 `json:"expiration,omitempty"`
}
