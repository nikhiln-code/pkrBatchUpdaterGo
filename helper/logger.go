package helper

import (
	"fmt"
	"io"
	"time"
	"os"
	"path/filepath"
	log "github.com/sirupsen/logrus"
)

func InitiateLogger(name string) {
	cwd, err := os.Getwd()
	fmt.Print(cwd)

	if err != nil {
		log.Fatalf("Failed to determine working directory: %s", err)
	}

	runID := time.Now().Format("run-2006-01-02-15-04-05")
	logLocation := filepath.Join(cwd, name+"_"+runID+".log")
	logFile, err := os.OpenFile(logLocation, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file %s for output: %s", logLocation, err)
	}
	log.SetOutput(io.MultiWriter(os.Stderr, logFile))

	exitHandler := func() {
		if logFile == nil {
			return
		}
		logFile.Close()
	}

	log.RegisterExitHandler(exitHandler)

}

func ExitLogger() {
	// perform actions
	log.Exit(0)
}
