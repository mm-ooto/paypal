package model

import (
	"fmt"
	"net/url"
	"strconv"
)

// ResEmpty 空的响应
type ResEmpty struct {
}

// OtherDefaultHeaderParameters 其它默认的请求头参数
type OtherDefaultHeaderParameters struct {
	PayPalRequestId     string
	Prefer              string
	PayPalAuthAssertion string
}

// ReqCreateOrder
// https://developer.paypal.com/docs/api/orders/v2/#orders_create
type ReqCreateOrder struct {
	Intent                string                   `json:"intent"`
	ProcessingInstruction string                   `json:"processing_instruction,omitempty"`
	Payer                 *Payer                   `json:"payer,omitempty"`
	PurchaseUnits         []*PurchaseUnitsRequest  `json:"purchase_units"`
	ApplicationContext    *OrderApplicationContext `json:"application_context,omitempty"`
}

// ReqUpdateOrder
// https://developer.paypal.com/docs/api/orders/v2/#orders_patch
type ReqUpdateOrder struct {
	Id           string  `json:"-"`
	PatchRequest []Patch `json:"patch_request,omitempty"`
}

// ReqShowOrderDetails
// https://developer.paypal.com/docs/api/orders/v2/#orders_get
type ReqShowOrderDetails struct {
	Id              string                           `json:"-"`
	QueryParameters *ShowOrderDetailsQueryParameters `json:"-"`
}
type ShowOrderDetailsQueryParameters struct {
	Fields string `json:"fields,omitempty"`
}

func (r *ReqShowOrderDetails) GetQueryParameters() string {
	queryParams := r.QueryParameters
	if queryParams == nil {
		return ""
	}

	urlValues := url.Values{}
	if len(queryParams.Fields) > 0 {
		urlValues.Add("fields", queryParams.Fields)
	}
	return urlValues.Encode()
}

// ReqOrderAuthorize
// https://developer.paypal.com/docs/api/orders/v2/#orders_authorize
type ReqOrderAuthorize struct {
	Id            string                              `json:"-"`
	PaymentSource *OrderAuthorizeRequestPaymentSource `json:"payment_source,omitempty"`
}

// ReqOrdersCapture
// https://developer.paypal.com/docs/api/orders/v2/#orders_capture
type ReqOrdersCapture struct {
	Id            string                             `json:"-"`
	PaymentSource *OrdersCaptureRequestPaymentSource `json:"payment_source,omitempty"`
}

// ReqConfirmOrder
// https://developer.paypal.com/docs/api/orders/v2/#orders_confirm
type ReqConfirmOrder struct {
	Id                    string                          `json:"-"`
	PaymentSource         *PaymentSource                  `json:"payment_source"`
	ProcessingInstruction string                          `json:"processing_instruction,omitempty"`
	ApplicationContext    *OrderConfirmApplicationContext `json:"application_context,omitempty"`
}

// ReqGetAuthorizedPaymentDetails
// https://developer.paypal.com/docs/api/payments/v2/#authorizations_get
type ReqGetAuthorizedPaymentDetails struct {
	AuthorizationId string `json:"-"`
}

// ReqCaptureAuthorizedPayment
// https://developer.paypal.com/docs/api/payments/v2/#authorizations_capture
type ReqCaptureAuthorizedPayment struct {
	AuthorizationId    string              `json:"-"`
	InvoiceId          string              `json:"invoice_id,omitempty"`
	NoteToPayer        string              `json:"note_to_payer,omitempty"`
	Amount             *Money              `json:"amount,omitempty"`
	FinalCapture       bool                `json:"final_capture,omitempty"`
	PaymentInstruction *PaymentInstruction `json:"payment_instruction,omitempty"`
	SoftDescriptor     string              `json:"soft_descriptor,omitempty"`
}

// ReqReauthorizeAuthorizedPayment
// https://developer.paypal.com/docs/api/payments/v2/#authorizations_reauthorize
type ReqReauthorizeAuthorizedPayment struct {
	AuthorizationId string `json:"-"`
	Amount          *Money `json:"amount,omitempty"`
}

// ReqVoidAuthorizedPayment
// https://developer.paypal.com/docs/api/payments/v2/#authorizations_void
type ReqVoidAuthorizedPayment struct {
	AuthorizationId string `json:"-"`
}

