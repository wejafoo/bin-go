package blib

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/willf/pad"
)

var pythonError error

func NewPython() bool {
	if Fd.FdVerbose {
		fmt.Printf("\n%s", pad.Right("Compiling MIFE DEPLOYMENT configuration ", 113, "."))
		fmt.Printf("%s %s", LogWin, Blue(Fd.FdBuildContext))
	}

	DeploymentHead()
	PipelineHead()

	if Fd.FdBuild {
		pythonBuild()
	} else {
		SkipStep("pythonBuild():")
	}

	return DeploymentFoot(PipelineFoot(pythonDeploy()))
}

func pythonBuild() bool {
	logPrefix := Yellow(pad.Right("\npythonBuild():", 20, " "))
	args := "build" + Fd.FdTargetAlias
	argsAbbrev := args

	return pythonRun(logPrefix, args, argsAbbrev)
}

func pythonDeploy() bool {
	success := true
	if success = NewDocker(); !success {
		pythonError = GetComposeError()
	}
	return success
}

func pythonRun(logPrefix string, cmdArgs string, cmdArgsAbbrev string) bool {
	if Fd.FdVerbose {
		logCommand := BlackOnGray(" python " + cmdArgs + " ")
		fmt.Printf("%s$ %s", logPrefix, logCommand)
		fmt.Printf("\n")
	} else {
		logCommand := "python " + cmdArgsAbbrev
		fmt.Printf("%s$ %s", logPrefix, logCommand)
	}

	command := exec.Command("python", strings.Split(cmdArgs, " ")...)
	setEnvironment()
	command.Env = os.Environ()

	stderr, _ := command.StderrPipe()
	pythonError = command.Start()
	if pythonError != nil {
		log.Printf("%s", Red(pythonError))
	}
	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		stderrText := scanner.Text()
		if Fd.FdVerbose {
			log.Printf("%s", Grey(stderrText))
		}
	}

	pythonError = command.Wait()
	if pythonError != nil {
		log.Printf("%s$  %s%s", logPrefix, command, WhiteOnRed(" X "))
		log.Fatalf("\n%s", Red(pythonError))
	}

	return true
}

func GetPythonError() error { return pythonError }
