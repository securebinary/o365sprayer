package constants

import (
	"reflect"
)

var RESOURCES = []string{
	"https://graph.windows.net",
	"https://graph.microsoft.com",
}

var SCOPES = []string{
	".default",
	"openid",
	"profile",
	"offline_access",
}

// https://github.com/secureworks/TokenMan#foci-application-client-id-map
var CLIENT_IDS = map[string]string{
	"Accounts Control UI":                      "a40d7d7d-59aa-447e-a655-679a4107e548",
	"Microsoft Authenticator App":              "4813382a-8fa7-425e-ab75-3b753aab3abb",
	"Microsoft Azure CLI":                      "04b07795-8ddb-461a-bbee-02f9e1bf7b46",
	"Microsoft Azure PowerShell":               "1950a258-227b-4e31-a9cf-717495945fc2",
	"Microsoft Bing Search for Microsoft Edge": "2d7f3606-b07d-41d1-b9d2-0d0c9296a6e8",
	"Microsoft Bing Search":                    "cf36b471-5b44-428c-9ce7-313bf84528de",
	"Microsoft Edge":                           "f44b1140-bc5e-48c6-8dc0-5cf5a53c0e34",
	"Microsoft Edge (1)":                       "e9c51622-460d-4d3d-952d-966a5b1da34c",
	"Microsoft Edge AAD BrokerPlugin":          "ecd6b820-32c2-49b6-98a6-444530e5a77a",
	"Microsoft Flow":                           "57fcbcfa-7cee-4eb1-8b25-12d2030b4ee0",
	"Microsoft Intune Company Portal":          "9ba1a5c7-f17a-4de9-a1f1-6178c8d51223",
	"Microsoft Office":                         "d3590ed6-52b3-4102-aeff-aad2292ab01c",
	"Microsoft Planner":                        "66375f6b-983f-4c2c-9701-d680650f588f",
	"Microsoft Power BI":                       "c0d2a505-13b8-4ae0-aa9e-cddd5eab0b12",
	"Microsoft Stream Mobile Native":           "844cca35-0656-46ce-b636-13f48b0eecbd",
	"Microsoft Teams - Device Admin Agent":     "87749df4-7ccf-48f8-aa87-704bad0e0e16",
	"Microsoft Teams":                          "1fec8e78-bce4-4aaf-ab1b-5451cc387264",
	"Microsoft To-Do client":                   "22098786-6e16-43cc-a27d-191a01a1e3b5",
	"Microsoft Tunnel":                         "eb539595-3fe1-474e-9c1d-feb3625d1be5",
	"Microsoft Whiteboard Client":              "57336123-6e14-4acc-8dcf-287b6088aa28",
	"Office 365 Management":                    "00b41c95-dab0-4487-9791-b9d2c32c80f2",
	"Office UWP PWA":                           "0ec893e0-5785-4de6-99da-4ed124e5296c",
	"OneDrive iOS App":                         "af124e86-4e96-495a-b70a-90f90ab96707",
	"OneDrive SyncEngine":                      "ab9b8c07-8f02-4f72-87fa-80105867a763",
	"OneDrive":                                 "b26aadf8-566f-4478-926f-589f601d9c74",
	"Outlook Mobile":                           "27922004-5251-4030-b22d-91ecd9a37ea4",
	"PowerApps":                                "4e291c71-d680-4d0e-9640-0a3358e31177",
	"SharePoint Android":                       "f05ff7c9-f75a-4acd-a3b5-f4b6a870245d",
	"SharePoint":                               "d326c1ce-6cc6-4de2-bebc-4591e5e13ef0",
	"Visual Studio":                            "872cd9fa-d31f-45e0-9eab-6e460a02d1f1",
	"Windows Search":                           "26a7ee05-5602-4d76-a7ba-eae8b7b67941",
	"Yammer iPhone":                            "a569458c-7f2b-45cb-bab9-b7dee514d112",
}

var GRANT_TYPE = "password"

// Azure AD Authentication Error Codes
// https://learn.microsoft.com/en-us/azure/active-directory/develop/reference-aadsts-error-codes
// AADSTS5XXXX series to be used for Auth validation - as per 16/11/2022
// Not neccessary to validate all AADST codes, only related to Auth is enough
var AADST_ERROR_CODES = map[string]string{
	"50034": "User Not Found",
	"50053": "Account Locked Out",
	"50126": "Account Exists, Invalid Password",
}

// Validate User Using
const GET_OAUTH_TOKEN = "https://login.microsoft.com/common/oauth2/token"

func GetMapItemRandKey(m map[string]string) string {
	return reflect.ValueOf(m).MapKeys()[0].String()
}

// GetOauthToken JSON Response
type GetOauthTokenResponseJSON struct {
	Error            string `json:"error"`
	ErrorDescription int    `json:"error_description"`
	ErrorCodes       []int  `json:"error_codes"`
	TimeStamp        string `json:"timestamp"`
	TraceId          string `json:"trace_id"`
	CorrelationId    string `json:"correlation_id"`
	ErrorURI         string `json:"error_uri"`
}
