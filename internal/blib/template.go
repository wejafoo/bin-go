

package blib

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var		templateBuildExit	int
var		templateDeployExit	int


func	Template() bool {
	
	ContextHead()
	templateBuild()
	
	return ContextFoot(templateDeploy())
}

func TemplateBuildExit() int { return templateBuildExit }

func TemplateDeployExit() int { return templateDeployExit }




func templateBuild() bool {
	success				:=	true
	templateBuildExit	=	0
	
	if Fd.FdLocal	{ BuildLocal()	}
	if Fd.FdRemote	{ BuildRemote()	}
	
	
	
	// Todo: Template Build Stuffz
	
	
	
	
	return success
}





func	templateDeploy() bool {
	success				:=	true
	templateDeployExit	=	0
	
	if Fd.FdLocal	{ DeployLocal()		}
	if Fd.FdRemote	{ DeployRemote()	}
	
	
	
	
	// Todo: Template Deploy Stuffz
	
	
	
	
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
