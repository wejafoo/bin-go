package templates

/*
	var	templateBuildExit	int
	var	templateDeployExit	int


	func NewTemplate() bool {
		fmt.Printf("%s %s", LogWin, Blue(Fd.FdBuildContext))
		DeploymentHead()
		PipelineHead()
		if Fd.FdBuild {templateBuild()}   else   {SkipStep("templateBuild():")}
		return DeploymentFoot(PipelineFoot(templateDeploy()))
	}


	func TemplateBuildExit() int { return templateBuildExit }
	func TemplateDeployExit() int { return templateDeployExit }


func templateBuild() bool {

	success := true
	templateBuildExit = 0

	if Fd.FdLocal {
		success = BuildLocal()
	} else if Fd.FdRemote {
		success = BuildRemote()
	}

	// Todo: Template Build Stuffz

	return success
}


func templateDeploy() bool {

	success := true
	templateDeployExit = 0

	if Fd.FdLocal {
		success = DeployLocal()
	} else if Fd.FdRemote {
		success = DeployRemote()
	}

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
*/
