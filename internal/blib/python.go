

package blib

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var	pythonBuildExit		int
var	pythonDeployExit	int


func Python() bool {
	
	ContextHead()
	pythonBuild()
	
	return ContextFoot(pythonDeploy())
}

func PythonBuildExit() int { return pythonBuildExit }

func PythonDeployExit() int { return pythonDeployExit }




func pythonBuild() bool {
	success				:=	true
	pythonBuildExit	=	0
	
	if Fd.FdLocal	{ BuildLocal()	}
	if Fd.FdRemote	{ BuildRemote()	}
	
	
	
	// Todo: Python Build Stuffz
	
	
	
	
	return success
}





func	pythonDeploy() bool {
	success				:=	true
	pythonDeployExit	=	0
	
	if Fd.FdLocal	{ DeployLocal()		}
	if Fd.FdRemote	{ DeployRemote()	}
	
	
	
	
	// Todo: Python Deploy Stuffz
	
	
	
	
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

