

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

var tsNpmError error


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
	argsAbbrev	:= args

	return tsNpmRun(logPrefix, args, argsAbbrev)
}


func typescriptDeploy() bool {
	success := true

	if Fd.FdLocal {
		if success = NewDocker(); !success { tsNpmError = GetDockerError() }
	}  else if Fd.FdRemote { success = NewGcp() }

	return success
}


func tsNpmRun(logPrefix string, cmdArgs string, cmdArgsAbbrev string) bool {
	if Fd.FdVerbose {
		logCommand	:= BlackOnGray(" tsNpm " + cmdArgs + " ")
		fmt.Printf("%s$ %s", logPrefix, logCommand)
		fmt.Printf("\n")
	} else {
		logCommand	:= "tsNpm " + cmdArgsAbbrev
		fmt.Printf("%s$ %s", logPrefix, logCommand)
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
		log.Printf("%s$  %s%s", logPrefix, command, WhiteOnRed(" X "))
		log.Fatalf("\n%s", Red(tsNpmError))
	}

	return true
}


func GetTypescriptError() error { return tsNpmError }
