package models

type VCifAPI struct {
	Channelid  string  `json:"channelid"`
	Appdtls    AppDtls `json:"appdtls"`
	Devicedtls struct {
		Deviceid string `json:"deviceid"`
	} `json:"devicedtls"`
	Deviceidentifier DeviceIdentifier `json:"deviceidentifier"`
	Custcreddtls     struct {
		Userid         string `json:"userid"`
		Usertype       string `json:"usertype"`
		Role           string `json:"role"`
		Cif            string `json:"cif"`
		FullName       string `json:"full_name"`
		FirstName      string `json:"first_name"`
		MiddleName     string `json:"middle_name"`
		LastName       string `json:"last_name"`
		Mobileno       string `json:"mobileno"`
		Mobcountrycode string `json:"mobcountrycode"`
		Emailid        string `json:"emailid"`
		Catcode        string `json:"catcode"`
	} `json:"custcreddtls"`
	Createindvcifdtl struct {
		ExtCustomerID         string `json:"extCustomerID"`
		IndividualCustomerDTO struct {
			Sex                        string `json:"sex"`
			MaritalStatus              string `json:"maritalStatus"`
			ProfessionCode             string `json:"professionCode"`
			IsStaff                    string `json:"isStaff"`
			CustomerId                 string `json:"customerId"`
			EmployeeId                 string `json:"employeeId"`
			CustomerEducation          string `json:"customerEducation"`
			MotherMaidenName           string `json:"motherMaidenName"`
			CountryOfResidence         string `json:"countryOfResidence"`
			CustomerMobilePhone        string `json:"customerMobilePhone"`
			EmailId                    string `json:"emailId"`
			Nationality                string `json:"nationality"`
			NationalIdentificationCode string `json:"nationalIdentificationCode"`
			IcType                     string `json:"icType"`
			Category                   string `json:"category"`
			Name                       struct {
				FirstName         string `json:"firstName"`
				FormattedFullName string `json:"formattedFullName"`
				FullName          string `json:"fullName"`
				LastName          string `json:"lastName"`
				MidName           string `json:"midName"`
				Prefix            string `json:"prefix"`
				ShortName         string `json:"shortName"`
				SingleFullName    string `json:"singleFullName"`
			} `json:"name"`
			DateOfBirthOrRegistration string `json:"dateOfBirthOrRegistration"`
			SignatureType             string `json:"signatureType"`
			Phone                     string `json:"phone"`
			Origin                    string `json:"origin"`
			Language                  string `json:"language"`
			HomeBranchCode            string `json:"homeBranchCode"`
			CifType                   string `json:"cifType"`
			Address                   struct {
				City    string `json:"city"`
				Country string `json:"country"`
				Line1   string `json:"line1"`
				Line2   string `json:"line2"`
				Line3   string `json:"line3"`
				State   string `json:"state"`
				Pincode string `json:"pincode"`
			} `json:"address"`
			PermanantAddress struct {
				City    string `json:"city"`
				Country string `json:"country"`
				Line1   string `json:"line1"`
				Line2   string `json:"line2"`
				Line3   string `json:"line3"`
				State   string `json:"state"`
				Pincode string `json:"pincode"`
			} `json:"permanantAddress"`
			Relation   string `json:"relation"`
			FatherName string `json:"fatherName"`
		} `json:"individualCustomerDTO"`
		MisClass       string `json:"misClass"`
		IsFCPB         string `json:"isFCPB"`
		MisCode        string `json:"misCode"`
		IncomeCategory string `json:"incomeCategory"`
	} `json:"createindvcifdtl"`
	Token  string `json:"token"`
	Signcs string `json:"signcs"`
}

type VCifAPIResponse struct {
	Response string `json:"response"`
	Respcode string `json:"respcode"`
}
