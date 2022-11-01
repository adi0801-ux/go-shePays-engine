package repositories

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"shepays/constants"
	"shepays/models"
	"shepays/utils"
	"time"
)

type HappyClient struct {
	Client   *http.Client
	BaseUrl  string
	AppId    string
	AppToken *string
	LogRep   *ApiLogsRepository
}

func CreateHttpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

func (h *HappyClient) SendPostRequest(endpoint string, body interface{}) (response *http.Response, errResp error) {
	method := http.MethodPost
	URL := h.BaseUrl + endpoint
	jsonBody, err := json.Marshal(body)
	if err != nil {
		utils.Log.Error(err)
		return nil, err
	}
	req, err := http.NewRequest(method, URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	//save to db
	apiLog := &models.APILog{
		RequestId: utils.GenerateID(),
		CreatedAt: time.Time{},
		Params:    req.URL.RawQuery,
		Payload:   string(jsonBody),
		Method:    method,
		Endpoint:  endpoint,
	}
	err = h.LogRep.CreateApiLog(apiLog)
	if err != nil {
		utils.Log.Error(err)
		return nil, err
	}

	response, errResp = h.callHappy(req, apiLog.RequestId)

	respBytes, _ := ioutil.ReadAll(response.Body)
	response.Body = ioutil.NopCloser(bytes.NewBuffer(respBytes))

	//	update in db
	apiLog.Response = string(respBytes)
	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		apiLog.Error = string(respBytes)
	}

	apiLog.ResponseCode = response.StatusCode
	apiLog.ResponseDate = time.Now().String()
	err = h.LogRep.UpdateApiLog(apiLog)
	if err != nil {
		utils.Log.Error(err)
		return response, err
	}
	return response, errResp

}
func (h *HappyClient) SendGetRequest(endpoint string, q url.Values) (*http.Response, error) {
	method := http.MethodGet
	URL := h.BaseUrl + endpoint

	req, err := http.NewRequest(method, URL, nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = q.Encode()

	//save to db
	apiLog := &models.APILog{
		RequestId: utils.GenerateID(),
		CreatedAt: time.Time{},
		Params:    req.URL.RawQuery,
		Method:    method,
		Endpoint:  endpoint,
	}
	err = h.LogRep.CreateApiLog(apiLog)
	if err != nil {
		utils.Log.Error(err)
		return nil, err
	}

	response, err := h.callHappy(req, apiLog.RequestId)

	respBytes, _ := ioutil.ReadAll(response.Body)
	response.Body = ioutil.NopCloser(bytes.NewBuffer(respBytes))

	//	update in db
	apiLog.Response = string(respBytes)
	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		apiLog.Error = string(respBytes)

	}
	apiLog.ResponseCode = response.StatusCode
	apiLog.ResponseDate = time.Now().String()
	err = h.LogRep.UpdateApiLog(apiLog)
	if err != nil {
		utils.Log.Error(err)
		return nil, err
	}

	return response, nil

}
func (h *HappyClient) SendPutRequest(endpoint string, body interface{}) (*http.Response, error) {
	method := http.MethodPut
	URL := h.BaseUrl + endpoint
	jsonBody, err := json.Marshal(body)
	if err != nil {
		utils.Log.Error(err)
		return nil, err
	}

	req, err := http.NewRequest(method, URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	//save to db
	apiLog := &models.APILog{
		RequestId: utils.GenerateID(),
		CreatedAt: time.Time{},
		Params:    req.URL.RawQuery,
		Payload:   string(jsonBody),
		Method:    method,
		Endpoint:  endpoint,
	}
	err = h.LogRep.CreateApiLog(apiLog)
	if err != nil {
		utils.Log.Error(err)
		return nil, err
	}

	response, err := h.callHappy(req, apiLog.RequestId)

	respBytes, _ := ioutil.ReadAll(response.Body)
	response.Body = ioutil.NopCloser(bytes.NewBuffer(respBytes))

	//	update in db
	apiLog.Response = string(respBytes)
	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		apiLog.Error = string(respBytes)

	}
	apiLog.ResponseCode = response.StatusCode
	apiLog.ResponseDate = time.Now().String()
	err = h.LogRep.UpdateApiLog(apiLog)
	if err != nil {
		utils.Log.Error(err)
		return nil, err
	}
	return response, nil
}

func (h *HappyClient) callHappy(req *http.Request, RequestId string) (*http.Response, error) {

	req.Header.Add("App-Id", h.AppId)
	req.Header.Add("App-Token", *h.AppToken)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Request-Id", RequestId)

	res, err := h.Client.Do(req)
	if err != nil {
		utils.Log.Error(err)
		return res, err
	}

	return res, nil

}

func (h *HappyClient) CallAuthHappy(key string, secret string) (*http.Response, error) {

	method := http.MethodPost
	URL := h.BaseUrl + constants.AuthHappayEndpoint

	req, err := http.NewRequest(method, URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("App-Id", h.AppId)
	req.Header.Add("key", key)
	req.Header.Add("secret", secret)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Request-Id", utils.GenerateID())

	res, err := h.Client.Do(req)
	if err != nil {
		utils.Log.Error(err)
		return res, err
	}

	return res, nil

}
