package paypal

/**
*	订单
*	https://developer.paypal.com/docs/api/orders/v2/#orders
 */
import (
	"fmt"
	"paypal/model"
)

// CreateOrder 创建订单
func (p *Client) CreateOrder(req *model.ReqCreateOrder) (res model.ResCreateOrder, err error) {
	if err = p.CallApiRequest(HttpMethodPost, OrderAPI, req, &res); err != nil {
		return
	}
	return
}

// UpdateOrder 更新订单
func (p *Client) UpdateOrder(req *model.ReqUpdateOrder) (err error) {
	api := fmt.Sprintf("%s/%s", OrderAPI, req.Id)
	if err = p.CallApiRequest(HttpMethodPatch, api, req, nil); err != nil {
		return
	}
	return
}

// ShowOrderDetails 获取订单详情
func (p *Client) ShowOrderDetails(req *model.ReqShowOrderDetails) (res model.ResShowOrderDetails, err error) {
	api := fmt.Sprintf("%s?%s", fmt.Sprintf("%s/%s", OrderAPI, req.Id), req.GetQueryParameters())
	if err = p.CallApiRequest(HttpMethodGet, api, nil, &res); err != nil {
		return
	}
	return
}

// OrderAuthorize 订单授权
func (p *Client) OrderAuthorize(req *model.ReqOrderAuthorize) (res model.ResOrderAuthorize, err error) {
	api := fmt.Sprintf(OrderAuthorizeAPI, req.Id)
	if err = p.CallApiRequest(HttpMethodPost, api, req, &res); err != nil {
		return
	}
	return
}

// OrdersCapture 订单捕获
func (p *Client) OrdersCapture(req *model.ReqOrdersCapture) (res model.ResOrdersCapture, err error) {
	api := fmt.Sprintf(OrderCaptureAPI, req.Id)
	if err = p.CallApiRequest(HttpMethodPost, api, req, &res); err != nil {
		return
	}
	return
}

// ConfirmOrder 订单确认
func (p *Client) ConfirmOrder(req *model.ReqConfirmOrder) (res model.ResConfirmOrder, err error) {
	api := fmt.Sprintf(OrderConfirmAPI, req.Id)
	if err = p.CallApiRequest(HttpMethodPost, api, req, &res); err != nil {
		return
	}
	return
}
