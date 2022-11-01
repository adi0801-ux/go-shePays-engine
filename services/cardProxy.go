package services

import (
	"encoding/json"
	"net/http"
	"shepays/constants"
	"shepays/models"
)

func (p *ServiceConfig) CreatePhysicalCard(createPhysicalCard *models.CreatePhysicalCard) (int, interface{}, error) {
	//call api --> repository for calling happay

	//get address Id from UserId
	address, err := p.UserAddressRep.ReadUserAddress(createPhysicalCard.UserId)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	//create models to call request
	createPhysicalCardApi := &models.CreatePhysicalCardApi{
		BinId:             createPhysicalCard.BinId,
		EmbossingNameList: []string{createPhysicalCard.EmbossingName},
		NumberOfCards:     constants.DeafultNumberOfCards,
		AddressId:         address.AddressID,
	}

	response, err := p.HappayClient.SendPostRequest(constants.CreatePhysicalCardEndpoint, *createPhysicalCardApi)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	if response.StatusCode != http.StatusCreated {
		return response.StatusCode, response.Body, err
	}

	var data []models.CreatePhysicalCardApiResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	data[0].UserID = createPhysicalCard.UserId

	//save card details to db
	err = p.CardRep.SavePhysicalCardResponse(&data[0])
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	//assign card

	assignCard := &models.AssignCard{UserId: createPhysicalCard.UserId}
	statusCode, resp, err := p.AssignCard(data[0].CardId, assignCard)
	if err != nil {
		return statusCode, resp, err
	}
	if statusCode != http.StatusOK {
		return statusCode, resp, err
	}

	//update the status in db
	data[0].Status = resp.(map[string]interface{})["status"].(string)
	err = p.CardRep.UpdatePhysicalCardResponse(&data[0])
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	//activate card
	statusCode, resp, err = p.ActivateCard(data[0].CardId)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	//update the card to activated

	data[0].Active = true
	err = p.CardRep.UpdatePhysicalCardResponse(&data[0])
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	data[0].Status = resp.(map[string]interface{})["status"].(string)

	return response.StatusCode, resp, nil

}

func (p *ServiceConfig) CreateVirtualCard(createVirtualCard *models.CreateVirtualCard) (int, interface{}, error) {
	//call api --> repository for calling happay

	//create models to call request
	createVirtualCardApi := &models.CreateVirtualCardApi{BinId: createVirtualCard.BinId}

	response, err := p.HappayClient.SendPostRequest(constants.CreateVirtualCardEndpoint, *createVirtualCardApi)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var data models.CreateVirtualCardApiResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	data.UserID = createVirtualCard.UserId

	err = p.CardRep.SaveVirtualCardResponse(&data)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	//assign card
	assignCard := &models.AssignCard{UserId: createVirtualCard.UserId}
	statusCode, resp, err := p.AssignCard(data.CardId, assignCard)
	if err != nil {
		return statusCode, resp, err
	}
	if statusCode != http.StatusOK {
		return statusCode, resp, err
	}

	//update the status in db
	data.Status = resp.(map[string]interface{})["status"].(string)
	err = p.CardRep.UpdateVirtualCardResponse(&data)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	//activate card
	statusCode, resp, err = p.ActivateCard(data.CardId)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	//update the card to activated

	data.Active = true
	err = p.CardRep.UpdateVirtualCardResponse(&data)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	data.Status = resp.(map[string]interface{})["status"].(string)

	return response.StatusCode, data, nil

}

func (p *ServiceConfig) AssignCard(cardId string, assignCard *models.AssignCard) (int, interface{}, error) {
	//call api --> repository for calling happay
	users, err := p.UserRep.ReadUserDetails(assignCard.UserId)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	assignCard.UserId = users.HappayUserId

	response, err := p.HappayClient.SendPutRequest(constants.GenerateAssignCardEndpoint(cardId), *assignCard)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, data, nil

}

