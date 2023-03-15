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
	Id                    string                 `json:"id,omitempty"`
	PaymentSource         *PaymentSourceResponse `json:"payment_source,omitempty"`
	Intent                string                 `json:"intent,omitempty"`
	ProcessingInstruction string                 `json:"processing_instruction,omitempty"`
	Payer                 *Payer                 `json:"payer,omitempty"`
	PurchaseUnits         []*PurchaseUnit        `json:"purchase_units,omitempty"`
	Status                string                 `json:"status,omitempty"`
	Links                 []*LinkDescription     `json:"links"`
	CreateTime            string                 `json:"create_time,omitempty"`
	UpdateTime            string                 `json:"update_time,omitempty"`
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
	Status           string                      `json:"status,omitempty"`
	StatusDetails    *AuthorizationStatusDetails `json:"status_details,omitempty"`
	Id               string                      `json:"id,omitempty"`
	Amount           *Money                      `json:"amount,omitempty"`
	InvoiceId        string                      `json:"invoice_id,omitempty"`
	CustomId         string                      `json:"custom_id,omitempty"`
	SellerProtection *SellerProtection           `json:"seller_protection,omitempty"`
	ExpirationTime   string                      `json:"expiration_time,omitempty"`
	Links            []*LinkDescription          `json:"links,omitempty"`
	CreateTime       string                      `json:"create_time,omitempty"`
	UpdateTime       string                      `json:"update_time,omitempty"`
}

// ResCaptureAuthorizedPayment
// https://developer.paypal.com/docs/api/payments/v2/#authorizations_capture-response
type ResCaptureAuthorizedPayment struct {
	Status                    string                      `json:"status,omitempty"`
	StatusDetails             *AuthorizationStatusDetails `json:"status_details,omitempty"`
	Id                        string                      `json:"id,omitempty"`
	Amount                    *Money                      `json:"amount,omitempty"`
	InvoiceId                 string                      `json:"invoice_id,omitempty"`
	CustomId                  string                      `json:"custom_id,omitempty"`
	SellerProtection          *SellerProtection           `json:"seller_protection,omitempty"`
	FinalCapture              bool                        `json:"final_capture,omitempty"`
	SellerReceivableBreakdown *SellerReceivableBreakdown  `json:"seller_receivable_breakdown,omitempty"`
	DisbursementMode          string                      `json:"disbursement_mode,omitempty"`
	Links                     []*LinkDescription          `json:"links,omitempty"`
	ProcessorResponse         *ProcessorResponse          `json:"processor_response,omitempty"`
	CreateTime                string                      `json:"create_time,omitempty"`
	UpdateTime                string                      `json:"update_time,omitempty"`
}

// ResReauthorizeAuthorizedPayment
// https://developer.paypal.com/docs/api/payments/v2/#authorizations_reauthorize-response
type ResReauthorizeAuthorizedPayment struct {
	Status           string                      `json:"status,omitempty"`
	StatusDetails    *AuthorizationStatusDetails `json:"status_details,omitempty"`
	Id               string                      `json:"id,omitempty"`
	Amount           *Money                      `json:"amount,omitempty"`
	InvoiceId        string                      `json:"invoice_id,omitempty"`
	CustomId         string                      `json:"custom_id,omitempty"`
	SellerProtection *SellerProtection           `json:"seller_protection,omitempty"`
	ExpirationTime   string                      `json:"expiration_time,omitempty"`
	Links            []*LinkDescription          `json:"links,omitempty"`
	CreateTime       string                      `json:"create_time,omitempty"`
	UpdateTime       string                      `json:"update_time,omitempty"`
}

// ResGetCapturedPaymentDetails
// https://developer.paypal.com/docs/api/payments/v2/#captures_get-response
type ResGetCapturedPaymentDetails struct {
	Status                    string                      `json:"status,omitempty"`
	StatusDetails             *AuthorizationStatusDetails `json:"status_details,omitempty"`
	Id                        string                      `json:"id,omitempty"`
	Amount                    *Money                      `json:"amount,omitempty"`
	InvoiceId                 string                      `json:"invoice_id,omitempty"`
	CustomId                  string                      `json:"custom_id,omitempty"`
	SellerProtection          *SellerProtection           `json:"seller_protection,omitempty"`
	FinalCapture              bool                        `json:"final_capture,omitempty"`
	SellerReceivableBreakdown *SellerReceivableBreakdown  `json:"seller_receivable_breakdown,omitempty"`
	DisbursementMode          string                      `json:"disbursement_mode,omitempty"`
	Links                     []*LinkDescription          `json:"links,omitempty"`
	ProcessorResponse         *ProcessorResponse          `json:"processor_response,omitempty"`
	CreateTime                string                      `json:"create_time,omitempty"`
	UpdateTime                string                      `json:"update_time,omitempty"`
}

