package model

// ResAccessToken 访问令牌响应
type ResAccessToken struct {
	Scope       string `json:"scope"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	AppId       string `json:"app_id"`
	ExpiresIn   int64  `json:"expires_in"`
	Nonce       string `json:"nonce"`
}

// ResOrderBase ResOrderBase
type ResOrderBase struct {
	Id                    string                `json:"id,omitempty"`
	PaymentSource         PaymentSourceResponse `json:"payment_source,omitempty"`
	Intent                string                `json:"intent,omitempty"`
	ProcessingInstruction string                `json:"processing_instruction,omitempty"`
	Payer                 Payer                 `json:"payer,omitempty"`
	PurchaseUnits         []PurchaseUnit        `json:"purchase_units,omitempty"`
	Status                string                `json:"status,omitempty"`
	Links                 []LinkDescription     `json:"links"`
	CreateTime            string                `json:"create_time,omitempty"`
	UpdateTime            string                `json:"update_time,omitempty"`
}

// ResCreateOrder
// https://developer.paypal.com/docs/api/orders/v2/#orders-create-response
type ResCreateOrder struct {
	ResOrderBase
}

// ResShowOrderDetails
// https://developer.paypal.com/docs/api/orders/v2/#orders-get-response
type ResShowOrderDetails struct {
	ResOrderBase
}

// ResOrderAuthorize
// https://developer.paypal.com/docs/api/orders/v2/#orders-authorize-response
type ResOrderAuthorize struct {
	ResOrderBase
}

// ResOrdersCapture
// https://developer.paypal.com/docs/api/orders/v2/#orders-capture-response
type ResOrdersCapture struct {
	ResOrderBase
}

// ResConfirmOrder
// https://developer.paypal.com/docs/api/orders/v2/#orders_confirm-response
type ResConfirmOrder struct {
	ResOrderBase
}

// ResGetAuthorizedPaymentDetails
// https://developer.paypal.com/docs/api/payments/v2/#authorizations_get-response
type ResGetAuthorizedPaymentDetails struct {
	Status           string                     `json:"status,omitempty"`
	StatusDetails    AuthorizationStatusDetails `json:"status_details,omitempty"`
	Id               string                     `json:"id,omitempty"`
	Amount           Money                      `json:"amount,omitempty"`
	InvoiceId        string                     `json:"invoice_id,omitempty"`
	CustomId         string                     `json:"custom_id,omitempty"`
	SellerProtection SellerProtection           `json:"seller_protection,omitempty"`
	ExpirationTime   string                     `json:"expiration_time,omitempty"`
	Links            []LinkDescription          `json:"links,omitempty"`
	CreateTime       string                     `json:"create_time,omitempty"`
	UpdateTime       string                     `json:"update_time,omitempty"`
}

// ResCaptureAuthorizedPayment
// https://developer.paypal.com/docs/api/payments/v2/#authorizations_capture-response
type ResCaptureAuthorizedPayment struct {
	Status                    string                     `json:"status,omitempty"`
	StatusDetails             AuthorizationStatusDetails `json:"status_details,omitempty"`
	Id                        string                     `json:"id,omitempty"`
	Amount                    Money                      `json:"amount,omitempty"`
	InvoiceId                 string                     `json:"invoice_id,omitempty"`
	CustomId                  string                     `json:"custom_id,omitempty"`
	SellerProtection          SellerProtection           `json:"seller_protection,omitempty"`
	FinalCapture              bool                       `json:"final_capture,omitempty"`
	SellerReceivableBreakdown SellerReceivableBreakdown  `json:"seller_receivable_breakdown,omitempty"`
	DisbursementMode          string                     `json:"disbursement_mode,omitempty"`
	Links                     []LinkDescription          `json:"links,omitempty"`
	ProcessorResponse         ProcessorResponse          `json:"processor_response,omitempty"`
	CreateTime                string                     `json:"create_time,omitempty"`
	UpdateTime                string                     `json:"update_time,omitempty"`
}

// ResReauthorizeAuthorizedPayment
// https://developer.paypal.com/docs/api/payments/v2/#authorizations_reauthorize-response
type ResReauthorizeAuthorizedPayment struct {
	Status           string                     `json:"status,omitempty"`
	StatusDetails    AuthorizationStatusDetails `json:"status_details,omitempty"`
	Id               string                     `json:"id,omitempty"`
	Amount           Money                      `json:"amount,omitempty"`
	InvoiceId        string                     `json:"invoice_id,omitempty"`
	CustomId         string                     `json:"custom_id,omitempty"`
	SellerProtection SellerProtection           `json:"seller_protection,omitempty"`
	ExpirationTime   string                     `json:"expiration_time,omitempty"`
	Links            []LinkDescription          `json:"links,omitempty"`
	CreateTime       string                     `json:"create_time,omitempty"`
	UpdateTime       string                     `json:"update_time,omitempty"`
}

// ResGetCapturedPaymentDetails
// https://developer.paypal.com/docs/api/payments/v2/#captures_get-response
type ResGetCapturedPaymentDetails struct {
	Status                    string                     `json:"status,omitempty"`
	StatusDetails             AuthorizationStatusDetails `json:"status_details,omitempty"`
	Id                        string                     `json:"id,omitempty"`
	Amount                    Money                      `json:"amount,omitempty"`
	InvoiceId                 string                     `json:"invoice_id,omitempty"`
	CustomId                  string                     `json:"custom_id,omitempty"`
	SellerProtection          SellerProtection           `json:"seller_protection,omitempty"`
	FinalCapture              bool                       `json:"final_capture,omitempty"`
	SellerReceivableBreakdown SellerReceivableBreakdown  `json:"seller_receivable_breakdown,omitempty"`
	DisbursementMode          string                     `json:"disbursement_mode,omitempty"`
	Links                     []LinkDescription          `json:"links,omitempty"`
	ProcessorResponse         ProcessorResponse          `json:"processor_response,omitempty"`
	CreateTime                string                     `json:"create_time,omitempty"`
	UpdateTime                string                     `json:"update_time,omitempty"`
}

// ResRefundCapturedPayment
// https://developer.paypal.com/docs/api/payments/v2/#captures_refund-response
type ResRefundCapturedPayment struct {
	Status                    string                     `json:"status,omitempty"`
	StatusDetails             AuthorizationStatusDetails `json:"status_details,omitempty"`
	Id                        string                     `json:"id,omitempty"`
	Amount                    Money                      `json:"amount,omitempty"`
	InvoiceId                 string                     `json:"invoice_id,omitempty"`
	NoteToPayer               string                     `json:"note_to_payer,omitempty"`
	SellerReceivableBreakdown SellerReceivableBreakdown  `json:"seller_receivable_breakdown,omitempty"`
	Links                     []LinkDescription          `json:"links,omitempty"`
	CreateTime                string                     `json:"create_time,omitempty"`
	UpdateTime                string                     `json:"update_time,omitempty"`
}

// ResGetRefundDetails
// https://developer.paypal.com/docs/api/payments/v2/#refunds_get-response
type ResGetRefundDetails struct {
	Status                    string                     `json:"status,omitempty"`
	StatusDetails             AuthorizationStatusDetails `json:"status_details,omitempty"`
	Id                        string                     `json:"id,omitempty"`
	Amount                    Money                      `json:"amount,omitempty"`
	InvoiceId                 string                     `json:"invoice_id,omitempty"`
	NoteToPayer               string                     `json:"note_to_payer,omitempty"`
	SellerReceivableBreakdown SellerReceivableBreakdown  `json:"seller_receivable_breakdown,omitempty"`
	Links                     []LinkDescription          `json:"links,omitempty"`
	CreateTime                string                     `json:"create_time,omitempty"`
	UpdateTime                string                     `json:"update_time,omitempty"`
}

// ResCreateBatchPayout
// https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts-post-request-body-response
type ResCreateBatchPayout struct {
	BatchHeader PayoutHeader      `json:"batch_header,omitempty"`
	Links       []LinkDescription `json:"links,omitempty"`
}

// ResGetPayoutBatchDetails
// https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts_get-response
type ResGetPayoutBatchDetails struct {
	TotalItems  int                 `json:"total_items,omitempty"`
	TotalPages  int                 `json:"total_pages,omitempty"`
	BatchHeader *PayoutBatchHeader  `json:"batch_header,omitempty"`
	Items       []*PayoutBatchItems `json:"items,omitempty"`
	Links       []*LinkDescription  `json:"links,omitempty"`
}

// ResGetPayoutItemDetails
// https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts-item_get-response
type ResGetPayoutItemDetails struct {
	PayoutItemId       string                    `json:"payout_item_id,omitempty"`
	TransactionId      string                    `json:"transaction_id,omitempty"`
	ActivityId         string                    `json:"activity_id,omitempty"`
	TransactionStatus  string                    `json:"transaction_status,omitempty"`
	PayoutItemFee      *Currency                 `json:"payout_item_fee,omitempty"`
	PayoutBatchId      string                    `json:"payout_batch_id,omitempty"`
	SenderBatchId      string                    `json:"sender_batch_id,omitempty"`
	PayoutItem         *PayoutItemDetail         `json:"payout_item,omitempty"`
	CurrencyConversion *PayoutCurrencyConversion `json:"currency_conversion,omitempty"`
	TimeProcessed      string                    `json:"time_processed,omitempty"`
	Erros              *Error                    `json:"erros,omitempty"`
	Links              []*LinkDescription        `json:"links,omitempty"`
}

// ResCancelUnclaimedPayoutItem
// https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts-item_cancel-response
type ResCancelUnclaimedPayoutItem struct {
	ResGetPayoutItemDetails
}

// ResGetTransactionsList
// https://developer.paypal.com/docs/api/transaction-search/v1/#transactions_get-response
type ResGetTransactionsList struct {
	TransactionDetails    []*TransactionDetail `json:"transaction_details,omitempty"`
	AccountNumber         string               `json:"account_number,omitempty"`
	StartDate             string               `json:"start_date,omitempty"`
	EndDate               string               `json:"end_date,omitempty"`
	LastRefreshedDatetime string               `json:"last_refreshed_datetime,omitempty"`
	Page                  int                  `json:"page,omitempty"`
	TotalItems            int                  `json:"total_items,omitempty"`
	TotalPages            int                  `json:"total_pages,omitempty"`
	Links                 []*LinkDescription   `json:"links,omitempty"`
}

// ResGetListAllBalances
// https://developer.paypal.com/docs/api/transaction-search/v1/#balances_get-response
type ResGetListAllBalances struct {
	Balances        []*BalanceDetail `json:"balances,omitempty"`
	AccountId       string           `json:"account_id,omitempty"`
	AsOfTime        string           `json:"as_of_time,omitempty"`
	LastRefreshTime string           `json:"last_refresh_time,omitempty"`
}

// ResGetListWebhooks
// https://developer.paypal.com/docs/api/webhooks/v1/#webhooks_list-response
type ResGetListWebhooks struct {
	Webhooks []*Webhook `json:"webhooks,omitempty"`
}

// ResCreateWebhook
// https://developer.paypal.com/docs/api/webhooks/v1/#webhooks_post-response
type ResCreateWebhook struct {
	Id        string             `json:"id,omitempty"`
	Url       string             `json:"url"`
	EventType []*EventType       `json:"event_type"`
	Links     []*LinkDescription `json:"links,omitempty"`
}

// ResUpdateWebhook
// https://developer.paypal.com/docs/api/webhooks/v1/#webhooks_update-response
type ResUpdateWebhook struct {
	ResCreateWebhook
}

// ResGetWebhookDetails
// https://developer.paypal.com/docs/api/webhooks/v1/#webhooks_get-response
type ResGetWebhookDetails struct {
	ResCreateWebhook
}

// ResGetListEventSubscriptionsForWebhook
// https://developer.paypal.com/docs/api/webhooks/v1/#webhooks_list-response
type ResGetListEventSubscriptionsForWebhook struct {
	EventType []*EventType `json:"event_type"`
}

// ResVerifyWebhookSignature
// https://developer.paypal.com/docs/api/webhooks/v1/#verify-webhook-signature-post-response
type ResVerifyWebhookSignature struct {
	VerificationStatus string `json:"verification_status"` // 签名验证的状态：SUCCESS,FAILURE
}
