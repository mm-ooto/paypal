package model

// NameCountryCodeBase NameCountryCodeBase
type NameCountryCodeBase struct {
	Name        string `json:"name"`
	CountryCode string `json:"country_code"`
}

// Details Details
type Details struct {
	Field       string `json:"field"`
	Value       string `json:"value"`
	Location    string `json:"location"`
	Issue       string `json:"issue"`
	Description string `json:"description"`
}

// Payer
// https://developer.paypal.com/docs/api/orders/v2/#definition-payer
type Payer struct {
	EmailAddress string `json:"email_address,omitempty"`
	PayerId      string `json:"payer_id,omitempty"`

	Name      *Name            `json:"name,omitempty"`
	Phone     *PhoneWithType   `json:"phone,omitempty"`
	BirthDate string           `json:"birth_date,omitempty"` // YYYY-MM-DD
	TaxInfo   *TaxInfo         `json:"tax_info,omitempty"`
	Address   *AddressPortable `json:"address,omitempty"`
}

// Name
// https://developer.paypal.com/docs/api/orders/v2/#definition-name
type Name struct {
	Name string `json:"name,omitempty"`
}

// PhoneWithType
// https://developer.paypal.com/docs/api/orders/v2/#definition-phone_with_type
type PhoneWithType struct {
	PhoneType   string              `json:"phone_type,omitempty"`
	PhoneNumber *PhoneWithTypePhone `json:"phone_number"`
}

// PhoneWithTypePhone
// https://developer.paypal.com/docs/api/orders/v2/#definition-phone_with_type.phone
type PhoneWithTypePhone struct {
	NationalNumber string `json:"national_number"`
}

// TaxInfo
// https://developer.paypal.com/docs/api/orders/v2/#definition-tax_info
type TaxInfo struct {
	TaxId     string `json:"tax_id"`
	TaxIdType string `json:"tax_id_type"`
}

// AddressPortable
// https://developer.paypal.com/docs/api/orders/v2/#definition-address_portable
type AddressPortable struct {
	AddressLine1   string          `json:"address_line_1,omitempty"`
	AddressLine2   string          `json:"address_line_2,omitempty"`
	AddressLine3   string          `json:"address_line_3,omitempty"`
	AdminArea4     string          `json:"admin_area_4,omitempty"`
	AdminArea3     string          `json:"admin_area_3,omitempty"`
	AdminArea2     string          `json:"admin_area_2,omitempty"`
	AdminArea1     string          `json:"admin_area_1,omitempty"`
	PostalCode     string          `json:"postal_code,omitempty"`
	CountryCode    string          `json:"country_code"`
	AddressDetails *AddressDetails `json:"address_details,omitempty"`
}

// AddressDetails
// https://developer.paypal.com/docs/api/orders/v2/#definition-address_details
type AddressDetails struct {
	StreetNumber    string `json:"street_number,omitempty"`
	StreetName      string `json:"street_name,omitempty"`
	StreetType      string `json:"street_type,omitempty"`
	DeliveryService string `json:"delivery_service,omitempty"`
	BuildingName    string `json:"building_name,omitempty"`
	SubBuilding     string `json:"sub_building,omitempty"`
}

// Item
// https://developer.paypal.com/docs/api/orders/v2/#definition-item
type Item struct {
	Name        string `json:"name"`
	UnitAmount  *Money `json:"unit_amount"`
	Tax         *Money `json:"tax,omitempty"`
	Quantity    string `json:"quantity"`
	Description string `json:"description,omitempty"`
	Sku         string `json:"sku,omitempty"`
	Category    string `json:"category,omitempty"`
}

// Money
// https://developer.paypal.com/docs/api/orders/v2/#definition-money
type Money struct {
	CurrencyCode string `json:"currency_code"`
	Value        string `json:"value"`
}

// ShippingDetail
// https://developer.paypal.com/docs/api/orders/v2/#definition-shipping_detail
type ShippingDetail struct {
	Name    *ShippingDetailName                `json:"name,omitempty"`
	Type    string                             `json:"type,omitempty"`
	Address *ShippingDetailNameAddressPortable `json:"address,omitempty"`
}

// ShippingDetailName
// https://developer.paypal.com/docs/api/orders/v2/#definition-shipping_detail.name
type ShippingDetailName struct {
	FullName string `json:"full_name,omitempty"`
}

// ShippingDetailNameAddressPortable
// https://developer.paypal.com/docs/api/orders/v2/#definition-shipping_detail.address_portable
type ShippingDetailNameAddressPortable struct {
	AddressLine1 string `json:"address_line_1,omitempty"`
	AddressLine2 string `json:"address_line_2,omitempty"`
	AdminArea2   string `json:"admin_area_2,omitempty"`
	AdminArea1   string `json:"admin_area_1,omitempty"`
	PostalCode   string `json:"postal_code,omitempty"`
	CountryCode  string `json:"country_code"`
}

