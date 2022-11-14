package models

type CardCreationAPI struct {
	Channelid  string  `json:"channelid"`
	Appdtls    AppDtls `json:"appdtls"`
	Devicedtls struct {
		Deviceid string `json:"deviceid"`
	} `json:"devicedtls"`
	Deviceidentifier DeviceIdentifier `json:"deviceidentifier"`
	Msg              struct {
		AccountNo   string `json:"accountNo"`
		NetworkType string `json:"network_type"`
		BinPrefix   string `json:"binPrefix"`
		AddressDtls struct {
			Address1 string `json:"address1"`
			Address2 string `json:"address2"`
			Address3 string `json:"address3"`
			City     string `json:"city"`
			State    string `json:"state"`
			Pincode  string `json:"pincode"`
			Country  string `json:"country"`
		} `json:"address_dtls"`
		SessionId string `json:"sessionId"`
	} `json:"msg"`
	Token  string `json:"token"`
	Signcs string `json:"signcs"`
}

type GetCardDetailsAPI struct {
	Channelid  string  `json:"channelid"`
	Appdtls    AppDtls `json:"appdtls"`
	Devicedtls struct {
		Deviceid string `json:"deviceid"`
	} `json:"devicedtls"`
	Deviceidentifier DeviceIdentifier `json:"deviceidentifier"`
	Token            string           `json:"token"`
	Signcs           string           `json:"signcs"`
}

type GetCardDetailsAPIResponse struct {
	AccntProdlist []struct {
		AccntProdCode       string `json:"accntProdCode"`
		AccntProdName       string `json:"accntProdName"`
		SweepMappedProdCode string `json:"sweepMappedProdCode"`
		SweepMappedProdName string `json:"sweepMappedProdName"`
		AccntWelcomeHeader  string `json:"accntWelcomeHeader"`
		AccntDetails        []struct {
			Details     string `json:"details"`
			Description string `json:"description"`
		} `json:"accntDetails"`
		MoreDetails []struct {
			Details     string `json:"details"`
			Description string `json:"description"`
		} `json:"moreDetails"`
		EligibleDebitCards []string `json:"eligibleDebitCards"`
		CardFlag           []string `json:"CardFlag"`
		TxnFlag            []string `json:"TxnFlag"`
	} `json:"accntProdlist"`
	Response     string `json:"response"`
	CardProdlist []struct {
		CardProdCode       string   `json:"cardProdCode"`
		CardProdName       string   `json:"cardProdName"`
		CardProdOffers     []string `json:"cardProdOffers"`
		CardProdOffersLink string   `json:"cardProdOffersLink"`
		Network            string   `json:"network"`
		BinPrifix          string   `json:"binPrifix"`
		NetworkiCon        string   `json:"networkiCon"`
	} `json:"cardProdlist"`
	Respcode            string `json:"respcode"`
	KycAccntprodmaplist []struct {
		KYCtypes      string   `json:"KYCtypes"`
		AccntProdCode []string `json:"accntProdCode"`
	} `json:"kycAccntprodmaplist"`
}
