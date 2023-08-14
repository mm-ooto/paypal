package paypal

import (
	"fmt"
	"github/mm-ooto/paypal/model"
)

/*
	纠纷/争议
*/

// DisputesList 争议列表
func (p *Client) DisputesList(req *model.ReqDisputesList) (res model.ResDisputesList, err error) {
	if err = p.CallApiRequest(HttpMethodGet, DisputesListAPI, req, &res); err != nil {
		return
	}
	return
}

// DisputesDetail 争议详情
func (p *Client) DisputesDetail(req *model.ReqDisputesDetail) (res model.ReqDisputesDetail, err error) {
	api := fmt.Sprintf(DisputesDetailAPI, req.Id)
	if err = p.CallApiRequest(HttpMethodGet, api, req, &res); err != nil {
		return
	}
	return
}
