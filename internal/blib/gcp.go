

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
	success = gcpDeploy()

	return success
}


func gcpDeploy() bool {

	success		:= false
	logPrefix	:= Yellow(pad.Right("\ngcpUp():", 20, " "))
	args		:= "--log-level " + Fd.FdTargetLogLevel + " up --detach --force-recreate " + Fd.FdServiceName
	logMessage	:= BlackOnGray(" gcp-gcp " + args)

	if !Fd.FdQuiet {
		fmt.Printf("%s$ %s", logPrefix, logMessage)
	} else {
		logMessage2 := "Spinning up local Gcp instance"
		fmt.Printf("%s%s", logPrefix, logMessage2)
	}
	if Fd.FdVerbose { fmt.Printf("\n") }

	cmd := exec.Command("gcp-gcp", strings.Split(args, " ")...)
	setEnvironment()
	cmd.Env = os.Environ()

	stderr, _ := cmd.StderrPipe()
	e1 := cmd.Start()
	if e1 != nil { log.Printf("%s", Red(e1)) }

	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		stderrText := scanner.Text()
		if Fd.FdVerbose { log.Printf("%s", Grey(stderrText)) }
	}
	e2 := cmd.Wait()

	commandSmall := pad.Right(BlackOnGray(" gcp-gcp up (...) "+Fd.FdServiceName+" "), 69, ".")
	if e2 != nil {
		fmt.Printf("%s$ %s%s", logPrefix, commandSmall, LogLose)
		fmt.Printf("\n")
		log.Fatal(Red(e2))
	} else {
		success = true
		fmt.Printf("%s$ %s%s", logPrefix, commandSmall, LogWin)
	}
	return success
}


func GetGcpError() error { return gcpError }
