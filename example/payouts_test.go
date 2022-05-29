package main

import (
	"encoding/json"
	"github/mm-ooto/paypal/model"
	"testing"
)

func TestCreateBatchPayout(t *testing.T) {
	req := &model.ReqCreateBatchPayout{
		SenderBatchHeader: &model.SenderBatchHeader{
			SenderBatchId: "Payouts_2018_100007",
			EmailSubject:  "You have a payout!",
			EmailMessage:  "You have received a payout! Thanks for using our service!",
		},
		Items: []*model.PayoutItem{
			{
				RecipientType: "EMAIL",
				Amount: &model.Currency{
					Currency: "USD",
					Value:    "9.87",
				},
				Note:            "Thanks for your patronage!",
				Receiver:        "receiver@example.com",
				SenderItemId:    "201403140001",
				AlternateNotificationMethod: &model.AlternateNotificationMethod{
					Phone: &model.Phone{
						CountryCode:     "91",
						NationalNumber:  "9999988888",
					}},
				NotificationLanguage: "fr-FR",
				ApplicationContext: &model.ApplicationContext{
					SocialFeedPrivacy: "",
					HollerUrl:         "",
					LogoUrl:           "",
				},
			},
		},
	}
	res, err := pClient.CreateBatchPayout(req)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}

func TestGetPayoutBatchDetails(t *testing.T) {
	req := &model.ReqGetPayoutBatchDetails{
		PayoutBatchId: "N6YCKGJ9XU4V4",
		QueryParameters: &model.GetPayoutBatchDetailsQueryParameters{
			Fields:        "",
			Page:          0,
			PageSize:      0,
			TotalRequired: true,
		},
	}
	res, err := pClient.GetPayoutBatchDetails(req)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}

func TestGetPayoutItemDetails(t *testing.T) {
	req := &model.ReqGetPayoutItemDetails{
		PayoutItemId: "8AELMXH8UB2P8",
	}
	res, err := pClient.GetPayoutItemDetails(req)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}

func TestCancelUnclaimedPayoutItem(t *testing.T) {
	req := &model.ReqCancelUnclaimedPayoutItem{
		PayoutItemId: "8AELMXH8UB2P8",
	}
	res, err := pClient.CancelUnclaimedPayoutItem(req)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}
