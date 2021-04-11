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
	success := true
	success = dockerBuild()
	success = dockerDeploy()
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
	args		:= " build --no-cache --pull " + Fd.FdServiceName
	success		:= composeRun(logPrefix, args)
	return success
	/*
		logMessage	:= BlackOnGray(" docker-compose " + args + " ")
		logVars		:= BlackOnGray(" TARGET_LOCAL_PORT=" + Fd.FdTargetLocalPort + " TARGET_PROJECT_ID=" + Fd.FdTargetProjectId + " SERVICE_NAME=" + Fd.FdServiceName + " TARGET_ALIAS=" + Fd.FdTargetAlias + " TARGET_IMAGE_TAG=" + Fd.FdTargetImageTag + " NICKNAME=" + Fd.FdNickname + " ")

		if !Fd.FdQuiet {
			fmt.Printf("%s$ %s", logPrefix, logVars)
			fmt.Printf("%s$ %s", logPrefix, logMessage)
		} else {
			logMessage := "Building Docker image"
			fmt.Printf("%s%s", logPrefix, logMessage)
		}

		if Fd.FdVerbose { fmt.Printf("\n") }
		cmd := exec.Command("docker-compose", strings.Split(args, " ")...)
		setEnvironment()
		cmd.Env = os.Environ()
		stderr, _	:= cmd.StderrPipe()
		e1			:= cmd.Start()
		if e1 != nil { log.Printf("%s", Red(e1)) }

		scanner := bufio.NewScanner(stderr)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			stderrText := scanner.Text()
			if Fd.FdVerbose { log.Printf("%s", Grey(stderrText)) }
		}

		cmdPinch := pad.Right(BlackOnGray(" docker-compose build (...) "+Fd.FdServiceName+" "), 69, ".")
		e2 := cmd.Wait()
		if e2 != nil {
			fmt.Printf("%s$ %s%s", logPrefix, cmdPinch, LogLose)
			fmt.Printf("\n")
			log.Fatal(Red(e2))
		} else {
			success = true
			fmt.Printf("%s$ %s%s", logPrefix, cmdPinch, LogWin)
		}

		return success
	*/
}


func composePush(success bool) bool {

	logPrefix	:= Yellow(pad.Right("\ncomposePush():", 20, " "))
	args		:= "push " + Fd.FdServiceName
	success		= composeRun(logPrefix, args)
	return success
	/*
		command		:= BlackOnGray("docker-compose " + args)
		fmt.Printf("%s$  %s", logPrefix, command)
		if Fd.FdVerbose { fmt.Printf("\n") }
		cmd			:= exec.Command("docker-compose", strings.Split(args, " ")...)
		setEnvironment()
		cmd.Env		= os.Environ()
		stderr, _	:= cmd.StderrPipe()
		e1			:= cmd.Start()
		if e1 != nil { log.Printf("%s", Red(e1)) }
		scanner		:= bufio.NewScanner(stderr)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			stderrText := scanner.Text()
			if Fd.FdVerbose { log.Printf("%s", Grey(stderrText)) }
		}
		e2			:= cmd.Wait()
		if e2 != nil {
			fmt.Printf("%s$  %s%s", logPrefix, command, LogWin)
			log.Fatal(Red(e2))
		}
		return success
	*/
}


func composePull() bool {

	logPrefix	:= Yellow(pad.Right("\ncomposePull():", 20, " "))
	args		:= "pull " + Fd.FdServiceName
	success		:= composeRun(logPrefix, args)
	return success
	/*
		command		:= BlackOnGray("docker-compose " + args)
		fmt.Printf("%s$  %s", logPrefix, command)
		if Fd.FdVerbose { fmt.Printf("\n") }
		cmd			:= exec.Command("docker-compose", strings.Split(args, " ")...)
		setEnvironment()
		cmd.Env		= os.Environ()
		stderr, _	:= cmd.StderrPipe()
		e1			:= cmd.Start()
		if e1 != nil { log.Printf("%s", Red(e1)) }
		scanner := bufio.NewScanner(stderr)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			stderrText := scanner.Text()
			if Fd.FdVerbose { log.Printf("%s", Grey(stderrText)) }
		}
		e2			:= cmd.Wait()
		if e2 != nil {
			fmt.Printf("%s$  %s%s", logPrefix, command, LogWin)
			log.Fatal(Red(e2))
		}
		return success
	*/
}


