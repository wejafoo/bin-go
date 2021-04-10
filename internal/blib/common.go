package blib

import (
	"fmt"
	"github.com/willf/pad"
	"os"
)

func BuildLocal() bool {

	// success 	:= true
	logPrefix := Yellow(pad.Right("\nBuildLocal():", 20, " "))

	action1 := pad.Right("Check for preexisting port bindings ", 58, ".")
	action2 := pad.Right("Turn on listener ", 58, ".")

	fmt.Printf("%s%s%s %s", logPrefix, action1, LogWin, Blue(Fd.FdTargetLocalPort))
	fmt.Printf("%s%s%s %s", logPrefix, action2, LogWin, Blue("STDERR"))

	return true
}

func BuildRemote() bool {

	// success		:= true
	logPrefix := Yellow(pad.Right("\nBuildRemote():", 20, " "))

	fmt.Printf("%s\t\t... DOING COMMON REMOTE BUILD STUFFZ  ...", logPrefix)

	return true
}

func DeployLocal() bool {

	// success		:= true
	logPrefix := Yellow(pad.Right("\nDeployLocal():", 20, " "))

	if Fd.FdVerbose {
		fmt.Printf(logPrefix)
	}
	if Fd.Success = NewDocker(); !Fd.Success {
		os.Exit(DockerBuildExit())
	}

	return true
}

func DeployRemote() bool {

	// success 	:= true
	logPrefix := Yellow(pad.Right("\nDeployRemote():", 20, " "))

	if Fd.FdTargetRemotePort != "8080" {
		fmt.Printf("%s%s", logPrefix, "Remote deployments require the remote port to be '8080'. Consider using '-force' to proceed beyond this message.")
		os.Exit(3)
	} else {
		println("\n\n\t\t... DOING COMMON REMOTE DEPLOY STUFFZ ...")
	}

	return true
}
