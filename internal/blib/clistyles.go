

package blib

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/willf/pad"
	"os"
)

var (
	Black  	= color.FgBlack.Render
	Blue	= color.FgBlue.Render
	Green  	= color.FgGreen.Render
	Red  	= color.FgRed.Render
	Yellow 	= color.FgYellow.Render
	// Magenta	= color.FgMagenta.Render
	// Grey		= color.FgDarkGray.Render
	// Cyan		= color.FgCyan.Render
	renderColor = Black
)

func ContextHead() bool {
	success	:= true
	
	switch Fd.FdBuildContext {
		case "ng", "angular":		fmt.Printf( Blue( "\n  ==  ANGULAR DEPLOYMENT START"			))
		case "ts", "typescript":	fmt.Printf( Blue( "\n  ==  TYPESCRIPT DEPLOYMENT START"			))
		case "go":					fmt.Printf( Blue( "\n  ==  GO DEPLOYMENT START"					))
		case "py", "python":		fmt.Printf( Blue( "\n  ==  PYTHON DEPLOYMENT START"				))
		case "do", "docker":		fmt.Printf( Blue( "\n  ==  DOCKER(GENERIC) DEPLOYMENT START"	))
		default:					fmt.Printf("I don't understand	:`(" )
									os.Exit(2)
	}

	if ! Fd.FdQuiet {
		fmt.Println()
		fmt.Printf("%s %s\n", 		pad.Left("Deployment context:",	29," "),Green(Fd.FdBuildContext		))
		fmt.Printf("\n    %s %s",	pad.Right("Local? ",			25,"."),Green(Fd.FdLocal			))
		fmt.Printf("\n    %s %s",	pad.Right("Remote? ",			25,"."),Green(Fd.FdRemote			))
		fmt.Printf("\n    %s %s",	pad.Right("Debug? ",			25,"."),Green(Fd.FdDebug			))
		fmt.Printf("\n    %s %s",	pad.Right("Quiet? ",			25,"."),Green(Fd.FdQuiet			))
		fmt.Printf("\n    %s %s\n",	pad.Right("Verbose? ",			25,"."),Green(Fd.FdVerbose			))
		fmt.Printf("\n    %s %s",	pad.Right("Nickname ",			25,"."),Green(Fd.FdNickname			))
		fmt.Printf("\n    %s %s",	pad.Right("Service Name ",		25,"."),Green(Fd.FdServiceName		))
		fmt.Printf("\n    %s %s",	pad.Right("Target Alias ",		25,"."),Green(Fd.FdTargetAlias		))
		fmt.Printf("\n    %s %s",	pad.Right("Target Domain ",		25,"."),Green(Fd.FdTargetDomain		))
		fmt.Printf("\n    %s %s",	pad.Right("Target Image Tag ",	25,"."),Green(Fd.FdTargetImageTag	))
		fmt.Printf("\n    %s %s",	pad.Right("Target Local Port ",	25,"."),Green(Fd.FdTargetLocalPort	))
		fmt.Printf("\n    %s %s",	pad.Right("Target Log Level ",	25,"."),Green(Fd.FdTargetLogLevel	))
		fmt.Printf("\n    %s %s",	pad.Right("Target Project ID  ",25,"."),Green(Fd.FdTargetProjectId	))
		fmt.Printf("\n    %s %s",	pad.Right("Target Realm ",		25,"."),Green(Fd.FdTargetRealm		))
		fmt.Printf("\n    %s %s",	pad.Right("Target Remote Port ",25,"."),Green(Fd.FdTargetRemotePort	))
		
		switch Fdc.FdBuildContext {
			case "ng", "angular":
				fmt.Printf( Blue( "\n\n    ==  ANGULAR PIPELINE START		"))
				color.Style{ color.Yellow, color.OpItalic }.Printf("\t<<<angular.go>>>\n\n"	)
			case "ts", "typescript":
				fmt.Printf( Blue( "\n\n    ==  TYPESCRIPT PIPELINE START	"))
				color.Style{color.Yellow, color.OpItalic}.Printf("\t<<<typescript.go>>>\n\n"	)
			case "go":
				fmt.Printf( Blue( "\n\n    ==  GO PIPELINE START			"))
				color.Style{color.Yellow, color.OpItalic}.Printf("\t<<<go.go>>>\n\n"			)
			case "py", "python":
				fmt.Printf( Blue( "\n\n    ==  PYTHON PIPELINE START		"))
				color.Style{color.Yellow, color.OpItalic}.Printf("\t<<<python.go>>>\n\n"		)
			case "do", "docker":
				fmt.Printf( Blue( "\n\n  ==  DOCKER(GENERIC) PIPELINE START	"))
				color.Style{color.Yellow, color.OpItalic}.Printf("\t<<<docker.go>>>\n\n"		)
			default:
				fmt.Printf("%s", Red("I don't understand this build context	:`(")		)
				os.Exit(2)
		}
	}
	
	return success
}

