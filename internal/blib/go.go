

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

var (
	goError error
)


func NewGo() bool {

	fmt.Printf("%s %s", LogWin, Blue(Fd.FdBuildContext))
	DeploymentHead()
	PipelineHead()
	if Fd.FdBuild {goBuild()}   else   {SkipStep("goBuild():")}

	return DeploymentFoot(PipelineFoot(goDeploy()))
}


func goBuild() bool {

	logPrefix	:= Yellow(pad.Right("\ngoBuild():", 20, " "))
	args		:= "run build:ngssc:" + Fd.FdTargetAlias
	success		:= goRun(logPrefix, args)

	return success
}


func goDeploy() bool {

	success := true
	if Fd.FdLocal {
		if success = NewDocker(); !success {
			success	= false
			goError	= GetDockerError()
		}
	}  else if Fd.FdRemote { success = NewGcp() }

	// Todo: Incorporate GoLang native Docker interface in lieu of clunky shell implementation

	return success
}


func GetGoError() error { return goError }


func goRun(prefix string, cmdArgs string) bool {

	success 	:= false
	logCommand	:= BlackOnGray("go " + cmdArgs)

	fmt.Printf("%s$  %s", prefix, logCommand)
	if Fd.FdVerbose { fmt.Printf("\n") }

	command		:= exec.Command("go", strings.Split(cmdArgs, " ")...)
	setEnvironment()
	command.Env	= os.Environ()

	stderr, _		:= command.StderrPipe()
	goError	= command.Start()
	if goError != nil { log.Printf("%s", Red(goError)) }
	scanner			:= bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		stderrText := scanner.Text()
		if Fd.FdVerbose { log.Printf("%s", Grey(stderrText)) }
	}

	goError = command.Wait()
	if goError != nil {
		log.Printf("%s$  %s%s", prefix, command, WhiteOnRed(" X "))
		log.Fatal("%s", Red(goError))
	}
	success = true

	return success
}
