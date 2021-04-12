

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
	tsNpmError error
)


func NewTypescript() bool {
	if Fd.FdVerbose { fmt.Printf("%s %s", LogWin, Blue(Fd.FdBuildContext)) }

	DeploymentHead()
	PipelineHead()

	if Fd.FdBuild {typescriptBuild()}   else   {SkipStep("typescriptBuild():")}

	return DeploymentFoot(PipelineFoot(typescriptDeploy()))
}


func typescriptBuild() bool {

	logPrefix	:= Yellow(pad.Right("\ntypescriptBuild():", 20, " "))
	args		:= "run build:ngssc:" + Fd.FdTargetAlias
	success		:= tsNpmRun(logPrefix, args)

	return success
}


func typescriptDeploy() bool {

	success := true
	if Fd.FdLocal {
		if success = NewDocker(); !success {
			success		= false
			tsNpmError	= GetDockerError()
		}
	}  else if Fd.FdRemote { success = NewGcp() }
	// Todo: Incorporate GoLang native Docker interface in lieu of clunky shell implementation

	return success
}


func tsNpmRun(prefix string, cmdArgs string) bool {

	success 	:= false

	if Fd.FdVerbose {
		logCommand	:= BlackOnGray(" tsNpm " + cmdArgs)
		fmt.Printf("%s$  %s", prefix, logCommand)
		fmt.Printf("\n")
	}

	command		:= exec.Command("tsNpm", strings.Split(cmdArgs, " ")...)
	setEnvironment()
	command.Env	= os.Environ()

	stderr, _	:= command.StderrPipe()
	tsNpmError	= command.Start()
	if tsNpmError != nil { log.Printf("%s", Red(tsNpmError)) }
	scanner		:= bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		stderrText := scanner.Text()
		if Fd.FdVerbose { log.Printf("%s", Grey(stderrText)) }
	}

	tsNpmError = command.Wait()
	if tsNpmError != nil {
		log.Printf("%s$  %s%s", prefix, command, WhiteOnRed(" X "))
		log.Fatalf("\n%s", Red(tsNpmError))
	}
	success = true

	return success
}


func GetTypescriptError() error { return tsNpmError }