// PaymentCollection
// https://developer.paypal.com/docs/api/orders/v2/#definition-payment_collection
type PaymentCollection struct {
	Authorizations []*AuthorizationWithAdditionalData `json:"authorizations,omitempty"`
	Captures       []*Capture                         `json:"captures,omitempty"`
	Refund         []string                           `json:"refund,omitempty"`
}

// AuthorizationWithAdditionalData
// https://developer.paypal.com/docs/api/orders/v2/#definition-authorization_with_additional_data
type AuthorizationWithAdditionalData struct {
	ProcessorResponse *ProcessorResponse `json:"processor_response,omitempty"`
}

// ProcessorResponse
// https://developer.paypal.com/docs/api/orders/v2/#definition-processor_response
type ProcessorResponse struct {
	AvsCode           string `json:"avs_code,omitempty"`
	CvvCode           string `json:"cvv_code,omitempty"`
	ResponseCode      string `json:"response_code,omitempty"`
	PaymentAdviceCode string `json:"payment_advice_code,omitempty"`
}

// Capture
// https://developer.paypal.com/docs/api/orders/v2/#definition-capture
type Capture struct {
	Status                    string                `json:"status,omitempty"`
	StatusDetails             *CaptureStatusDetails `json:"status_details,omitempty"`
	Id                        string                `json:"id,omitempty"`
	Amount                    *Money                `json:"amount,omitempty"`
	InvoiceId                 string                `json:"invoice_id,omitempty"`
	CustomId                  string                `json:"custom_id,omitempty"`
	SellerProtection          *SellerProtection     `json:"seller_protection,omitempty"`
	FinalCapture              bool                  `json:"final_capture"`
	SellerReceivableBreakdown `json:"seller_receivable_breakdown,omitempty"`
	DisbursementMode          string             `json:"disbursement_mode,omitempty"`
	Link                      []*LinkDescription `json:"link,omitempty"`
	ProcessorResponse         *ProcessorResponse `json:"processor_response,omitempty"`
}

// CaptureStatusDetails
// https://developer.paypal.com/docs/api/orders/v2/#definition-capture_status_details
type CaptureStatusDetails struct {
	Reason string `json:"reason,omitempty"`
}

// SellerProtection
// https://developer.paypal.com/docs/api/orders/v2/#definition-seller_protection
type SellerProtection struct {
	Status            string             `json:"status,omitempty"`
	DisputeCategories []*DisputeCategory `json:"dispute_categories,omitempty"`
}

// DisputeCategory
// https://developer.paypal.com/docs/api/orders/v2/#definition-dispute_category
type DisputeCategory struct {
	DisputeCategory string `json:"dispute_category,omitempty"`
}

// SellerReceivableBreakdown
// https://developer.paypal.com/docs/api/orders/v2/#definition-seller_receivable_breakdown
type SellerReceivableBreakdown struct {
	GrossAmount                   *Money         `json:"gross_amount"`
	PayPalFee                     *Money         `json:"pay_pal_fee,omitempty"`
	PaypalFeeInReceivableCurrency *Money         `json:"paypal_fee_in_receivable_currency,omitempty"`
	NetAmount                     *Money         `json:"net_amount,omitempty"`
	ReceivableAmount              *Money         `json:"receivable_amount,omitempty"`
	ExchangeRate                  *ExchangeRate  `json:"exchange_rate,omitempty"`
	PlatformFees                  []*PlatformFee `json:"platform_fees,omitempty"`
}

// ExchangeRate
// https://developer.paypal.com/docs/api/orders/v2/#definition-exchange_rate
type ExchangeRate struct {
	SourceCurrency string `json:"source_currency,omitempty"`
	TargetCurrency string `json:"target_currency,omitempty"`
	Value          string `json:"value,omitempty"`
}

// PlatformFee
// https://developer.paypal.com/docs/api/orders/v2/#definition-platform_fee
type PlatformFee struct {
	Amount *Money `json:"amount"`
	Payee  *Payee `json:"payee,omitempty"`
}

// Payee
// https://developer.paypal.com/docs/api/orders/v2/#definition-payee_base
type Payee struct {
	EmailAddress string `json:"email_address,omitempty"`
	MerchantId   string `json:"merchant_id,omitempty"`
}

// LinkDescription
// https://developer.paypal.com/docs/api/orders/v2/#definition-link_description
type LinkDescription struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method,omitempty"`
}

// PurchaseUnitsRequest
// https://developer.paypal.com/docs/api/orders/v2/#definition-purchase_unit_request
type PurchaseUnitsRequest struct {
	ReferenceId string               `json:"reference_ids,omitempty"`
	Amount      *AmountWithBreakdown `json:"amount"`
}

// AmountWithBreakdown
// https://developer.paypal.com/docs/api/orders/v2/#definition-amount_with_breakdown
type AmountWithBreakdown struct {
	CurrencyCode string           `json:"currency_code"` // 币种代码
	Value        string           `json:"value"`         // 金额
	Breakdown    *AmountBreakdown `json:"breakdown,omitempty"`
}

