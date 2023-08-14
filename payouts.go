package paypal

import (
	"fmt"
	"paypal/model"
)

/*
*	支出
*	https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts
 */

// CreateBatchPayout 批量支出创建
func (p *Client) CreateBatchPayout(req *model.ReqCreateBatchPayout) (res model.ResCreateBatchPayout, err error) {
	if err = p.CallApiRequest(HttpMethodPost, CreateBatchPayoutAPI, req, &res); err != nil {
		return
	}
	return
}

// GetPayoutBatchDetails 批量获取支出详情
func (p *Client) GetPayoutBatchDetails(req *model.ReqGetPayoutBatchDetails) (res model.ResGetPayoutBatchDetails, err error) {
	api := fmt.Sprintf("%s?%s", fmt.Sprintf(ShowPayoutBatchDetailsAPI, req.PayoutBatchId), req.GetQueryParameters())
	if err = p.CallApiRequest(HttpMethodGet, api, nil, &res); err != nil {
		return
	}
	return
}

// GetPayoutItemDetails 获取支出的项目详情
func (p *Client) GetPayoutItemDetails(req *model.ReqGetPayoutItemDetails) (res model.ResGetPayoutItemDetails, err error) {
	api := fmt.Sprintf(ShowPayoutItemDetailsAPI, req.PayoutItemId)
	if err = p.CallApiRequest(HttpMethodGet, api, nil, &res); err != nil {
		return
	}
	return
}

// CancelUnclaimedPayoutItem 取消无人认领的支出项目
func (p *Client) CancelUnclaimedPayoutItem(req *model.ReqCancelUnclaimedPayoutItem) (res model.ResCancelUnclaimedPayoutItem, err error) {
	api := fmt.Sprintf(CancelUnclaimedPayoutItemAPI, req.PayoutItemId)
	if err = p.CallApiRequest(HttpMethodPost, api, nil, &res); err != nil {
		return
	}
	return
}
