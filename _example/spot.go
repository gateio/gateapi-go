package main

import (
	"context"

	"github.com/antihax/optional"
	"github.com/gateio/gateapi-go/v5"
	"github.com/shopspring/decimal"
)

func SpotDemo(config *RunConfig) {
	client := gateapi.NewAPIClient(gateapi.NewConfiguration())
	// Setting host is optional. It defaults to https://api.gateio.ws/api/v4
	client.ChangeBasePath(config.BaseUrl)
	ctx := context.WithValue(context.Background(), gateapi.ContextGateAPIV4, gateapi.GateAPIV4{
		Key:    config.ApiKey,
		Secret: config.ApiSecret,
	})

	currencyPair := "GT_USDT"
	currency := "USDT"
	cp, _, err := client.SpotApi.GetCurrencyPair(ctx, currencyPair)
	if err != nil {
		panicGateError(err)
	}
	logger.Printf("testing against currency pair: %s\n", cp.Id)
	minAmount := cp.MinBaseAmount

	tickers, _, err := client.SpotApi.ListTickers(ctx, &gateapi.ListTickersOpts{CurrencyPair: optional.NewString(cp.Id)})
	if err != nil {
		panicGateError(err)
	}
	lastPrice := tickers[0].Last

	// better avoid using float, take the following decimal library for example
	// `go get github.com/shopspring/decimal`
	orderAmount := decimal.RequireFromString(minAmount).Mul(decimal.NewFromInt32(2))

	balance, _, err := client.SpotApi.ListSpotAccounts(ctx, &gateapi.ListSpotAccountsOpts{Currency: optional.NewString(currency)})
	if err != nil {
		panicGateError(err)
	}
	if decimal.RequireFromString(balance[0].Available).Cmp(orderAmount) < 0 {
		logger.Fatal("balance not enough")
	}

	newOrder := gateapi.Order{
		Text:         "t-my-custom-id", // optional custom order ID
		CurrencyPair: cp.Id,
		Type:         "limit",
		Account:      "spot", // create spot order. set to "margin" if creating margin orders
		Side:         "buy",
		Amount:       orderAmount.String(),
		Price:        lastPrice, // use last price
		TimeInForce:  "gtc",
		AutoBorrow:   false,
	}
	logger.Printf("place a spot %s order in %s with amount %s and price %s\n", newOrder.Side, newOrder.CurrencyPair, newOrder.Amount, newOrder.Price)
	createdOrder, _, err := client.SpotApi.CreateOrder(ctx, newOrder)
	if err != nil {
		panicGateError(err)
	}
	logger.Printf("order created with ID: %s, status: %s\n", createdOrder.Id, createdOrder.Status)
	if createdOrder.Status == "open" {
		order, _, err := client.SpotApi.GetOrder(ctx, createdOrder.Id, createdOrder.CurrencyPair)
		if err != nil {
			panicGateError(err)
		}
		logger.Printf("order %s filled: %s, left: %s\n", order.Id, order.FilledTotal, order.Left)
		result, _, err := client.SpotApi.CancelOrder(ctx, createdOrder.Id, createdOrder.CurrencyPair)
		if err != nil {
			panicGateError(err)
		}
		if result.Status == "cancelled" {
			logger.Printf("order %s cancelled\n", createdOrder.Id)
		}
	} else {
		// order finished
		trades, _, err := client.SpotApi.ListMyTrades(ctx, createdOrder.CurrencyPair,
			&gateapi.ListMyTradesOpts{OrderId: optional.NewString(createdOrder.Id)})
		if err != nil {
			panicGateError(err)
		}
		for _, t := range trades {
			logger.Printf("order %s filled %s with price: %s\n", t.OrderId, t.Amount, t.Price)
		}
	}
}
