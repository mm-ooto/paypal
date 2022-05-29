package main

import (
	"encoding/json"
	"github/mm-ooto/paypal"
	"github/mm-ooto/paypal/consts"
	"github/mm-ooto/paypal/model"
	"testing"
)

var pClient *paypal.PClient

func TestMain(m *testing.M) {
	clientId := "AcSU4Cbl7gbGZqI1xTNiJwkkpEzLpb0QntUmotlM239mz_8N5CYJlGfbByidGWXuoo6lDI06hgGfpyy5"
	clientSecret := "EBOvatSnBAMBls3YdsFKzvuXljxFxb3sX9ZNXZNkhJpdyLS_IKIbVV3Xkh04JAmDE2meJfeMO4ZMnRPS"
	pClient = paypal.NewPClient(clientId, clientSecret, false)
	m.Run()
}

func TestGetAccessToken(t *testing.T) {
	err := pClient.GetAccessToken()
	if err != nil {
		t.Log(err)
		return
	}
	bytes, _ := json.Marshal(pClient)
	t.Log(string(bytes))
}

func TestCreateOrder(t *testing.T) {
	req := &model.ReqCreateOrder{
		Intent: consts.IntentCapture,
		PurchaseUnits: []*model.PurchaseUnitsRequest{
			{
				Amount: &model.AmountWithBreakdown{
					CurrencyCode: "CAD",
					Value:        "10.00",
				},
			},
		},
	}
	res, err := pClient.CreateOrder(req)
	if err != nil {
		t.Log(err.Error())
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}

func TestShowOrderDetails(t *testing.T) {
	req := &model.ReqShowOrderDetails{
		Id: "16462243JA7662122",
	}
	res, err := pClient.ShowOrderDetails(req)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}

func TestOrderAuthorize(t *testing.T) {
	req := &model.ReqOrderAuthorize{
		Id: "16462243JA7662122",
	}
	res, err := pClient.OrderAuthorize(req)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}

func TestOrdersCapture(t *testing.T) {
	req := &model.ReqOrdersCapture{
		Id: "16462243JA7662122",
	}
	res, err := pClient.OrdersCapture(req)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}

func TestConfirmOrder(t *testing.T) {
	req := &model.ReqConfirmOrder{
		Id: "16462243JA7662122",
		PaymentSource: &model.PaymentSource{
			Token:      nil,
			Bancontact: nil,
			Blik:       nil,
			Eps:        nil,
			Giropay:    nil,
			Ideal:      model.IdealRequest{},
			Mybank:     nil,
			P24:        nil,
			Sofort:     nil,
			Trustly:    nil,
		},
	}
	res, err := pClient.ConfirmOrder(req)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}
