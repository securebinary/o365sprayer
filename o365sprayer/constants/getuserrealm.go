package constants

// ------------------------------------------

// GetUserRealm.srf

// XML - https://login.microsoftonline.com/getuserrealm.srf?login=username@domain.com&xml=1

// JSON
const GET_USER_REALM = "https://login.microsoftonline.com/getuserrealm.srf?json=1&login="

type GetUserRealmJSON struct {
	State                   int
	UserState               int
	Login                   string
	NameSpaceType           string
	DomainName              string
	FederationGlobalVersion int
	AuthURL                 string
	FederationBrandName     string
	CloudInstanceName       string
	CloudInstanceIssuerUri  string
}

// ----------------------------------------------------------

// Get Tenant ID

const GET_TENANT_ID = "https://login.microsoftonline.com/$DOMAIN/v2.0/.well-known/openid-configuration"

type GetTenantIdJSON struct {
	TokenEndpoint         string `json:"token_endpoint"`
	JwksUri               string `json:"jwks_uri"`
	KerberosEndpoint      string `json:"kerberos_endpoint"`
	AuthorizationEndpoint string `json:"authorization_endpoint"`
}