func ContextFoot( Success bool ) bool {
	renderColor = Red
	if Success { renderColor = Green }
	switch Fdc.FdBuildContext {
		case "ng", "angular":		fmt.Printf(renderColor("\n\n    ==  ANGULAR PIPELINE END			"))
		case "ts", "typescript":	fmt.Printf(renderColor("\n\n    ==  TYPESCRIPT PIPELINE END			"))
		case "go":					fmt.Printf(renderColor("\n\n    ==  GO PIPELINE END					"))
		case "py", "python":		fmt.Printf(renderColor("\n\n    ==  PYTHON PIPELINE END				"))
		case "do", "docker":		fmt.Printf(renderColor("\n\n    ==  DOCKER(GENERIC) PIPELINE END	"))
		default:
			fmt.Printf("\n\n", pad.Left(Red("I don't understand the context :`("),8," "))
			os.Exit(2 )
	}
	color.Style{ color.Yellow, color.OpItalic }.Printf( "<<<Success: %v", Black( Success ))
	color.Style{ color.Yellow, color.OpItalic }.Printf( ">>>")
	fmt.Printf("\n    Cleaning up...\n")
	
	
	// Todo: Add pipeline cleanup stuffz
	
	
	fmt.Printf("    done\n")
	renderColor := Red
	if Success { renderColor = Green }
	
	switch Fdc.FdBuildContext {
		case "ng", "angular":		fmt.Printf( renderColor("\n  ==  ANGULAR DEPLOYMENT END			"))
		case "ts", "typescript":	fmt.Printf( renderColor("\n  ==  TYPESCRIPT DEPLOYMENT END		"))
		case "go":					fmt.Printf( renderColor("\n  ==  GO DEPLOYMENT END				"))
		case "py", "python":		fmt.Printf( renderColor("\n  ==  PYTHON DEPLOYMENT END			"))
		case "do", "docker":		fmt.Printf( renderColor("\n  ==  DOCKER(GENERIC) DEPLOYMENT END	"))
		default:
			fmt.Printf("I don't understand	:`(")
			os.Exit(2)
	}
	
	color.Style{ color.Yellow, color.OpItalic }.Printf( "<<<Success: %v", Black( Success ))
	color.Style{ color.Yellow, color.OpItalic }.Printf( ">>>")
	
	return Success
}

func DeployHead() {
	fmt.Printf	( Blue		(pad.Right("\n", 80, "=")))
	fmt.Printf	( Blue		("\n==  FLEX DEPLOYMENT START" 		))
	fmt.Printf	( Yellow	("\t\t\t<<<bingo.go>>>"				))
	fmt.Printf	( Blue		(pad.Right("\n", 80, "=")))
	fmt.Println	("\n  Compiling FLEX DEPLOY configuration... done." )
}

func DeployFoot(Success bool) {
	fmt.Printf("\n  Cleaning up...\n")
	
	
	// Todo: Add deployment cleanup stuff
	
	
	fmt.Printf("  done")
	renderColor := Red
	if Success { renderColor = Green }
	fmt.Printf( renderColor( pad.Right("\n", 80, "=" )))
	fmt.Printf( renderColor("\n==  FLEX DEPLOYMENT END" ))
	color.Style{ color.Yellow, color.OpItalic }.Printf("\t\t\t\t<<<Success: %v", Black( Success ))
	color.Style{ color.Yellow, color.OpItalic }.Printf(">>>")
	fmt.Printf( renderColor( pad.Right("\n", 80, "=" )))
	fmt.Println()
}

/*  !!!!!!!!!! POTENTIAL SYNCHRONOUS LOGGING APPROACH !!!!!!!!!!
import ( "bytes" "fmt" "strconv" "sync" )
func yourfunc( message string, w *sync.WaitGroup ) {
	defer w.Done()
	b := &bytes.Buffer{}
	defer fmt.Print( b )
	fmt.Fprintf( b, "starting yourfunc with %s\n", message)
	fmt.Fprintf( b, "message is %s\n", message)
	fmt.Fprintf( b, "finished yourfunc with %s\n", message)
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