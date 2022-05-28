package paypal

import (
	"fmt"
	"github/mm-ooto/paypal/consts"
	"github/mm-ooto/paypal/model"
)

/*
	交易查询
*/

// GetTransactionsList 根据条件获取交易列表
func (p *PClient) GetTransactionsList(req *model.ReqGetTransactionsList) (res model.ResGetTransactionsList, err error) {
	api := fmt.Sprintf("%s?%s", consts.ListTransactionsAPI, req.GetQueryParameters())
	if err = p.CallApiRequest(consts.HttpMethodGet, api, nil, &res); err != nil {
		return
	}
	return
}

// GetListAllBalances 根据条件获取所有余额
func (p *PClient) GetListAllBalances(req *model.ReqGetListAllBalances) (res model.ResGetListAllBalances, err error) {
	api := fmt.Sprintf("%s?%s", consts.ListAllBalancesAPI, req.GetQueryParameters())
	if err = p.CallApiRequest(consts.HttpMethodGet, api, nil, &res); err != nil {
		return
	}
	return
}
