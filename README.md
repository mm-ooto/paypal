# Paypal for Golang

## 初始化Client需要的数据：
`clientId`, `clientSecret`

## 示例
```go
var pClient *paypal.PClient

func TestMain(m *testing.M) {
	clientId := "AcSU4Cbl7gbGZqI1xTNiJwkkpEzLpb0QntUmotlM239mz_8N5CYJlGfbByidGWXuoo6lDI06hgGfpyy5"
	clientSecret := "EBOvatSnBAMBls3YdsFKzvuXljxFxb3sX9ZNXZNkhJpdyLS_IKIbVV3Xkh04JAmDE2meJfeMO4ZMnRPS"
	pClient = paypal.NewPClient(clientId, clientSecret, false)
	m.Run()
}

func TestGetAccessToken(t *testing.T) {
    err := pClient.GetAccessToken()
    if err != nil {
        t.Log(err)
        return
    }
    bytes, _ := json.Marshal(pClient)
    t.Log(string(bytes))
}

func TestCreateOrder(t *testing.T) {
    req := &model.ReqCreateOrder{
        Intent: consts.IntentCapture,
        PurchaseUnits: []*model.PurchaseUnitsRequest{
            {
                Amount: &model.AmountWithBreakdown{
                    CurrencyCode: "CAD",
                    Value:        "10.00",
                },
            },
        },
    }
    res, err := pClient.CreateOrder(req)
    if err != nil {
        t.Log(err.Error())
        return
    }
    bytes, _ := json.Marshal(res)
    t.Log(string(bytes))
}
```

### 已支持的API

#### AccessToken
* 获取访问令牌（Get AccessToken：`POST /v1/oauth2/token`）：`pClient.GetAccessToken()`

#### 订单
* 创建订单（Create order：`POST /v2/checkout/orders`）：`pClient.CreateOrder()`
* 更新订单（Update order：`PATCH /v2/checkout/orders/{id}`）：`pClient.UpdateOrder()`
* 获取订单详情（Show order details：`GET /v2/checkout/orders/{id}`）：`pClient.ShowOrderDetails()`
* 订单授权（Authorize payment for order：`POST /v2/checkout/orders/{id}/authorize`）：`pClient.OrderAuthorize()`
* 订单捕获（Capture payment for order：`POST /v2/checkout/orders/{id}/capture`）：`pClient.OrdersCapture()`
* 订单确认（Confirm the Order：`POST /v2/checkout/orders/{id}/confirm-payment-source`）：`pClient.ConfirmOrder()`

#### 支出
* 批量支出创建（Create batch payout：`POST /v1/payments/payouts`）：`pClient.CreateBatchPayout()`
* 批量获取支出详情（Show payout batch details：`GET /v1/payments/payouts/{payout_batch_id}`）：`pClient.GetPayoutBatchDetails()`
* 获取支出的项目详情（Show payout item details：`GET /v1/payments/payouts-item/{payout_item_id}`）：`pClient.GetPayoutItemDetails()`
* 取消无人认领的支出项目（Cancel unclaimed payout item：`POST /v1/payments/payouts-item/{payout_item_id}/cancel`）：`pClient.CancelUnclaimedPayoutItem()`

#### 付款
* 获取付款授权详情（Show details for authorized payment：`GET /v2/payments/authorizations/{authorization_id}`）：`pClient.GetAuthorizedPaymentDetails()`
* 付款授权捕获（Capture authorized payment：`POST /v2/payments/authorizations/{authorization_id}/capture`）：`pClient.CaptureAuthorizedPayment()`
* 付款重新授权（Reauthorize authorized payment：`POST /v2/payments/authorizations/{authorization_id}/reauthorize`）：`pClient.ReauthorizeAuthorizedPayment()`
* 无效的授权付款（Void authorized payment：`POST /v2/payments/authorizations/{authorization_id}/void`）：`pClient.VoidAuthorizedPayment()`
* 获取捕获的付款详情（Show captured payment details：`GET /v2/payments/captures/{capture_id}`）：`pClient.GetCapturedPaymentDetails()`
* 退还被捕获的付款（Refund captured payment：`GET /v2/payments/captures/{capture_id}/refund`）：`pClient.RefundCapturedPayment()`
* 获取退款详情（Show refund details：`GET /v2/payments/refunds/{refund_id}`）：`pClient.GetRefundDetails()`

#### 交易查询
* 获取交易列表（List transactions：`GET /v1/reporting/transactions`）：`pClient.GetTransactionsList()`
* 获取所有余额（List all balances：`GET /v1/reporting/balances`）：`pClient.GetListAllBalances()`

#### 网络钩子Webhooks
* 获取所有的钩子（List webhooks：`GET /v1/notifications/webhooks`）：`pClient.GetListWebhooks()`
* 创建钩子（Create webhook：`POST /v1/notifications/webhooks`）：`pClient.CreateWebhook()`
* 删除钩子（Delete webhook：`DELETE /v1/notifications/webhooks/{webhook_id}`）：`pClient.DeleteWebhook()`
* 更新钩子（Update webhook：`PATCH /v1/notifications/webhooks/{webhook_id}`）：`pClient.UpdateWebhook()`
* 获取钩子详情（Show webhook details：`GET /v1/notifications/webhooks/{webhook_id}`）：`pClient.GetWebhookDetails()`
* 获取钩子的所有订阅事件（List event subscriptions for webhook：`GET /v1/notifications/webhooks/{webhook_id}/event-types`）：`pClient.GetListEventSubscriptionsForWebhook()`
* 验证钩子签名（Verify webhook signature：`POST /v1/notifications/verify-webhook-signature`）： `pClient.VerifyWebhookSignature()`


### 参考：
* 注册地址：https://www.paypal.com/us/webapps/mpp/account-selection
* 开发者网站：https://developer.paypal.com/
* 沙箱环境账号：https://developer.paypal.com/developer/accounts/  
* Demo地址：https://demo.paypal.com/c2/demo/code_samples