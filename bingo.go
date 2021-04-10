package main

import (
	"flag"
	"fmt"
	"os"
	"weja.us/bingo/internal/blib"
)

const InitFile = ".fd.json"

var (
	Fd  blib.FDC // final	config	( passed to builders via blib.SetConfig())
	Fdc blib.FDC // project	config	( from .fd.*.json and ENV vars, serves as defaults for "Fda" )
	Fda blib.FDA // instance	config	( from CLI args )
)

func main() {
	mockEnv()
	blib.FlexHead()
	Fdc = blib.NewInstance(InitFile)
	Fda = blib.FDA{
		BuildPtr:            flag.Bool("build", Fdc.FdBuild, "DEFAULT: true - Turns on application-specific builds"),
		DebugPtr:            flag.Bool("debug", Fdc.FdDebug, "Turns on detailed logging and enables a debugger if identified"),
		LocalPtr:            flag.Bool("local", Fdc.FdLocal, "Identifies a build target as local(i.e. not remote)"),
		QuietPtr:            flag.Bool("quiet", Fdc.FdQuiet, "Turns off all logging to STDOUT "),
		RemotePtr:           flag.Bool("remote", Fdc.FdRemote, "Identifies a build target as remote(i.e. not local)"),
		VerbosePtr:          flag.Bool("verbose", Fdc.FdVerbose, "Verbose execution output"),
		BuildContextPtr:     flag.String("BuildContext", Fdc.FdBuildContext, "REQUIRED - Boolean that indicates local(-local) or cloud(-remote) deploy"),
		NicknamePtr:         flag.String("Nickname", Fdc.FdNickname, "Provides the route for mife"),
		ServiceNamePtr:      flag.String("ServiceName", Fdc.FdServiceName, "DEFAULT: $PWD - Working directory and/or Docker Compose service directive."),
		TargetAliasPtr:      flag.String("TargetAlias", Fdc.FdTargetAlias, "Recognizable label added to viewable instance name"),
		TargetDomainPtr:     flag.String("TargetDomain", Fdc.FdTargetDomain, "The domain within which the target service will be mapped."),
		TargetImageTagPtr:   flag.String("TargetImageTag", Fdc.FdTargetImageTag, "The default tag of a newly minted build images."),
		TargetLocalPortPtr:  flag.String("TargetLocalPort", Fdc.FdTargetLocalPort, "The host port accessible by a user and mapped to the service port"),
		TargetLogLevelPtr:   flag.String("TargetLogLevel", Fdc.FdTargetLogLevel, "DEBUG, INFO, WARNING, ERROR, CRITICAL (default: INFO)"),
		TargetProjectIdPtr:  flag.String("TargetProjectID", Fdc.FdTargetProjectId, "The project ID used for a cloud-based deployments."),
		TargetRealmPtr:      flag.String("Target Realm", Fdc.FdTargetRealm, "Prefix that, when prepended to root domain, serves as the app OAuth realm."),
		TargetRemotePortPtr: flag.String("TargetRemotePort", Fdc.FdTargetRemotePort, "The actual service port of a running container, rarely available to users."),
	}

	flag.Parse()
	applyConfigRules()

	Fd = blib.FDC{
		*Fda.BuildPtr,
		*Fda.DebugPtr,
		*Fda.LocalPtr,
		*Fda.QuietPtr,
		*Fda.RemotePtr,
		*Fda.VerbosePtr,
		*Fda.BuildContextPtr,
		*Fda.NicknamePtr,
		*Fda.ServiceNamePtr,
		*Fda.TargetAliasPtr,
		*Fda.TargetDomainPtr,
		*Fda.TargetImageTagPtr,
		*Fda.TargetLocalPortPtr,
		*Fda.TargetLogLevelPtr,
		*Fda.TargetProjectIdPtr,
		*Fda.TargetRealmPtr,
		*Fda.TargetRemotePortPtr,
		true,
	}

	blib.SetConfig(Fd)

	if Fd.FdDebug {
		blib.ShowGlobalDefaults()
	}

	switch Fd.FdBuildContext {
	case "ng", "angular":
		if Fd.Success = blib.NewAngular(); !Fd.Success {
			os.Exit(blib.AngularBuildExit())
		}
	case "ts", "typescript":
		if Fd.Success = blib.NewTypescript(); !Fd.Success {
			os.Exit(blib.TypescriptBuildExit())
		}
	case "go":
		if Fd.Success = blib.NewGo(); !Fd.Success {
			os.Exit(blib.GoBuildExit())
		}
	case "py", "python":
		if Fd.Success = blib.NewPython(); !Fd.Success {
			os.Exit(blib.PythonBuildExit())
		}
	case "do", "docker":
		if Fd.Success = blib.NewDocker(); !Fd.Success {
			os.Exit(blib.DockerBuildExit())
		}
	default:
		fmt.Printf("%s %s\n\nBuild context: %s was not found... quitting\n\n", blib.LogLose, blib.Red("ALL Bad!"), blib.Red(*Fda.BuildContextPtr))
		os.Exit(1)
	}

	blib.FlexFoot(true)
}

func applyConfigRules() {

	if *Fda.RemotePtr {
		*Fda.LocalPtr = false
	} else if !*Fda.LocalPtr {
		fmt.Println(blib.Red("\n\nEither '-local' or '-remote' must be true... quitting"))
		fmt.Printf("\n\n")
		os.Exit(2)
	}

	if (*Fda.DebugPtr || *Fda.VerbosePtr) && *Fda.QuietPtr {
		fmt.Println(blib.Red("\n\nNeither '-debug' nor '-verbose' can be true when '-quiet' is true... quitting"))
		fmt.Printf("\n\n")
		os.Exit(2)
	}
}

func mockEnv() {
	// if e := os.Setenv("FD_TARGET_LOG_LEVEL","INFO"	); e != nil { fmt.Println("WAAAAAAAAT:", e)}
	// if e := os.Setenv("FD_TARGET_LOCAL_PORT","3000"		); e != nil { fmt.Println("WAAAAAAAAT:", e)}
	// if e := os.Setenv("FD_TARGET_PROJECT_ID", "ENV-weja-us"		); e != nil { fmt.Println(	"WAAAAAAAAT:", e)}
	// if e := os.Setenv("FD_TARGET_IMAGE_TAG",	"ENV-latest"		); e != nil { fmt.Println(	"WAAAAAAAAT:", e)}
	// if e := os.Setenv("FD_SERVICE_NAME",		"ENV-public"		); e != nil { fmt.Println(	"WAAAAAAAAT:", e)}
	// if e := os.Setenv("FD_TARGET",			"REMOTE"			); e != nil { fmt.Println(	"WAAAAAAAAT:", e)}
	// if e := os.Setenv("FD_TARGET_DOMAIN",	"ENV-weja-us"		); e != nil { fmt.Println(	"WAAAAAAAAT:", e)}
	// if e := os.Setenv("FD_TARGET_ALIAS",		"ENV-wes"			); e != nil { fmt.Println(	"WAAAAAAAAT:", e)}
	// if e := os.Setenv("FD_BUILD_CONTEXT",	"ENV-ANGULAR"		); e != nil { fmt.Println(	"WAAAAAAAAT:", e)}
}
