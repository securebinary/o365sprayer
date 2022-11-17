package constants

// -----------------------------------------------------

// Validate User Using GetCredentialType
const GET_CREDENTIAL_TYPE = "https://login.microsoftonline.com/common/GetCredentialType"

// GetCredentialType JSON Request
type GetCredentialTypeRequestJSON struct {
	Username string `json:"Username"`
}

// GetCredentialType JSON Response
type GetCredentialTypeResponseJSON struct {
	Username       string `json:"Username"`
	IfExistsResult int    `json:"IfExistsResult"`
	// More to be added for additional info
}

// ----------------------------------------------------

// IfExistsResult
// # 1 - User Does Not Exist on Azure as Identity Provider
// # 0 - Account exists for domain using Azure as Identity Provider
// # 5 - Account exists but uses different IdP other than Microsoft
// # 6 - Account exists and is setup to use the domain and an IdP other than Microsoft

// ----------------------------------------------------