// AmountBreakdown
// https://developer.paypal.com/docs/api/orders/v2/#definition-amount_breakdown
type AmountBreakdown struct {
	ItemTotal        *Money `json:"item_total,omitempty"`
	Shipping         *Money `json:"shipping,omitempty"`
	Handling         *Money `json:"handling,omitempty"`
	TaxTotal         *Money `json:"tax_total,omitempty"`
	Insurance        *Money `json:"insurance,omitempty"`
	ShippingDiscount *Money `json:"shipping_discount,omitempty"`
	Discount         *Money `json:"discount,omitempty"`
}

// OrderApplicationContext
// https://developer.paypal.com/docs/api/orders/v2/#definition-order_application_context
type OrderApplicationContext struct {
	BrandName           string               `json:"brand_name,omitempty"`
	Locale              string               `json:"locale,omitempty"`
	LandingPage         string               `json:"landing_page,omitempty"`
	ShippingPreference  string               `json:"shipping_preference,omitempty"`
	UserAction          string               `json:"user_action,omitempty"`
	PaymentMethod       *PaymentMethod       `json:"payment_method,omitempty"`
	ReturnUrl           string               `json:"return_url,omitempty"`
	CancelUrl           string               `json:"cancel_url,omitempty"`
	StoredPaymentSource *StoredPaymentSource `json:"stored_payment_source,omitempty"`
}

// PaymentMethod
// https://developer.paypal.com/docs/api/orders/v2/#definition-payment_method
type PaymentMethod struct {
	PayeePreferred         string `json:"payee_preferred,omitempty"`
	StandardEntryClassCode string `json:"standard_entry_class_code,omitempty"`
}

// StoredPaymentSource
// https://developer.paypal.com/docs/api/orders/v2/#definition-stored_payment_source
type StoredPaymentSource struct {
	PaymentInitiator                    string                       `json:"payment_initiator"`
	PaymentType                         string                       `json:"payment_type"`
	Usage                               string                       `json:"usage,omitempty"`
	PreviousNetworkTransactionTeference *NetworkTransactionTeference `json:"previous_network_transaction_teference"`
}

// NetworkTransactionTeference
// https://developer.paypal.com/docs/api/orders/v2/#definition-network_transaction_reference
type NetworkTransactionTeference struct {
	Id      string `json:"id"`
	Date    string `json:"date,omitempty"`
	Network string `json:"network"`
}

// PaymentSourceResponse
// https://developer.paypal.com/docs/api/orders/v2/#definition-payment_source_response
type PaymentSourceResponse struct {
	Card       *CardResponse `json:"card,omitempty"`
	Bancontact *Bancontact   `json:"bancontact,omitempty"`
	Blik       *Blik         `json:"blik,omitempty"`
	Eps        *Eps          `json:"eps,omitempty"`
	Giropay    *Giropay      `json:"giropay,omitempty"`
	Ideal      *Ideal        `json:"ideal,omitempty"`
	Mybank     *Mybank       `json:"mybank,omitempty"`
	P24        *P24          `json:"p24,omitempty"`
	Sofort     *Sofort       `json:"sofort,omitempty"`
	Trustly    *Trustly      `json:"trustly,omitempty"`
}

// CardResponse
// https://developer.paypal.com/docs/api/orders/v2/#definition-card_response
type CardResponse struct {
	Name                 string                       `json:"name,omitempty"`
	BillingAddress       *CardResponseAddressPortable `json:"billing_address,omitempty"`
	LastDigits           string                       `json:"last_digits,omitempty"`
	Brand                string                       `json:"brand,omitempty"`
	Type                 string                       `json:"type,omitempty"`
	AuthenticationResult *AuthenticationResponse      `json:"authentication_result,omitempty"`
}

// CardResponseAddressPortable
// https://developer.paypal.com/docs/api/orders/v2/#definition-card_response.address_portable
type CardResponseAddressPortable struct {
	AddressLine1 string `json:"address_line_1,omitempty"`
	AddressLine2 string `json:"address_line_2,omitempty"`
	AdminArea2   string `json:"admin_area_2,omitempty"`
	AdminArea1   string `json:"admin_area_1,omitempty"`
	PostalCode   string `json:"postal_code,omitempty"`
	CountryCode  string `json:"country_code"`
}

// AuthenticationResponse
// https://developer.paypal.com/docs/api/orders/v2/#definition-authentication_response
type AuthenticationResponse struct {
	LiabilityShift string                              `json:"liability_shift,omitempty"`
	ThreeDSecure   *ThreeDSecureAuthenticationResponse `json:"three_d_secure,omitempty"`
}

// ThreeDSecureAuthenticationResponse
// https://developer.paypal.com/docs/api/orders/v2/#definition-three_d_secure_authentication_response
type ThreeDSecureAuthenticationResponse struct {
	AuthenticationStatus string `json:"authentication_status,omitempty"`
	EnrollmentStatus     string `json:"enrollment_status,omitempty"`
}

