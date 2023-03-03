package main

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/antihax/optional"
	"github.com/gateio/gateapi-go/v6"
	"github.com/shopspring/decimal"
)

func FuturesDemo(config *RunConfig) {
	settle := "usdt"
	contract := "BTC_USDT"

	client := gateapi.NewAPIClient(gateapi.NewConfiguration())
	client.ChangeBasePath(config.BaseUrl)
	ctx := context.WithValue(context.Background(), gateapi.ContextGateAPIV4, gateapi.GateAPIV4{
		Key:    config.ApiKey,
		Secret: config.ApiSecret,
	})

	// update position leverage
	leverage := "3"
	_, _, err := client.FuturesApi.UpdatePositionLeverage(ctx, settle, contract, leverage, nil)
	if err != nil {
		panicGateError(err)
	}

	// retrieve position size
	positionSize := int64(0)
	position, _, err := client.FuturesApi.GetPosition(ctx, settle, contract)
	if err != nil {
		if e, ok := err.(gateapi.GateAPIError); ok {
			// ignore position not found error
			if e.Label != "POSITION_NOT_FOUND" {
				panicGateError(e)
			}
		} else {
			panicGateError(err)
		}
	} else {
		positionSize = position.Size
	}

	futuresContract, _, err := client.FuturesApi.GetFuturesContract(ctx, settle, contract)
	if err != nil {
		panicGateError(err)
	}
	orderSize := int64(10)
	if orderSize < futuresContract.OrderSizeMin {
		orderSize = futuresContract.OrderSizeMin
	}
	if positionSize < 0 {
		orderSize = 0 - orderSize
	}

	// example to update risk limit
	riskLimit := decimal.RequireFromString(futuresContract.RiskLimitBase).Add(decimal.RequireFromString(futuresContract.RiskLimitStep))
	_, _, err = client.FuturesApi.UpdatePositionRiskLimit(ctx, settle, contract, riskLimit.String())
	if err != nil {
		panicGateError(err)
	}

	// retrieve last price to calculate margin needed
	tickers, _, err := client.FuturesApi.ListFuturesTickers(ctx, settle, &gateapi.ListFuturesTickersOpts{
		Contract: optional.NewString(contract),
	})
	lastPrice := tickers[0].Last
	logger.Printf("last price of contract %s: %s", contract, lastPrice)

	margin := decimal.NewFromInt(orderSize).Mul(
		decimal.RequireFromString(lastPrice)).Mul(
		decimal.RequireFromString(futuresContract.QuantoMultiplier)).DivRound(
		decimal.RequireFromString(leverage), 8).Mul(
		decimal.RequireFromString("1.1")).Round(8)
	logger.Printf("needs margin amount: %s\n", margin.String())

	// if balance is not enough, transfer from spot account
	available := "0"
	futuresAccount, _, err := client.FuturesApi.ListFuturesAccounts(ctx, settle)
	if err != nil {
		if e, ok := err.(gateapi.GateAPIError); ok {
			if e.Label != "USER_NOT_FOUND" {
				panicGateError(e)
			}
		} else {
			panicGateError(err)
		}
	} else {
		available = futuresAccount.Available
	}
	logger.Printf("futures account available: %s %s\n", available, strings.ToUpper(settle))
	if margin.GreaterThan(decimal.RequireFromString(available)) {
		if config.UseTestNet {
			logger.Fatal("testnet account balance not enough, make a transferal on web")
		}
		transfer := gateapi.Transfer{
			Currency: strings.ToUpper(settle),
			From:     "spot",
			To:       "futures",
			Amount:   margin.String(),
			Settle:   settle,
		}
		tx, _, err := client.WalletApi.Transfer(ctx, transfer)
		if err != nil {
			panicGateError(err)
		}
		logger.Printf("transferred %s %s to futures %s account, transfer ID %d\n", transfer.Amount, transfer.Currency, transfer.Settle, tx.TxId)
	}

	// example to cancel all open orders in contract
	_, _, err = client.FuturesApi.CancelFuturesOrders(ctx, settle, contract, nil)
	if err != nil {
		panicGateError(err)
	}

	// order using market price
	order := gateapi.FuturesOrder{Contract: contract, Size: orderSize, Price: "0", Tif: "ioc"}
	orderResponse, _, err := client.FuturesApi.CreateFuturesOrder(ctx, settle, order)
	if err != nil {
		panicGateError(err)
	}
	logger.Printf("order %d created with status: %s\n", orderResponse.Id, orderResponse.Status)
	if orderResponse.Status == "open" {
		futuresOrder, _, err := client.FuturesApi.GetFuturesOrder(ctx, settle, strconv.FormatInt(orderResponse.Id, 10))
		if err != nil {
			panicGateError(err)
		}
		logger.Printf("order %d status %s, total size %d, left %d\n",
			futuresOrder.Id, futuresOrder.Status, futuresOrder.Size, futuresOrder.Left)
		_, _, err = client.FuturesApi.CancelFuturesOrder(ctx, settle, strconv.FormatInt(orderResponse.Id, 10))
		if err != nil {
			panicGateError(err)
		}
		logger.Printf("order %d cancelled\n", orderResponse.Id)
	} else {
		time.Sleep(time.Millisecond * 200)
		orderTrades, _, err := client.FuturesApi.GetMyTrades(ctx, settle, &gateapi.GetMyTradesOpts{
			Contract: optional.NewString(contract),
			Order:    optional.NewInt64(orderResponse.Id),
		})
		if err != nil {
			panicGateError(err)
		}
		for _, t := range orderTrades {
			logger.Printf("order %s filled size %d with price %s\n", t.OrderId, t.Size, t.Price)
		}
	}

	// example to update position margin
	_, _, err = client.FuturesApi.UpdatePositionMargin(ctx, settle, contract, "0.01")
	if err != nil {
		panicGateError(err)
	}
}
