package accounts

import (
	"strconv"

	"github.com/Sylar007/form3-httpclient/customhttp"
	"github.com/google/uuid"
)

// delete account
func Delete(account_id uuid.UUID, version int) (*customhttp.Response, error) {

	var combineUrlString string = apiURL + "/" + account_id.String() + "?version=" + strconv.Itoa(version)
	response, err := httpClient.Delete(combineUrlString, nil)
	if err != nil {
		return response, err
	}
	return response, nil
}
