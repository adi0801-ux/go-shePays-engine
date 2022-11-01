package models

type AuthResponse struct {
	ResStr  string `json:"res_str"`
	ResData struct {
		AppToken string `json:"app_token"`
	} `json:"res_data"`
}