// ResRefundCapturedPayment
// https://developer.paypal.com/docs/api/payments/v2/#captures_refund-response
type ResRefundCapturedPayment struct {
	Status                    string                      `json:"status,omitempty"`
	StatusDetails             *AuthorizationStatusDetails `json:"status_details,omitempty"`
	Id                        string                      `json:"id,omitempty"`
	Amount                    *Money                      `json:"amount,omitempty"`
	InvoiceId                 string                      `json:"invoice_id,omitempty"`
	NoteToPayer               string                      `json:"note_to_payer,omitempty"`
	SellerReceivableBreakdown *SellerReceivableBreakdown  `json:"seller_receivable_breakdown,omitempty"`
	Links                     []*LinkDescription          `json:"links,omitempty"`
	CreateTime                string                      `json:"create_time,omitempty"`
	UpdateTime                string                      `json:"update_time,omitempty"`
}

// ResGetRefundDetails
// https://developer.paypal.com/docs/api/payments/v2/#refunds_get-response
type ResGetRefundDetails struct {
	Status                    string                      `json:"status,omitempty"`
	StatusDetails             *AuthorizationStatusDetails `json:"status_details,omitempty"`
	Id                        string                      `json:"id,omitempty"`
	Amount                    *Money                      `json:"amount,omitempty"`
	InvoiceId                 string                      `json:"invoice_id,omitempty"`
	NoteToPayer               string                      `json:"note_to_payer,omitempty"`
	SellerReceivableBreakdown *SellerReceivableBreakdown  `json:"seller_receivable_breakdown,omitempty"`
	Links                     []*LinkDescription          `json:"links,omitempty"`
	CreateTime                string                      `json:"create_time,omitempty"`
	UpdateTime                string                      `json:"update_time,omitempty"`
}

// ResCreateBatchPayout
// https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts-post-request-body-response
type ResCreateBatchPayout struct {
	BatchHeader *PayoutHeader      `json:"batch_header,omitempty"`
	Links       []*LinkDescription `json:"links,omitempty"`
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
	Id         string             `json:"id,omitempty"`
	Url        string             `json:"url"`
	EventTypes []*EventType       `json:"event_types"`
	Links      []*LinkDescription `json:"links,omitempty"`
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
	EventTypes []*EventType `json:"event_types"`
}

// ResResendWebhookEventNotification 重发事件通知
type ResResendWebhookEventNotification struct {
	WebhookIds []struct {
		WebhookId string `json:"webhook_id"` // 要重发的webhook事件通知的ID
	} `json:"webhook_ids"`
}

// ResVerifyWebhookSignature
// https://developer.paypal.com/docs/api/webhooks/v1/#verify-webhook-signature-post-response
type ResVerifyWebhookSignature struct {
	VerificationStatus string `json:"verification_status"` // 签名验证的状态：SUCCESS,FAILURE
}

// ResDisputesList 争议列表
// https://developer.paypal.com/docs/api/customer-disputes/v1/#disputes-list-response
type ResDisputesList struct {
	Items []DisputesInfo    `json:"items"` // 匹配筛选条件的争议数组。按照从最近到最早的创建时间顺序排序。
	Links []LinkDescription `json:"links"`
}

// DisputesInfo 包含争议信息对象
type DisputesInfo struct {
	DisputeId     string            `json:"dispute_id"`     //
	CreateTime    string            `json:"create_time"`    // 争议产生的日期和时间，以互联网日期和时间格式。例如:yyyy-MM-ddTHH:mm:ss.SSSZ。
	UpdateTime    string            `json:"update_time"`    //
	Status        string            `json:"status"`         //
	Reason        string            `json:"reason"`         //
	DisputeState  string            `json:"dispute_state"`  //
	DisputeAmount Money             `json:"dispute_amount"` // 交易中客户最初争议的金额。由于客户有时只能对部分付款提出异议，因此有争议的金额可能与原始交易的总毛额或净额不同。
	Links         []LinkDescription `json:"links"`
}

