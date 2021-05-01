

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
	gcpError error
)


func NewGcp() bool {
	success := true

	if Fd.FdLocal {
		fmt.Printf("UNDER CONSTRUCTION:  Gcp(Kubernetes local deployment")
	}  else {
		success = gcpDeploy()
	}

	return success
}


func gcpDeploy() bool {
	logPrefix	:= Yellow(pad.Right("\ngcloudDeploy():", 20, " "))
	gcpCbFile	:= "cloudbuild.json"
	args		:= "builds submit --no-source --config=" + gcpCbFile				+
						" --substitutions="											+
						"_DEBUG="				+ strconv.FormatBool(Fd.FdDebug)	+
						",_LOGS="				+ strconv.FormatBool(Fd.FdVerbose)	+
						",_SERVICE="			+ Fd.FdService						+
						",_ROUTE_BASE="			+ Fd.FdRouteBase					+
						",_REPO="				+ Fd.FdRepo							+
						",_TITLE=\""			+ Fd.FdTitle						+
						"\",_TARGET_ALIAS="		+ Fd.FdTargetAlias					+
						",_TARGET_LOG_LEVEL="	+ Fd.FdTargetLogLevel				+
						",_TARGET_IMAGE_TAG="	+ Fd.FdTargetImageTag
	// 					",_TARGET_REMOTE_PORT="	+ Fd.FdTargetRemotePort				+
	// 					",_TEST="				+ strconv.FormatBool(Fd.FdTest)		+  // future implementation supporting remote push-right testing
	argsAbbrev	:= "builds submit (...) --config=" + gcpCbFile + " --substitutions=(...)"

	return gcloudRun(logPrefix, args, argsAbbrev)
}


func gcloudRun(logPrefix string, cmdArgs string, cmdArgsAbbrev string) bool {
	stderrText1	:= ""
	stderrText2	:= ""
	stderrText3	:= ""
	logCommand	:= ""
	
	if Fd.FdVerbose {
		logCommand	= BlackOnGray(" gcloud " + cmdArgs + " ")
		fmt.Printf("%s$ %s", logPrefix, logCommand)
		fmt.Printf("\n")
	} else {
		logCommand	= "gcloud " + cmdArgsAbbrev
		fmt.Printf("%s$ %s", logPrefix, logCommand)
	}
	
	command			:= exec.Command("gcloud", strings.Split(cmdArgs, " ")...)
	command.Env		= os.Environ()
	stderr, _		:= command.StderrPipe()
	gcpError		= command.Start()
	if gcpError != nil { log.Printf("%s", Red(gcpError)) }

	scanner			:= bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		stderrText1 = scanner.Text()
		if Fd.FdVerbose { log.Printf("%s", Grey(stderrText1)) }
		stderrText3 = stderrText2
		stderrText2 = stderrText1
	}
	
	gcpError = command.Wait()
	if gcpError != nil  {
		fmt.Printf("%s$ %s%s\n", logPrefix, logCommand, LogLose)
		if ! Fd.FdVerbose {
			log.Printf("%s", stderrText3)
			log.Printf("%s", stderrText2)
			log.Printf("%s", stderrText1)
		}
		log.Fatalf("%s", Red(gcpError))
	}

	return true
}


func GetGcpError() error { return gcpError }
