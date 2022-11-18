package models

type CustomerDocsAPI struct {
	Channelid  string  `json:"channelid"`
	Appdtls    AppDtls `json:"appdtls"`
	Devicedtls struct {
		Deviceid string `json:"deviceid"`
	} `json:"devicedtls"`
	Deviceidentifier DeviceIdentifier `json:"deviceidentifier"`
	Docdtls          struct {
		DocType      string `json:"doc_type"`
		DocProofType string `json:"doc_proof_type"`
		DocId        string `json:"doc_id"`
		DocVerDt     string `json:"doc_ver_dt"`
		DocVerChanid string `json:"doc_ver_chanid"`
		DocStatus    string `json:"doc_status"`
		DocIssDt     string `json:"doc_iss_dt"`
		DocXml       string `json:"doc_xml"`
		Name         string `json:"name"`
		FName        string `json:"f_name"`
		MName        string `json:"m_name"`
		LName        string `json:"l_name"`
	} `json:"docdtls"`
	Token  string `json:"token"`
	Signcs string `json:"signcs"`
}
