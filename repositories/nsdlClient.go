package repositories

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"shepays/models"
	"shepays/utils"
	"time"
)

type NSDLClient struct {
	Client           *http.Client
	BaseUrl          string
	LogRep           *ApiLogsRepository
	ChannelId        string
	AppDtls          *models.AppDtls
	DeviceIdentifier *models.DeviceIdentifier
}

func CreateHttpClient() *http.Client {
	client := &http.Client{Timeout: 50 * time.Second}
	return client
}

func (h *NSDLClient) SendPostRequest(endpoint string, body interface{}) (response *http.Response, errResp error) {

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

	response, errResp = h.callNSDL(req, apiLog.RequestId)

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

func (h *NSDLClient) SendPostRequestWithoutResponseLog(endpoint string, body interface{}) (response *http.Response, errResp error) {

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

	response, errResp = h.callNSDL(req, apiLog.RequestId)

	apiLog.ResponseCode = response.StatusCode
	apiLog.ResponseDate = time.Now().String()
	err = h.LogRep.UpdateApiLog(apiLog)
	if err != nil {
		utils.Log.Error(err)
		return response, err
	}
	return response, errResp

}

func (h *NSDLClient) callNSDL(req *http.Request, RequestId string) (*http.Response, error) {

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Request-Id", RequestId)

	res, err := h.Client.Do(req)
	if err != nil {
		utils.Log.Error(err)
		return res, err
	}

	return res, nil

}

func (h *NSDLClient) callPostFormNSDL(req *http.Request, RequestId string, header string) (*http.Response, error) {

	req.Header.Add("Content-Type", header)
	req.Header.Add("Request-Id", RequestId)

	res, err := h.Client.Do(req)
	if err != nil {
		utils.Log.Error(err)
		return res, err
	}

	return res, nil

}

func (h *NSDLClient) SendPostFormRequest(endpoint string, body *bytes.Buffer, header string) (response *http.Response, errResp error) {

	URL := h.BaseUrl + endpoint
	method := http.MethodPost
	//payload, _ := body.ReadString(0)

	//save to db
	apiLog := &models.APILog{
		RequestId: utils.GenerateID(),
		CreatedAt: time.Time{},
		Payload:   "not storing aadhar api payload",
		Method:    method,
		Endpoint:  endpoint,
	}
	err := h.LogRep.CreateApiLog(apiLog)
	if err != nil {
		utils.Log.Error(err)
		return nil, err
	}

	req, err := http.NewRequest(method, URL, body)
	if err != nil {
		return nil, err
	}

	response, errResp = h.callPostFormNSDL(req, apiLog.RequestId, header)

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
