package blib

import (
	"encoding/json"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"io/ioutil"
	"os"
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
	InitFile		= file
	ConfigIsValid	= false
	Pwd, _			= os.Getwd()
	Fdg = FDC {
		true,
		false,
		false,
		false,
		false,
		false,
		"ng",
		"public",
		"public",
		"Weja Too",
		"dev",
		"example.com",
		"latest",
		"9999",
		"info",
		"default-project",
		"default.realm.",
		"8080",
		true,
	}

	Fdc = loadInstance(fmt.Sprintf(".fd.%s.json", initInstance(InitFile).InitTargetDomain))

	if !Fdc.FdBuild						{ Fdc.FdBuild				= Fdg.FdBuild				}
	if !Fdc.FdDebug						{ Fdc.FdDebug				= Fdg.FdDebug				}
	if !Fdc.FdLocal						{ Fdc.FdLocal				= Fdg.FdLocal				}
	if !Fdc.FdQuiet						{ Fdc.FdQuiet				= Fdg.FdQuiet				}
	if !Fdc.FdRemote					{ Fdc.FdRemote				= Fdg.FdRemote				}
	if !Fdc.FdVerbose					{ Fdc.FdVerbose				= Fdg.FdVerbose				}
	if Fdc.FdBuildContext		== ""	{ Fdc.FdBuildContext		= Fdg.FdBuildContext		}
	if Fdc.FdNickname			== ""	{ Fdc.FdNickname			= Fdg.FdNickname			}
	if Fdc.FdServiceName		== ""	{ Fdc.FdServiceName			= Fdg.FdServiceName			}
	if Fdc.FdSiteNickname		== ""	{ Fdc.FdSiteNickname		= Fdg.FdSiteNickname		}
	if Fdc.FdTargetAlias		== ""	{ Fdc.FdTargetAlias			= Fdg.FdTargetAlias			}
	if Fdc.FdTargetDomain		== ""	{ Fdc.FdTargetDomain		= Fdg.FdTargetDomain		}
	if Fdc.FdTargetImageTag		== ""	{ Fdc.FdTargetImageTag		= Fdg.FdTargetImageTag		}
	if Fdc.FdTargetLocalPort	== ""	{ Fdc.FdTargetLocalPort		= Fdg.FdTargetLocalPort		}
	if Fdc.FdTargetLogLevel		== ""	{ Fdc.FdTargetLogLevel		= Fdg.FdTargetLogLevel		}
	if Fdc.FdTargetProjectId	== ""	{ Fdc.FdTargetProjectId		= Fdg.FdTargetProjectId		}
	if Fdc.FdTargetRealm		== ""	{ Fdc.FdTargetRealm			= Fdg.FdTargetRealm			}
	if Fdc.FdTargetRemotePort	== ""	{ Fdc.FdTargetRemotePort	= Fdg.FdTargetRemotePort	}

	e := envconfig.Process("", &Fdc)
	if e != nil { panic(e) } else { return Fdc }
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