// Bancontact
// https://developer.paypal.com/docs/api/orders/v2/#definition-bancontact
type Bancontact struct {
	Name           string `json:"name"`
	CountryCode    string `json:"country_code"`
	Bic            string `json:"bic"`
	IbanLastChars  string `json:"iban_last_chars"`
	CardLastDigits string `json:"card_last_digits"`
}

// Blik
// https://developer.paypal.com/docs/api/orders/v2/#definition-blik
type Blik struct {
	Name        string `json:"name"`
	CountryCode string `json:"country_code"`
	Email       string `json:"email"`
}

// Eps
// https://developer.paypal.com/docs/api/orders/v2/#definition-eps
type Eps struct {
	Name        string `json:"name"`
	CountryCode string `json:"country_code"`
	Bic         string `json:"bic"`
}

// Giropay
// https://developer.paypal.com/docs/api/orders/v2/#definition-giropay
type Giropay struct {
	Name        string `json:"name"`
	CountryCode string `json:"country_code"`
	Bic         string `json:"bic"`
}

// Ideal
// https://developer.paypal.com/docs/api/orders/v2/#definition-ideal
type Ideal struct {
	Name          string `json:"name"`
	CountryCode   string `json:"country_code"`
	Bic           string `json:"bic"`
	IbanLastChars string `json:"iban_last_chars"`
}

// Mybank
// https://developer.paypal.com/docs/api/orders/v2/#definition-mybank
type Mybank struct {
	Name          string `json:"name"`
	CountryCode   string `json:"country_code"`
	Bic           string `json:"bic"`
	IbanLastChars string `json:"iban_last_chars"`
}

// P24
// https://developer.paypal.com/docs/api/orders/v2/#definition-p24
type P24 struct {
	Name              string `json:"name"`
	Email             string `json:"email"`
	CountryCode       string `json:"country_code"`
	PaymentDescriptor string `json:"payment_descriptor"`
	MethodId          string `json:"method_id"`
	MethodDescription string `json:"method_description"`
}

// Sofort
// https://developer.paypal.com/docs/api/orders/v2/#definition-sofort
type Sofort struct {
	Name          string `json:"name"`
	CountryCode   string `json:"country_code"`
	Bic           string `json:"bic"`
	IbanLastChars string `json:"iban_last_chars"`
}

// Trustly
// https://developer.paypal.com/docs/api/orders/v2/#definition-trustly
type Trustly struct {
	Name          string `json:"name"`
	CountryCode   string `json:"country_code"`
	Bic           string `json:"bic"`
	IbanLastChars string `json:"iban_last_chars"`
}

// PurchaseUnit
// https://developer.paypal.com/docs/api/orders/v2/#definition-purchase_unit
type PurchaseUnit struct {
	ReferenceId        string               `json:"reference_id,omitempty"`
	Amount             *AmountWithBreakdown `json:"amount,omitempty"`
	Payee              *Payee               `json:"payee,omitempty"`
	PaymentInstruction *PaymentInstruction  `json:"payment_instruction,omitempty"`
}

// PaymentInstruction
// https://developer.paypal.com/docs/api/orders/v2/#definition-payment_instruction
type PaymentInstruction struct {
	PlatformFees            *PlatformFee `json:"platform_fees,omitempty"`
	DisbursementMode        string       `json:"disbursement_mode,omitempty"`
	PayeePricingTierId      string       `json:"payee_pricing_tier_id,omitempty"`
	PayeeReceivableFxRateId string       `json:"payee_receivable_fx_rate_id,omitempty"`
}

// Patch
// https://developer.paypal.com/docs/api/orders/v2/#definition-patch
type Patch struct {
	Op    string      `json:"op"`
	Path  string      `json:"path,omitempty"`
	Value interface{} `json:"value,omitempty"` // number,integer,string,boolean,null,array,object
	From  string      `json:"from,omitempty"`
}

// OrderAuthorizeRequestPaymentSource
// https://developer.paypal.com/docs/api/orders/v2/#definition-order_authorize_request.payment_source
type OrderAuthorizeRequestPaymentSource struct {
	Token *Token `json:"token,omitempty"`
}

// Token
// https://developer.paypal.com/docs/api/orders/v2/#definition-token
type Token struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

// OrdersCaptureRequestPaymentSource
// https://developer.paypal.com/docs/api/orders/v2/#definition-order_capture_request.payment_source
type OrdersCaptureRequestPaymentSource struct {
	Token *Token `json:"token,omitempty"`
}

// PaymentSource
// https://developer.paypal.com/docs/api/orders/v2/#definition-payment_source
type PaymentSource struct {
	Token      *Token             `json:"token,omitempty"`
	Bancontact *BancontactRequest `json:"bancontact,omitempty"`
	Blik       *BlikRequest       `json:"blik,omitempty"`
	Eps        *EpsRequest        `json:"eps,omitempty"`
	Giropay    *GiropayRequest    `json:"giropay,omitempty"`
	Ideal      IdealRequest       `json:"ideal,omitempty"`
	Mybank     *MybankRequest     `json:"mybank,omitempty"`
	P24        *P24Request        `json:"p24,omitempty"`
	Sofort     *SofortRequest     `json:"sofort,omitempty"`
	Trustly    *TrustlyRequest    `json:"trustly,omitempty"`
}

