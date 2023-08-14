package main

import (
	"encoding/json"
	"paypal/model"
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
	res, err := client.DisputesList(req)
	if err != nil {
		t.Log(err.Error())
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}

func TestDisputesDetail(t *testing.T) {
	req := &model.ReqDisputesDetail{Id: ""}
	res, err := client.DisputesDetail(req)
	if err != nil {
		t.Log(err.Error())
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}
