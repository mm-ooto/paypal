package paypal

var (
	SandboxGateWayUrl    = "https://api.sandbox.paypal.com" // 沙箱网关地址
	ProductionGateWayUrl = "https://api.paypal.com"         // 生产环境网关地址

	V1Version = "v1" // rest api v1版本
	V2Version = "v2" // rest api v2版本

	ContentTypeJson           = "application/json"
	ContentTypeFormUrlencoded = "application/x-www-form-urlencoded"

	VerificationStatusSuccess = "SUCCESS" // 签名验证成功
	VerificationStatusFailure = "FAILURE" // 签名验证失败

)
var (
	HttpMethodGet    = "GET"
	HttpMethodPost   = "POST"
	HttpMethodPut    = "PUT"
	HttpMethodPatch  = "PATCH"
	HttpMethodDelete = "DELETE"
)

var (
	/*
	* ResourceType：与webhook通知事件相关的资源名
	 */
	ResourceTypeDispute  = "dispute"  // 资源类型：纠纷
	ResourceTypeRefund   = "refund"   // 资源类型：退款
	ResourceTypeSale     = "sale"     // 资源类型：交易
	ResourceTypeInvoices = "invoices" // 资源类型：发票
	/*
	*	EventType：触发webhook事件通知的事件类型
	 */
	EventTypeRiskDisputeCreate       = "RISK.DISPUTE.CREATED"      // 事件类型：创建纠纷
	EventTypeCustomerDisputeCreate   = "CUSTOMER.DISPUTE.CREATED"  // 事件类型：顾客纠纷创建
	EventTypeCustomerDisputeUpdated  = "CUSTOMER.DISPUTE.UPDATED"  // 事件类型：顾客修改纠纷
	EventTypeCustomerDisputeResolved = "CUSTOMER.DISPUTE.RESOLVED" // 事件类型：顾客解决纠纷

	EventTypePaymentSaleRefund = "PAYMENT.SALE.REFUND" // 事件类型：退款完成

	EventTypePaymentSaleCompleted = "PAYMENT.SALE.COMPLETED" // 事件类型：付款完成
	EventTypePaymentSalePending   = "PAYMENT.SALE.PENDING"   // 事件类型：付款中
	EventTypePaymentSaleDenied    = "PAYMENT.SALE.DENIED"    // 事件类型：拒绝付款
)

var (
	IntentCapture   = "CAPTURE"
	IntentAuthorize = "AUTHORIZE"
)

var (
	AccessTokenAPI = "/v1/oauth2/token"

	OrderAPI          = "/v2/checkout/orders"
	OrderAuthorizeAPI = "/v2/checkout/orders/%s/authorize"
	OrderCaptureAPI   = "/v2/checkout/orders/%s/capture"
	OrderConfirmAPI   = "/v2/checkout/orders/%s/confirm-payment-source"

	AuthorizedPaymentDetailsAPI     = "/v2/payments/authorizations/%s"
	CaptureAuthorizedPaymentAPI     = "/v2/payments/authorizations/%s/capture"
	ReauthorizeAuthorizedPaymentAPI = "/v2/payments/authorizations/%s/reauthorize"
	VoidAuthorizedPaymentAPI        = "/v2/payments/authorizations/%s/void"

	CapturedPaymentDetailsAPI = "/v2/payments/captures/%s"
	RefundCapturedPaymentAPI  = "/v2/payments/captures/%s/refund"

	RefundDetailsAPI = "/v2/payments/refunds/%s"

	CreateBatchPayoutAPI      = "/v1/payments/payouts"
	ShowPayoutBatchDetailsAPI = "/v1/payments/payouts/%s"

	ShowPayoutItemDetailsAPI     = "/v1/payments/payouts-item/%s"
	CancelUnclaimedPayoutItemAPI = "/v1/payments/payouts-item/%s/cancel"

	ListTransactionsAPI = "/v1/reporting/transactions"
	ListAllBalancesAPI  = "/v1/reporting/balances"

	WebhooksAPI                         = "/v1/notifications/webhooks"
	ListEventSubscriptionsForWebhookAPI = "/v1/notifications/webhooks/%s/event-types"
	VerifyWebhookSignatureAPI           = "/v1/notifications/verify-webhook-signature"
	ShowEventNotificationDetailAPI      = "/v1/notifications/webhooks-events/%s"
	ResendWebhookEventNotificationAPI   = "/v1/notifications/webhooks-events/%s/resend"

	DisputesListAPI   = "/v1/customer/disputes"
	DisputesDetailAPI = "/v1/customer/disputes/%s"
)
