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
	pythonError error
)


func NewPython() bool {

	fmt.Printf("%s %s", LogWin, Blue(Fd.FdBuildContext))
	DeploymentHead()
	PipelineHead()
	if Fd.FdBuild {pythonBuild()}   else   {SkipStep("pythonBuild():")}

	return DeploymentFoot(PipelineFoot(pythonDeploy()))
}


func pythonBuild() bool {

	logPrefix	:= Yellow(pad.Right("\npythonBuild():", 20, " "))
	args		:= "build" + Fd.FdTargetAlias
	success		:= pythonRun(logPrefix, args)
	
	return success
}


func pythonDeploy() bool {

	success := true
	if Fd.FdLocal {
		if success = NewDocker(); !success {
			success		= false
			pythonError	= GetDockerError()
		}
	}  else if Fd.FdRemote { success = NewGcp() }

	return success
}


func pythonRun(prefix string, cmdArgs string) bool {

	success 	:= false
	logCommand	:= BlackOnGray("python " + cmdArgs)

	fmt.Printf("%s$  %s", prefix, logCommand)
	if Fd.FdVerbose { fmt.Printf("\n") }

	command		:= exec.Command("python", strings.Split(cmdArgs, " ")...)
	setEnvironment()
	command.Env	= os.Environ()

	stderr, _	:= command.StderrPipe()
	pythonError	= command.Start()
	if pythonError != nil { log.Printf("%s", Red(pythonError)) }
	scanner		:= bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		stderrText := scanner.Text()
		if Fd.FdVerbose { log.Printf("%s", Grey(stderrText)) }
	}

	pythonError = command.Wait()
	if pythonError != nil {
		log.Printf("%s$  %s%s", prefix, command, WhiteOnRed(" X "))
		log.Fatal("%s", Red(pythonError))
	}
	success = true

	return success
}


func GetPythonError() error { return pythonError }
