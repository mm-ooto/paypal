package main

import (
	"encoding/json"
	"github/mm-ooto/paypal/model"
	"testing"
)

func TestGetListWebhooks(t *testing.T) {
	req := &model.ReqGetListWebhooks{
		QueryParameters: &model.GetListWebhooksQueryParameters{
			AnchorType: "APPLICATION",
		},
	}
	res, err := pClient.GetListWebhooks(req)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}

func TestCreateWebhook(t *testing.T) {
	req := &model.ReqCreateWebhook{
		Url: "https://example.com/example_webhook",
		EventTypes: []*model.EventType{
			{
				Name: "PAYMENT.AUTHORIZATION.CREATED",
			},
			{
				Name: "PAYMENT.AUTHORIZATION.VOIDED",
			},
		},
	}
	res, err := pClient.CreateWebhook(req)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}

func TestDeleteWebhook(t *testing.T) {
	req := &model.ReqDeleteWebhook{
		WebhookId: "09Y1313102281883H",
	}
	err := pClient.DeleteWebhook(req)
	if err != nil {
		t.Log(err)
		return
	}
}

// ...
func TestUpdateWebhook(t *testing.T) {
	req := &model.ReqUpdateWebhook{
		WebhookId: "9FM75813N28895501",
		PatchRequest: []*model.Patch{
			{
				Op:    "replace",
				Path:  "/url",
				Value: "https://example.com/example_webhook_2",
			},
			{
				Op:   "replace",
				Path: "/event_types",
				Value: []*model.EventType{
					{
						Name: "PAYMENT.AUTHORIZATION.VOIDED",
					},
				},
			},
		},
	}
	res, err := pClient.UpdateWebhook(req)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}

func TestGetWebhookDetails(t *testing.T) {
	req := &model.ReqGetWebhookDetails{
		WebhookId: "9FM75813N28895501",
	}
	res, err := pClient.GetWebhookDetails(req)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}

func TestGetListEventSubscriptionsForWebhook(t *testing.T) {
	req := &model.ReqGetListEventSubscriptionsForWebhook{
		WebhookId: "9FM75813N28895501",
	}
	res, err := pClient.GetListEventSubscriptionsForWebhook(req)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}
