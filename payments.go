package paypal

/*
*	付款
*	https://developer.paypal.com/docs/api/payments/v2/#authorizations
*/
import (
	"fmt"
	"github/mm-ooto/paypal/consts"
	"github/mm-ooto/paypal/model"
)

// GetAuthorizedPaymentDetails 获取付款授权详情
func (p *PClient) GetAuthorizedPaymentDetails(req *model.ReqGetAuthorizedPaymentDetails) (res model.ResGetAuthorizedPaymentDetails, err error) {
	api := fmt.Sprintf(consts.AuthorizedPaymentDetailsAPI, req.AuthorizationId)
	if err = p.CallApiRequest(consts.HttpMethodGet, api, nil, &res); err != nil {
		return
	}
	return
}

// CaptureAuthorizedPayment 付款授权捕获
func (p *PClient) CaptureAuthorizedPayment(req *model.ReqCaptureAuthorizedPayment) (res model.ResCaptureAuthorizedPayment, err error) {
	api := fmt.Sprintf(consts.CaptureAuthorizedPaymentAPI, req.AuthorizationId)
	if err = p.CallApiRequest(consts.HttpMethodPost, api, req, &res); err != nil {
		return
	}
	return
}

// ReauthorizeAuthorizedPayment 付款重新授权
func (p *PClient) ReauthorizeAuthorizedPayment(req *model.ReqReauthorizeAuthorizedPayment) (res model.ResReauthorizeAuthorizedPayment, err error) {
	api := fmt.Sprintf(consts.ReauthorizeAuthorizedPaymentAPI, req.AuthorizationId)
	if err = p.CallApiRequest(consts.HttpMethodPost, api, req, &res); err != nil {
		return
	}
	return
}

// VoidAuthorizedPayment 无效的授权付款
func (p *PClient) VoidAuthorizedPayment(req *model.ReqVoidAuthorizedPayment) (err error) {
	api := fmt.Sprintf(consts.VoidAuthorizedPaymentAPI, req.AuthorizationId)
	if err = p.CallApiRequest(consts.HttpMethodPost, api, req, nil); err != nil {
		return
	}
	return
}

// GetCapturedPaymentDetails 获取捕获的付款详情
func (p *PClient) GetCapturedPaymentDetails(req *model.ReqGetCapturedPaymentDetails) (res model.ResGetCapturedPaymentDetails, err error) {
	api := fmt.Sprintf(consts.CapturedPaymentDetailsAPI, req.CaptureId)
	if err = p.CallApiRequest(consts.HttpMethodGet, api, nil, &res); err != nil {
		return
	}
	return
}

// RefundCapturedPayment 退还被捕获的付款
func (p *PClient) RefundCapturedPayment(req *model.ReqRefundCapturedPayment) (res model.ResRefundCapturedPayment, err error) {
	api := fmt.Sprintf(consts.RefundCapturedPaymentAPI, req.CaptureId)
	if err = p.CallApiRequest(consts.HttpMethodPost, api, req, &res); err != nil {
		return
	}
	return
}

// GetRefundDetails 获取退款详情
func (p *PClient) GetRefundDetails(req *model.ReqGetRefundDetails) (res model.ResGetRefundDetails, err error) {
	api := fmt.Sprintf(consts.RefundDetailsAPI, req.RefundId)
	if err = p.CallApiRequest(consts.HttpMethodPost, api, nil, &res); err != nil {
		return
	}
	return
}