// BancontactRequest
// https://developer.paypal.com/docs/api/orders/v2/#definition-bancontact_request
type BancontactRequest struct {
	NameCountryCodeBase
}

// BlikRequest
// https://developer.paypal.com/docs/api/orders/v2/#definition-blik_request
type BlikRequest struct {
	NameCountryCodeBase
	Email string `json:"email,omitempty"`
}

// EpsRequest
// https://developer.paypal.com/docs/api/orders/v2/#definition-eps_request
type EpsRequest struct {
	NameCountryCodeBase
}

// GiropayRequest
// https://developer.paypal.com/docs/api/orders/v2/#definition-giropay_request
type GiropayRequest struct {
	NameCountryCodeBase
}

// IdealRequest
// https://developer.paypal.com/docs/api/orders/v2/#definition-ideal_request
type IdealRequest struct {
	NameCountryCodeBase
	Bic string `json:"bic,omitempty"`
}

// MybankRequest
// https://developer.paypal.com/docs/api/orders/v2/#definition-mybank_request
type MybankRequest struct {
	NameCountryCodeBase
}

// P24Request
// https://developer.paypal.com/docs/api/orders/v2/#definition-p24_request
type P24Request struct {
	NameCountryCodeBase
	Email string `json:"email"`
}

// SofortRequest
// https://developer.paypal.com/docs/api/orders/v2/#definition-sofort_request
type SofortRequest struct {
	NameCountryCodeBase
}

// TrustlyRequest
// https://developer.paypal.com/docs/api/orders/v2/#definition-trustly_request
type TrustlyRequest struct {
	NameCountryCodeBase
}

// OrderConfirmApplicationContext
// https://developer.paypal.com/docs/api/orders/v2/#definition-order_confirm_application_context
type OrderConfirmApplicationContext struct {
	BrandName string `json:"brand_name,omitempty"`
	Locale    string `json:"locale,omitempty"`
	ReturnUrl string `json:"return_url,omitempty"`
	CancelUrl string `json:"cancel_url,omitempty"`
}

// AuthorizationStatusDetails
// https://developer.paypal.com/docs/api/payments/v2/#definition-authorization_status_details
type AuthorizationStatusDetails struct {
	Reason string `json:"reason,omitempty"`
}

// SenderBatchHeader
// https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#definition-sender_batch_header
type SenderBatchHeader struct {
	SenderBatchId string `json:"sender_batch_id,omitempty"`
	RecipientType string `json:"recipient_type,omitempty"`
	EmailSubject  string `json:"email_subject,omitempty"`
	EmailMessage  string `json:"email_message,omitempty"`
	Note          string `json:"note,omitempty"`
}

// PayoutItem
// https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#definition-payout_item
type PayoutItem struct {
	RecipientType               string                       `json:"recipient_type,omitempty"`
	Amount                      *Currency                    `json:"amount"`
	Note                        string                       `json:"note,omitempty"`
	Receiver                    string                       `json:"receiver"`
	SenderItemId                string                       `json:"sender_item_id,omitempty"`
	RecipientWallet             string                       `json:"recipient_wallet,omitempty"`
	AlternateNotificationMethod *AlternateNotificationMethod `json:"alternate_notification_method,omitempty"`
	NotificationLanguage        string                       `json:"notification_language,omitempty"`
	ApplicationContext          *ApplicationContext          `json:"application_context,omitempty"`
}

// Currency
// https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#definition-currency
type Currency struct {
	Currency string `json:"currency"`
	Value    string `json:"value"`
}

// AlternateNotificationMethod
// https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#definition-alternate_notification_method
type AlternateNotificationMethod struct {
	Phone *Phone `json:"phone,omitempty"`
}

// Phone
// https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#definition-phone
type Phone struct {
	CountryCode     string `json:"country_code"`
	NationalNumber  string `json:"national_number"`
	ExtensionNumber string `json:"extension_number,omitempty"`
}

// ApplicationContext
// https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#definition-application_context
type ApplicationContext struct {
	SocialFeedPrivacy string `json:"social_feed_privacy,omitempty"`
	HollerUrl         string `json:"holler_url,omitempty"`
	LogoUrl           string `json:"logo_url,omitempty"`
}

// PayoutHeader
// https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#definition-payout_header
type PayoutHeader struct {
	PayoutBatchId     string                   `json:"payout_batch_id"`
	BatchStatus       string                   `json:"batch_status"`
	TimeCreated       string                   `json:"time_created"`
	SenderBatchHeader *PayoutSenderBatchHeader `json:"sender_batch_header"`
}

