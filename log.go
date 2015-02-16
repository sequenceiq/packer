package main

import (
	"io"
	"io/ioutil"
	"os"
)

// These are the environmental variables that determine if we log, and if
// we log whether or not the log should go to a file.
const EnvLog = "PACKER_LOG"          //Set to True
const EnvLogFile = "PACKER_LOG_PATH" //Set to a file

// logOutput determines where we should send logs (if anywhere).
func logOutput() (logOutput io.Writer, err error) {
	logOutput = ioutil.Discard
	if os.Getenv(EnvLog) != "" {
		logOutput = os.Stderr

		if logPath := os.Getenv(EnvLogFile); logPath != "" {
			var err error
			logOutput, err = os.Create(logPath)
			if err != nil {
				return nil, err
			}
		}
	}
	//fmt.Println("[LOG-HACK] output:", logOutput)
	return
}
