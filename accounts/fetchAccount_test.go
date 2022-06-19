package accounts

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestFetchAccount_WhenNonExistAccountId_ReturnMessageRecordNotExist(t *testing.T) {
	// Initialization
	mockId := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dd"
	message := "record " + mockId + " does not exist"
	mockMessage := &ResponseError{
		ErrorMessage: message,
	}
	// Execution
	account, response, err := Fetch(uuid.MustParse(mockId))
	if err != nil {
		t.Errorf("Returned err: got %v want %v",
			account, nil)
	}
	var result ResponseError
	if err := response.UnmarshalJson(&result); err != nil {
		t.Errorf("error message conversion error")
	}

	// Validation
	if response.StatusCode != http.StatusNotFound {
		t.Errorf("test is not successful because returning response code is not status not found(404)")
	}

	if !reflect.DeepEqual(mockMessage, &result) {
		t.Errorf("test is not successful because not returning the expected message")
	}
}

func TestFetchAccount_WhenValidAccountId_ReturnAccountData(t *testing.T) {
	// Initialization
	mockType := "accounts"
	mockId := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	mockOrganisationId := "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
	mockCountry := "GB"
	mockBaseCurrency := "GBP"
	mockBankId := "400300"
	mockBankIdCode := "GBDSC"
	mockBic := "NWBKGB22"
	mockName := []string{"Name1", "Name2", "Name3", "Name4"} //Name of the account holder, up to four lines possible.
	mockVersion := 0
	mockSelf := "/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"

	// Execution
	account, response, err := Fetch(uuid.MustParse(mockId))
	if err != nil {
		t.Errorf("Returned err: got %v want %v",
			account, nil)
	}

	// Validation
	if response.StatusCode != http.StatusOK {
		t.Errorf("test is not successful because returning response code not status ok(200)")
	}

	if _, err := json.Marshal(account); err != nil {
		t.Errorf("test is not successful because returning data cannot be map to json format")
	}

	if account.Data.Id != mockId {
		t.Errorf("test is not successful because unexpected Id: return %s mock %s",
			account.Data.Id, mockId)
	}

	//validation - on data returned
	if account.Data.Id != mockId {
		t.Errorf("test is not successful because unexpected Id: return %s mock %s",
			account.Data.Id, mockId)
	}

	if account.Data.OrganisationId != mockOrganisationId {
		t.Errorf("test is not successful because unexpected OrganisationId: return %s mock %s",
			account.Data.OrganisationId, mockOrganisationId)
	}

	if account.Data.Type != mockType {
		t.Errorf("test is not successful because unexpected Type: return %s mock %s",
			account.Data.Type, mockType)
	}

	if account.Data.Version != mockVersion {
		t.Errorf("test is not successful because unexpected Version: return %d mock %d",
			account.Data.Version, mockVersion)
	}

	// validation - on attributes returned

	if account.Data.Attributes.BankId != mockBankId {
		t.Errorf("test is not successful because unexpected BankId: return %s mock %s",
			account.Data.Attributes.AccountClassification, mockBankId)
	}

	if account.Data.Attributes.BankIdCode != mockBankIdCode {
		t.Errorf("test is not successful because unexpected BankIdCode: return %s mock %s",
			account.Data.Attributes.BankIdCode, mockBankIdCode)
	}

	if account.Data.Attributes.Country != mockCountry {
		t.Errorf("test is not successful because unexpected Country: return %s mock %s",
			account.Data.Attributes.Country, mockCountry)
	}

	if account.Data.Attributes.BaseCurrency != mockBaseCurrency {
		t.Errorf("test is not successful because unexpected Currency: return %s mock %s",
			account.Data.Attributes.BaseCurrency, mockBaseCurrency)
	}

	if account.Data.Attributes.Bic != mockBic {
		t.Errorf("test is not successful because unexpected Bic: return %s mock %s",
			account.Data.Attributes.Bic, mockBic)
	}

	if !compareArrayString(account.Data.Attributes.Name, mockName) {
		t.Errorf("test is not successful because unexpected Name: return %s mock %s",
			account.Data.Attributes.Name, mockName)
	}

	// validation - on links returned

	if account.Links.Self != mockSelf {
		t.Errorf("test is not successful because unexpected Links: return %s mock %s",
			account.Links.Self, mockSelf)
	}
}

// func TestFetchAccount_WhenInvalidFormatAccountId_ReturnMessageIdInvalidFormat(t *testing.T) {

// 	// Initialization
// 	mockId := "00000000-0000-0000-0000000000000000"
// 	message := "id is not a valid uuid"
// 	mockMessage := &ResponseError{
// 		ErrorMessage: message,
// 	}
// 	// Execution
// 	account, response, err := Fetch(uuid.MustParse(mockId))
// 	if err != nil {
// 		t.Errorf("Returned err: got %v want %v",
// 			account, nil)
// 	}
// 	var result ResponseError
// 	if err := response.UnmarshalJson(&result); err != nil {
// 		t.Errorf("error message conversion error")
// 	}

// 	// Validation
// 	if response.StatusCode != http.StatusBadRequest {
// 		t.Errorf("request is not successful because returning response code is not status bad request(400)")
// 	}

// 	if !reflect.DeepEqual(mockMessage, &result) {
// 		t.Errorf("request is not returning the expected message")
// 	}
// }
