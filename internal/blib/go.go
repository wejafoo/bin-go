package blib

import (
	"fmt"
)

var goBuildExit int

// var goDeployExit int

func NewGo() bool {

	fmt.Printf("%s %s", LogWin, Blue(Fd.FdBuildContext))

	ContextHead()
	goBuild()

	return PipelineFoot(goDeploy())
}

func GoBuildExit() int { return goBuildExit }

// func GoDeployExit() int { return goDeployExit }

func goBuild() bool {
	success := true
	goBuildExit = 0
	if Fd.FdLocal {
		success = BuildLocal()
	} else if Fd.FdRemote {
		success = BuildRemote()
	}

	// Todo: Go Build Stuffz

	return success
}

func goDeploy() bool {

	success := true
	// goDeployExit = 0

	if Fd.FdLocal {
		success = DeployLocal()
	} else if Fd.FdRemote {
		success = DeployRemote()
	}

	// Todo: Go Deploy Stuffz

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
