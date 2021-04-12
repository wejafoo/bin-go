package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"weja.us/bingo/internal/blib"
)

const InitFile = ".fd.json"

var (
	Fda blib.FDA
)


func main() {
	mockEnv()

	Fd := blib.SetConfig(getCleanConfig(blib.NewInstance(InitFile)))

	blib.FlexHead()

	if Fd.FdDebug { blib.ShowGlobalDefaults()}
	switch Fd.FdBuildContext {
		case "ng", "angular":		if Fd.Success = blib.NewAngular();		!Fd.Success { log.Fatalf("\n%s", blib.GetAngularError()		)}
		case "ts", "typescript":	if Fd.Success = blib.NewTypescript();	!Fd.Success { log.Fatalf("\n%s", blib.GetTypescriptError()	)}
		case "go":					if Fd.Success = blib.NewGo();			!Fd.Success { log.Fatalf("\n%s", blib.GetGoError()			)}
		case "py", "python":		if Fd.Success = blib.NewPython();		!Fd.Success { log.Fatalf("\n%s", blib.GetPythonError()		)}
		case "do", "docker":		if Fd.Success = blib.NewDocker();		!Fd.Success { log.Fatalf("\n%s", blib.GetDockerError()		)}
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

func getCleanConfig(Fdc blib.FDC) blib.FDC{
	Fda = blib.FDA{
		BuildPtr:            flag.Bool("build", Fdc.FdBuild, "DEFAULT: true - Turns on application-specific builds"),
		DebugPtr:            flag.Bool("debug", Fdc.FdDebug, "Turns on detailed logging and enables a debugger if identified"),
		LocalPtr:            flag.Bool("local", Fdc.FdLocal, "Identifies a build target as local(i.e. not remote)"),
		QuietPtr:            flag.Bool("quiet", Fdc.FdQuiet, "Turns off all logging to STDOUT "),
		RemotePtr:           flag.Bool("remote", Fdc.FdRemote, "Identifies a build target as remote(i.e. not local)"),
		VerbosePtr:          flag.Bool("verbose", Fdc.FdVerbose, "Verbose execution output"),
		BuildContextPtr:     flag.String("context", Fdc.FdBuildContext, "REQUIRED - Boolean that indicates local(-local) or cloud(-remote) deploy"),
		NicknamePtr:         flag.String("nickname", Fdc.FdNickname, "Provides the route for mife"),
		ServiceNamePtr:      flag.String("service", Fdc.FdServiceName, "DEFAULT: $PWD - Working directory and/or Docker Compose service directive."),
		SiteNicknamePtr:     flag.String("site", Fdc.FdSiteNickname, "Provides the route for mife"),
		TargetAliasPtr:      flag.String("alias", Fdc.FdTargetAlias, "Recognizable label added to viewable instance name"),
		TargetDomainPtr:     flag.String("domain", Fdc.FdTargetDomain, "The domain within which the target service will be mapped."),
		TargetImageTagPtr:   flag.String("image", Fdc.FdTargetImageTag, "The default tag of a newly minted build images."),
		TargetLocalPortPtr:  flag.String("port", Fdc.FdTargetLocalPort, "The host port accessible by a user and mapped to the service port"),
		TargetLogLevelPtr:   flag.String("Log", Fdc.FdTargetLogLevel, "DEBUG, INFO, WARNING, ERROR, CRITICAL (default: INFO)"),
		TargetProjectIdPtr:  flag.String("pid", Fdc.FdTargetProjectId, "The project ID used for a cloud-based deployments."),
		TargetRealmPtr:      flag.String("realm", Fdc.FdTargetRealm, "Prefix that, when prepended to root domain, serves as the app OAuth realm."),
		TargetRemotePortPtr: flag.String("port2", Fdc.FdTargetRemotePort, "The actual service port of a running container, rarely available to users."),
	}
	flag.Parse()
	finalConfig := blib.FDC {
		*Fda.BuildPtr,
		*Fda.DebugPtr,
		*Fda.LocalPtr,
		*Fda.QuietPtr,
		*Fda.RemotePtr,
		*Fda.VerbosePtr,
		*Fda.BuildContextPtr,
		*Fda.NicknamePtr,
		*Fda.ServiceNamePtr,
		*Fda.SiteNicknamePtr,
		*Fda.TargetAliasPtr,
		*Fda.TargetDomainPtr,
		*Fda.TargetImageTagPtr,
		*Fda.TargetLocalPortPtr,
		*Fda.TargetLogLevelPtr,
		*Fda.TargetProjectIdPtr,
		*Fda.TargetRealmPtr,
		*Fda.TargetRemotePortPtr,
		false,
	}
	applyConfigRules()

	return finalConfig
}

func mockEnv() {
	// if e := os.Setenv("FD_TARGET_LOG_LEVEL",	"INFO"			); e != nil { fmt.Println("WAAAAAAAAT:", e)}
	// if e := os.Setenv("FD_TARGET_LOCAL_PORT","3000"			); e != nil { fmt.Println("WAAAAAAAAT:", e)}
	// if e := os.Setenv("FD_TARGET_PROJECT_ID","ENV-weja-us"	); e != nil { fmt.Println("WAAAAAAAAT:", e)}
	// if e := os.Setenv("FD_TARGET_IMAGE_TAG",	"ENV-latest"	); e != nil { fmt.Println("WAAAAAAAAT:", e)}
	// if e := os.Setenv("FD_SERVICE_NAME",		"ENV-public"	); e != nil { fmt.Println("WAAAAAAAAT:", e)}
	// if e := os.Setenv("FD_TARGET",			"REMOTE"		); e != nil { fmt.Println("WAAAAAAAAT:", e)}
	// if e := os.Setenv("FD_TARGET_DOMAIN",	"ENV-weja-us"	); e != nil { fmt.Println("WAAAAAAAAT:", e)}
	// if e := os.Setenv("FD_TARGET_ALIAS",		"ENV-wes"		); e != nil { fmt.Println("WAAAAAAAAT:", e)}
	// if e := os.Setenv("FD_BUILD_CONTEXT",	"ENV-ANGULAR"	); e != nil { fmt.Println("WAAAAAAAAT:", e)}
}

// Todo: Consider synchronous logging approach
/*
	import ( "bytes" "fmt" "strconv" "sync" )
	func yourfunc( message string, w *sync.WaitGroup ) {
		defer w.Done()
		b := &bytes.Buffer{}
		defer fmt.Print( b )
		fmt.Fprintf( b, "starting yourfunc with %s", message)
		fmt.Fprintf( b, "message is %s", message)
		fmt.Fprintf( b, "finished yourfunc with %s", message)
	}
	func main() {
		w			:= &sync.WaitGroup{}
		messages	:= make( []string, 0 )
		for i := 0; i < 100; i++ { messages = append(messages, strconv.Itoa( i))}
		for _, m := range messages {
			w.Add( 1 )
			go yourfunc(m, w)
		}
		w.Wait()
	}
*/

