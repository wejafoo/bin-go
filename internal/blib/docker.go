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

var dockerBuildExit int

// var	dockerDeployExit int

func NewDocker() bool {

	success := true

	success = dockerBuild()
	success = dockerDeploy()

	return success
}

func DockerBuildExit() int { return dockerBuildExit }

// func DockerDeployExit() int { return dockerDeployExit }

func dockerBuild() bool {

	success := true
	dockerBuildExit = 0

	success = composeBuild()

	return success
}

func dockerDeploy() bool {

	success := true

	// dockerDeployExit = 0
	// composeStop()

	if Fd.FdLocal {
		success = composeUp()
	}

	return success
}

func composeBuild() bool {

	success := false
	logPrefix := Yellow(pad.Right("\ncomposeBuild():", 20, " "))
	args := "--log-level " + Fd.FdTargetLogLevel + " build --no-cache --progress auto --pull " + Fd.FdServiceName
	logMessage := BlackOnGray(" docker-compose " + args + " ")
	logVars := BlackOnGray(" TARGET_LOCAL_PORT=" + Fd.FdTargetLocalPort + " TARGET_PROJECT_ID=" + Fd.FdTargetProjectId + " SERVICE_NAME=" + Fd.FdServiceName + " TARGET_ALIAS=" + Fd.FdTargetAlias + " TARGET_IMAGE_TAG=" + Fd.FdTargetImageTag + " NICKNAME=" + Fd.FdNickname + " ")

	if !Fd.FdQuiet {
		fmt.Printf("%s$ %s", logPrefix, logVars)
		fmt.Printf("%s$ %s", logPrefix, logMessage)
	} else {
		logMessage := "Building Docker image"
		fmt.Printf("%s$ %s", logPrefix, logMessage)
	}

	if Fd.FdVerbose {
		fmt.Printf("\n")
	}

	cmd := exec.Command("docker-compose", strings.Split(args, " ")...)
	setEnvironment()
	cmd.Env = os.Environ()

	stderr, _ := cmd.StderrPipe()
	e1 := cmd.Start()
	if e1 != nil {
		log.Printf("%s", Red(e1))
	}

	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		if Fd.FdVerbose {
			log.Printf("%s", Grey(m))
		}
	}

	commandSmall := pad.Right(BlackOnGray(" docker-compose build (...) "+Fd.FdServiceName+" "), 69, ".")

	e2 := cmd.Wait()
	if e2 != nil {
		fmt.Printf("%s$ %s%s", logPrefix, commandSmall, LogLose)
		fmt.Printf("\n")
		log.Fatal(Red(e2))
	} else {
		success = true
		fmt.Printf("%s$ %s%s", logPrefix, commandSmall, LogWin)
	}

	return success
}

func composeUp() bool {

	success := false
	logPrefix := Yellow(pad.Right("\ncomposeUp():", 20, " "))
	args := "--log-level " + Fd.FdTargetLogLevel + " up --detach --force-recreate " + Fd.FdServiceName
	logMessage := BlackOnGray(" docker-compose " + args)

	if !Fd.FdQuiet {
		fmt.Printf("%s$ %s", logPrefix, logMessage)
	} else {
		logMessage2 := "Spinning up local Docker instance"
		fmt.Printf("%s%s", logPrefix, logMessage2)
	}

	if Fd.FdVerbose {
		fmt.Printf("\n")
	}

	cmd := exec.Command("docker-compose", strings.Split(args, " ")...)
	setEnvironment()
	cmd.Env = os.Environ()

	stderr, _ := cmd.StderrPipe()
	e1 := cmd.Start()

	if e1 != nil {
		log.Printf("%s", Red(e1))
	}

	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		if Fd.FdVerbose {
			log.Printf("%s", Grey(m))
		}
	}

	e2 := cmd.Wait()

	commandSmall := pad.Right(BlackOnGray(" docker-compose up (...) "+Fd.FdServiceName+" "), 69, ".")
	if e2 != nil {
		fmt.Printf("%s$ %s%s", logPrefix, commandSmall, LogLose)
		fmt.Printf("\n")
		log.Fatal(Red(e2))
	} else {
		success = true
		fmt.Printf("%s$ %s%s", logPrefix, commandSmall, LogWin)
	}

	return success
}

func composePull() { // docker-compose pull "${service_name}"  >>/dev/null 2>&1

	logPrefix := Yellow(pad.Right("\ncomposePull():", 20, " "))
	args := "pull " + Fd.FdServiceName
	command := BlackOnGray("docker-compose " + args)

	fmt.Printf("%s$  %s", logPrefix, command)
	if Fd.FdVerbose {
		fmt.Printf("\n")
	}

	cmd := exec.Command("docker-compose", strings.Split(args, " ")...)
	setEnvironment()
	cmd.Env = os.Environ()

	stderr, _ := cmd.StderrPipe()
	e1 := cmd.Start()

	if e1 != nil {
		log.Printf("%s", Red(e1))
	}

	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		if Fd.FdVerbose {
			log.Printf("%s", Grey(m))
		}
	}

	e2 := cmd.Wait()
	if e2 != nil {
		fmt.Printf("%s$  %s%s", logPrefix, command, LogWin)
		log.Fatal(Red(e2))
	}
}

func composeStop() { // docker-compose stop "${service_name}" >> /dev/null 2>&1  && docker-compose rm --force "${service_name}" >> /dev/null 2>&1

	logPrefix := Yellow(pad.Right("\ncomposeStop():", 20, " "))
	args := "stop " + Fd.FdServiceName
	command := BlackOnGray("docker-compose " + args)

	fmt.Printf("%s$  %s", logPrefix, command)
	if Fd.FdVerbose {
		fmt.Printf("\n")
	}

	cmd := exec.Command("docker-compose", strings.Split(args, " ")...)
	setEnvironment()
	cmd.Env = os.Environ()

	stderr, _ := cmd.StderrPipe()
	e1 := cmd.Start()

	if e1 != nil {
		log.Printf("%s", Red(e1))
	}

	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		if Fd.FdVerbose {
			log.Printf("%s", Grey(m))
		}
	}

	e2 := cmd.Wait()
	if e2 != nil {
		log.Printf("%s$  %s%s", logPrefix, command, WhiteOnRed(" X "))
		log.Fatal("\n%s", Red(e2))
	}
}

func setEnvironment() {
	err := os.Setenv("TARGET_LOCAL_PORT", Fd.FdTargetLocalPort)
	if err != nil {
		return
	}
	err = os.Setenv("TARGET_PROJECT_ID", Fd.FdTargetProjectId)
	if err != nil {
		return
	}
	err = os.Setenv("SERVICE_NAME", Fd.FdServiceName)
	if err != nil {
		return
	}
	err = os.Setenv("TARGET_ALIAS", Fd.FdTargetAlias)
	if err != nil {
		return
	}
	err = os.Setenv("TARGET_IMAGE_TAG", Fd.FdTargetImageTag)
	if err != nil {
		return
	}
	err = os.Setenv("NICKNAME", Fd.FdNickname)
	if err != nil {
		return
	}
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
