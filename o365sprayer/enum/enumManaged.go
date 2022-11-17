package enum

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/securebinary/o365sprayer/o365sprayer/constants"
	"github.com/securebinary/o365sprayer/o365sprayer/logging"

	"github.com/fatih/color"
)

var countManaged = 0

func counterManaged() {
	countManaged += 1
}

func ValidateEmailManagedO365(command string, email string, file *os.File) {
	getCredentialTypeRequestJSON := constants.GetCredentialTypeRequestJSON{
		Username: email,
	}
	jsonData, _ := json.Marshal(getCredentialTypeRequestJSON)
	client := &http.Client{}
	req, err := http.NewRequest("POST", constants.GET_CREDENTIAL_TYPE, bytes.NewBuffer(jsonData))
	req.Header.Add("User-Agent", constants.USER_AGENTS[rand.Intn(len(constants.USER_AGENTS))])
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var getCredentialTypeResponseJSON constants.GetCredentialTypeResponseJSON
	json.Unmarshal(body, &getCredentialTypeResponseJSON)
	if getCredentialTypeResponseJSON.IfExistsResult == 0 {
		go counterManaged()
		color.Green("[*] Valid User : " + email)
		logging.LogEnumeratedAccount(file, email)
	}
	if command == "standalone" && getCredentialTypeResponseJSON.IfExistsResult != 0 {
		color.Red("[-] Invalid User : " + email)
	}
}

func EnumEmailsManagedO365(domainName string, command string, email string, filepath string, delay float64) {
	color.Yellow("[+] Enumerating Valid O365 Emails - Managed")
	timeStamp := strconv.Itoa(time.Now().Year()) + strconv.Itoa(int(time.Now().Month())) + strconv.Itoa(int(time.Now().Hour())) + strconv.Itoa(int(time.Now().Minute())) + strconv.Itoa(int(time.Now().Second()))
	fileName := domainName + "_enum_" + timeStamp
	logFile, err := os.OpenFile((fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Could not open " + fileName)
		return
	}
	defer logFile.Close()
	if command == "standalone" {
		ValidateEmailManagedO365(command, email, logFile)
	}
	if command == "file" {
		file, err := os.Open(filepath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ValidateEmailManagedO365(command, scanner.Text(), logFile)
			time.Sleep(time.Duration(delay))
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		if countManaged > 0 {
			color.Yellow("[+] " + strconv.Itoa(countManaged) + " Valid O365 Emails Found !")
		} else {
			color.Red("[-] No Valid O365 Email Found !")
		}
	}
}
