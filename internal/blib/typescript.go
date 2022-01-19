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

var tsNpmError error

func NewTypescript() bool {
	if Fd.FdVerbose {
		fmt.Printf("\n%s", pad.Right("Compiling MIFE DEPLOYMENT configuration ", 113, "."))
		fmt.Printf("%s %s", LogWin, Blue(Fd.FdBuildContext))
	}
	DeploymentHead()
	PipelineHead()
	if Fd.FdBuild {
		typescriptBuild()
	} else {
		SkipStep("typescriptBuild():")
	}
	return DeploymentFoot(PipelineFoot(typescriptDeploy()))
}

func typescriptBuild() bool {
	logPrefix := Yellow(pad.Right("\ntypescriptBuild():", 20, " "))
	args := "run build:" + Fd.FdTargetAlias
	argsAbbrev := args
	return tsNpmRun(logPrefix, args, argsAbbrev)
}

// func typescriptTest() bool {
// 	logPrefix := Yellow(pad.Right("\ntypescriptTest():", 20, " "))
// 	args := "run test:" + Fd.FdTargetAlias
// 	argsAbbrev := args
// 
// 	return ngNpmRun(logPrefix, args, argsAbbrev)
// }

func typescriptDeploy() bool {
	success := NewDocker()
	if !success {
		tsNpmError = GetComposeError()
	}
	return success
}

func tsNpmRun(logPrefix string, cmdArgs string, cmdArgsAbbrev string) bool {
	if Fd.FdVerbose {
		logCommand := BlackOnGray(" npm " + cmdArgs)
		fmt.Printf("%s$ %s", logPrefix, logCommand)
		fmt.Printf("\n")
	} else {
		logCommand := "npm " + cmdArgsAbbrev
		fmt.Printf("%s$ %s", logPrefix, logCommand)
	}

	command := exec.Command("npm", strings.Split(cmdArgs, " ")...)
	setEnvironment()
	command.Env = os.Environ()

	stderr, _ := command.StderrPipe()
	tsNpmError = command.Start()
	if tsNpmError != nil {
		log.Printf("%s", Red(tsNpmError))
	}

	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		stderrText := scanner.Text()
		if Fd.FdVerbose {
			log.Printf("%s", Grey(stderrText))
		}
	}

	tsNpmError = command.Wait()
	if tsNpmError != nil {
		log.Printf("%s$  %s%s", logPrefix, command, WhiteOnRed(" X "))
		log.Fatalf("\n%s", Red(tsNpmError))
	}
	return true
}

func GetTypescriptError() error { return tsNpmError }
