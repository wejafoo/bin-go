

package blib

import (
	"bufio"
	"fmt"
	"github.com/willf/pad"
	"log"
	"os"
	"os/exec"
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
	args		:= "builds submit --config="	+ gcpCbFile				+ " --substitutions" +
				" _NICKNAME="					+ Fd.FdNickname			+
				",_SERVICE_NAME="				+ Fd.FdServiceName		+
				",_TARGET_ALIAS="				+ Fd.FdTargetAlias		+
				",_TARGET_IMAGE_TAG="			+ Fd.FdTargetImageTag
	argsSmall	:= "builds submit --config="	+ gcpCbFile
	success		:= gcloudRun(logPrefix, args, argsSmall)

	return success
}


func gcloudRun(logPrefix string, cmdArgs string, cmdArgsSmall string) bool {
	success 	:= false

	if Fd.FdVerbose {
		logCommand	:= BlackOnGray(" gcloud " + cmdArgs + " ")
		fmt.Printf("%s$ %s\n", logPrefix, logCommand)
		fmt.Printf("\n")
	}

	command		:= exec.Command("gcloud", strings.Split(cmdArgs, " ")...)
	command.Env	= os.Environ()
	stderr, _	:= command.StderrPipe()
	gcpError	= command.Start()
	if gcpError != nil { log.Printf("%s", Red(gcpError)) }

	scanner			:= bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		stderrText := scanner.Text()
		if Fd.FdVerbose { log.Printf("%s", Grey(stderrText)) }
	}

	logCommandSmall := BlackOnGray(pad.Right("gcloud " + cmdArgsSmall,70,"." ))

	gcpError = command.Wait()
	if gcpError != nil {
		fmt.Printf("%s$ %s%s", logPrefix, logCommandSmall, LogLose)
		log.Fatalf("\n%s", Red(gcpError))
	}
	success = true

	return success
}


func GetGcpError() error { return gcpError }
