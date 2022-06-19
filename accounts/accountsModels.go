package accounts

type Account struct {
	Data  *Data  `json:"data"`
	Links *Links `json:"links"`
}

type Data struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	Id             string             `json:"id,omitempty"`
	OrganisationId string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        int                `json:"version,omitempty"`
}

type AccountAttributes struct {
	AccountClassification   string            `json:"account_classification,omitempty"`
	AccountNumber           string            `json:"account_number,omitempty"`
	AcceptanceQualifier     string            `json:"acceptance_qualifier,omitempty"`
	AlternativeNames        []string          `json:"alternative_names,omitempty"`
	BankId                  string            `json:"bank_id,omitempty"`
	BankIdCode              string            `json:"bank_id_code,omitempty"`
	BaseCurrency            string            `json:"base_currency,omitempty"`
	Bic                     string            `json:"bic,omitempty"`
	Country                 string            `json:"country" binding:"required"`
	Iban                    string            `json:"iban,omitempty"`
	JointAccount            bool              `json:"joint_account,omitempty"`
	Name                    []string          `json:"name" binding:"required"`
	NameMatchingStatus      string            `json:"name_matching_status,omitempty"`
	ReferenceMask           string            `json:"reference_mask,omitempty"`
	SecondaryIdentification string            `json:"secondary_identification,omitempty"`
	Status                  string            `json:"status,omitempty"`
	StatusReason            string            `json:"status_reason,omitempty"`
	UserDefinedData         []UserDefinedData `json:"user_defined_data"`
	ValidationType          string            `json:"validation_type,omitempty"`
}

type UserDefinedData struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type Links struct {
	Self string `json:"self,omitempty"`
}

type ResponseError struct {
	ErrorMessage string `json:"error_message,omitempty"`
}
