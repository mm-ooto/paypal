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


### 参考：
* 注册地址：https://www.paypal.com/us/webapps/mpp/account-selection
* 开发者网站：https://developer.paypal.com/
* 沙箱环境账号：https://developer.paypal.com/developer/accounts/  
* Demo地址：https://demo.paypal.com/c2/demo/code_samples