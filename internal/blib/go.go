

package blib

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var	goBuildExit		int
var	goDeployExit	int


func	Go() bool {
	
	ContextHead()
	goBuild()
	
	return ContextFoot(goDeploy())
}

func GoBuildExit() int { return goBuildExit }

func GoDeployExit() int { return goDeployExit }




func goBuild() bool {
	success				:=	true
	goBuildExit	=	0
	
	if Fd.FdLocal	{ BuildLocal()	}
	if Fd.FdRemote	{ BuildRemote()	}
	
	
	
	// Todo: Go Build Stuffz
	
	
	
	
	return success
}





func	goDeploy() bool {
	success				:=	true
	goDeployExit	=	0
	
	if Fd.FdLocal	{ DeployLocal()		}
	if Fd.FdRemote	{ DeployRemote()	}
	
	
	
	
	// Todo: Go Deploy Stuffz
	
	
	
	
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

