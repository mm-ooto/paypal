package paypal

import (
	"fmt"
	"paypal/model"
)

/*
*	交易查询
*	https://developer.paypal.com/docs/api/transaction-search/v1/#transactions
 */

// GetTransactionsList 获取交易列表
func (p *Client) GetTransactionsList(req *model.ReqGetTransactionsList) (res model.ResGetTransactionsList, err error) {
	api := fmt.Sprintf("%s?%s", ListTransactionsAPI, req.GetQueryParameters())
	if err = p.CallApiRequest(HttpMethodGet, api, nil, &res); err != nil {
		return
	}
	return
}

// GetListAllBalances 获取所有余额
func (p *Client) GetListAllBalances(req *model.ReqGetListAllBalances) (res model.ResGetListAllBalances, err error) {
	api := fmt.Sprintf("%s?%s", ListAllBalancesAPI, req.GetQueryParameters())
	if err = p.CallApiRequest(HttpMethodGet, api, nil, &res); err != nil {
		return
	}
	return
}
