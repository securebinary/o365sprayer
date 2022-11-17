package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/securebinary/o365sprayer/o365sprayer/constants"

	"github.com/fatih/color"
)

type EnumResults struct {
	DomainName          string
	FederationBrandName string
	AuthURL             string
	TenandId            string
	NameSpaceType       string
}

func CheckO365(domainname string) EnumResults {
	color.Yellow("[+] Gathering O365 Information For - " + domainname)
	client := &http.Client{}
	req, err := http.NewRequest("GET", constants.GET_USER_REALM+"UsErNaMe@"+domainname, nil)
	req.Header.Add("User-Agent", constants.USER_AGENTS[rand.Intn(len(constants.USER_AGENTS))])
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var getUserRealmJSON constants.GetUserRealmJSON
	json.Unmarshal(body, &getUserRealmJSON)
	domainName := getUserRealmJSON.DomainName
	federationBrandName := getUserRealmJSON.FederationBrandName
	nameSpaceType := getUserRealmJSON.NameSpaceType
	authURL := getUserRealmJSON.AuthURL
	req, err = http.NewRequest("GET", strings.Replace(constants.GET_TENANT_ID, "$DOMAIN", domainName, 1), nil)
	req.Header.Add("User-Agent", constants.USER_AGENTS[rand.Intn(len(constants.USER_AGENTS))])
	resp, err = client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	var getTenantIdJSON constants.GetTenantIdJSON
	json.Unmarshal(body, &getTenantIdJSON)
	tenandId := strings.TrimLeft(strings.TrimRight(getTenantIdJSON.AuthorizationEndpoint, "/oauth2/v2.0/authorize"), "https://login.microsoftonline.com/")
	var enumResult EnumResults
	enumResult.DomainName = domainName
	enumResult.FederationBrandName = federationBrandName
	enumResult.TenandId = tenandId
	enumResult.AuthURL = authURL
	enumResult.NameSpaceType = nameSpaceType
	return enumResult
}
