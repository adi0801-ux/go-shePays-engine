package models

type NameValidationAPIResponse struct {
	BadMatchFlag           interface{} `json:"bad match flag"`
	Response               string      `json:"response"`
	Respcode               string      `json:"respcode"`
	Panfullname            string      `json:"panfullname"`
	Namesimilarityresponse struct {
		Respmsg  string `json:"respmsg"`
		Name2    string `json:"name2"`
		Respcode string `json:"respcode"`
		Name1    string `json:"name1"`
	} `json:"namesimilarityresponse"`
	Adhaarfullname string `json:"adhaarfullname"`
}
