package example

import (
	"encoding/json"
	"github/mm-ooto/paypal/model"
	"testing"
)

func TestGetAuthorizedPaymentDetails(t *testing.T) {
	req := &model.ReqGetAuthorizedPaymentDetails{
		AuthorizationId: "16462243JA7662122",
	}
	res, err := pClient.GetAuthorizedPaymentDetails(req)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}

func TestCaptureAuthorizedPayment(t *testing.T) {
	req := &model.ReqCaptureAuthorizedPayment{
		AuthorizationId: "16462243JA7662122",
	}
	res, err := pClient.CaptureAuthorizedPayment(req)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}

func TestReauthorizeAuthorizedPayment(t *testing.T) {
	req := &model.ReqReauthorizeAuthorizedPayment{
		AuthorizationId: "16462243JA7662122",
	}
	res, err := pClient.ReauthorizeAuthorizedPayment(req)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}

func TestGetCapturedPaymentDetails(t *testing.T) {
	req := &model.ReqGetCapturedPaymentDetails{
		CaptureId: "16462243JA7662122",
	}
	res, err := pClient.GetCapturedPaymentDetails(req)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}

func TestRefundCapturedPayment(t *testing.T) {
	req := &model.ReqRefundCapturedPayment{
		CaptureId:   "",
		Amount:      nil,
		InvoiceId:   "",
		NoteToPayer: "",
	}
	res, err := pClient.RefundCapturedPayment(req)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}

func TestGetRefundDetails(t *testing.T) {
	req := &model.ReqGetRefundDetails{
		RefundId: "",
	}
	res, err := pClient.GetRefundDetails(req)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}
