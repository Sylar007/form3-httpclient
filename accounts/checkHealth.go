package accounts

import (
	"github.com/Sylar007/form3-httpclient/customhttp"
)

// get account service Health
func Check() (*customhttp.Response, error) {
	response, err := httpClient.Get(apiHealthURL, nil)
	if err != nil {
		return response, err
	}
	return response, nil
}
