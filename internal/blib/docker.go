

package blib

import (
	"bufio"
	"fmt"
	"github.com/willf/pad"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var (
	composeError	error
	dockerError		error
)


func NewDocker() bool {
	if Fd.FdBuildContext == "docker" {
		DeploymentHead()
		PipelineHead()
	}
	return dockerDeploy(dockerBuild())
}


func dockerBuild() bool {
	success := true
	if Fd.FdClean { dockerClean() }
	if Fd.FdLocal { success = composeBuild() } else { success = composePush(composeBuild())}
	return success
}


func dockerDeploy(prevSuccess bool) bool {
	success := prevSuccess
	if prevSuccess {
		if Fd.FdLocal {
			success = composeUp(composeRemove(composeStop()))
		} else {
			if success = NewGcp();	!success { dockerError = GetGcpError() }
		}
	}
	return success
}


func composeBuild() bool {
	logPrefix	:= Yellow(pad.Right("\ncomposeBuild():", 20, " "))
	// args		:=	"--verbose build --no-cache --pull"	+ " "  Todo: REMOVE THIS COMMENT ON NEXT COMMIT -- UPGRADE TO DOCKER COMPOSE V2
	args		:=	"build --no-cache --pull"	+ " "						// Add new build args here
	args		+=	"--build-arg REPO" 					+ " "
	args		+=	"--build-arg ROUTE_BASE" 			+ " "
	args		+=	"--build-arg TARGET_ALIAS"			+ " "
	args		+=	"--build-arg EXECUTABLE"			+ " "
	argsAbbrev	:= "build (...)"						+ " "

	if Fd.FdService == "" {
		args		+= Fd.FdRepo
		argsAbbrev	+= Fd.FdRepo
	} else {
		args		+= "--build-arg SERVICE" + " "
		args		+= Fd.FdService
		argsAbbrev	+= Fd.FdService
	}

	return composeRun(logPrefix, args, argsAbbrev)
}


func composePush(prevSuccess bool) bool {
	if !prevSuccess{ return false }

	logPrefix	:= Yellow(pad.Right("\ncomposePush():", 20, " "))
	args		:= "push" + " "
	argsAbbrev	:= args

	if Fd.FdService == "" { args += Fd.FdRepo } else { args += Fd.FdService }

	return composeRun(logPrefix, args, argsAbbrev)
}


func composePull() bool {
	logPrefix	:= Yellow(pad.Right("\ncomposePull():", 20, " "))
	args		:= "pull" + " "
	argsAbbrev	:= args

	if Fd.FdService == "" { args += Fd.FdRepo } else { args += Fd.FdService }

	return composeRun(logPrefix, args, argsAbbrev)
}


func composeUp(prevSuccess bool) bool {
	if !prevSuccess { return false }

	logPrefix	:= Yellow(pad.Right("\ncomposeUp():", 20, " "))
//	args		:= "--log-level " + Fd.FdTargetLogLevel + " up --detach --force-recreate "
	args		:= "up --detach --force-recreate "
	argsAbbrev	:= "(...) up (...) "

	if Fd.FdService == "" {
		args		+= Fd.FdRepo
		argsAbbrev	+= Fd.FdRepo
	} else {
		args		+= Fd.FdService
		argsAbbrev	+= Fd.FdService
	}

	return composeRun(logPrefix, args, argsAbbrev)
}


// func composeStop(prevSuccess bool) bool {
// 	if !prevSuccess { return false }
func composeStop() bool {
	logPrefix	:= Yellow(pad.Right("\ncomposeStop():", 20, " "))
	args		:= "stop" + " "
	if Fd.FdService == "" { args += Fd.FdRepo } else { args += Fd.FdService }
	argsAbbrev	:= args

	return composeRun(logPrefix, args, argsAbbrev)
}


func composeRemove(prevSuccess bool) bool {
	if !prevSuccess { return false }

	logPrefix	:= Yellow(pad.Right("\ncomposeRemove():", 20, " "))
	args		:= "rm --force" + " "

	if Fd.FdService == "" {
		args += Fd.FdRepo
	} else {
		args += Fd.FdService
	}

	argsAbbrev	:= args

	return composeRun(logPrefix, args, argsAbbrev)
}


func composeRun(prefix string, cmdArgs string, cmdArgsAbbrev string) bool {
	if Fd.FdVerbose {
		logCommand	:= BlackOnGray("docker-compose " + cmdArgs + " ")

		fmt.Printf("%s$ %s", prefix, logCommand)
		fmt.Printf("\n")
	} else {
		logCommand	:= "docker-compose " + cmdArgsAbbrev
		fmt.Printf("%s$ %s", prefix, logCommand)
	}

	command		:= exec.Command("docker-compose", strings.Split(cmdArgs, " ")...)
	if success	:= setEnvironment(); !success { log.Println("Curious issue with setting the environment :'(")}
	command.Env	 = os.Environ()
	if Fd.FdDebug {
		if Fd.FdService == "" { log.Printf("UNDEFINED... using REPO\n") } else { log.Printf("SERVICE=%s", os.Getenv("SERVICE")) }
		log.Printf("DEBUG=%s\n",					os.Getenv("DEBUG"				))
		log.Printf("LOGS=%s\n",					os.Getenv("LOGS"				))
		log.Printf("IMAGE_URL=%s\n",				os.Getenv("IMAGE_URL"			))
		log.Printf("CONTAINER=%s\n",				os.Getenv("CONTAINER"			))
		log.Printf("REPO=%s\n",					os.Getenv("REPO"				))
		log.Printf("ROUTE_BASE=%s\n",			os.Getenv("ROUTE_BASE"			))
		log.Printf("TITLE=%s\n",					os.Getenv("TITLE"				))
		log.Printf("TARGET_ALIAS=%s\n",			os.Getenv("TARGET_ALIAS"		))
		log.Printf("TARGET_IMAGE_TAG=%s\n",		os.Getenv("TARGET_IMAGE_TAG"	))
		log.Printf("TARGET_LOCAL_PORT=%s\n",		os.Getenv("TARGET_LOCAL_PORT"	))
		log.Printf("TARGET_LOG_LEVEL=%s\n",		os.Getenv("TARGET_LOG_LEVEL"	))
		log.Printf("TARGET_PROJECT_ID=%s\n",		os.Getenv("TARGET_PROJECT_ID"	))
		log.Printf("TARGET_REMOTE_PORT=%s\n",	os.Getenv("TARGET_REMOTE_PORT"	))
	}

	stderr, _		:= command.StderrPipe()
	composeError	 = command.Start()
	if composeError != nil { log.Printf("%s", Red(composeError)) }

	scanner			:= bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		stderrText := scanner.Text()
		if Fd.FdVerbose { log.Printf("%s", Grey(stderrText)) }
	}

	composeError = command.Wait()
	if composeError != nil {
		log.Printf("%s$ %s%s", prefix, command, WhiteOnRed(" X "))
		log.Fatalf("\n%s", Red(composeError))
	}

	return true
}