// ReqGetCapturedPaymentDetails
// https://developer.paypal.com/docs/api/payments/v2/#captures_get
type ReqGetCapturedPaymentDetails struct {
	CaptureId string `json:"-"`
}

// ReqRefundCapturedPayment
// https://developer.paypal.com/docs/api/payments/v2/#captures_refund
type ReqRefundCapturedPayment struct {
	CaptureId   string `json:"-"`
	Amount      *Money `json:"amount,omitempty"`
	InvoiceId   string `json:"invoice_id,omitempty"`
	NoteToPayer string `json:"note_to_payer,omitempty"`
}

// ReqGetRefundDetails
// https://developer.paypal.com/docs/api/payments/v2/#refunds_get
type ReqGetRefundDetails struct {
	RefundId string `json:"-"`
}

// ReqCreateBatchPayout
// https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts-post-request-body
type ReqCreateBatchPayout struct {
	SenderBatchHeader *SenderBatchHeader `json:"sender_batch_header"`
	Items             []*PayoutItem      `json:"items"`
}

// ReqGetPayoutBatchDetails
// https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts_get
type ReqGetPayoutBatchDetails struct {
	PayoutBatchId   string                                `json:"-"`
	QueryParameters *GetPayoutBatchDetailsQueryParameters `json:"-"`
}

type GetPayoutBatchDetailsQueryParameters struct {
	Fields        string `json:"fields,omitempty"`
	Page          int    `json:"page,omitempty"`
	PageSize      int    `json:"page_size,omitempty"`
	TotalRequired bool   `json:"total_required"`
}

func (r *ReqGetPayoutBatchDetails) GetQueryParameters() string {
	queryParams := r.QueryParameters

	if queryParams == nil {
		return ""
	}
	urlValues := url.Values{}
	if len(queryParams.Fields) > 0 {
		urlValues.Add("fields", queryParams.Fields)
	}
	if queryParams.Page > 0 {
		urlValues.Add("page", strconv.Itoa(queryParams.Page))
	}
	if queryParams.PageSize > 0 {
		urlValues.Add("page_size", strconv.Itoa(queryParams.PageSize))
	}
	urlValues.Add("total_required", fmt.Sprintf("%t", queryParams.TotalRequired))
	return urlValues.Encode()
}

// ReqGetPayoutItemDetails
// https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts-item_get
type ReqGetPayoutItemDetails struct {
	PayoutItemId string `json:"-"`
}

// ReqCancelUnclaimedPayoutItem
// https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts-item_cancel
type ReqCancelUnclaimedPayoutItem struct {
	PayoutItemId string `json:"-"`
}

// ReqGetTransactionsList
// https://developer.paypal.com/docs/api/transaction-search/v1/#transactions_get
type ReqGetTransactionsList struct {
	QueryParameters *GetTransactionsListQueryParameters `json:"-"`
}
type GetTransactionsListQueryParameters struct {
	TransactionId               string `json:"transaction_id,omitempty"`
	TransactionType             string `json:"transaction_type,omitempty"`
	TransactionStatus           string `json:"transaction_status,omitempty"`
	TransactionCurrency         string `json:"transaction_currency,omitempty"`
	StartDate                   string `json:"start_date"`
	EenDate                     string `json:"een_date"`
	PaymentInstrumentType       string `json:"payment_instrument_type,omitempty"`
	StoreId                     string `json:"store_id,omitempty"`
	TerminalId                  string `json:"terminal_id,omitempty"`
	Fields                      string `json:"fields,omitempty"`
	BalanceAffectingRecordsOnly string `json:"balance_affecting_records_only,omitempty"`
	Page                        int    `json:"page,omitempty"`
	PageSize                    int    `json:"page_size,omitempty"`
}

