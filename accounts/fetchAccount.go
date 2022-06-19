package accounts

import (
	"github.com/Sylar007/form3-httpclient/customhttp"
	"github.com/google/uuid"
)

// fetch account
func Fetch(account_id uuid.UUID) (*Account, *customhttp.Response, error) {

	var combineUrlString string = apiURL + "/" + account_id.String()
	response, err := httpClient.Get(combineUrlString, nil)
	if err != nil {
		return nil, response, err
	}
	var result Account
	if err := response.UnmarshalJson(&result); err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
