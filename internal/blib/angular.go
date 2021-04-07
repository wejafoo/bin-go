

package blib

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var		angularBuildExit	int
var		angularDeployExit	int


func	Angular() bool {
	
	ContextHead()
	angularBuild()
	
	return ContextFoot(angularDeploy())
}

func AngularBuildExit() int { return angularBuildExit }

func AngularDeployExit() int { return angularDeployExit }




func angularBuild() bool {
	success				:=	true
	angularBuildExit	=	0

	if Fd.FdLocal	{ BuildLocal()	}
	if Fd.FdRemote	{ BuildRemote()	}

	
	
	// Todo: Angular Build Stuffz
	
	
	
	
	return success
}





func	angularDeploy() bool {
	success				:=	true
	angularDeployExit	=	0
	
	if Fd.FdLocal	{ DeployLocal()		}
	if Fd.FdRemote	{ DeployRemote()	}
	
	
	
	
	// Todo: Angular Deploy Stuffz
	
	
	
	
	ctx				:= context.Background()
	clientPtr, e	:= client.NewEnvClient()
	if e != nil { panic( e )}
	fmt.Printf(" ctx: %T \n %v \n %v", clientPtr, clientPtr, *clientPtr)
	
	containers, e := clientPtr.ContainerList(ctx, types.ContainerListOptions{})
	if e != nil { panic( e )}
	
	fmt.Printf("\n\n")
	for _, container := range containers { fmt.Printf("\n Found container ID: %s  %T", container.ID, container.ID)}
	
	return success
}

