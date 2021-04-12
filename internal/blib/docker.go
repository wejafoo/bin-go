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
	composeError error
)


func NewDocker() bool {
	success := dockerBuild()

	if Fd.FdLocal && success { success = dockerDeploy()}

	return success
}


func dockerBuild() bool {
	success := composePush(composeBuild())

	return success
}


func dockerDeploy() bool {
	success := true
	if Fd.FdLocal { success = composeUp(composeRemove(composeStop(composePull()))) } else { fmt.Printf("UNDER CONSTRUCTION:  Pure Docker remote deployment ")}

	return success
}


func composeBuild() bool {

	logPrefix	:= Yellow(pad.Right("\ncomposeBuild():", 20, " "))
	// args		:= "--log-level " + Fd.FdTargetLogLevel + " build --no-cache --progress auto --pull " + Fd.FdServiceName
	args		:= "build --no-cache --pull " + Fd.FdServiceName
	success		:= composeRun(logPrefix, args)
	return success
}


func composePush(success bool) bool {

	logPrefix	:= Yellow(pad.Right("\ncomposePush():", 20, " "))
	args		:= "push " + Fd.FdServiceName
	success		= composeRun(logPrefix, args)
	return success
}


func composePull() bool {

	logPrefix	:= Yellow(pad.Right("\ncomposePull():", 20, " "))
	args		:= "pull " + Fd.FdServiceName
	success		:= composeRun(logPrefix, args)
	return success
}


func composeUp(success bool) bool {

	logPrefix	:= Yellow(pad.Right("\ncomposeUp():", 20, " "))
	args		:= "--log-level " + Fd.FdTargetLogLevel + " up --detach --force-recreate " + Fd.FdServiceName
	success		= composeRun(logPrefix, args)
	return success
}


func composeStop(success bool) bool {

	logPrefix	:= Yellow(pad.Right("\ncomposeStop():", 20, " "))
	args		:= "stop " + Fd.FdServiceName
	success		= composeRun(logPrefix, args)
	return success
}


func composeRemove(success bool) bool {

	logPrefix	:= Yellow(pad.Right("\ncomposeRemove():", 20, " "))
	args		:= "rm --force " + Fd.FdServiceName
	success		= composeRun(logPrefix, args)
	return success
}


func composeRun(prefix string, cmdArgs string) bool {

	success 	:= false
	logCommand	:= BlackOnGray("docker-compose " + cmdArgs)

	fmt.Printf("%s$  %s", prefix, logCommand)
	if Fd.FdVerbose { fmt.Printf("\n") }

	command		:= exec.Command("docker-compose", strings.Split(cmdArgs, " ")...)
	setEnvironment()
	command.Env	= os.Environ()

	stderr, _		:= command.StderrPipe()
	composeError	= command.Start()
	if composeError != nil { log.Printf("%s", Red(composeError)) }

	scanner			:= bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		stderrText := scanner.Text()
		if Fd.FdVerbose { log.Printf("%s", Grey(stderrText)) }
	}

	composeError = command.Wait()
	if composeError != nil {
		log.Printf("%s$  %s%s", prefix, command, WhiteOnRed(" X "))
		log.Fatalf("\n%s", Red(composeError))
	}
	success = true

	return success
}


func GetDockerError() error { return composeError }


func setEnvironment() {
	if err := os.Setenv("TARGET_LOCAL_PORT",	Fd.FdTargetLocalPort	); err != nil { return }
	if err := os.Setenv("TARGET_PROJECT_ID",	Fd.FdTargetProjectId	); err != nil { return }
	if err := os.Setenv("TARGET_ALIAS",		Fd.FdTargetAlias		); err != nil { return }
	if err := os.Setenv("SERVICE_NAME",		Fd.FdServiceName		); err != nil { return }
	if err := os.Setenv("TARGET_IMAGE_TAG",	Fd.FdTargetImageTag		); err != nil { return }
	if err := os.Setenv("NICKNAME",			Fd.FdNickname			); err != nil { return }
	if err := os.Setenv("SITE_NICKNAME",		Fd.FdSiteNickname		); err != nil { return }
}