func (r *ReqGetTransactionsList) GetQueryParameters() string {
	queryParams := r.QueryParameters
	if queryParams == nil {
		return ""
	}

	urlValues := url.Values{}
	if len(queryParams.TransactionId) > 0 {
		urlValues.Add("transaction_id", queryParams.TransactionId)
	}
	if len(queryParams.TransactionType) > 0 {
		urlValues.Add("transaction_type", queryParams.TransactionType)
	}
	if len(queryParams.TransactionStatus) > 0 {
		urlValues.Add("transaction_status", queryParams.TransactionStatus)
	}
	if len(queryParams.TransactionCurrency) > 0 {
		urlValues.Add("transaction_currency", queryParams.TransactionCurrency)
	}
	urlValues.Add("start_date", queryParams.StartDate)
	urlValues.Add("een_date", queryParams.EenDate)

	if len(queryParams.PaymentInstrumentType) > 0 {
		urlValues.Add("payment_instrument_type", queryParams.PaymentInstrumentType)
	}
	if len(queryParams.StoreId) > 0 {
		urlValues.Add("store_id", queryParams.StoreId)
	}
	if len(queryParams.TerminalId) > 0 {
		urlValues.Add("terminal_id", queryParams.TerminalId)
	}

	if len(queryParams.Fields) > 0 {
		urlValues.Add("fields", queryParams.Fields)
	}
	if len(queryParams.BalanceAffectingRecordsOnly) > 0 {
		urlValues.Add("balance_affecting_records_only", queryParams.BalanceAffectingRecordsOnly)
	}
	if queryParams.Page > 0 {
		urlValues.Add("page", strconv.Itoa(queryParams.Page))
	}
	if queryParams.PageSize > 0 {
		urlValues.Add("page_size", strconv.Itoa(queryParams.PageSize))
	}
	return urlValues.Encode()
}

// ReqGetListAllBalances
// https://developer.paypal.com/docs/api/transaction-search/v1/#balances_get
type ReqGetListAllBalances struct {
	QueryParameters *GetListAllBalancesQueryParameters `json:"-"`
}

type GetListAllBalancesQueryParameters struct {
	AsOfTime     string `json:"as_of_time,omitempty"`
	CurrencyCode string `json:"currency_code,omitempty"`
}

func (r *ReqGetListAllBalances) GetQueryParameters() string {
	queryParams := r.QueryParameters
	if queryParams == nil {
		return ""
	}
	urlValues := url.Values{}
	if len(queryParams.AsOfTime) > 0 {
		urlValues.Add("as_of_time", queryParams.AsOfTime)
	}
	if len(queryParams.CurrencyCode) > 0 {
		urlValues.Add("currency_code", queryParams.CurrencyCode)
	}
	return urlValues.Encode()
}

// ReqGetListWebhooks
// https://developer.paypal.com/docs/api/webhooks/v1/#webhooks_list
type ReqGetListWebhooks struct {
	QueryParameters *GetListWebhooksQueryParameters `json:"-"`
}

// GetListWebhooksQueryParameters
// https://developer.paypal.com/docs/api/webhooks/v1/#webhooks-list-query-parameters
type GetListWebhooksQueryParameters struct {
	AnchorType string `json:"anchor_type,omitempty"` // 可能的值:APPLICATION,ACCOUNT
}

func (r *ReqGetListWebhooks) GetQueryParameters() string {
	queryParams := r.QueryParameters
	if queryParams == nil {
		return ""
	}
	urlValues := url.Values{}
	if len(queryParams.AnchorType) > 0 {
		urlValues.Add("anchor_type", queryParams.AnchorType)
	}
	return urlValues.Encode()
}

// ReqCreateWebhook
// https://developer.paypal.com/docs/api/webhooks/v1/#webhooks_post
type ReqCreateWebhook struct {
	Id         string             `json:"id,omitempty"`
	Url        string             `json:"url"`
	EventTypes []*EventType       `json:"event_types"`
	Links      []*LinkDescription `json:"links,omitempty"`
}

// ReqDeleteWebhook
// https://developer.paypal.com/docs/api/webhooks/v1/#webhooks_delete
type ReqDeleteWebhook struct {
	WebhookId string `json:"-"`
}

// ReqUpdateWebhook
// https://developer.paypal.com/docs/api/webhooks/v1/#webhooks_update
type ReqUpdateWebhook struct {
	WebhookId    string   `json:"-"`
	PatchRequest []*Patch `json:"patch_request,omitempty"`
}

// ReqGetWebhookDetails
// https://developer.paypal.com/docs/api/webhooks/v1/#webhooks_get
type ReqGetWebhookDetails struct {
	WebhookId string `json:"-"`
}

