

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
	ngNpmError error
)


func NewAngular() bool {

	fmt.Printf("%s %s", LogWin, Blue(Fd.FdBuildContext))
	DeploymentHead()
	PipelineHead()
	if Fd.FdBuild {angularBuild()}   else   {SkipStep("angularBuild():")}

	return DeploymentFoot(PipelineFoot(angularDeploy()))
}


func angularBuild() bool {

	logPrefix	:= Yellow(pad.Right("\nangularBuild():", 20, " "))
	args		:= "run build:ngssc:" + Fd.FdTargetAlias
	success		:= ngNpmRun(logPrefix, args)
	
	return success
	/*
		logMessage			:= BlackOnGray(" ngNpm " + args + " ")
		success				:= false
		if !Fd.FdQuiet {
			fmt.Printf("%s$ %s", logPrefix, logMessage)
		} else {
			logMessage2 := "Building Angular distribution"
			fmt.Printf("%s%s", logPrefix, logMessage2)
		}
		if Fd.FdVerbose { fmt.Printf("\n") }
		cmd			:= exec.Command("ngNpm", strings.Split(args, " ")...)
		stderr, _	:= cmd.StderrPipe()
		e1			:= cmd.Start()
		if e1 != nil { log.Printf("%s", Red(e1)) }
		scanner := bufio.NewScanner(stderr)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			stderrText := scanner.Text()
			if Fd.FdVerbose { log.Printf("%s", Grey(stderrText)) }
		}
		e2 := cmd.Wait()
		if e2 != nil {
			fmt.Printf("%s$ %s%s", logPrefix, pad.Right(logMessage, 69, "."), LogLose)
			fmt.Printf("\n")
			log.Fatalf("%s", Red(e2))
		} else {
			success = true
			fmt.Printf("%s$ %s%s", logPrefix, pad.Right(logMessage, 69, "."), LogWin)
		}
		return success
	*/
}


func angularDeploy() bool {

	success := true
	if Fd.FdLocal {
		if success = NewDocker(); !success {
			success		= false
			ngNpmError	= GetDockerError()
		}
	}  else if Fd.FdRemote {
		if success = NewGcp(); !success {
			success		= false
			ngNpmError	= GetGcpError()
		}
	}
	// Todo: Incorporate GoLang native Docker interface in lieu of clunky shell implementation

	return success
}


func ngNpmRun(prefix string, cmdArgs string) bool {

	success 	:= false
	logCommand	:= BlackOnGray("npm " + cmdArgs)

	fmt.Printf("%s$  %s", prefix, logCommand)
	if Fd.FdVerbose { fmt.Printf("\n") }

	command		:= exec.Command("npm", strings.Split(cmdArgs, " ")...)
	setEnvironment()
	command.Env	= os.Environ()

	stderr, _	:= command.StderrPipe()
	ngNpmError	= command.Start()
	if ngNpmError != nil { log.Printf("%s", Red(ngNpmError)) }
	scanner		:= bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		stderrText := scanner.Text()
		if Fd.FdVerbose { log.Printf("%s", Grey(stderrText)) }
	}

	ngNpmError = command.Wait()
	if ngNpmError != nil {
		log.Printf("%s$  %s%s", prefix, command, WhiteOnRed(" X "))
		log.Fatal("%s", Red(ngNpmError))
	}
	success = true

	return success
}


func GetAngularError() error { return ngNpmError }
