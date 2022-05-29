package paypal

import (
	"encoding/json"
	"fmt"
	"github/mm-ooto/paypal/consts"
	"github/mm-ooto/paypal/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type PClient struct {
	gateWayUrl string // 网关地址
	//version              string // 版本号：v1/v2
	clientId             string //
	clientSecret         string //
	isProduction         bool   // 是否是生产环境
	client               *http.Client
	accessToken          string    // 访问令牌
	accessTokenExpiresAt time.Time // 访问令牌过期时间
}

type OptionFunc func(p *PClient)

func AddClient(client *http.Client) OptionFunc {
	return func(p *PClient) {
		p.client = client
	}
}

var pClientV *PClient

func NewPClient(clientId, clientSecret string, isProduction bool, opts ...OptionFunc) (pClient *PClient) {
	pClient = &PClient{
		gateWayUrl:   consts.SandboxGateWayUrl,
		clientId:     clientId,
		clientSecret: clientSecret,
		isProduction: isProduction,
		client:       http.DefaultClient,
	}
	if isProduction {
		pClient.gateWayUrl = consts.ProductionGateWayUrl
	}
	for _, opt := range opts {
		opt(pClient)
	}
	// 访问令牌
	err := pClient.GetAccessToken()
	if err != nil {
		panic(err)
	}
	pClientV = pClient
	return
}

func GetPClient() *PClient {
	return pClientV
}

// GenerateApiUrl 生成api url
func (p *PClient) GenerateApiUrl(api string) string {
	return fmt.Sprintf("%s%s", p.gateWayUrl, api)
}

// GetAccessToken 获取访问令牌
func (p *PClient) GetAccessToken() error {
	value := url.Values{}
	value.Add("grant_type", "client_credentials")
	req, err := http.NewRequest(consts.HttpMethodPost, p.GenerateApiUrl(consts.AccessTokenAPI), strings.NewReader(value.Encode()))
	if err != nil {
		return err
	}
	req.SetBasicAuth(p.clientId, p.clientSecret)
	req.Header.Set("Content-Type", consts.ContentTypeFormUrlencoded)
	res, err := p.client.Do(req)
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var resAccessToken model.ResAccessToken
	if err = json.Unmarshal(bytes, &resAccessToken); err != nil {
		return err
	}
	p.accessToken = resAccessToken.AccessToken
	p.accessTokenExpiresAt = time.Now().Add(time.Duration(resAccessToken.ExpiresIn) * time.Second)
	return nil
}

func (p *PClient) CallApiRequest(httpMethod, api string, apiReq interface{}, apiResp interface{}) error {
	if p.accessToken == "" || p.accessTokenExpiresAt.Before(time.Now()) {
		// 过期重新获取访问令牌
		err := p.GetAccessToken()
		if err != nil {
			panic(err)
		}
	}

	httpMethod = strings.ToUpper(httpMethod)
	reqData, err := json.Marshal(apiReq)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(httpMethod, p.GenerateApiUrl(api), strings.NewReader(string(reqData)))
	if err != nil {
		return err
	}

	req.Header.Set("Accept-Language", "en_US")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", p.accessToken))
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", consts.ContentTypeJson)
	}

	res, err := p.client.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	//fmt.Println("statusCode=", res.StatusCode)
	// 请求成功
	switch res.StatusCode {
	case http.StatusOK, http.StatusCreated:
		if apiResp != nil {
			if err = json.Unmarshal(body, &apiResp); err != nil {
				return err
			}
		}
	case http.StatusNoContent:
		return nil
	case http.StatusUnauthorized:
		var resValidationError = &ValidationError{
			RequestUrl: res.Request.URL.String(),
		}
		if err = json.Unmarshal(body, resValidationError); err != nil {
			return err
		}
		//bytes, _ := json.Marshal(resValidationError)
		//fmt.Println(string(bytes))
		return resValidationError
	case http.StatusForbidden:
		var requestForbiddenErr = &ForbiddenErr{
			RequestUrl: res.Request.URL.String(),
		}
		if err = json.Unmarshal(body, requestForbiddenErr); err != nil {
			return err
		}
		//bytes, _ := json.Marshal(requestForbiddenErr)
		//fmt.Println(string(bytes))
		return requestForbiddenErr
	default:
		var resRequestFailed = &ResponseError{
			RequestUrl: res.Request.URL.String(),
		}
		if err = json.Unmarshal(body, resRequestFailed); err != nil {
			return err
		}
		//bytes, _ := json.Marshal(resRequestFailed)
		//fmt.Println(string(bytes))
		return resRequestFailed
	}

	return nil
}