// ReqGetListEventSubscriptionsForWebhook
// https://developer.paypal.com/docs/api/webhooks/v1/#webhooks_list
type ReqGetListEventSubscriptionsForWebhook struct {
	WebhookId string `json:"-"`
}

// ReqResendWebhookEventNotification 重发事件通知
// https://developer.paypal.com/docs/api/webhooks/v1/#webhooks-events_resend
type ReqResendWebhookEventNotification struct {
	EventId string `json:"event_id"` // 事件id
}

// ReqShowEventNotificationDetail 事件通知详情
// https://developer.paypal.com/docs/api/webhooks/v1/#webhooks-events_get
type ReqShowEventNotificationDetail struct {
	EventId string `json:"event_id"` // 事件id
}

// ReqVerifyWebhookSignature
// https://developer.paypal.com/docs/api/webhooks/v1/#verify-webhook-signature_post
type ReqVerifyWebhookSignature struct {
	AuthAlgo         string `json:"auth_algo"`         // PayPal用来生成签名的算法，以及你可以用来验证签名的算法。从PAYPAL-AUTH-ALGO响应头中提取该值，该响应头与webhook通知一起接收。
	CertUrl          string `json:"cert_url"`          // X.509公钥证书。从这个URL下载证书并使用它来验证签名。从PAYPAL-CERT-URL响应头中提取这个值，它与webhook通知一起接收。
	TransmissionId   string `json:"transmission_id"`   // HTTP传输ID。包含在通知消息的PAYPAL-TRANSMISSION-ID报头中。
	TransmissionSig  string `json:"transmission_sig"`  // paypal生成的非对称签名。出现在通知消息的PAYPAL-TRANSMISSION-SIG报头中。
	TransmissionTime string `json:"transmission_time"` // HTTP传输的日期和时间，以Internet日期和时间格式。出现在通知消息的PAYPAL-TRANSMISSION-TIME报头中。
	WebhookId        string `json:"webhook_id"`        // 在开发者门户帐户中配置的webhook的ID。
	WebhookEvent     *Event `json:"webhook_event"`     // 一个webhook事件通知。
}

// ReqDisputesList 争议列表
// https://developer.paypal.com/docs/api/customer-disputes/v1/#disputes_list
type ReqDisputesList struct {
	DisputeState          string `json:"dispute_state,omitempty"`           // 过滤状态响应中的争议。多个值之间用“，”分隔。当指定多个争议状态时，响应将列出属于任何指定的争议状态的争议
	DisputedTransactionId string `json:"disputed_transaction_id,omitempty"` // 根据事务(ID)过滤响应中的争议，可以指定start_time和argumted_transaction_id查询参数之一，但不能同时指定。
	NextPageToken         string `json:"next_page_token,omitempty"`         // 描述要获取的下一页结果的标记。列表争议调用在响应中的HATEOAS链接中返回此令牌。
	PageSize              int    `json:"page_size,omitempty"`               // 限制对此值的响应中的争议数量。
	StartTime             string `json:"start_time,omitempty"`              // 根据创建日期和时间过滤响应中的争议。开始时间必须在最近180天内。值为Internet日期和时间格式。例如:yyyy-MM-ddTHH:mm:ss.SSSZ。 可以指定start_time和argumted_transaction_id查询参数之一，但不能同时指定。
	UpdateTimeAfter       string `json:"update_time_after,omitempty"`       // 争议最后一次更新的日期和时间，以互联网日期和时间格式。例如:yyyy-MM-ddTHH:mm:ss.SSSZ。Update_time_after必须在最近180天内，默认是支持的最大时间(180天)。
	UpdateTimeBefore      string `json:"update_time_before,omitempty"`      // 争议最后一次更新的日期和时间，以互联网日期和时间格式。例如:yyyy-MM-ddTHH:mm:ss.SSSZ。Update_time_before必须在最近180天内，默认为当前时间。
}

// ReqDisputesDetail 争议详情
// https://developer.paypal.com/docs/api/customer-disputes/v1/#disputes_get
type ReqDisputesDetail struct {
	Id string `json:"id"` // 要显示其详细信息的争议的ID。
}
