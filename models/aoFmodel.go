package models

type AoFModel struct {
	ProductCodeString string `json:"product_code_string"`
	CardProdCode      string `json:"card_prod_code"`
	CardProdName      string `json:"card_prod_name"`
	Network           string `json:"network"`
	UserId            string `json:"user_id"`
}

type AoFModelAPI struct {
	Channelid  string  `json:"channelid"`
	Appdtls    AppDtls `json:"appdtls"`
	Devicedtls struct {
		Deviceid string `json:"deviceid"`
	} `json:"devicedtls"`
	Deviceidentifier DeviceIdentifier `json:"deviceidentifier"`
	Custcreddtls     struct {
		Userid   string `json:"userid"`
		Usertype string `json:"usertype"`
		Role     string `json:"role"`
	} `json:"custcreddtls"`
	Aofdtls struct {
		Prefix                     string `json:"prefix"`
		FirstName                  string `json:"firstName"`
		MidName                    string `json:"midName"`
		LastName                   string `json:"lastName"`
		MotherMaidenName           string `json:"motherMaidenName"`
		FatherName                 string `json:"fatherName"`
		MaritalStatus              string `json:"maritalStatus"`
		CustomerMobilePhone        string `json:"customerMobilePhone"`
		DateOfBirth                string `json:"dateOfBirth"`
		IncomeCategory             string `json:"incomeCategory"`
		EmailId                    string `json:"emailId"`
		NationalIdentificationCode string `json:"nationalIdentificationCode"`
		ProfessionCode             string `json:"professionCode"`
		Relation                   string `json:"relation"`
		Sex                        string `json:"sex"`
		ProductCodeString          string `json:"productCodeString"`
		Address                    struct {
			City    string `json:"city"`
			Country string `json:"country"`
			Line1   string `json:"line1"`
			Line2   string `json:"line2"`
			Line3   string `json:"line3"`
			State   string `json:"state"`
			Pincode string `json:"pincode"`
		} `json:"address"`
		AccountTitle          string `json:"accountTitle"`
		AcctCurrencyString    string `json:"acctCurrencyString"`
		BranchCode            string `json:"branchCode"`
		CustomerIDString      string `json:"customerIDString"`
		FlgJointHolderString  string `json:"flgJointHolderString"`
		FlgRestrictAcctString string `json:"flgRestrictAcctString"`
		FlgSCWaiveString      string `json:"flgSCWaiveString"`
		FlgTransactionType    string `json:"flgTransactionType"`
		MinorAcctStatusString string `json:"minorAcctStatusString"`
		Nationality           string `json:"nationality"`
		CustomerEducation     string `json:"customerEducation"`
		IsStaff               string `json:"isStaff"`
		IcType                string `json:"icType"`
		CifType               string `json:"cifType"`
		Origin                string `json:"origin"`
		Name                  struct {
			FirstName         string `json:"firstName"`
			FormattedFullName string `json:"formattedFullName"`
			FullName          string `json:"fullName"`
			LastName          string `json:"lastName"`
			MidName           string `json:"midName"`
			Prefix            string `json:"prefix"`
			ShortName         string `json:"shortName"`
			SingleFullName    string `json:"singleFullName"`
		} `json:"name"`
	} `json:"aofdtls"`
	Carddtls struct {
		CardProdCode string `json:"cardProdCode"`
		CardProdName string `json:"cardProdName"`
		Network      string `json:"network"`
	} `json:"carddtls"`
	Token  string `json:"token"`
	Signcs string `json:"signcs"`
}

type AoFModelAPIResponse struct {
	Refno    string `json:"refno"`
	Response string `json:"response"`
	Aofpdf   string `json:"aofpdf"`
	Respcode string `json:"respcode"`
}
