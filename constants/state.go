package constants

const (
	AuthHappayEndpoint           = "/app/v1/authorize/"
	CreateSavingsAccountEndpoint = "/api/v2/account"
	CheckTransferStatusEndpoint  = "/neobank/v1/fund/transfer/status/"
	InitiateNEFTTransferEndpoint = "/neobank/v1/fund/initiate"
	ValidateNEFTTransferEndpoint = "/neobank/v1/fund/transfer"

	CreatePhysicalCardEndpoint = "/api/v1/card/physical"

	CreateVirtualCardEndpoint = "/api/v1/card/virtual"
	CardEndpointFirst         = "/api/v1/card/"
	AssignCardEndpoint        = "/assign"
	ActivateCardEndpoint      = "/activate"

	BlockCardEndpoint          = "/block"
	UnblockCardEndpoint        = "/unblock"
	InitiateCardPinSetEndpoint = "/api/v1/card/pin-set/"
	SetCardPinEndpoint         = "/api/v1/card/pin/set"
	ReplaceCardEndpoint        = "/api/v1/card/replace̵"
	CheckCkycEndpoint          = "/ckyc/v1/search"
	CreateUserEndpoint         = "/api/v1/user"
	CreateUserAddressEndpoint  = "/api/v1/address"

	StartDateTime = "0001-01-01 00:00:00 +0000 UTC"

	NoDisputesFound      = "no disputes found for the user"
	DefaultAccountType   = "SAVINGS"
	DeafultNumberOfCards = 1
)

func GenerateAssignCardEndpoint(CardId string) string {
	return CardEndpointFirst + CardId + AssignCardEndpoint
}

func GenerateActivateCardEndpoint(CardId string) string {
	return CardEndpointFirst + CardId + ActivateCardEndpoint
}

func GenerateCardEndpoint(CardId string) string {
	return CardEndpointFirst + CardId
}

func GenerateBlockCardEndpoint(CardId string) string {
	return CardEndpointFirst + CardId + BlockCardEndpoint
}

func GenerateUnblockCardEndpoint(CardId string) string {
	return CardEndpointFirst + CardId + UnblockCardEndpoint
}

func GenerateUpdateUserEndpoint(UserId string) string {
	return CreateUserEndpoint + "/" + UserId
}
