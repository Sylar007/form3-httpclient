package accounts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCreateAccount_WhenArrayOfNamesMoreThan4_ReturnValidationErrorMessage(t *testing.T) {
	// Initialization
	//remove mockCountry which is the required data
	mockType := "accounts"
	mockId := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	mockVersion := 0
	mockOrganisationId := "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
	mockCountry := "GB"
	mockBaseCurrency := "GBP"
	mockBankId := "400300"
	mockAccountNumber := "41426819"
	mockBankIdCode := "GBDSC"
	mockBic := "NWBKGB22"
	mockIban := "GB11NWBK40030041426819"
	mockStatus := "confirmed"
	mockName := []string{"Name1", "Name2", "Name3", "Name4", "Name5"} //Name of the account holder are more than 4
	mockValidationType := "card"
	mockReferenceMask := "############"
	mockAcceptanceQualifier := "same_day"
	mockSelf := "/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"

	mockAccount := &Account{
		Data: &Data{
			Type:           mockType,
			Id:             mockId,
			Version:        mockVersion,
			OrganisationId: mockOrganisationId,
			Attributes: &AccountAttributes{
				Country:             mockCountry,
				BaseCurrency:        mockBaseCurrency,
				AccountNumber:       mockAccountNumber,
				BankId:              mockBankId,
				BankIdCode:          mockBankIdCode,
				Bic:                 mockBic,
				Iban:                mockIban,
				Name:                mockName,
				Status:              mockStatus,
				ValidationType:      mockValidationType,
				ReferenceMask:       mockReferenceMask,
				AcceptanceQualifier: mockAcceptanceQualifier,
			},
		},
		Links: &Links{
			Self: mockSelf,
		},
	}

	message := "validation failure list:\nvalidation failure list:\nvalidation failure list:\nname in body should have at most 4 items"
	mockMessage := &ResponseError{
		ErrorMessage: message,
	}

	// Execution
	_, response, _ := Create(mockAccount)

	var result ResponseError
	if err := response.UnmarshalJson(&result); err != nil {
		t.Errorf("error message conversion error")
	}

	// Validation
	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("request is not successful because returning response code is not status bad request(400)")
	}

	if !reflect.DeepEqual(mockMessage, &result) {
		fmt.Println(&result)
		t.Errorf("request is not returning the expected message")
	}
}

func TestCreateAccount_WhenMissingRequiredData_ReturnValidationErrorMessage(t *testing.T) {
	// Initialization
	//remove mockCountry which is the required data
	mockType := "accounts"
	mockId := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	mockVersion := 0
	mockOrganisationId := "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
	mockBaseCurrency := "GBP"
	mockBankId := "400300"
	mockAccountNumber := "41426819"
	mockBankIdCode := "GBDSC"
	mockBic := "NWBKGB22"
	mockIban := "GB11NWBK40030041426819"
	mockStatus := "confirmed"
	mockName := []string{"Name1", "Name2", "Name3", "Name4"} //Name of the account holder, up to four lines possible.
	mockValidationType := "card"
	mockReferenceMask := "############"
	mockAcceptanceQualifier := "same_day"
	mockSelf := "/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"

	mockAccount := &Account{
		Data: &Data{
			Type:           mockType,
			Id:             mockId,
			Version:        mockVersion,
			OrganisationId: mockOrganisationId,
			Attributes: &AccountAttributes{
				BaseCurrency:        mockBaseCurrency,
				AccountNumber:       mockAccountNumber,
				BankId:              mockBankId,
				BankIdCode:          mockBankIdCode,
				Bic:                 mockBic,
				Iban:                mockIban,
				Name:                mockName,
				Status:              mockStatus,
				ValidationType:      mockValidationType,
				ReferenceMask:       mockReferenceMask,
				AcceptanceQualifier: mockAcceptanceQualifier,
			},
		},
		Links: &Links{
			Self: mockSelf,
		},
	}

	message := "validation failure list:\nvalidation failure list:\nvalidation failure list:\ncountry in body should match '^[A-Z]{2}$'"
	mockMessage := &ResponseError{
		ErrorMessage: message,
	}

	// Execution
	_, response, _ := Create(mockAccount)

	var result ResponseError
	if err := response.UnmarshalJson(&result); err != nil {
		t.Errorf("error message conversion error")
	}

	// Validation
	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("request is not successful because returning response code is not status bad request(400)")
	}

	if !reflect.DeepEqual(mockMessage, &result) {
		fmt.Println(&result)
		t.Errorf("request is not returning the expected message")
	}
}

