package paypal

import (
	"encoding/json"
	"fmt"
	"github/mm-ooto/paypal/model"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	gateWayUrl string // 网关地址
	//version              string // 版本号：v1/v2
	clientId             string // 应用编号
	clientSecret         string // 应用密钥
	isProduction         bool   // 是否是生产环境
	client               *http.Client
	accessToken          string    // 访问令牌
	accessTokenExpiresAt time.Time // 访问令牌过期时间，单位秒
}

type OptionFunc func(p *Client)

func AddClient(client *http.Client) OptionFunc {
	return func(p *Client) {
		p.client = client
	}
}

var pClientV *Client

func NewPClient(clientId, clientSecret string, isProduction bool, opts ...OptionFunc) (client *Client) {
	client = &Client{
		gateWayUrl:   SandboxGateWayUrl,
		clientId:     clientId,
		clientSecret: clientSecret,
		isProduction: isProduction,
		client:       http.DefaultClient,
	}
	if isProduction {
		client.gateWayUrl = ProductionGateWayUrl
	}
	for _, opt := range opts {
		opt(client)
	}
	// 访问令牌
	err := client.GetAccessToken()
	if err != nil {
		panic(err)
	}
	pClientV = client
	return
}

func GetPClient() *Client {
	return pClientV
}

// GenerateApiUrl 生成api url
func (p *Client) GenerateApiUrl(api string) string {
	return fmt.Sprintf("%s%s", p.gateWayUrl, api)
}

// GetAccessToken 获取访问令牌
func (p *Client) GetAccessToken() error {
	value := url.Values{}
	value.Add("grant_type", "client_credentials")
	req, err := http.NewRequest(HttpMethodPost, p.GenerateApiUrl(AccessTokenAPI), strings.NewReader(value.Encode()))
	if err != nil {
		return err
	}
	req.SetBasicAuth(p.clientId, p.clientSecret)
	req.Header.Set("Content-Type", ContentTypeFormUrlencoded)
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

// CallApiRequest 调用API请求
func (p *Client) CallApiRequest(httpMethod, api string, apiReq interface{}, apiResp interface{}) error {
	if p.accessToken == "" || p.accessTokenExpiresAt.Before(time.Now()) {
		// 过期重新获取访问令牌
		err := p.GetAccessToken()
		if err != nil {
			return err
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
	req.Header.Set("Content-Type", ContentTypeJson)

	res, err := p.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
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
