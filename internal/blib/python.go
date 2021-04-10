package blib

import (
	"fmt"
)

var pythonBuildExit int

// var	pythonDeployExit int

func NewPython() bool {

	fmt.Printf("%s %s", LogWin, Blue(Fd.FdBuildContext))

	ContextHead()
	pythonBuild()
	return PipelineFoot(pythonDeploy())
}

func pythonBuild() bool {

	success := true
	pythonBuildExit = 0

	if Fd.FdLocal {
		success = BuildLocal()
	} else if Fd.FdRemote {
		success = BuildRemote()
	}

	// Todo: Python Build Stuffz

	return success
}

func pythonDeploy() bool {

	success := true
	// pythonDeployExit = 0

	if Fd.FdLocal {
		success = DeployLocal()
	} else if Fd.FdRemote {
		success = DeployRemote()
	}

	// Todo: Python Deploy Stuffz
	/*
		ctx				:= context.Background()
		clientPtr, e	:= client.NewEnvClient()
		if e != nil { panic( e )}
		fmt.Printf(" ctx: %T \n %v \n %v", clientPtr, clientPtr, *clientPtr)
		containers, e := clientPtr.ContainerList(ctx, types.ContainerListOptions{})
		if e != nil { panic( e )}
		fmt.Printf("\n\n")
		for _, container := range containers { fmt.Printf("\n Found container ID: %s  %T", container.ID, container.ID)}
	*/

	return success
}

func PythonBuildExit() int { return pythonBuildExit }

// func PythonDeployExit() int { return pythonDeployExit }
