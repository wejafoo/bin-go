

package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gookit/color"
	"github.com/willf/pad"
	"os"
	"strings"
)

func main() {
	
	fmt.Printf("\n==  DOCKER CLEAN START												")
	green	:= color.FgGreen.Render
	envBase	:= "AIO"

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if strings.Contains(pair[0], envBase) {
			fmt.Printf("\n %s %s ", pad.Right(pair[0]+" ", 20, "."), green(pair[1]))
		}
	}

	ctx				:= context.Background()
	clientPtr, err	:= client.NewEnvClient()
	
	if err != nil { panic(err) }

	color.Style{color.Yellow, color.OpItalic}.Printf("\n\n Hello, logger!")
	fmt.Printf("\n\n ctx: %T \n %v \n %v", clientPtr, clientPtr, *clientPtr)
	color.Style{color.Yellow, color.OpBold}.Printf("\n\n Goodbye, logger!")

	containers, err := clientPtr.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil { panic(err) }

	fmt.Printf("\n")
	for _, container := range containers {
		fmt.Printf("\n Found container ID: %s  %T", container.ID, container.ID)
	}
	fmt.Printf("\n==  DOCKER CLEAN END													")
	fmt.Println()
}

// docker container	prune	--force
// docker image		prune 	--force	--all
// docker network	prune	--force
// docker volume	prune	--force

// NewClientWithOpts(client.WithVersion("1.37"))