func dockerClean() bool {
	logPrefix	:= Yellow(pad.Right("\ndockerClean():", 20, " "))
	args		:= "container prune --force"
	argsAbbrev	:= args

	dockerRun(logPrefix, args, argsAbbrev)

	args		= "image prune --all --force"
	argsAbbrev	= args
	dockerRun(logPrefix, args, argsAbbrev)

	args		= "network prune --force"
	argsAbbrev	= args
	dockerRun(logPrefix, args, argsAbbrev)

	args		= "volume prune --force"
	argsAbbrev	= args
	dockerRun(logPrefix, args, argsAbbrev)

	return true
}


func dockerRun(prefix string, cmdArgs string, cmdArgsAbbrev string) bool {
	if Fd.FdVerbose {
		logCommand	:= BlackOnGray(" docker " + cmdArgs + " ")
		fmt.Printf("%s$ %s", prefix, logCommand)
	} else {
		logCommand	:= "docker " + cmdArgsAbbrev
		fmt.Printf("%s$ %s", prefix, logCommand)
	}

	command		:= exec.Command("docker", strings.Split(cmdArgs, " ")...)
	stderr, _	:= command.StderrPipe()
	dockerError	 = command.Start(); if dockerError != nil { log.Printf("%s", Red(dockerError)) }
	scanner		:= bufio.NewScanner(stderr)

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		stderrText := scanner.Text()
		if Fd.FdVerbose { log.Printf("%s", Grey(stderrText)) }
	}

	dockerError		= command.Wait()
	if dockerError != nil {
		log.Printf("%s$ %s%s", prefix, command, WhiteOnRed(" X "))
		log.Fatalf("\n%s", Red(dockerError))
	}

	return true
}


