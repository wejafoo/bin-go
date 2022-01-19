package blib

import (
	"fmt"

	"github.com/gookit/color"
	"github.com/willf/pad"
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

func DeploymentHead() {
	logPrefix := Blue(pad.Right("\n==  Deployment Start", 25, " "))
	dockerLogInfo := Yellow(pad.Left("<docker>", 95, " "))
	googleLogInfo := Yellow(pad.Left("<google>", 95, " "))
	if Fd.FdLocal {
		fmt.Printf("%s%s\n", logPrefix, dockerLogInfo)
	} else {
		fmt.Printf("%s%s\n", logPrefix, googleLogInfo)
	}
	if !Fd.FdQuiet {
		fmt.Printf(" +%s\n", pad.Right("", 118, "-"))
		fmt.Printf(" | %s %s\n", pad.Right("Build? ", 25, "."), Green(Fd.FdBuild))
		fmt.Printf(" | %s %s\n", pad.Right("Clean? ", 25, "."), Green(Fd.FdClean))
		fmt.Printf(" | %s %s\n", pad.Right("Debug? ", 25, "."), Green(Fd.FdDebug))
		fmt.Printf(" | %s %s\n", pad.Right("Local? ", 25, "."), Green(Fd.FdLocal))
		fmt.Printf(" | %s %s\n", pad.Right("Quiet? ", 25, "."), Green(Fd.FdQuiet))
		fmt.Printf(" | %s %s\n", pad.Right("Remote? ", 25, "."), Green(Fd.FdRemote))
		fmt.Printf(" | %s %s\n", pad.Right("Test? ", 25, "."), Green(Fd.FdTest))
		fmt.Printf(" | %s %s\n", pad.Right("Verbose? ", 25, "."), Green(Fd.FdVerbose))
		fmt.Printf(" +%s\n", pad.Right("", 118, "-"))
		fmt.Printf(" | %s %s\n", pad.Right("ADC ", 25, "."), Green(Fd.FdAdc))
		fmt.Printf(" | %s ", pad.Right("Service ", 25, "."))

		if Fd.FdService == "" {
			fmt.Printf("%s\n", Yellow("N/A"))
		} else {
			fmt.Printf("%s\n", Green(Fd.FdService))
		}

		fmt.Printf(" | %s ", pad.Right("Route Base ", 25, "."))

		if Fd.FdRouteBase == "" {
			fmt.Printf("%s\n", Yellow("N/A"))
		} else {
			fmt.Printf("%s\n", Green(Fd.FdRouteBase))
		}

		fmt.Printf(" | %s %s\n", pad.Right("Code Repository ", 25, "."), Green(Fd.FdRepo))
		fmt.Printf(" | %s %s\n", pad.Right("Service Title ", 25, "."), Green(Fd.FdTitle))
		fmt.Printf(" | %s %s\n", pad.Right("Target Alias ", 25, "."), Green(Fd.FdTargetAlias))
		fmt.Printf(" | %s %s\n", pad.Right("Target Domain ", 25, "."), Green(Fd.FdTargetDomain))
		fmt.Printf(" | %s %s\n", pad.Right("Target Image Tag ", 25, "."), Green(Fd.FdTargetImageTag))
		fmt.Printf(" | %s %s\n", pad.Right("Target Local Port ", 25, "."), Green(Fd.FdTargetLocalPort))
		fmt.Printf(" | %s %s\n", pad.Right("Target Log Level ", 25, "."), Green(Fd.FdTargetLogLevel))
		fmt.Printf(" | %s %s\n", pad.Right("Target Project ID ", 25, "."), Green(Fd.FdTargetProjectId))
		fmt.Printf(" | %s %s\n", pad.Right("Target Realm ", 25, "."), Green(Fd.FdTargetRealm))
		fmt.Printf(" | %s %s\n", pad.Right("Target Remote Port ", 25, "."), Green(Fd.FdTargetRemotePort))
		if Fd.FdInit != "" {
			fmt.Printf("\n%s %s\n", pad.Right("Init ", 25, "."), Green(Fd.FdInit))
		}
		fmt.Printf(" +%s", pad.Right("", 118, "-"))

	}
}

func DeploymentFoot(success bool) bool {
	logPrefix := pad.Right("==  Deployment End ", 110, ".")
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

func PipelineHead() {
	fmt.Printf(Blue(pad.Right("\n==  Pipeline Start", 25, " ")))
	fmt.Printf("%s", Yellow(pad.Left("<"+Fdc.FdBuildContext+">", 95, " ")))
}

func PipelineFoot(success bool) bool {
	logPrefix := pad.Right("==  Pipeline End ", 110, ".")
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

func SkipStep(skippedFunc string) {
	logPrefix := Yellow(pad.Right(skippedFunc, 20, " "))
	logMessage := pad.Right("Pipeline step explicitly skipped ", 57, ".")
	logInfo := Blue("--build=false")
	fmt.Printf("\n%s%s%s %s", logPrefix, logMessage, LogLose, logInfo)
}

func FlexHead() {
	fmt.Printf("\n%s%s", Blue(pad.Right("==  MIFE FLEX DEPLOYMENT START", 113, " ")), Yellow("<bingo>"))
	fmt.Printf(Blue(pad.Right("\n", 121, "=")))
	if Fd.FdVerbose {
		logInfo := Blue(".fd." + Fd.FdTargetDomain + ".json")
		fmt.Printf("\n%s", pad.Right("Validating MIFE DEPLOYMENT configuration ", 97, "."))
		if ConfigIsValid {
			fmt.Printf("%s %s", LogWin, logInfo)
		} else {
			fmt.Printf("%s %s", LogLose, logInfo)
		}
	}
}

func FlexFoot(success bool) {
	logPrefix := pad.Right("Cleaning up ", 110, ".")
	logInfo := Blue("done")

	if Fd.FdVerbose {
		fmt.Printf("\n%s%s %s", logPrefix, LogWin, logInfo)
	}
	// Todo: Add deployment cleanup stuff
	if success {
		renderColor = Green
	} else {
		renderColor = Red
	}

	fmt.Printf(renderColor(pad.Right("\n", 121, "=")))
	fmt.Printf(renderColor("\n==  FLEX MIFE DEPLOYMENT END"))

	if Fd.FdRemote {
		fmt.Printf(
			"\n\n%s  %s", Red("REMINDER:"), "This is a remote deployment that may require an update to the mifedom config(s) to enable use.",
		)
	}
	logMessage := "*** Success! Visit your handiwork:"
	if success {
		if Fd.FdLocal {
			fmt.Printf("\n%s  %s%s", logMessage, "http://localhost:", Fd.FdTargetLocalPort)
		} else if Fd.FdTargetAlias == "prod" {
			fmt.Printf("\n%s  %s", logMessage, "https://foo."+Fd.FdTargetDomain)
		} else {
			fmt.Printf("\n%s  %s", logMessage, "https://too."+Fd.FdTargetDomain)
		}
		if Fd.FdRouteBase != "" && !Fd.FdLocal {
			fmt.Printf("%s", Fd.FdRouteBase)
		}
		fmt.Printf(" ***\n\n")
	}
}

func ShowGlobalDefaults() {
	fmt.Printf("\n +%s\n", pad.Right("", 118, "-"))
	fmt.Printf(" | DIAGNOSTIC INFO:\n")
	fmt.Printf(" +%s\n", pad.Right("", 118, "-"))
	fmt.Printf(" | %s %s%s%s\n", pad.Left("Valid FD init file:", 19, " "), Blue(Pwd), Blue("/"), Blue(InitFile))
	fmt.Printf(" | %s %s%s%s\n", pad.Left("Valid FD cfg file:", 19, " "), Blue(Pwd), Blue("/"), Blue(fmt.Sprintf(".fd.%s.json", Fdg.FdTargetDomain)))
	fmt.Printf(" +%s\n", pad.Right("", 118, "-"))
	fmt.Printf(" | %s %s\n", pad.Right("Build? ", 25, "."), Blue(Fdg.FdBuild))
	fmt.Printf(" | %s %s\n", pad.Right("Clean? ", 25, "."), Blue(Fdg.FdClean))
	fmt.Printf(" | %s %s\n", pad.Right("Debug? ", 25, "."), Blue(Fdg.FdDebug))
	fmt.Printf(" | %s %s\n", pad.Right("Local? ", 25, "."), Blue(Fdg.FdLocal))
	fmt.Printf(" | %s %s\n", pad.Right("Quiet? ", 25, "."), Blue(Fdg.FdQuiet))
	fmt.Printf(" | %s %s\n", pad.Right("Remote? ", 25, "."), Blue(Fdg.FdRemote))
	fmt.Printf(" | %s %s\n", pad.Right("Test? ", 25, "."), Blue(Fdg.FdTest))
	fmt.Printf(" | %s %s\n", pad.Right("Verbose? ", 25, "."), Blue(Fdg.FdVerbose))
	fmt.Printf(" +%s\n", pad.Right("", 118, "-"))
	fmt.Printf(" | %s %s\n", pad.Right("ADC", 25, "."), Blue(Fdg.FdAdc))
	fmt.Printf(" | %s %s\n", pad.Right("Build Context", 25, "."), Blue(Fdg.FdBuildContext))
	fmt.Printf(" | %s %s\n", pad.Right("Init", 25, "."), Blue(Fdg.FdInit))
	fmt.Printf(" | %s %s\n", pad.Right("Repository", 25, "."), Blue(Fdg.FdRepo))
	fmt.Printf(" | %s", pad.Right("Service", 25, "."))

	if Fdg.FdService == "" {
		fmt.Printf("%s\n", Yellow(" N/A"))
	} else {
		fmt.Printf("%s\n", Blue(Fdg.FdService))
	}

	fmt.Printf(" | %s ", pad.Right("Route Base", 25, "."))

	if Fdg.FdService == "" {
		fmt.Printf("%s\n", Yellow("N/A"))
	} else {
		fmt.Printf("%s\n", Blue(Fdg.FdService))
	}

	fmt.Printf(" | %s %s\n", pad.Right("Service Title", 25, "."), Blue(Fdg.FdTitle))
	fmt.Printf(" | %s %s\n", pad.Right("Target Alias", 25, "."), Blue(Fdg.FdTargetAlias))
	fmt.Printf(" | %s %s\n", pad.Right("Target Domain", 25, "."), Blue(Fdg.FdTargetDomain))
	fmt.Printf(" | %s %s\n", pad.Right("Target Image Tag", 25, "."), Blue(Fdg.FdTargetImageTag))
	fmt.Printf(" | %s %s\n", pad.Right("Target Local Port", 25, "."), Blue(Fdg.FdTargetLocalPort))
	fmt.Printf(" | %s %s\n", pad.Right("Target Log Level", 25, "."), Blue(Fdg.FdTargetLogLevel))
	fmt.Printf(" | %s %s\n", pad.Right("Target Project ID", 25, "."), Blue(Fdg.FdTargetProjectId))
	fmt.Printf(" | %s %s\n", pad.Right("Target Remote Port", 25, "."), Blue(Fdg.FdTargetRemotePort))
	fmt.Printf(" | %s %s\n", pad.Right("Target Realm", 25, "."), Blue(Fdg.FdTargetRealm))
	fmt.Printf(" +%s", pad.Right("", 118, "-"))
	fmt.Println()
}