func TestCreateAccount_WhenValidData_ReturnAccountCreate(t *testing.T) {
	// Initialization
	mockType := "accounts"
	mockId := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	mockVersion := 0
	mockOrganisationId := "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
	mockCountry := "GB"
	mockBaseCurrency := "GBP"
	mockBankId := "400300"
	mockAccountNumber := "41426819"
	mockBankIdCode := "GBDSC"
	mockBic := "NWBKGB22"
	mockIban := "GB11NWBK40030041426819"
	mockStatus := "confirmed"
	mockName := []string{"Name1", "Name2", "Name3", "Name4"} //Name of the account holder, up to four lines possible.
	mockValidationType := "card"
	mockReferenceMask := "############"
	mockAcceptanceQualifier := "same_day"
	mockUserDefined := []UserDefinedData{
		{
			Key:   "Some account related key",
			Value: "Some account related value",
		},
	}
	mockSelf := "/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"

	mockAccount := &Account{
		Data: &Data{
			Type:           mockType,
			Id:             mockId,
			Version:        mockVersion,
			OrganisationId: mockOrganisationId,
			Attributes: &AccountAttributes{
				Country:             mockCountry,
				BaseCurrency:        mockBaseCurrency,
				AccountNumber:       mockAccountNumber,
				BankId:              mockBankId,
				BankIdCode:          mockBankIdCode,
				Bic:                 mockBic,
				Iban:                mockIban,
				Name:                mockName,
				Status:              mockStatus,
				UserDefinedData:     mockUserDefined,
				ValidationType:      mockValidationType,
				ReferenceMask:       mockReferenceMask,
				AcceptanceQualifier: mockAcceptanceQualifier,
			},
		},
		Links: &Links{
			Self: mockSelf,
		},
	}

	// Execution
	account, response, err := Create(mockAccount)
	if err != nil {
		t.Errorf("test is not successful with this error: " + err.Error())
	}

	// Validation
	if response.StatusCode != http.StatusCreated {
		t.Errorf("test is not successful because returning response code not status created(201)")
	}

	if _, err := json.Marshal(account); err != nil {
		t.Errorf("test is not successful because returning data cannot be map to json format")
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

	if account.Data.Version != mockVersion {
		t.Errorf("test is not successful because unexpected Version: return %d mock %d",
			account.Data.Version, mockVersion)
	}

	if account.Data.Type != mockType {
		t.Errorf("test is not successful because unexpected Type: return %s mock %s",
			account.Data.Type, mockType)
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

	if account.Data.Attributes.AccountNumber != mockAccountNumber {
		t.Errorf("test is not successful because unexpected AccountNumber: return %s mock %s",
			account.Data.Attributes.AccountNumber, mockAccountNumber)
	}

	if account.Data.Attributes.BaseCurrency != mockBaseCurrency {
		t.Errorf("test is not successful because unexpected Currency: return %s mock %s",
			account.Data.Attributes.BaseCurrency, mockBaseCurrency)
	}

	if account.Data.Attributes.Bic != mockBic {
		t.Errorf("test is not successful because unexpected Bic: return %s mock %s",
			account.Data.Attributes.Bic, mockBic)
	}

	// validation - on links returned

	if account.Links.Self != mockSelf {
		t.Errorf("test is not successful because unexpected Links: return %s mock %s",
			account.Links.Self, mockSelf)
	}
}