// PayoutSenderBatchHeader
// https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#definition-payout_sender_batch_header
type PayoutSenderBatchHeader struct {
	SenderBatchId string `json:"sender_batch_id,omitempty"`
	RecipientType string `json:"recipient_type,omitempty"`
	EmailSubject  string `json:"email_subject,omitempty"`
	EmailMessage  string `json:"email_message,omitempty"`
}

// PayoutBatchHeader
// https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#definition-payout_batch_header
type PayoutBatchHeader struct {
	PayoutBatchId     string                   `json:"payout_batch_id"`
	BatchStatus       string                   `json:"batch_status"`
	TimeCreated       string                   `json:"time_created,omitempty"`
	TimeComplete      string                   `json:"time_complete,omitempty"`
	TimeClosed        string                   `json:"time_closed,omitempty"`
	SenderBatchHeader *PayoutSenderBatchHeader `json:"sender_batch_header"`
	FundingSource     string                   `json:"funding_source,omitempty"`
	Amount            *Currency                `json:"amount,omitempty"`
	Fees              *Currency                `json:"fees,omitempty"`
}

// PayoutBatchItems
// https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#definition-payout_batch_items
type PayoutBatchItems struct {
	PayoutItemId      string            `json:"payout_item_id"`
	TransactionId     string            `json:"transaction_id,omitempty"`
	ActivityId        string            `json:"activity_id,omitempty"`
	TransactionStatus string            `json:"transaction_status,omitempty"`
	PayoutItemFee     *Currency         `json:"payout_item_fee,omitempty"`
	PayoutBatchId     string            `json:"payout_batch_id"`
	PayoutItem        *PayoutItemDetail `json:"payout_item"`
}

// PayoutItemDetail
// https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#definition-payout_item_detail
type PayoutItemDetail struct {
	RecipientType   string      `json:"recipient_type,omitempty"`
	Amount          *Currency   `json:"amount"`
	Note            string      `json:"note,omitempty"`
	Receiver        string      `json:"receiver"`
	SenderItemId    string      `json:"sender_item_id,omitempty"`
	RecipientName   *PayoutName `json:"recipient_name,omitempty"`
	RecipientWallet string      `json:"recipient_wallet,omitempty"`
}

// PayoutName
// https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#definition-name
type PayoutName struct {
	Prefix            string `json:"prefix,omitempty"`
	GiveName          string `json:"give_name,omitempty"`
	Surname           string `json:"surname,omitempty"`
	MiddleName        string `json:"middle_name,omitempty"`
	Suffix            string `json:"suffix,omitempty"`
	AlternateFullName string `json:"alternate_full_name,omitempty"`
	FullName          string `json:"full_name,omitempty"`
}

// PayoutCurrencyConversion
// https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#definition-payout_currency_conversion
type PayoutCurrencyConversion struct {
	FromAmount   *Currency `json:"from_amount,omitempty"`
	ToAmount     *Currency `json:"to_amount,omitempty"`
	ExchangeRate string    `json:"exchange_rate"`
}

// Error
// https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#definition-error
type Error struct {
	Name            string             `json:"name"`
	Message         string             `json:"message"`
	DebugId         string             `json:"debug_id"`
	InformationLink string             `json:"information_link,omitempty"`
	Details         []*ErrorDetails    `json:"details,omitempty"`
	Links           []*LinkDescription `json:"links,omitempty"`
}

// ErrorDetails
// https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#definition-error_details
type ErrorDetails struct {
	Field       string `json:"field,omitempty"`
	Value       string `json:"value,omitempty"`
	Location    string `json:"location,omitempty"`
	Issue       string `json:"issue"`
	Description string `json:"description,omitempty"`
}

// TransactionDetail
// https://developer.paypal.com/docs/api/transaction-search/v1/#definition-transaction_detail
type TransactionDetail struct {
	TransactionInfo *TransactionInfo `json:"transaction_info,omitempty"`
	PayerInfo       *PayerInfo       `json:"payer_info,omitempty"`
	ShippingInfo    *ShippingInfo    `json:"shipping_info,omitempty"`
	CartInfo        *CartInfo        `json:"cart_info,omitempty"`
	StoreInfo       *StoreInfo       `json:"store_info,omitempty"`
	AuctionInfo     *AuctionInfo     `json:"auction_info,omitempty"`
	IncentiveInfo   *IncentiveInfo   `json:"incentive_info,omitempty"`
}

