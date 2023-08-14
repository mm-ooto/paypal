package paypal

import (
	"errors"
	"fmt"
	"paypal/model"
)

var (
	signatureVerifyFailureErr = errors.New("signature verify failure")
)

// ResponseError ResponseError
// https://developer.paypal.com/api/rest/responses/#link-failedrequests
type ResponseError struct {
	RequestUrl string
	Name       string                  `json:"name"`
	Message    string                  `json:"message"`
	DebugId    string                  `json:"debug_id"`
	Details    []model.Details         `json:"details"`
	Links      []model.LinkDescription `json:"links"`
}

func (r *ResponseError) Error() string {
	return fmt.Sprintf("RequestUrl:%s Error:%s Message:%s", r.RequestUrl, r.Name, r.Message)
}

// ValidationError ValidationError
// https://developer.paypal.com/api/rest/responses/#link-authorizationerrors
type ValidationError struct {
	RequestUrl       string
	Err              string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("RequestUrl:%s Error:%s Description:%s", v.RequestUrl, v.Err, v.ErrorDescription)
}

// ForbiddenErr ForbiddenErr
type ForbiddenErr struct {
	RequestUrl       string
	LocalizedMessage string          `json:"localizedMessage"`
	Name             string          `json:"name"`
	Message          string          `json:"message"`
	Details          []model.Details `json:"details"`
	InformationLink  string          `json:"information_link"`
	DebugId          string          `json:"debug_id"`
}

func (f *ForbiddenErr) Error() string {
	return fmt.Sprintf("RequestUrl:%s Error:%s Description:%s", f.RequestUrl, f.Name, f.Message)
}
