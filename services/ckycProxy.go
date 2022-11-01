package services

import (
	"encoding/json"
	"net/http"
	"shepays/constants"
	"shepays/models"
)

func (p *ServiceConfig) CheckCkyc(checkCkyc *models.UserCkyc) (int, interface{}, error) {
	//call api --> repository for calling happay

	response, err := p.HappayClient.SendPostRequest(constants.CheckCkycEndpoint, *checkCkyc)
	if err != nil || response.StatusCode != http.StatusOK {
		return http.StatusBadRequest, nil, err
	}

	var data map[string]interface{}

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	if response.StatusCode != http.StatusOK {
		return response.StatusCode, data, err
	}

	//save to db

	checkCkyc.CkycName = data["ckyc_name"].(string)
	checkCkyc.CkycDate = data["ckyc_date"].(string)
	checkCkyc.CkycNumber = data["ckyc_no"].(string)

	err = p.CkycRep.CreateUserCkyc(checkCkyc)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, data, nil
}