// TransactionInfo
// https://developer.paypal.com/docs/api/transaction-search/v1/#definition-transaction_info
type TransactionInfo struct {
	PaypalAccountId           string `json:"paypal_account_id,omitempty"`
	TransactionId             string `json:"transaction_id,omitempty"`
	PaypalReferenceId         string `json:"paypal_reference_id,omitempty"`
	PaypalReferenceIdType     string `json:"paypal_reference_id_type,omitempty"`
	TransactionEventCode      string `json:"transaction_event_code,omitempty"`
	TransactionInitiationDate string `json:"transaction_initiation_date,omitempty"`
	TransactionUpdatedDate    string `json:"transaction_updated_date,omitempty"`
	TransactionAmount         *Money `json:"transaction_amount,omitempty"`
	FeeAmount                 *Money `json:"fee_amount,omitempty"`
	DiscountAmount            *Money `json:"discount_amount,omitempty"`
	InsuranceAmount           *Money `json:"insurance_amount,omitempty"`
	ShippingAmount            *Money `json:"shipping_amount,omitempty"`
	ShippingDiscountAmount    *Money `json:"shipping_discount_amount,omitempty"`
	ShippingTaxAmount         *Money `json:"shipping_tax_amount,omitempty"`
	OtherAmount               *Money `json:"other_amount,omitempty"`
	TipAmount                 *Money `json:"tip_amount,omitempty"`
	TransactionStatus         string `json:"transaction_status,omitempty"`
	TransactionSubject        string `json:"transaction_subject,omitempty"`
	TransactionNote           string `json:"transaction_note,omitempty"`
	PaymentTrackingId         string `json:"payment_tracking_id,omitempty"`
	BankReferenceId           string `json:"bank_reference_id,omitempty"`
	EndingBalance             *Money `json:"ending_balance,omitempty"`
	AvailableBalance          *Money `json:"available_balance,omitempty"`
	InvoiceId                 string `json:"invoice_id,omitempty"`
	CustomField               string `json:"custom_field,omitempty"`
	ProtectionEligibility     string `json:"protection_eligibility,omitempty"`
	CreditTerm                string `json:"credit_term,omitempty"`
	CreditTransactionalFee    *Money `json:"credit_transactional_fee,omitempty"`
	CreditPromotionalFee      *Money `json:"credit_promotional_fee,omitempty"`
	AnnualPercentageRate      string `json:"annual_percentage_rate,omitempty"`
	PaymentMethodType         string `json:"payment_method_type,omitempty"`
	InstrumentType            string `json:"instrument_type,omitempty"`
	InstrumentSubType         string `json:"instrument_sub_type,omitempty"`
}

// PayerInfo
// https://developer.paypal.com/docs/api/transaction-search/v1/#definition-payer_info
type PayerInfo struct {
	AccountId     string   `json:"account_id,omitempty"`
	EmailAddress  string   `json:"email_address,omitempty"`
	PhoneNumber   string   `json:"phone_number,omitempty"`
	AddressStatus string   `json:"address_status,omitempty"`
	PaterStatus   string   `json:"pater_status,omitempty"`
	PayerName     *Name    `json:"payer_name,omitempty"`
	CountryCode   string   `json:"country_code,omitempty"`
	Address       *Address `json:"address,omitempty"`
}

// Address
// https://developer.paypal.com/docs/api/transaction-search/v1/#definition-address
type Address struct {
	Line1       string `json:"line_1"`
	Line2       string `json:"line_2,omitempty"`
	City        string `json:"city"`
	State       string `json:"state,omitempty"`
	CountryCode string `json:"country_code"`
	PostalCode  string `json:"postal_code,omitempty"`
}

// ShippingInfo
// https://developer.paypal.com/docs/api/transaction-search/v1/#definition-shipping_info
type ShippingInfo struct {
	Name                     string   `json:"name,omitempty"`
	Method                   string   `json:"method,omitempty"`
	Address                  *Address `json:"address,omitempty"`
	SecondaryShippingAddress *Address `json:"secondary_shipping_address,omitempty"`
}

// CartInfo
// https://developer.paypal.com/docs/api/transaction-search/v1/#definition-cart_info
type CartInfo struct {
	ItemDetails     []*ItemDetail `json:"item_details,omitempty"`
	TaxInclusive    bool          `json:"tax_inclusive,omitempty"`
	PaypalInvoiceId string        `json:"paypal_invoice_id,omitempty"`
}

// ItemDetail
// https://developer.paypal.com/docs/api/transaction-search/v1/#definition-item_detail
type ItemDetail struct {
	ItemCode            string                 `json:"item_code,omitempty"`
	ItemName            string                 `json:"item_name,omitempty"`
	ItemDescription     string                 `json:"item_description,omitempty"`
	ItemOptions         string                 `json:"item_options,omitempty"`
	ItemQuantity        string                 `json:"item_quantity,omitempty"`
	ItemUnitPrice       *Money                 `json:"item_unit_price,omitempty"`
	ItemAmount          *Money                 `json:"item_amount,omitempty"`
	DiscountAmount      *Money                 `json:"discount_amount,omitempty"`
	AdjustmentAmount    *Money                 `json:"adjustment_amount,omitempty"`
	GiftWrapAmount      *Money                 `json:"gift_wrap_amount,omitempty"`
	TaxPercentage       string                 `json:"tax_percentage,omitempty"`
	TaxAmounts          []*ItemDetailTaxAmount `json:"tax_amounts,omitempty"`
	BasicShippingAmount *Money                 `json:"basic_shipping_amount,omitempty"`
	ExtraShippingAmount *Money                 `json:"extra_shipping_amount,omitempty"`
	HandlingAmount      *Money                 `json:"handling_amount,omitempty"`
	InsuranceAmount     *Money                 `json:"insurance_amount,omitempty"`
	TotalItemAmount     *Money                 `json:"total_item_amount,omitempty"`
	InvoiceNumber       string                 `json:"invoice_number,omitempty"`
	CheckoutOptions     []CheckoutOption       `json:"checkout_options,omitempty"`
}

