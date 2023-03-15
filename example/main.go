package main

import (
	"github/mm-ooto/paypal"
	"net/http"
)

var client *paypal.PClient

func init() {
	clientId := "AcSU4Cbl7gbGZqI1xTNiJwkkpEzLpb0QntUmotlM239mz_8N5CYJlGfbByidGWXuoo6lDI06hgGfpyy5"
	clientSecret := "EBOvatSnBAMBls3YdsFKzvuXljxFxb3sX9ZNXZNkhJpdyLS_IKIbVV3Xkh04JAmDE2meJfeMO4ZMnRPS"
	client = paypal.NewPClient(clientId, clientSecret, false)
}
func main() {
	http.HandleFunc("/notify", func(writer http.ResponseWriter, request *http.Request) {
		// 通知处理
		if err := client.HandleWebhookAsyncNotify("9FM75813N28895501", request); err != nil {
			writer.Write([]byte(err.Error()))
			return
		}
		writer.WriteHeader(http.StatusOK)
	})
	http.ListenAndServe(":8099", nil)
}
