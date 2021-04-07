

package blib

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var		dockerBuildExit	int
var		dockerDeployExit	int


func	Docker() bool {
	
	ContextHead()
	dockerBuild()
	
	return ContextFoot(dockerDeploy())
}

func DockerBuildExit() int { return dockerBuildExit }

func DockerDeployExit() int { return dockerDeployExit }




func dockerBuild() bool {
	success				:=	true
	dockerBuildExit	=	0
	
	if Fd.FdLocal	{ BuildLocal()	}
	if Fd.FdRemote	{ BuildRemote()	}
	
	
	
	// Todo: Docker Build Stuffz
	
	
	
	
	return success
}





func	dockerDeploy() bool {
	success				:=	true
	dockerDeployExit	=	0
	
	if Fd.FdLocal	{ DeployLocal()		}
	if Fd.FdRemote	{ DeployRemote()	}
	
	
	
	
	// Todo: Docker Deploy Stuffz
	
	
	
	
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

