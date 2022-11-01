package models

type CreateSavingsAccount struct {
	DocDetails     DocDetails               `json:"doc_details"`
	NomineeDetails []NomineesSavingsAccount `json:"nominee_details"`
	RelatedDetails RelatedDetails           `json:"related_details"`
	UserId         string                   `json:"user_id"`
	AccountType    string                   `json:"account_type" validate:"required,oneof='PREPAID' 'LIMIT' 'SAVINGS'"`
}

type RelatedDetails struct {
	EmploymentStatus string `json:"employment_status"` //what all validations -- ?
	MaritalStatus    string `json:"marital_status"`    //what all validations -- ?
	Nationality      string `json:"nationality"`
}

type DocDetails struct {
	DocNumber    string `json:"doc_number"`
	IssueDate    string `json:"issue_date" validate:"required"`
	PlaceOfIssue string `json:"place_of_issue"`
}

type NomineesSavingsAccount struct {
	Address      AddressSavingsAccount `json:"address"`
	NomineeName  string                `json:"nominee_name"`
	RelationType string                `json:"relation_type"`
	Value        int                   `json:"value"`
}
type AddressSavingsAccount struct {
	AddressLine1 string `json:"address_line_1"`
	City         string `json:"city"`
	Country      string `json:"country"`
	PostalCode   string `json:"postal_code"`
	State        string `json:"state"`
}

type InitiateNeftTransfer struct {
	Remarks         string `json:"remarks"`
	AccountId       string `json:"account_id"`
	Amount          int    `json:"amount"`
	CreditorDetails struct {
		AccountNumber string `json:"account_number"`
		Ifsc          string `json:"ifsc"`
		Name          string `json:"name"`
	} `json:"creditor_details"`
	PaymentType string `json:"payment_type"`
}

type ValidateNeftTransafer struct {
	Otp                     string `json:"otp" validate:"required"`
	TxnIdentificationNumber string `json:"txn_identification_number" validate:"required"`
}

type CheckCKyc struct {
	DocId   string `json:"doc_id" validate:"required"`
	DocType string `json:"doc_type" validate:"required"`
}

type SavingsAccountApiResponse struct {
	Ifsc              string `json:"ifsc"`
	UserId            string `json:"user_id"`
	HappayUserId      string `json:"happay_user_id"`
	AccountId         string `json:"account_id"`
	BankAccountNumber string `json:"bank_account_number"`
}
