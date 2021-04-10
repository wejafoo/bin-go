package blib

import (
	"encoding/json"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/willf/pad"
	"io/ioutil"
	"os"
)

// Reconciles a config file with any overriding environment variables and
// returns an FDC that serves as default values for the flag parser.

var (
	Fd       FDC
	Fdg      FDC
	Fdc      FDC
	Pwd      string
	InitFile string
)

func NewInstance(file string) FDC {
	InitFile = file
	Pwd, _ = os.Getwd()
	Fdg = FDC{
		true, false, false, false, false, false, "ng", "public",
		"public", "dev", "example.com", "latest", "9999",
		"info", "default-project", "default.realm.", "8080", true,
	}

	Fdc = loadInstance(fmt.Sprintf(".fd.%s.json", initInstance(InitFile).InitTargetDomain))
	if !Fdc.FdBuild {
		Fdc.FdBuild = Fdg.FdBuild
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
	if !Fdc.FdVerbose {
		Fdc.FdVerbose = Fdg.FdVerbose
	}
	if Fdc.FdBuildContext == "" {
		Fdc.FdBuildContext = Fdg.FdBuildContext
	}
	if Fdc.FdNickname == "" {
		Fdc.FdNickname = Fdg.FdNickname
	}
	if Fdc.FdServiceName == "" {
		Fdc.FdServiceName = Fdg.FdServiceName
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
}

func SetConfig(fd FDC) { Fd = fd }

func ShowGlobalDefaults() {
	fmt.Printf("\n      %s", pad.Right("", 25, "-"))
	fmt.Printf("\n    | DIAGNOSTIC INFO:")
	fmt.Printf("\n    | %s", pad.Right("", 25, "-"))
	fmt.Printf("\n    | %s %s%s%s", pad.Left("Valid FD init file:", 25, " "), Blue(Pwd), Blue("/"), Blue(InitFile))
	fmt.Printf("\n    | %s %s%s%s", pad.Left("Valid FD cfg file:", 25, " "), Blue(Pwd), Blue("/"), Blue(fmt.Sprintf(".fd.%s.json", Fdg.FdTargetDomain)))
	fmt.Printf("\n    | %s", pad.Right("", 25, "-"))
	fmt.Printf("\n    | %s %s", pad.Right("Build? ", 25, "."), Blue(Fdg.FdBuild))
	fmt.Printf("\n    | %s %s", pad.Right("Debug? ", 25, "."), Blue(Fdg.FdDebug))
	fmt.Printf("\n    | %s %s", pad.Right("Local? ", 25, "."), Blue(Fdg.FdLocal))
	fmt.Printf("\n    | %s %s", pad.Right("Quiet? ", 25, "."), Blue(Fdg.FdQuiet))
	fmt.Printf("\n    | %s %s", pad.Right("Remote? ", 25, "."), Blue(Fdg.FdRemote))
	fmt.Printf("\n    | %s %s", pad.Right("Verbose? ", 25, "."), Blue(Fdg.FdVerbose))
	fmt.Printf("\n    | %s", pad.Right("", 25, "-"))
	fmt.Printf("\n    | %s %s", pad.Right("Build Context", 25, "."), Blue(Fdg.FdBuildContext))
	fmt.Printf("\n    | %s %s", pad.Right("Nickname", 25, "."), Blue(Fdg.FdNickname))
	fmt.Printf("\n    | %s %s", pad.Right("Service Name", 25, "."), Blue(Fdg.FdServiceName))
	fmt.Printf("\n    | %s %s", pad.Right("Target Alias", 25, "."), Blue(Fdg.FdTargetAlias))
	fmt.Printf("\n    | %s %s", pad.Right("Target Domain", 25, "."), Blue(Fdg.FdTargetDomain))
	fmt.Printf("\n    | %s %s", pad.Right("Target Image Tag", 25, "."), Blue(Fdg.FdTargetImageTag))
	fmt.Printf("\n    | %s %s", pad.Right("Target Local Port", 25, "."), Blue(Fdg.FdTargetLocalPort))
	fmt.Printf("\n    | %s %s", pad.Right("Target Log Level", 25, "."), Blue(Fdg.FdTargetLogLevel))
	fmt.Printf("\n    | %s %s", pad.Right("Target Project ID", 25, "."), Blue(Fdg.FdTargetProjectId))
	fmt.Printf("\n    | %s %s", pad.Right("Target Remote Port", 25, "."), Blue(Fdg.FdTargetRemotePort))
	fmt.Printf("\n    | %s %s", pad.Right("Target Realm", 25, "."), Blue(Fdg.FdTargetRealm))
	fmt.Printf("\n      %s", pad.Right("", 25, "_"))
	fmt.Println()
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
	if e != nil {
		panic(e)
	}
	c := FDC{}
	e = json.Unmarshal(d, &c)
	if e != nil {
		panic(e)
	}
	return c
}
