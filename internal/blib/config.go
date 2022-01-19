package blib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/kelseyhightower/envconfig"
)

// Reconciles a config file with any overriding environment variables and
// returns an FDC that serves as default values for the flag parser.

var (
	Fd            FDC
	Fdg           FDC
	Fdc           FDC
	Pwd           string
	InitFile      string
	ConfigIsValid bool
)

func NewInstance(file string) FDC {
	InitFile = file
	ConfigIsValid = false
	Pwd, _ = os.Getwd()
	Fdg = FDC{
		FdBuild:            true,
		FdClean:            false,
		FdDebug:            false,
		FdLocal:            false,
		FdQuiet:            false,
		FdRemote:           false,
		FdTest:             false,
		FdVerbose:          false,
		FdAdc:              "/.secrets/credentials.json",
		FdBuildContext:     "ng",
		FdInit:             "",
		FdService:          "",
		FdRouteBase:        "",
		FdRepo:             "public",
		FdTitle:            "Weja Too",
		FdTargetAlias:      "dev",
		FdTargetDomain:     "weja.us",
		FdTargetImageTag:   "latest",
		FdTargetLocalPort:  "9999",
		FdTargetLogLevel:   "info",
		FdTargetProjectId:  "default-project",
		FdTargetRealm:      "default.realm.",
		FdTargetRemotePort: "8080",
	}

	Fdc = loadInstance(fmt.Sprintf(".fd.%s.json", initInstance(InitFile).InitTargetDomain))

	if !Fdc.FdBuild {
		Fdc.FdBuild = Fdg.FdBuild
	}
	if !Fdc.FdClean {
		Fdc.FdClean = Fdg.FdClean
	}
	if !Fdc.FdDebug {
		Fdc.FdDebug = Fdg.FdDebug
	}
	if !Fdc.FdLocal {
		Fdc.FdLocal = Fdg.FdLocal
	}
	if !Fdc.FdQuiet {
		Fdc.FdQuiet = Fdg.FdQuiet
	}
	if !Fdc.FdRemote {
		Fdc.FdRemote = Fdg.FdRemote
	}
	if !Fdc.FdTest {
		Fdc.FdTest = Fdg.FdTest
	}
	if !Fdc.FdVerbose {
		Fdc.FdVerbose = Fdg.FdVerbose
	}
	if Fdc.FdAdc == "" {
		Fdc.FdAdc = Fdg.FdAdc
	}
	if Fdc.FdBuildContext == "" {
		Fdc.FdBuildContext = Fdg.FdBuildContext
	}
	if Fdc.FdInit == "" {
		Fdc.FdInit = Fdg.FdInit
	}
	if Fdc.FdRepo == "" {
		Fdc.FdRepo = Fdg.FdRepo
	}
	if Fdc.FdService == "" {
		Fdc.FdService = Fdg.FdService
	}
	if Fdc.FdRouteBase == "" {
		Fdc.FdRouteBase = Fdg.FdRouteBase
	}
	if Fdc.FdTitle == "" {
		Fdc.FdTitle = Fdg.FdTitle
	}
	if Fdc.FdTargetAlias == "" {
		Fdc.FdTargetAlias = Fdg.FdTargetAlias
	}
	if Fdc.FdTargetDomain == "" {
		Fdc.FdTargetDomain = Fdg.FdTargetDomain
	}
	if Fdc.FdTargetImageTag == "" {
		Fdc.FdTargetImageTag = Fdg.FdTargetImageTag
	}
	if Fdc.FdTargetLocalPort == "" {
		Fdc.FdTargetLocalPort = Fdg.FdTargetLocalPort
	}
	if Fdc.FdTargetLogLevel == "" {
		Fdc.FdTargetLogLevel = Fdg.FdTargetLogLevel
	}
	if Fdc.FdTargetProjectId == "" {
		Fdc.FdTargetProjectId = Fdg.FdTargetProjectId
	}
	if Fdc.FdTargetRealm == "" {
		Fdc.FdTargetRealm = Fdg.FdTargetRealm
	}
	if Fdc.FdTargetRemotePort == "" {
		Fdc.FdTargetRemotePort = Fdg.FdTargetRemotePort
	}

	e := envconfig.Process("", &Fdc)
	if e != nil {
		panic(e)
	} else {
		return Fdc
	}

	return Fdc
}

func SetConfig(fd FDC) FDC {
	Fd = fd

	return Fd
}

func initInstance(f string) FDI {
	d, e := ioutil.ReadFile(f)
	if e != nil {
		panic(e)
	}
	i := FDI{}
	e = json.Unmarshal(d, &i)
	if e != nil {
		panic(e)
	}

	return i
}

func loadInstance(f string) FDC {
	d, e := ioutil.ReadFile(f)
	// d, e := ioutil.ReadAll(f)
	if e != nil {
		panic(e)
	}
	c := FDC{}
	e = json.Unmarshal(d, &c)
	if e != nil {
		panic(e)
	}
	ConfigIsValid = json.Valid([]byte(d))

	return c
}

// Todo:  Add "--init <DOMAIN_NAME>" argument that passes the full complement of configurations/templates
//  		to an existing non-bingo project (e.g .fd.json, .fd.<DOMAIN_NAME>.json, Dockerfile, docker-compose.yml,
//  			cloudbuild.json, docker-entrypoint.sh, .env.local.yml

func InitConfig(domainName string) {
	if matched, _ := regexp.MatchString("^[0-9a-zA-Z]+(.[0-9a-zA-Z]+)*$", domainName); matched {
		fmt.Println("Valid domain")
	} else {
		fmt.Println("Booooooo! Aaah c'mon, you at least need to provide a valid domain name!!")
		os.Exit(3)
	}

	currentFile := ".fd." + Fd.FdTargetDomain + ".json"
	newFile := ".fd." + domainName + ".json"

	if _, err := os.Stat(InitFile); os.IsNotExist(err) {
		fmt.Println("Okay, looks like this project is all new to BinGo... welcome! ", newFile, " using presets from: ", currentFile)
		fmt.Printf("\nSetting your default working domain in: %s", "test.fd.json")

		defaultWorkingDomain := "{\"InitTargetDomain\": \"test.fd.json\"}"

		fmt.Printf("\nAdding the following JSON representation: %s", defaultWorkingDomain)

		fmt.Printf("Cool, an existing project to work with.", newFile)

		if _, err := os.Stat(newFile); os.IsNotExist(err) {
			fmt.Println("Got it, looks good to add to your new configuration: ", newFile, " using presets from: ", currentFile)
		} else {
			fmt.Printf("Hmm, looks like this domain is already init'd: Could %s be unintended?", newFile)
		}
	} else {
		fmt.Printf("\nHmm, I see we have an existing project to work with... nice.")
		if _, err := os.Stat(newFile); os.IsNotExist(err) {
			fmt.Println("\nRockin' a new configuration init file: ", newFile, " using presets from: ", currentFile)
		} else {
			fmt.Printf("\nHmm, looks like this domain is already init'd: %s  Unintended???", newFile)
		}
	}

}
