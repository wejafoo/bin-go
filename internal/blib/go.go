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

var goError error

func NewGo() bool {
	if Fd.FdVerbose {
		fmt.Printf("\n%s", pad.Right("Compiling MIFE DEPLOYMENT configuration ", 113, "."))
		fmt.Printf("%s %s", LogWin, Blue(Fd.FdBuildContext))
	}
	DeploymentHead()
	PipelineHead()
	if Fd.FdBuild {
		goBuild()
	} else {
		SkipStep("goBuild():")
	}
	return DeploymentFoot(PipelineFoot(goDeploy()))
}

func goBuild() bool {
	logPrefix := Yellow(pad.Right("\ngoBuild():", 20, " "))
	args := "version"
	goRun(logPrefix, args, "")

	args = "build -v -o .dist"

	if Fd.FdRouteBase == "" {
		if Fd.FdService == "" {
			args += Fd.FdRepo
		} else {
			args += Fd.FdService
		}
	} else {
		args += Fd.FdRouteBase
	}
	argsAbbrev := args
	if Fd.FdTest {
		return goTest(goRun(logPrefix, args, argsAbbrev))
	} else {
		return goRun(logPrefix, args, argsAbbrev)
	}
}

func goTest(success bool) bool {
	if !success {
		return false
	}
	logPrefix := Yellow(pad.Right("\ngoTest():", 20, " "))
	args := "test -json ./..."
	argsAbbrev := args
	return goRun(logPrefix, args, argsAbbrev)
}

func goDeploy() bool {
	success := true
	if success = NewDocker(); !success {
		goError = GetComposeError()
	}
	return success
}

func goRun(logPrefix string, cmdArgs string, cmdArgsAbbrev string) bool {
	if Fd.FdVerbose {
		logCommand := BlackOnGray(" go " + cmdArgs + " ")
		fmt.Printf("%s$ %s", logPrefix, logCommand)
		fmt.Printf("\n")
	} else {
		logCommand := "go " + cmdArgsAbbrev
		fmt.Printf("%s$ %s", logPrefix, logCommand)
	}

	command := exec.Command("go", strings.Split(cmdArgs, " ")...)
	setEnvironment()
	command.Env = os.Environ()
	command.Stdout = os.Stderr
	stderr, _ := command.StderrPipe()
	goError = command.Start()
	if goError != nil {
		log.Printf("%s", Red(goError))
	}
	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		stderrText := scanner.Text()
		if Fd.FdVerbose {
			log.Printf("%s", Grey(stderrText))
		}
	}

	goError = command.Wait()
	if goError != nil {
		log.Printf("%s$  %s%s", logPrefix, command, WhiteOnRed(" X "))
		log.Fatalf("%s", Red(goError))
	}
	return true
}

func GetGoError() error { return goError }
