package paypal

import (
	"fmt"
	"github/mm-ooto/paypal/consts"
	"github/mm-ooto/paypal/model"
)

/*
*	支出
*	https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts
*/

// CreateBatchPayout 批量支出创建
func (p *PClient) CreateBatchPayout(req *model.ReqCreateBatchPayout) (res model.ResCreateBatchPayout, err error) {
	if err = p.CallApiRequest(consts.HttpMethodPost, consts.CreateBatchPayoutAPI, req, &res); err != nil {
		return
	}
	return
}

// GetPayoutBatchDetails 批量获取支出详情
func (p *PClient) GetPayoutBatchDetails(req *model.ReqGetPayoutBatchDetails) (res model.ResGetPayoutBatchDetails, err error) {
	api := fmt.Sprintf("%s?%s", fmt.Sprintf(consts.ShowPayoutBatchDetailsAPI, req.PayoutBatchId), req.GetQueryParameters())
	if err = p.CallApiRequest(consts.HttpMethodGet, api, nil, &res); err != nil {
		return
	}
	return
}

// GetPayoutItemDetails 获取支出的项目详情
func (p *PClient) GetPayoutItemDetails(req *model.ReqGetPayoutItemDetails) (res model.ResGetPayoutItemDetails, err error) {
	api := fmt.Sprintf(consts.ShowPayoutItemDetailsAPI, req.PayoutItemId)
	if err = p.CallApiRequest(consts.HttpMethodGet, api, nil, &res); err != nil {
		return
	}
	return
}

// CancelUnclaimedPayoutItem 取消无人认领的支出项目
func (p *PClient) CancelUnclaimedPayoutItem(req *model.ReqCancelUnclaimedPayoutItem) (res model.ResCancelUnclaimedPayoutItem, err error) {
	api := fmt.Sprintf(consts.CancelUnclaimedPayoutItemAPI, req.PayoutItemId)
	if err = p.CallApiRequest(consts.HttpMethodPost, api, nil, &res); err != nil {
		return
	}
	return
}