func (p *ServiceConfig) ActivateCard(cardId string) (int, interface{}, error) {
	//call api --> repository for calling happay

	response, err := p.HappayClient.SendPutRequest(constants.GenerateActivateCardEndpoint(cardId), nil)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, data, nil

}

func (p *ServiceConfig) GetCardDetailsCard(cardId string) (int, interface{}, error) {
	//call api --> repository for calling happay

	response, err := p.HappayClient.SendGetRequest(constants.GenerateCardEndpoint(cardId), nil)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, data, nil

}

func (p *ServiceConfig) BlockCard(cardId string) (int, interface{}, error) {
	//call api --> repository for calling happay

	response, err := p.HappayClient.SendPutRequest(constants.GenerateBlockCardEndpoint(cardId), nil)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, data, nil

}

func (p *ServiceConfig) UnBlockCard(cardId string) (int, interface{}, error) {
	//call api --> repository for calling happay

	response, err := p.HappayClient.SendPutRequest(constants.GenerateUnblockCardEndpoint(cardId), nil)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, data, nil

}

func (p *ServiceConfig) InitiateCardPinSet(UserId string) (int, interface{}, error) {
	//call api --> repository for calling happay
	virtualCard, err := p.CardRep.ReadVirtualCardDetails(UserId)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	initiateCardPinSet := &models.InitiateCardPinSet{CardToken: virtualCard.CardId}

	response, err := p.HappayClient.SendPostRequest(constants.InitiateCardPinSetEndpoint, *initiateCardPinSet)
	if err != nil {
		return http.StatusBadRequest, models.InitiateCardPinSetApiResponse{}, err
	}

	var data models.InitiateCardPinSetApiResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return http.StatusBadRequest, models.InitiateCardPinSetApiResponse{}, err
	}

	return response.StatusCode, data, nil

}

//func (p *ServiceConfig) EncryptPin(data models.InitiateCardPinSetApiResponse, pin string) string {
//	//pinBytes := []byte(pin)
//	bKey := utils.Base64Decode(data.PublicEncryptionKey)
//	key := utils.ExportPEMStrToPubKey(bKey)
//	rng := rand.Reader
//	oaep, _ := rsa.EncryptPKCS1v15(rng, key, []byte(pin))
//	return utils.Base64Encode(oaep)
//}

//func (p *ServiceConfig) SetCardPin(setCardPin *models.SetCardPin) (int, interface{}, error) {
//	//call api --> repository for calling happay
//
//	//get virtual card details
//	virtualCard, err := p.CardRep.ReadVirtualCardDetails(setCardPin.UserId)
//	if err != nil {
//		return http.StatusBadRequest, nil, err
//	}
//
//	statusCode, pinSetResponse, err := p.InitiateCardPinSet(virtualCard.CardId)
//	if err != nil || statusCode != http.StatusOK {
//		return statusCode, pinSetResponse, err
//	}
//
//	pin := p.EncryptPin(pinSetResponse, setCardPin.Pin)
//
//	setCardPinApi := &models.SetCardPinApi{
//		Pin:         pin,
//		PinSetToken: pinSetResponse.PinSetToken,
//	}
//
//	response, err := p.HappayClient.SendPostRequest(constants.SetCardPinEndpoint, *setCardPinApi)
//	if err != nil {
//		return http.StatusBadRequest, nil, err
//	}
//
//	var data map[string]interface{}
//	err = json.NewDecoder(response.Body).Decode(&data)
//	if err != nil {
//		return http.StatusBadRequest, nil, err
//	}
//
//	return response.StatusCode, data, nil
//
//}

func (p *ServiceConfig) ReplaceCard(replaceCard *models.ReplaceCard) (int, interface{}, error) {
	//call api --> repository for calling happay

	response, err := p.HappayClient.SendPostRequest(constants.ReplaceCardEndpoint, *replaceCard)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, data, nil

}
