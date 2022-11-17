package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/securebinary/o365sprayer/o365sprayer/core"
)

func main() {
	domain := flag.String("d", "", "Target domain")
	validateEmail := flag.Bool("enum", false, "Validate O365 emails")
	sprayCheck := flag.Bool("spray", false, "Spray passwords on O365 emails")
	email := flag.String("u", "", "Email to validate")
	emailFile := flag.String("U", "", "Path to email list")
	password := flag.String("p", "", "Password to spray")
	passwordFile := flag.String("P", "", "Path to password list")
	delay := flag.Float64("delay", 0.25, "Delay between requests")
	lockout := flag.Int("lockout", 5, "Number of incorrect attempts for account lockout")
	lockoutDelay := flag.Int("lockoutDelay", 15, "Lockout cool down time")
	maxLockouts := flag.Int("max-lockout", 10, "Maximum number of lockout accounts")
	flag.Usage = func() {
		flagSet := flag.CommandLine
		order := []string{"d", "u", "p", "U", "P", "enum", "spray", "delay", "lockout", "lockoutDelay", "max-lockout"}
		for _, name := range order {
			flag := flagSet.Lookup(name)
			fmt.Printf("  -%s ", flag.Name)
			if len(flag.DefValue) > 0 {
				fmt.Printf("[DEFAULT : %s]", flag.DefValue)
			}
			fmt.Printf("\n      %s\n", flag.Usage)
		}
	}
	flag.Parse()
	fmt.Print(core.BANNER)
	if *domain == "" {
		// flag.PrintDefaults()
		flag.Usage()
		os.Exit(-1)
	}
	// Need to add domain validation
	if len(*domain) > 0 {
		core.StartO365Sprayer(
			*domain,
			*validateEmail,
			*sprayCheck,
			*email,
			*emailFile,
			*password,
			*passwordFile,
			*delay,
			*lockout,
			*lockoutDelay,
			*maxLockouts,
		)

	}
}
