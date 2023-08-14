package main

import (
	"encoding/json"
	"github/mm-ooto/paypal/model"
	"testing"
)

func TestGetTransactionsList(t *testing.T) {
	req := &model.ReqGetTransactionsList{
		QueryParameters: &model.GetTransactionsListQueryParameters{
			TransactionId:               "",
			TransactionType:             "",
			TransactionStatus:           "",
			TransactionCurrency:         "",
			StartDate:                   "",
			EenDate:                     "",
			PaymentInstrumentType:       "",
			StoreId:                     "",
			TerminalId:                  "",
			Fields:                      "",
			BalanceAffectingRecordsOnly: "",
			Page:                        1,
			PageSize:                    2,
		},
	}
	res, err := client.GetTransactionsList(req)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}

func TestGetListAllBalances(t *testing.T) {
	req := &model.ReqGetListAllBalances{
		QueryParameters: &model.GetListAllBalancesQueryParameters{
			AsOfTime:     "ALL",
			CurrencyCode: "",
		}}
	res, err := client.GetListAllBalances(req)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}
