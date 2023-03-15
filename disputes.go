package paypal

import (
	"fmt"
	"github/mm-ooto/paypal/consts"
	"github/mm-ooto/paypal/model"
)

/*
	纠纷/争议
*/

// DisputesList 争议列表
func (p *PClient) DisputesList(req *model.ReqDisputesList) (res model.ResDisputesList, err error) {
	if err = p.CallApiRequest(consts.HttpMethodGet, consts.DisputesListAPI, req, &res); err != nil {
		return
	}
	return
}

// DisputesDetail 争议详情
func (p *PClient) DisputesDetail(req *model.ReqDisputesDetail) (res model.ReqDisputesDetail, err error) {
	api := fmt.Sprintf(consts.DisputesDetailAPI, req.Id)
	if err = p.CallApiRequest(consts.HttpMethodGet, api, req, &res); err != nil {
		return
	}
	return
}