// ResDisputesDetail 争议详情
// https://developer.paypal.com/docs/api/customer-disputes/v1/#disputes-get-response
type ResDisputesDetail struct {
	DisputeId            string                    `json:"dispute_id"`  //
	CreateTime           string                    `json:"create_time"` // 争议产生的日期和时间，以互联网日期和时间格式。例如:yyyy-MM-ddTHH:mm:ss.SSSZ。
	UpdateTime           string                    `json:"update_time"` //
	DisputedTransactions []DisputesTransactionInfo `json:"disputed_transactions"`
	Reason               string                    `json:"reason"`         //
	Status               string                    `json:"status"`         //
	DisputeAmount        Money                     `json:"dispute_amount"` //
	DisputeOutcome       struct {
		OutcomeCode    string `json:"outcome_code"`
		AmountRefunded Money  `json:"amount_refunded"` //
	} `json:"dispute_outcome"`
	DisputeLifeCycleStage string `json:"dispute_life_cycle_stage"`
	DisputeChannel        string `json:"dispute_channel"`
	Messages              []struct {
		PostedBy   string `json:"posted_by"`
		TimePosted string `json:"time_posted"`
		Content    string `json:"content"`
		Documents  []struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"documents"`
	} `json:"messages"`
	Extensions struct {
		MerchandizeDisputeProperties struct {
			IssueType      string `json:"issue_type"`
			ServiceDetails []struct {
				SubReasons  []string `json:"sub_reasons"`
				PurchaseUrl string   `json:"purchase_url"`
			} `json:"service_details"`
		} `json:"merchandize_dispute_properties"`
	} `json:"extensions"`
	Offer struct {
		BuyerRequestedAmount Money  `json:"buyer_requested_amount"`
		OfferType            string `json:"offer_type"`
		History              struct {
			OfferTime             string `json:"offer_time"`
			Actor                 string `json:"actor"`
			EventType             string `json:"event_type"`
			OfferType             string `json:"offer_type"`
			OfferAmount           Money  `json:"offer_amount"`
			Notes                 string `json:"notes"`
			DisputeLifeCycleStage string `json:"dispute_life_cycle_stage"`
		} `json:"history"`
	} `json:"offer"`
	Links []LinkDescription `json:"links"`
}

// DisputesTransactionInfo 争议交易详情
type DisputesTransactionInfo struct {
	Buyer struct {
		Name string `json:"name"`
	}
	BuyerTransactionId string `json:"buyer_transaction_id"` // 客户看到的此事务的ID。
	CreateTime         string `json:"create_time"`          // 创建时间
	Custom             string `json:"custom"`               // 由商家在结帐时输入的自由文本字段。
	GrossAmount        Money  `json:"gross_amount"`         // 交易总额。
	GrossAsset         struct {
		AssetSymbol string `json:"asset_symbol"` // 由流动性提供者(交易所)分配的加密货币股票代码。
		Quantity    string `json:"quantity"`     // 加密货币资产数量。 这是一个十进制数字，由创始人为每个加密货币定义了刻度。例如, 比特币(BTC)的规模为8，以太坊(ETH)有18个规模。PayPal加密货币平台处理比特币及其分支和以太坊的8位数规模。
	} `json:"gross_asset"` // 交易的总资产。
	InvoiceNumber string `json:"invoice_number"` // 付款发票的ID。
	Items         []struct {
		DisputeAmount        Money  `json:"dispute_amount"`         // 争议事项金额
		ItemDescription      string `json:"item_description"`       // 争议事项描述
		ItemId               string `json:"item_id"`                // 争议事项ID
		ItemName             string `json:"item_name"`              // 争议事项名字
		ItemQuantity         string `json:"item_quantity"`          // 争议事项:争议事项的数量一定是个整数。
		ItemType             string `json:"item_type"`              // 争议事项类型，可选值：PRODUCT，SERVICE，BOOKING，DIGITAL_DOWNLOAD
		Notes                string `json:"notes"`                  // 该争议事项提供的任何备注。
		PartnerTransactionId string `json:"partner_transaction_id"` // 合作伙伴系统中事务的ID。合作伙伴事务ID在项目级别返回，因为合作伙伴可能为购物车中的不同项目显示不同的事务。
		Reason               string `json:"reason"`                 // 项目层面争议的原因。有关每个争议原因和相关证据类型所需信息的信息，请参见争议原因。可选值：MERCHANDISE_OR_SERVICE_NOT_RECEIVED，
		// MERCHANDISE_OR_SERVICE_NOT_AS_DESCRIBED，UNAUTHORISED，CREDIT_NOT_PROCESSED，DUPLICATE_TRANSACTION，INCORRECT_AMOUNT，PAYMENT_BY_OTHER_MEANS，
		// CANCELED_RECURRING_BILLING，PROBLEM_WITH_REMITTANCE，OTHER

	} `json:"items"`
	ReferenceId string `json:"reference_id"` // 合作伙伴看到的此事务的ID。
	Seller      struct {
		Email      string `json:"email"`
		MerchantId string `json:"merchant_id"`
		Name       string `json:"name"`
	} `json:"seller"`
	SellerTransactionId string `json:"seller_transaction_id"` // 卖家交易id
	TransactionStatus   string `json:"transaction_status"`    // 交易状态
}
