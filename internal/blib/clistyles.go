package blib

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/willf/pad"
	"os"
)

var (
	Black        = color.FgBlack.Render
	Blue         = color.FgBlue.Render
	Green        = color.FgGreen.Render
	Grey         = color.FgDarkGray.Render
	Red          = color.FgRed.Render
	Yellow       = color.FgYellow.Render
	BlackOnGray  = color.Style{color.FgBlack, color.BgGray}.Render
	WhiteOnRed   = color.Style{color.FgLightWhite, color.BgRed}.Render
	BlackOnGreen = color.Style{color.FgBlack, color.BgGreen}.Render
	renderColor  = Black
	LogWin       = BlackOnGreen(" √ ")
	LogLose      = WhiteOnRed(" X ")
)

func ContextHead() bool {

	// success	:= true

	fmt.Printf(Blue(pad.Right("\n==  Deployment Start", 25, " ")))

	if Fd.FdLocal {
		switch Fd.FdBuildContext {
		case "ng", "angular", "ts", "typescript", "go", "py", "python", "do", "docker":
			fmt.Printf(Yellow(pad.Left("<docker>", 56, " ")))
		default:
			fmt.Printf("I don't understand	:`(")
			os.Exit(2)
		}
	} else {
		switch Fd.FdBuildContext {
		case "ng", "angular", "ts", "typescript", "go", "py", "python", "do", "docker":
			fmt.Printf(Yellow(pad.Left("<google>", 56, " ")))
		default:
			fmt.Printf("I don't understand	:`(")
			os.Exit(2)
		}
	}

	if !Fd.FdQuiet {
		fmt.Printf("  \n    %s %s", pad.Right("Build? ", 25, "."), Green(Fd.FdBuild))
		fmt.Printf("  \n    %s %s", pad.Right("Debug? ", 25, "."), Green(Fd.FdDebug))
		fmt.Printf("  \n    %s %s", pad.Right("Local? ", 25, "."), Green(Fd.FdLocal))
		fmt.Printf("  \n    %s %s", pad.Right("Quiet? ", 25, "."), Green(Fd.FdQuiet))
		fmt.Printf("  \n    %s %s", pad.Right("Remote? ", 25, "."), Green(Fd.FdRemote))
		fmt.Printf("  \n    %s %s", pad.Right("Verbose? ", 25, "."), Green(Fd.FdVerbose))
		fmt.Printf("\n\n    %s %s", pad.Right("Nickname ", 25, "."), Green(Fd.FdNickname))
		fmt.Printf("  \n    %s %s", pad.Right("Service Name ", 25, "."), Green(Fd.FdServiceName))
		fmt.Printf("  \n    %s %s", pad.Right("Target Alias ", 25, "."), Green(Fd.FdTargetAlias))
		fmt.Printf("  \n    %s %s", pad.Right("Target Domain ", 25, "."), Green(Fd.FdTargetDomain))
		fmt.Printf("  \n    %s %s", pad.Right("Target Image Tag ", 25, "."), Green(Fd.FdTargetImageTag))
		fmt.Printf("  \n    %s %s", pad.Right("Target Local Port ", 25, "."), Green(Fd.FdTargetLocalPort))
		fmt.Printf("  \n    %s %s", pad.Right("Target Log Level ", 25, "."), Green(Fd.FdTargetLogLevel))
		fmt.Printf("  \n    %s %s", pad.Right("Target Project ID  ", 25, "."), Green(Fd.FdTargetProjectId))
		fmt.Printf("  \n    %s %s", pad.Right("Target Realm ", 25, "."), Green(Fd.FdTargetRealm))
		fmt.Printf("  \n    %s %s", pad.Right("Target Remote Port ", 25, "."), Green(Fd.FdTargetRemotePort))

		fmt.Printf(Blue(pad.Right("\n==  Pipeline Start", 25, " ")))

		switch Fdc.FdBuildContext {
		case "ng", "angular", "ts", "typescript", "go", "py", "python", "do", "docker":
			fmt.Printf("%s", Yellow(pad.Left("<"+Fdc.FdBuildContext+">", 56, " ")))
		default:
			fmt.Printf("I don't understand	:`(")
			os.Exit(2)
		}
	}

	return true
}

func PipelineFoot(success bool) bool {

	logPrefix := pad.Right("==  Pipeline End ", 77, ".")
	logInfo := Blue(Fdc.FdBuildContext)

	if success {
		renderColor = Green
		fmt.Printf("\n%s%s %s", renderColor(logPrefix), LogWin, logInfo)
	} else {
		renderColor = Red
		fmt.Printf("\n%s%s %s", renderColor(logPrefix), LogLose, logInfo)
	}

	// Todo: Add pipeline cleanup stuffz

	return success
}

func DeploymentFoot(success bool) bool {

	logPrefix := pad.Right("==  Deployment End ", 77, ".")
	logInfo := Blue("")

	if Fd.FdLocal {
		logInfo = Blue("docker")
	} else {
		logInfo = Blue("google")
	}
	if success {
		renderColor = Green
		fmt.Printf("\n%s%s %s", renderColor(logPrefix), LogWin, logInfo)
	} else {
		renderColor = Red
		fmt.Printf("\n%s%s %s", renderColor(logPrefix), LogLose, logInfo)
	}

	// Todo: Add pipeline cleanup stuffz

	return success
}

func FlexHead() {

	fmt.Printf(Blue(pad.Right("\n", 81, "=")))
	fmt.Printf("\n%s%s", Blue(pad.Right("==  FLEX DEPLOYMENT START", 73, " ")), Yellow("<bingo>"))
	fmt.Printf(Blue(pad.Right("\n", 81, "=")))
	fmt.Printf("\n%s", pad.Right("Compiling FLEX DEPLOY configuration ", 77, "."))
}

func FlexFoot(success bool) {

	logPrefix := pad.Right("Cleaning up ", 77, ".")
	logInfo := Blue("done")

	// Todo: Add deployment cleanup stuff

	fmt.Printf("\n%s%s %s", logPrefix, LogWin, logInfo)

	if success {
		renderColor = Green
	} else {
		renderColor = Red
	}

	fmt.Printf(renderColor(pad.Right("\n", 81, "=")))
	fmt.Printf(renderColor("\n==  FLEX DEPLOYMENT END"))
	fmt.Printf(renderColor(pad.Right("\n", 81, "=")))
	fmt.Println()

	logMessage := "*** Success! Visit your handiwork on:"
	if success {
		if Fd.FdLocal {
			fmt.Printf("%s  %s ***\n\n", logMessage, "http://localhost:"+Fd.FdTargetLocalPort+"/"+Fd.FdNickname)
		} else if Fd.FdTargetAlias == "prod" {
			fmt.Printf("%s  %s ***\n\n", logMessage, "https://foo.fb."+Fd.FdTargetDomain+"/"+Fd.FdNickname)
		} else {
			fmt.Printf("%s  %s ***\n\n", logMessage, "https://too.fb."+Fd.FdTargetDomain+"/"+Fd.FdNickname)
		}
	}
}

func SkipStep(skippedFunc string) {

	logPrefix := Yellow(pad.Right(skippedFunc, 20, " "))
	logMessage := pad.Right("Pipeline step explicitly skipped ", 57, ".")
	logInfo := Blue("--build=false")

	fmt.Printf("\n%s%s%s %s", logPrefix, logMessage, LogLose, logInfo)
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
