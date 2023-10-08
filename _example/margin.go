package main

import (
	"context"
	"math/rand"
	"strings"

	"github.com/antihax/optional"
	"github.com/gateio/gateapi-go/v5"
	"github.com/shopspring/decimal"
)

func MarginDemo(config *RunConfig) {
	currencyPair := "BTC_USDT"
	currency := strings.Split(currencyPair, "_")[1]

	client := gateapi.NewAPIClient(gateapi.NewConfiguration())
	client.ChangeBasePath(config.BaseUrl)
	ctx := context.WithValue(context.Background(), gateapi.ContextGateAPIV4, gateapi.GateAPIV4{
		Key:    config.ApiKey,
		Secret: config.ApiSecret,
	})

	tickers, _, err := client.SpotApi.ListTickers(ctx, &gateapi.ListTickersOpts{
		CurrencyPair: optional.NewString(currencyPair),
	})
	if err != nil {
		panicGateError(err)
	}
	lastPrice := tickers[0].Last
	logger.Printf("currency pair %s last price %s\n", currencyPair, lastPrice)

	pairs, _, err := client.MarginApi.ListMarginCurrencyPairs(ctx)
	if err != nil {
		panicGateError(err)
	}
	var pair gateapi.MarginCurrencyPair
	for _, p := range pairs {
		if p.Id == currencyPair {
			pair = p
			break
		}
	}
	var loanAmount decimal.Decimal
	if pair.MinQuoteAmount == "" {
		loanAmount = decimal.Zero
	} else {
		loanAmount, err = decimal.NewFromString(pair.MinQuoteAmount)
		if err != nil {
			panicGateError(err)
		}
	}

	decimal.DivisionPrecision = 8

	// example to lend
	fundingAccounts, _, err := client.MarginApi.ListFundingAccounts(ctx, &gateapi.ListFundingAccountsOpts{
		Currency: optional.NewString(currency),
	})
	if err != nil {
		panicGateError(err)
	}
	lendAmount := loanAmount.Add(decimal.NewFromFloat32(rand.Float32())).Round(8)
	if len(fundingAccounts) == 1 &&
		lendAmount.LessThanOrEqual(decimal.RequireFromString(fundingAccounts[0].Available)) {
		lendingLoan := gateapi.Loan{
			Amount:    lendAmount.String(),
			AutoRenew: false,
			Days:      10,
			Currency:  currency,
			Rate:      "0.002",
			Side:      "lend",
		}
		createdLoan, _, err := client.MarginApi.CreateLoan(ctx, lendingLoan)
		if err != nil {
			panicGateError(err)
		}
		logger.Printf("place a lending loan %s with currency %s, rate %s, amount %s\n",
			createdLoan.Id, createdLoan.Currency, createdLoan.Rate, createdLoan.Amount)
		loanResult, _, err := client.MarginApi.GetLoan(ctx, createdLoan.Id, "lend")
		if err != nil {
			panicGateError(err)
		}
		if loanResult.Status == "loaned" {
			records, _, err := client.MarginApi.ListLoanRecords(ctx, loanResult.Id, nil)
			if err != nil {
				panicGateError(err)
			}
			for _, r := range records {
				logger.Printf("loan %s is borrowed with record id %s, amount %s, status: %s\n",
					r.LoanId, r.Id, r.Amount, r.Status)
			}
		} else {
			_, _, err := client.MarginApi.CancelLoan(ctx, createdLoan.Id, currency)
			if err != nil {
				panicGateError(err)
			}
		}
	}

	margin := loanAmount.DivRound(decimal.NewFromInt32(pair.Leverage-1), 8)
	accounts, _, err := client.MarginApi.ListMarginAccounts(ctx, &gateapi.ListMarginAccountsOpts{
		CurrencyPair: optional.NewString(currencyPair),
	})
	if err != nil {
		panicGateError(err)
	}
	available := decimal.RequireFromString(accounts[0].Quote.Available)
	logger.Printf("available margin balance of currency %s in currency pair %s: %s\n",
		currency, currencyPair, available.String())
	if margin.GreaterThan(available) {
		transfer := gateapi.Transfer{
			Currency:     currency,
			From:         "spot",
			To:           "margin",
			Amount:       margin.Sub(available).Round(8).String(),
			CurrencyPair: currencyPair,
		}
		_, err := client.WalletApi.Transfer(ctx, transfer)
		if err != nil {
			panicGateError(err)
		}
		logger.Printf("transferred %s %s to margin account\n", transfer.Amount, transfer.Amount)
	}

	// borrow with minimum amount
	borrowAmount := loanAmount.Add(decimal.NewFromFloat32(rand.Float32())).Round(8)
	fundingBook, _, err := client.MarginApi.ListFundingBook(ctx, currency)
	if err != nil {
		panicGateError(err)
	}
	var minRateItem *gateapi.FundingBookItem
	for _, item := range fundingBook {
		if item.Rate != "" && decimal.RequireFromString(item.Amount).GreaterThan(borrowAmount) {
			if minRateItem == nil {
				minRateItem = &item
			} else {
				if decimal.RequireFromString(item.Rate).LessThan(decimal.RequireFromString(minRateItem.Rate)) {
					minRateItem = &item
				}
			}
		}
	}
	if minRateItem == nil {
		minRateItem = &gateapi.FundingBookItem{Rate: "0.002", Days: 10}
	}
	loan := gateapi.Loan{
		Side:         "borrow",
		CurrencyPair: currencyPair,
		Currency:     currency,
		Rate:         minRateItem.Rate,
		Amount:       borrowAmount.String(),
		Days:         minRateItem.Days,
	}

	borrowed, _, err := client.MarginApi.CreateLoan(ctx, loan)
	if err != nil {
		panicGateError(err)
	}
	logger.Printf("borrowed %s %s in currency pair %s with rate %s, id %s\n",
		borrowed.Amount, borrowed.Currency, borrowed.CurrencyPair, borrowed.Rate, borrowed.Id)

	// create margin order
	cp, _, err := client.SpotApi.GetCurrencyPair(ctx, currencyPair)
	if err != nil {
		panicGateError(err)
	}
	orderAmount := cp.MinQuoteAmount
	if orderAmount == "" {
		orderAmount = "0"
	}
	order := gateapi.Order{
		Account:      "margin",
		CurrencyPair: currencyPair,
		Price:        lastPrice,
		Amount:       orderAmount,
		Side:         "sell",
	}
	createdOrder, _, err := client.SpotApi.CreateOrder(ctx, order)
	if err != nil {
		logger.Printf("failed to create margin order: %v\n", err)
	} else {
		logger.Printf("margin order created with id %s, status %s\n", createdOrder.Id, createdOrder.Status)
	}

	repayRequest := gateapi.RepayRequest{
		CurrencyPair: currencyPair,
		Currency:     currency,
		Mode:         "all",
	}
	_, _, err = client.MarginApi.RepayLoan(ctx, borrowed.Id, repayRequest)
	if err != nil {
		panicGateError(err)
	}
	repayments, _, err := client.MarginApi.ListLoanRepayments(ctx, borrowed.Id)
	if err != nil {
		panicGateError(err)
	}
	for _, r := range repayments {
		logger.Printf("loan %s repaid %s with interest %s", borrowed.Id, r.Principal, r.Interest)
	}
}
