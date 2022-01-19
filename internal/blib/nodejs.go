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

var njsNpmError error

func NewNodejs() bool {
	if Fd.FdVerbose {
		fmt.Printf("\n%s", pad.Right("Compiling MIFE DEPLOYMENT configuration ", 113, "."))
		fmt.Printf("%s %s", LogWin, Blue(Fd.FdBuildContext))
	}

	DeploymentHead()
	PipelineHead()

	if Fd.FdBuild {
		nodejsBuild()
	} else {
		SkipStep("nodejsBuild():")
	}

	return DeploymentFoot(PipelineFoot(nodejsDeploy()))
}

func nodejsBuild() bool {
	logPrefix := Yellow(pad.Right("\nnodejsBuild():", 20, " "))
	args := "run build:" + Fd.FdTargetAlias
	argsAbbrev := args

	return njsNpmRun(logPrefix, args, argsAbbrev)
}

// func nodejsTest() bool {
// 	logPrefix := Yellow(pad.Right("\nnodejsTest():", 20, " "))
// 	args := "run test:" + Fd.FdTargetAlias
// 	argsAbbrev := args
// 
// 	return ngNpmRun(logPrefix, args, argsAbbrev)
// }

func nodejsDeploy() bool {
	success := NewDocker()
	if !success {
		njsNpmError = GetComposeError()
	}

	return success
}

func njsNpmRun(logPrefix string, cmdArgs string, cmdArgsAbbrev string) bool {
	if Fd.FdVerbose {
		logCommand := BlackOnGray(" npm " + cmdArgs)
		fmt.Printf("%s$ %s", logPrefix, logCommand)
		fmt.Printf("\n")
	} else {
		logCommand := "npm " + cmdArgsAbbrev
		fmt.Printf("%s$ %s", logPrefix, logCommand)
		// fmt.Printf("\n")
	}

	command := exec.Command("npm", strings.Split(cmdArgs, " ")...)
	setEnvironment()
	command.Env = os.Environ()

	stderr, _ := command.StderrPipe()
	njsNpmError = command.Start()
	if njsNpmError != nil {
		log.Printf("%s", Red(njsNpmError))
	}

	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		stderrText := scanner.Text()
		if Fd.FdVerbose {
			log.Printf("%s", Grey(stderrText))
		}
	}

	njsNpmError = command.Wait()
	if njsNpmError != nil {
		log.Printf("%s$  %s%s", logPrefix, command, WhiteOnRed(" X "))
		log.Fatalf("\n%s", Red(njsNpmError))
	}

	return true
}

func GetNodejsError() error { return njsNpmError }