func setEnvironment() bool {
	if err := os.Setenv("DEBUG",				strconv.FormatBool(Fd.FdDebug)	); err != nil { println("derp"); return false }
	if err := os.Setenv("TEST",					strconv.FormatBool(Fd.FdTest)	); err != nil { println("derp"); return false }
	if err := os.Setenv("LOGS",					strconv.FormatBool(Fd.FdVerbose)); err != nil { println("derp"); return false }
	if err := os.Setenv("TARGET_LOCAL_PORT",	Fd.FdTargetLocalPort			); err != nil { println("derp"); return false }
	if err := os.Setenv("TARGET_LOG_LEVEL",		Fd.FdTargetLogLevel				); err != nil { println("derp"); return false }
	if err := os.Setenv("TARGET_REMOTE_PORT",	Fd.FdTargetRemotePort			); err != nil { println("derp"); return false }
	if err := os.Setenv("TARGET_PROJECT_ID",	Fd.FdTargetProjectId			); err != nil { println("derp"); return false }
	if err := os.Setenv("TARGET_ALIAS",			Fd.FdTargetAlias				); err != nil { println("derp"); return false }
	if err := os.Setenv("REPO",					Fd.FdRepo						); err != nil { println("derp"); return false }
	if err := os.Setenv("TARGET_IMAGE_TAG",		Fd.FdTargetImageTag				); err != nil { println("derp"); return false }
	if err := os.Setenv("ROUTE_BASE",			Fd.FdRouteBase					); err != nil { println("derp"); return false }
	if err := os.Setenv("TITLE",				Fd.FdTitle						); err != nil { println("derp"); return false }

	if Fd.FdService == "" {
		if Fd.FdBuildContext == "go" { if err := os.Setenv("EXECUTABLE", Fd.FdRepo );	err != nil { println("derp"); return false }}
		if err := os.Setenv("IMAGE_URL", "us.gcr.io/"+Fd.FdTargetProjectId+"/"+Fd.FdRepo+"/"+Fd.FdTargetAlias+":"+Fd.FdTargetImageTag);	err != nil { println("derp"); return false }
		if err := os.Setenv("CONTAINER", Fd.FdRepo+"--"+Fd.FdTargetProjectId+"--"+Fd.FdTargetAlias);	err != nil { println("derp"); return false }

	} else {
		if err := os.Setenv("SERVICE", Fd.FdService);	err != nil { println("derp"); return false }
		if Fd.FdBuildContext == "go" { if err := os.Setenv("EXECUTABLE", Fd.FdService);	err != nil { println("derp"); return false }}
		if err := os.Setenv(
			"IMAGE_URL", "us.gcr.io/"+Fd.FdTargetProjectId+"/"+Fd.FdRepo+"/"+Fd.FdService+"/"+Fd.FdTargetAlias+":"+Fd.FdTargetImageTag,
		);	err != nil { println("derp"); return false }
		if err := os.Setenv(
			"CONTAINER", Fd.FdService+"--"+Fd.FdRepo+"--"+Fd.FdTargetProjectId+"--"+Fd.FdTargetAlias,
		);	err != nil { println("derp"); return false }
	}

	return true
}


func GetComposeError()	error { return composeError	}