// ItemDetailTaxAmount
// https://developer.paypal.com/docs/api/transaction-search/v1/#definition-item_detail_tax_amount
type ItemDetailTaxAmount struct {
	TaxAmount *Money `json:"tax_amount,omitempty"`
}

// CheckoutOption
// https://developer.paypal.com/docs/api/transaction-search/v1/#definition-checkout_option
type CheckoutOption struct {
	CheckoutOptionName  string `json:"checkout_option_name,omitempty"`
	CheckoutOptionValue string `json:"checkout_option_value,omitempty"`
}

// StoreInfo
// https://developer.paypal.com/docs/api/transaction-search/v1/#definition-store_info
type StoreInfo struct {
	StoreId    string `json:"store_id,omitempty"`
	TerminalId string `json:"terminal_id,omitempty"`
}

// AuctionInfo
// https://developer.paypal.com/docs/api/transaction-search/v1/#definition-auction_info
type AuctionInfo struct {
	AuctionSite        string `json:"auction_site,omitempty"`
	AuctionItemSite    string `json:"auction_item_site,omitempty"`
	AuctionBuyerId     string `json:"auction_buyer_id,omitempty"`
	AuctionClosingDate string `json:"auction_closing_date,omitempty"`
}

// IncentiveInfo
// https://developer.paypal.com/docs/api/transaction-search/v1/#definition-incentive_info
type IncentiveInfo struct {
	IncentiveDetails []*IncentiveDetail `json:"incentive_details,omitempty"`
}

// IncentiveDetail
// https://developer.paypal.com/docs/api/transaction-search/v1/#definition-incentive_detail
type IncentiveDetail struct {
	IncentiveType        string `json:"incentive_type,omitempty"`
	IncentiveCode        string `json:"incentive_code,omitempty"`
	IncentiveAmount      *Money `json:"incentive_amount,omitempty"`
	IncentiveProgramCode string `json:"incentive_program_code,omitempty"`
}

// BalanceDetail
// https://developer.paypal.com/docs/api/transaction-search/v1/#definition-balance_detail
type BalanceDetail struct {
	Currency         string `json:"currency"`
	Primary          bool   `json:"primary,omitempty"`
	TotalBalance     *Money `json:"total_balance"`
	AvailableBalance *Money `json:"available_balance,omitempty"`
	WithheldBalance  *Money `json:"withheld_balance,omitempty"`
}

// Webhook
// https://developer.paypal.com/docs/api/webhooks/v1/#definition-webhook
type Webhook struct {
	Id         string             `json:"id,omitempty"`
	Url        string             `json:"url"`
	EventTypes []*EventType       `json:"event_types"`
	Links      []*LinkDescription `json:"links,omitempty"`
}

// EventType
// https://developer.paypal.com/docs/api/webhooks/v1/#definition-event_type
type EventType struct {
	Name             string             `json:"name"`
	Description      string             `json:"description,omitempty"`
	Status           string             `json:"status,omitempty"`
	ResourceVersions []*ResourceVersion `json:"resource_versions,omitempty"`
}

// ResourceVersion
// https://developer.paypal.com/docs/api/webhooks/v1/#definition-resource_version
type ResourceVersion struct {
	ResourceVersion string `json:"resource_version"`
}

// Event
// https://developer.paypal.com/docs/api/webhooks/v1/#definition-event
type Event struct {
	Id              string             `json:"id,omitempty"`
	CreateTime      string             `json:"create_time,omitempty"`
	ResourceType    string             `json:"resource_type,omitempty"`
	EventVersion    string             `json:"event_version,omitempty"`
	EventType       string             `json:"event_type,omitempty"`
	Summary         string             `json:"summary,omitempty"`
	ResourceVersion string             `json:"resource_version,omitempty"`
	Resource        *Resource          `json:"resource,omitempty"`
	Links           []*LinkDescription `json:"links,omitempty"`
}

// Resource
// https://developer.paypal.com/docs/api/webhooks/v1/#definition-resource
type Resource struct {
	Id            string             `json:"id"`
	CreateTime    string             `json:"create_time"`
	UpdateTime    string             `json:"update_time"`
	State         string             `json:"state"`
	Amount        *Amount            `json:"amount"`
	ParentPayment string             `json:"parent_payment"`
	ValidUntil    string             `json:"valid_until"`
	Links         []*LinkDescription `json:"links,omitempty"`
}
type Amount struct {
	Total    string  `json:"total"`
	Currency string  `json:"currency"`
	Details  *Detail `json:"details"`
}
type Detail struct {
	Subtotal string `json:"subtotal"`
}
