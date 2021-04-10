package blib

import (
	"fmt"
)

var typescriptBuildExit int

// var	typescriptDeployExit int

func NewTypescript() bool {

	fmt.Printf("%s %s", LogWin, Blue(Fd.FdBuildContext))

	ContextHead()
	typescriptBuild()
	return PipelineFoot(typescriptDeploy())
}

func typescriptBuild() bool {

	success := true
	typescriptBuildExit = 0

	if Fd.FdLocal {
		success = BuildLocal()
	} else if Fd.FdRemote {
		success = BuildRemote()
	}

	// Todo: Typescript Build Stuffz

	return success
}

func typescriptDeploy() bool {

	success := true
	// typescriptDeployExit = 0

	if Fd.FdLocal {
		success = DeployLocal()
	} else if Fd.FdRemote {
		success = DeployRemote()
	}

	// Todo: Typescript Deploy Stuffz
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

func TypescriptBuildExit() int { return typescriptBuildExit }

// func TypescriptDeployExit() int { return typescriptDeployExit }
