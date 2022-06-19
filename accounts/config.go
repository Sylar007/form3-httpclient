package accounts

import (
	"net/http"
	"time"

	"github.com/Sylar007/form3-httpclient/customhttp"
)

var apiURL = "http://accountapi:8080/v1/organisation/accounts"
var apiHealthURL = "http://accountapi:8080/v1/health"

var (
	httpClient = getHttpClient()
)

func getHttpClient() customhttp.CustomClient {
	headers := make(http.Header)
	headers.Set("Content-Type", "application/json")

	client := customhttp.NewBuilder().
		SetHeaders(headers).
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		Build()
	return client
}

func compareArrayString(firstArray, secondArray []string) bool {
	if len(firstArray) != len(secondArray) {
		return false
	}
	for i, v := range firstArray {
		if v != secondArray[i] {
			return false
		}
	}
	return true
}
