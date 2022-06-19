package accounts

import (
	"errors"
	"net/http"

	"github.com/Sylar007/form3-httpclient/customhttp"
)

// Create new account
func Create(account *Account) (*Account, *customhttp.Response, error) {

	response, err := httpClient.Post(apiURL, account, nil)
	if err != nil {
		return nil, response, err
	}

	if response.StatusCode != http.StatusCreated {
		return nil, response, errors.New(string(response.Body))
	}

	var result Account
	if err := response.UnmarshalJson(&result); err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