func composeUp(success bool) bool {

	logPrefix	:= Yellow(pad.Right("\ncomposeUp():", 20, " "))
	args		:= "--log-level " + Fd.FdTargetLogLevel + " up --detach --force-recreate " + Fd.FdServiceName
	success		= composeRun(logPrefix, args)
	return success
	/*	logMessage	:= BlackOnGray(" docker-compose " + args)

		if ! Fd.FdQuiet {
			fmt.Printf("%s$ %s", logPrefix, logMessage)
		} else {
			logMessage2 := "Spinning up local Docker instance"
			fmt.Printf("%s%s", logPrefix, logMessage2)
		}
		if Fd.FdVerbose { fmt.Printf("\n") }

		cmd			:= exec.Command("docker-compose", strings.Split(args, " ")...)

		setEnvironment()
		cmd.Env		= os.Environ()

		stderr, _	:= cmd.StderrPipe()
		e1			:= cmd.Start()
		if e1 != nil { log.Printf("%s", Red(e1)) }

		scanner		:= bufio.NewScanner(stderr)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			stderrText := scanner.Text()
			if Fd.FdVerbose { log.Printf("%s", Grey(stderrText)) }
		}

		e2			:= cmd.Wait()
		cmdPinch	:= pad.Right(BlackOnGray(" docker-compose up (...) "+Fd.FdServiceName+" "), 69, ".")
		if e2 != nil {
			fmt.Printf("%s$ %s%s", logPrefix, cmdPinch, LogLose)
			fmt.Printf("\n")
			log.Fatal(Red(e2))
		} else {
			success = true
			fmt.Printf("%s$ %s%s", logPrefix, cmdPinch, LogWin)
		}
	*/
}


func composeStop(success bool) bool {										// docker-compose stop "${service_name}" >> /dev/null 2>&1

	logPrefix	:= Yellow(pad.Right("\ncomposeStop():", 20, " "))
	args		:= "stop " + Fd.FdServiceName
	success		= composeRun(logPrefix, args)
	return success
	/*
		return success
		command		:= BlackOnGray("docker-compose " + args)
		fmt.Printf("%s$  %s", logPrefix, command)
		if Fd.FdVerbose { fmt.Printf("\n") }
		cmd			:= exec.Command("docker-compose", strings.Split(args, " ")...)
		setEnvironment()
		cmd.Env		= os.Environ()
		stderr, _	:= cmd.StderrPipe()
		e1			:= cmd.Start()
		if e1 != nil { log.Printf("%s", Red(e1)) }
		scanner		:= bufio.NewScanner(stderr)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			stderrText := scanner.Text()
			if Fd.FdVerbose { log.Printf("%s", Grey(stderrText)) }
		}
		e2			:= cmd.Wait()
		if e2 != nil {
			log.Printf("%s$  %s%s", logPrefix, command, WhiteOnRed(" X "))
			log.Fatal("\n%s", Red(e2))
		}
	*/
}


func composeRemove(success bool) bool { 								// docker-compose rm --force "${service_name}" >> /dev/null 2>&1

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
		log.Fatal("%s", Red(composeError))
	}
	success = true

	return success
}


func GetDockerError() error { return composeError }

func setEnvironment() {
	
	if err := os.Setenv("TARGET_LOCAL_PORT",	Fd.FdTargetLocalPort);	err != nil { return }
	if err := os.Setenv("TARGET_PROJECT_ID",	Fd.FdTargetProjectId);	err != nil { return }
	if err := os.Setenv("TARGET_ALIAS",		Fd.FdTargetAlias);		err != nil { return }
	if err := os.Setenv("SERVICE_NAME",		Fd.FdServiceName);		err != nil { return }
	if err := os.Setenv("TARGET_IMAGE_TAG",	Fd.FdTargetImageTag);	err != nil { return }
	if err := os.Setenv("NICKNAME",			Fd.FdNickname);			err != nil { return }
}


/*
	func testcontainerComposeUp() {
		composeFilePaths	:= []string {"./docker-compose.yml"}
		identifier			:= strings.ToLower(uuid.New().String())
		compose				:= testcontainers.NewLocalDockerCompose(composeFilePaths, identifier)
		execError1	:= compose.Down()
		e1			:= execError1.Error
		if e1 != nil { fmt.Printf("Could not run compose file: %v - %v", composeFilePaths, e1) }
		execError2			:= compose.WithCommand([]string{
			"--log-level", Fd.FdTargetLogLevel, "up", "--detach", "--force-recreate", Fd.FdServiceName,
		}).WithEnv(map[string]string {
			"TARGET_PROJECT_ID":	Fd.FdTargetProjectId,
			"SERVICE_NAME":			Fd.FdServiceName,
			"TARGET_ALIAS":			Fd.FdTargetAlias,
			"TARGET_IMAGE_TAG":		Fd.FdTargetImageTag,
			"NICKNAME":				Fd.FdNickname,
			"TARGET_LOCAL_PORT":	Fd.FdTargetLocalPort,
		}).Invoke()
		e2 := execError2.Error
		if e2 != nil { log.Fatalf("Could not run compose file: %v - %v", composeFilePaths, e2)}
	}

	func dockerLookup() {
		ctx				:= context.Background()
		clientPtr, e	:= client.NewEnvClient()
		if e != nil { panic( e )}
		fmt.Printf(" ctx: %T \n %v \n %v", clientPtr, clientPtr, *clientPtr)
		containers, e := clientPtr.ContainerList(ctx, types.ContainerListOptions{})
		if e != nil { panic( e )}
		fmt.Printf("\n\n")
		for _, container := range containers { fmt.Printf("\n Found container ID: %s  %T", container.ID, container.ID)}

	}
*/
