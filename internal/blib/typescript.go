

package blib

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var		typescriptBuildExit	int
var		typescriptDeployExit	int


func	Typescript() bool {
	
	ContextHead()
	typescriptBuild()
	
	return ContextFoot(typescriptDeploy())
}

func TypescriptBuildExit() int { return typescriptBuildExit }

func TypescriptDeployExit() int { return typescriptDeployExit }




func typescriptBuild() bool {
	success				:=	true
	typescriptBuildExit	=	0
	
	if Fd.FdLocal	{ BuildLocal()	}
	if Fd.FdRemote	{ BuildRemote()	}
	
	
	
	// Todo: Typescript Build Stuffz
	
	
	
	
	return success
}





func	typescriptDeploy() bool {
	success				:=	true
	typescriptDeployExit	=	0
	
	if Fd.FdLocal	{ DeployLocal()		}
	if Fd.FdRemote	{ DeployRemote()	}
	
	
	
	
	// Todo: Typescript Deploy Stuffz
	
	
	
	
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

