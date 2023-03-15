package paypal

import (
	"encoding/json"
	"fmt"
	"github/mm-ooto/paypal/consts"
	"github/mm-ooto/paypal/model"
	"io/ioutil"
	"net/http"
)

/*
*	网络钩子：使用Webhooks进行事件通知
*    https://developer.paypal.com/docs/api/webhooks/v1/#webhooks
 */

// GetListWebhooks 列出应用程序的webhooks
func (p *PClient) GetListWebhooks(req *model.ReqGetListWebhooks) (res model.ResGetListWebhooks, err error) {
	api := fmt.Sprintf("%s?%s", consts.WebhooksAPI, req.GetQueryParameters())
	if err = p.CallApiRequest(consts.HttpMethodGet, api, nil, &res); err != nil {
		return
	}
	return
}

// CreateWebhook 创建webhook
func (p *PClient) CreateWebhook(req *model.ReqCreateWebhook) (res model.ResCreateWebhook, err error) {
	if err = p.CallApiRequest(consts.HttpMethodPost, consts.WebhooksAPI, req, &res); err != nil {
		return
	}
	return
}

// DeleteWebhook 删除指定的webhook
func (p *PClient) DeleteWebhook(req *model.ReqDeleteWebhook) (err error) {
	api := fmt.Sprintf("%s/%s", consts.WebhooksAPI, req.WebhookId)
	if err = p.CallApiRequest(consts.HttpMethodDelete, api, nil, nil); err != nil {
		return
	}
	return
}

// UpdateWebhook 更新指定的webhook
func (p *PClient) UpdateWebhook(req *model.ReqUpdateWebhook) (res model.ResUpdateWebhook, err error) {
	api := fmt.Sprintf("%s/%s", consts.WebhooksAPI, req.WebhookId)
	if err = p.CallApiRequest(consts.HttpMethodPatch, api, req, &res); err != nil {
		return
	}
	return
}

// GetWebhookDetails 通过ID显示webhook的详细信息
func (p *PClient) GetWebhookDetails(req *model.ReqGetWebhookDetails) (res model.ResGetWebhookDetails, err error) {
	api := fmt.Sprintf("%s/%s", consts.WebhooksAPI, req.WebhookId)
	if err = p.CallApiRequest(consts.HttpMethodGet, api, nil, &res); err != nil {
		return
	}
	return
}

// GetListEventSubscriptionsForWebhook 按ID列出一个webhook的事件订阅。
func (p *PClient) GetListEventSubscriptionsForWebhook(req *model.ReqGetListEventSubscriptionsForWebhook) (res model.ResGetListEventSubscriptionsForWebhook, err error) {
	api := fmt.Sprintf(consts.ListEventSubscriptionsForWebhookAPI, req.WebhookId)
	if err = p.CallApiRequest(consts.HttpMethodGet, api, nil, &res); err != nil {
		return
	}
	return
}

// ShowEventNotificationDetail 事件通知详情
func (p *PClient) ShowEventNotificationDetail(req *model.ReqShowEventNotificationDetail) (res model.Event, err error) {
	api := fmt.Sprintf(consts.ShowEventNotificationDetailAPI, req.EventId)
	if err = p.CallApiRequest(consts.HttpMethodGet, api, req, &res); err != nil {
		return
	}
	return
}

// ResendWebhookEventNotification 重发事件通知
func (p *PClient) ResendWebhookEventNotification(req *model.ReqResendWebhookEventNotification) (res model.ResResendWebhookEventNotification, err error) {
	api := fmt.Sprintf(consts.ResendWebhookEventNotificationAPI, req.EventId)
	if err = p.CallApiRequest(consts.HttpMethodPost, api, req, &res); err != nil {
		return
	}
	return
}

// VerifyWebhookSignature 验证webhook签名
func (p *PClient) VerifyWebhookSignature(req *model.ReqVerifyWebhookSignature) (res model.ResVerifyWebhookSignature, err error) {
	if err = p.CallApiRequest(consts.HttpMethodPost, consts.VerifyWebhookSignatureAPI, req, &res); err != nil {
		return
	}
	return
}

// HandleWebhookAsyncNotify 处理异步通知，处理成功后需要返回StatusCode=http.StatusOK
func (p *PClient) HandleWebhookAsyncNotify(webhookId string, req *http.Request) (webhookEvent *model.Event, err error) {

	requestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(requestBody, &webhookEvent); err != nil {
		return
	}

	// 验证签名
	verifyReq := &model.ReqVerifyWebhookSignature{
		AuthAlgo:         req.Header.Get("auth_algo"),
		CertUrl:          req.Header.Get("cert_url"),
		TransmissionId:   req.Header.Get("transmission_id"),
		TransmissionSig:  req.Header.Get("transmission_sig"),
		TransmissionTime: req.Header.Get("transmission_time"),
		WebhookId:        webhookId,
		WebhookEvent:     webhookEvent,
	}
	verifyRes, err := p.VerifyWebhookSignature(verifyReq)
	if err != nil {
		return
	}
	if verifyRes.VerificationStatus == consts.VerificationStatusFailure {
		err = signatureVerifyFailureErr
		return
	}

	// 签名验证成功后，根据不同的资源类型对数据进行处理
	switch webhookEvent.ResourceType {
	case consts.ResourceTypeDispute: // 纠纷
	// TODO
	case consts.ResourceTypeRefund: // 退款
	// TODO
	case consts.ResourceTypeSale: // 交易
	// TODO
	case consts.ResourceTypeInvoices: // 发票
	// TODO
	default:
		return
	}
	return
}
