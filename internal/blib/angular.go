package blib

import (
	"bufio"
	"fmt"
	"github.com/willf/pad"
	"log"
	"os/exec"
	"strings"
)

var (
	angularBuildExit int
	// angularDeployExit	int
)

func NewAngular() bool {
	fmt.Printf("%s %s", LogWin, Blue(Fd.FdBuildContext))
	ContextHead()
	if Fd.FdLocal && Fd.FdBuild {
		BuildLocal()
		angularBuild()
	} else if Fd.FdRemote && Fd.FdBuild {
		BuildRemote()
		angularBuild()
	} else {
		SkipStep("angularBuild():")
	}

	return DeploymentFoot(PipelineFoot(angularDeploy()))
}

func angularBuild() bool {

	logPrefix := Yellow(pad.Right("\nangularBuild():", 20, " "))
	args := "run build:ngssc:" + Fd.FdTargetAlias
	logMessage := BlackOnGray(" npm " + args + " ")
	success := false
	angularBuildExit = 0

	if !Fd.FdQuiet {
		fmt.Printf("%s$ %s", logPrefix, logMessage)
	} else {
		logMessage2 := "Building Angular distribution"
		fmt.Printf("%s%s", logPrefix, logMessage2)
	}

	if Fd.FdVerbose {
		fmt.Printf("\n")
	}

	cmd := exec.Command("npm", strings.Split(args, " ")...)
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
		fmt.Printf("%s$ %s%s", logPrefix, pad.Right(logMessage, 69, "."), LogLose)
		fmt.Printf("\n")
		log.Fatalf("%s", Red(e2))
	} else {
		success = true
		fmt.Printf("%s$ %s%s", logPrefix, pad.Right(logMessage, 69, "."), LogWin)
	}

	return success
}

func angularDeploy() bool {

	success := true
	// angularDeployExit = 0

	if Fd.FdLocal {
		success = DeployLocal()
	} else if Fd.FdRemote {
		success = DeployRemote()
	}

	// Todo: Incorporate GoLang native Docker interface in lieu of clunky shell implementation
	/*
		ctx				:= context.Background()
		clientPtr, e	:= client.NewEnvClient()
		if e != nil { panic( e )}
		fmt.Printf(" ctx: %T \n %v \n %v", clientPtr, clientPtr, *clientPtr)
		containers, e := clientPtr.ContainerList(ctx, types.ContainerListOptions{})
		if e != nil { panic( e )}
		for _, container := range containers { fmt.Printf("\n Found container ID: %s  %T", container.ID, container.ID)}
	*/

	return success
}

func AngularBuildExit() int { return angularBuildExit }

// func AngularDeployExit() int { return angularDeployExit	}
