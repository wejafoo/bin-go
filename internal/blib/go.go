

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

var goError error

func NewGo() bool {
	if Fd.FdVerbose { fmt.Printf("%s %s", LogWin, Blue(Fd.FdBuildContext)) }

	DeploymentHead()
	PipelineHead()

	if Fd.FdBuild {goBuild()}   else   {SkipStep("goBuild():")}

	return DeploymentFoot(PipelineFoot(goDeploy()))
}


func goBuild() bool {
	logPrefix	:= Yellow(pad.Right("\ngoBuild():", 20, " "))
	args		:= "build -v -o dist/" + Fd.FdNickname
	argsAbbrev	:= args

	if Fd.FdTest {
		return goTest(goRun(logPrefix, args, argsAbbrev))
	} else {
		return goRun(logPrefix, args, argsAbbrev)
	}

}


func goTest(success bool) bool {
	if !success { return false }
	logPrefix	:= Yellow(pad.Right("\ngoTest():", 20, " "))
	args		:= "test -json ./..."
	argsAbbrev	:= args

	return goRun(logPrefix, args, argsAbbrev)
}


func goDeploy() bool {
	success := true

	if Fd.FdLocal {
		if success = NewDocker();	!success { goError	= GetComposeError() }
	}  else if Fd.FdRemote {
		if success = NewDocker();	!success { ngNpmError = GetComposeError() }
		if success = NewGcp();		!success { ngNpmError = GetGcpError() }
	}

	return success
}


func goRun(logPrefix string, cmdArgs string, cmdArgsAbbrev string) bool {
	if Fd.FdVerbose {
		logCommand	:= BlackOnGray(" go " + cmdArgs + " ")
		fmt.Printf("%s$ %s", logPrefix, logCommand)
		fmt.Printf("\n")
	} else {
		logCommand	:= "go " + cmdArgsAbbrev
		fmt.Printf("%s$ %s", logPrefix, logCommand)
	}

	command	:= exec.Command("go", strings.Split(cmdArgs, " ")...)
	setEnvironment()
	command.Env		 = os.Environ()
	command.Stdout	 = os.Stderr
	stderr, _		:= command.StderrPipe()
	goError			 = command.Start()
	if goError != nil { log.Printf("%s", Red(goError)) }

	scanner		:= bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		stderrText := scanner.Text()
		if Fd.FdVerbose { log.Printf("%s", Grey(stderrText)) }
	}

	goError = command.Wait()
	if goError != nil {
		log.Printf("%s$  %s%s", logPrefix, command, WhiteOnRed(" X "))
		log.Fatalf("%s", Red(goError))
	}

	return true
}


func GetGoError() error { return goError }
