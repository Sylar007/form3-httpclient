package accounts

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestDeleteAccount_WhenNonExistAccountId_ReturnsNotFoundStatus(t *testing.T) {

	// Initialization
	mockInvalidAccountId := "ad27e265-9605-4b4b-a0e5-3003ea9cc4d1"
	mockValidVersion := 0

	// Execution
	response, err := Delete(uuid.MustParse(mockInvalidAccountId), mockValidVersion)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// Validation
	if response.StatusCode != http.StatusNotFound {
		t.Errorf("test is not successful because returning response code is not status not found(404)")
	}
}

func TestDeleteAccount_WhenInvalidVersion_ReturnMessageInvalidVersion(t *testing.T) {

	// Initialization
	mockValidAccountId := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	mockInvalidVersion := 1
	message := "invalid version"
	mockMessage := &ResponseError{
		ErrorMessage: message,
	}

	// Execution
	response, err := Delete(uuid.MustParse(mockValidAccountId), mockInvalidVersion)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	var result ResponseError
	if err := response.UnmarshalJson(&result); err != nil {
		t.Errorf("error message conversion error")
	}

	// Validation
	if response.StatusCode != http.StatusConflict {
		t.Errorf("test is not successful because returning response code is not status conflict(409)")
	}

	if !reflect.DeepEqual(mockMessage, &result) {
		t.Errorf("test is not successful because not returning the expected message")
	}
}

func TestDeleteAccount_WhenAccountIdExist_ReturnsNoContentStatus(t *testing.T) {

	// Initialization
	mockValidAccountId := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	mockValidVersion := 0

	// Execution
	response, err := Delete(uuid.MustParse(mockValidAccountId), mockValidVersion)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// Validation
	if response.StatusCode != http.StatusNoContent {
		t.Errorf("test is not successful because returning response code is not status no content(204)")
	}
}
