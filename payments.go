package paypal

/*
*	付款
*	https://developer.paypal.com/docs/api/payments/v2/#authorizations
 */
import (
	"fmt"
	"github/mm-ooto/paypal/model"
)

// GetAuthorizedPaymentDetails 获取付款授权详情
func (p *Client) GetAuthorizedPaymentDetails(req *model.ReqGetAuthorizedPaymentDetails) (res model.ResGetAuthorizedPaymentDetails, err error) {
	api := fmt.Sprintf(AuthorizedPaymentDetailsAPI, req.AuthorizationId)
	if err = p.CallApiRequest(HttpMethodGet, api, nil, &res); err != nil {
		return
	}
	return
}

// CaptureAuthorizedPayment 付款授权捕获
func (p *Client) CaptureAuthorizedPayment(req *model.ReqCaptureAuthorizedPayment) (res model.ResCaptureAuthorizedPayment, err error) {
	api := fmt.Sprintf(CaptureAuthorizedPaymentAPI, req.AuthorizationId)
	if err = p.CallApiRequest(HttpMethodPost, api, req, &res); err != nil {
		return
	}
	return
}

// ReauthorizeAuthorizedPayment 付款重新授权
func (p *Client) ReauthorizeAuthorizedPayment(req *model.ReqReauthorizeAuthorizedPayment) (res model.ResReauthorizeAuthorizedPayment, err error) {
	api := fmt.Sprintf(ReauthorizeAuthorizedPaymentAPI, req.AuthorizationId)
	if err = p.CallApiRequest(HttpMethodPost, api, req, &res); err != nil {
		return
	}
	return
}

// VoidAuthorizedPayment 无效的授权付款
func (p *Client) VoidAuthorizedPayment(req *model.ReqVoidAuthorizedPayment) (err error) {
	api := fmt.Sprintf(VoidAuthorizedPaymentAPI, req.AuthorizationId)
	if err = p.CallApiRequest(HttpMethodPost, api, req, nil); err != nil {
		return
	}
	return
}

// GetCapturedPaymentDetails 获取捕获的付款详情
func (p *Client) GetCapturedPaymentDetails(req *model.ReqGetCapturedPaymentDetails) (res model.ResGetCapturedPaymentDetails, err error) {
	api := fmt.Sprintf(CapturedPaymentDetailsAPI, req.CaptureId)
	if err = p.CallApiRequest(HttpMethodGet, api, nil, &res); err != nil {
		return
	}
	return
}

// RefundCapturedPayment 退还被捕获的付款
func (p *Client) RefundCapturedPayment(req *model.ReqRefundCapturedPayment) (res model.ResRefundCapturedPayment, err error) {
	api := fmt.Sprintf(RefundCapturedPaymentAPI, req.CaptureId)
	if err = p.CallApiRequest(HttpMethodPost, api, req, &res); err != nil {
		return
	}
	return
}

// GetRefundDetails 获取退款详情
func (p *Client) GetRefundDetails(req *model.ReqGetRefundDetails) (res model.ResGetRefundDetails, err error) {
	api := fmt.Sprintf(RefundDetailsAPI, req.RefundId)
	if err = p.CallApiRequest(HttpMethodPost, api, nil, &res); err != nil {
		return
	}
	return
}
