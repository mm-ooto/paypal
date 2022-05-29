package paypal

/**
*	订单
*	https://developer.paypal.com/docs/api/orders/v2/#orders
 */
import (
	"fmt"
	"github/mm-ooto/paypal/consts"
	"github/mm-ooto/paypal/model"
)

// CreateOrder 创建订单
func (p *PClient) CreateOrder(req *model.ReqCreateOrder) (res model.ResCreateOrder, err error) {
	if err = p.CallApiRequest(consts.HttpMethodPost, consts.OrderAPI, req, &res); err != nil {
		return
	}
	return
}

// UpdateOrder 更新订单
func (p *PClient) UpdateOrder(req *model.ReqUpdateOrder) (err error) {
	api := fmt.Sprintf("%s/%s", consts.OrderAPI, req.Id)
	if err = p.CallApiRequest(consts.HttpMethodPatch, api, req, nil); err != nil {
		return
	}
	return
}

// ShowOrderDetails 获取订单详情
func (p *PClient) ShowOrderDetails(req *model.ReqShowOrderDetails) (res model.ResShowOrderDetails, err error) {
	api := fmt.Sprintf("%s?%s", fmt.Sprintf("%s/%s", consts.OrderAPI, req.Id), req.GetQueryParameters())
	if err = p.CallApiRequest(consts.HttpMethodGet, api, nil, &res); err != nil {
		return
	}
	return
}

// OrderAuthorize 订单授权
func (p *PClient) OrderAuthorize(req *model.ReqOrderAuthorize) (res model.ResOrderAuthorize, err error) {
	api := fmt.Sprintf(consts.OrderAuthorizeAPI, req.Id)
	if err = p.CallApiRequest(consts.HttpMethodPost, api, req, &res); err != nil {
		return
	}
	return
}

// OrdersCapture 订单捕获
func (p *PClient) OrdersCapture(req *model.ReqOrdersCapture) (res model.ResOrdersCapture, err error) {
	api := fmt.Sprintf(consts.OrderCaptureAPI, req.Id)
	if err = p.CallApiRequest(consts.HttpMethodPost, api, req, &res); err != nil {
		return
	}
	return
}

// ConfirmOrder 订单确认
func (p *PClient) ConfirmOrder(req *model.ReqConfirmOrder) (res model.ResConfirmOrder, err error) {
	api := fmt.Sprintf(consts.OrderConfirmAPI, req.Id)
	if err = p.CallApiRequest(consts.HttpMethodPost, api, req, &res); err != nil {
		return
	}
	return
}
