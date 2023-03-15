package main

import (
	"encoding/json"
	"github/mm-ooto/paypal/model"
	"testing"
)

func TestDisputesList(t *testing.T) {
	req := &model.ReqDisputesList{
		DisputeState:          "",
		DisputedTransactionId: "",
		NextPageToken:         "",
		PageSize:              0,
		StartTime:             "",
		UpdateTimeAfter:       "",
		UpdateTimeBefore:      "",
	}
	res, err := pClient.DisputesList(req)
	if err != nil {
		t.Log(err.Error())
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}

func TestDisputesDetail(t *testing.T) {
	req := &model.ReqDisputesDetail{Id: ""}
	res, err := pClient.DisputesDetail(req)
	if err != nil {
		t.Log(err.Error())
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}
