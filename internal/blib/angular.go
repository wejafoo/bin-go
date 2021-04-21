

package blib

import (
	"bufio"
	"fmt"
	"github.com/willf/pad"
	"log"
	"os"
	"os/exec"
	"strings"
)

var ngNpmError error

func NewAngular() bool {
	if Fd.FdVerbose { fmt.Printf("%s %s", LogWin, Blue(Fd.FdBuildContext)) }

	DeploymentHead()
	PipelineHead()

	if Fd.FdBuild {angularBuild()}   else   {SkipStep("angularBuild():")}

	return DeploymentFoot(PipelineFoot(angularDeploy()))
}


func angularBuild() bool {
	logPrefix	:= Yellow(pad.Right("\nangularBuild():", 20, " "))
	args		:= "run build:ngssc:" + Fd.FdTargetAlias
	argsAbbrev	:= args

	return ngNpmRun(logPrefix, args, argsAbbrev)
}


func angularTest() bool {
	logPrefix	:= Yellow(pad.Right("\nangularTest():", 20, " "))
	args		:= "run test:" + Fd.FdTargetAlias
	argsAbbrev	:= args

	return ngNpmRun(logPrefix, args, argsAbbrev)
}


func angularDeploy() bool {
	success := true

	if Fd.FdLocal {
		if success = NewDocker();	!success { ngNpmError = GetComposeError() }
	}  else if Fd.FdRemote {
		if success = NewDocker();	!success { ngNpmError = GetComposeError() }
		if success = NewGcp();		!success { ngNpmError = GetGcpError() }
	}
	// Todo: Investigate Go native Docker in lieu of this clunky OS implementation
	return success
}


func ngNpmRun(logPrefix string, cmdArgs string, cmdArgsAbbrev string) bool {
	if Fd.FdVerbose {
		logCommand	:= BlackOnGray(" npm " + cmdArgs)
		fmt.Printf("%s$ %s", logPrefix, logCommand)
		fmt.Printf("\n")
	} else {
		logCommand	:= "npm " + cmdArgsAbbrev
		fmt.Printf("%s$ %s", logPrefix, logCommand)
		fmt.Printf("\n")
	}

	command			:= exec.Command("npm", strings.Split(cmdArgs, " ")...)
	setEnvironment()
	command.Env	= os.Environ()

	stderr, _		:= command.StderrPipe()
	ngNpmError		= command.Start()
	if ngNpmError != nil { log.Printf("%s", Red(ngNpmError)) }

	scanner		:= bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		stderrText := scanner.Text()
		if Fd.FdVerbose { log.Printf("%s", Grey(stderrText)) }
	}

	ngNpmError = command.Wait()
	if ngNpmError != nil {
		log.Printf("%s$  %s%s", logPrefix, command, WhiteOnRed(" X "))
		log.Fatalf("%s", Red(ngNpmError))
	}

	return true
}


func GetAngularError() error { return ngNpmError }
