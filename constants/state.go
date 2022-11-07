package constants

const (
	SuccessResponseMessage = "SUCCESS"
	StartDateTime          = "0001-01-01 00:00:00 +0000 UTC"

	NoDisputesFound            = "no disputes found for the user"
	DefaultAccountType         = "SAVINGS"
	DeafultNumberOfCards       = 1
	DefaultTokenValue          = "NA"
	DefaultSuccessResponseCode = "00"
	DefaultErrorResponseCode   = "99"
	DefaultPincodeValidated    = "PIN Code City List Available"

	VersionCheckEndpoint              = "/PBcheckdeviceversionME"
	RegisterDeviceEndpoint            = "/PBregisterdeviceME"
	CustomerDeviceRegisterEndpoint    = "/PBcreCustdeviceprofME"
	SMSGenRequestEndpoint             = "/PBsmsactgenreqME"
	SMSStatusActionEndpoint           = "/PBgetsmsactstatusME"
	CheckCustomerMobileNumber         = "/PBgetcustdataME"
	ValidatePincodeEndpoint           = "/PBvalidatepincodeME"
	SetAdditionalInfomrmationEndpoint = "/PBcustmindtlsME"
	ConsentAskingEndpoint             = "/PBcustpepconsentME"
	CheckPanExistsEndpoint            = "/PBcheckPANexistME"
	PANVerifyEndpoint                 = "/PBverifyPANME"
	SetMPINEndpoint                   = "/PBcrecustprofPINME"
	VerifyAadharEndpoint              = "/PBverifyAadharXmlME"
	ValidateNameEndpoint              = "/PBvalidatenameME"
	VerifySelfieEndpoint              = "/PBfacematchME"
	ValidateDOBEndpoint               = "/PBverifyvalidDOBME"
	ValidatePANEndpoint               = "/PBvalidatePANME"
)
