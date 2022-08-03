package service

import (
	"os/exec"
	"time"

	"github.com/labstack/echo/v4"
)

func alwaysFlush(ch chan int, c echo.Context) {
	/*
		Start a go routine to untimately flush the
		Writer (c.Response()) and quit if the cmd.Run()
		finished
	*/
	for {
		select {
		// Quit this go routine
		case <-ch:
			return
		default:
			c.Response().Flush()
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func runBash(command string, c echo.Context) string {
	writeString("Running "+command, c)
	ch := make(chan int)
	go alwaysFlush(ch, c)
	cmd := exec.Command("/bin/bash", "-c", command)
	// Redirect the stdout and stderr to ioWriter of echo
	cmd.Stdout = c.Response()
	cmd.Stderr = c.Response()
	err := cmd.Run()
	if err != nil {
		c.Response().Write([]byte(err.Error()))
		// Stop the goroutine
		ch <- 1
		//log.Fatalf("Error %s", command, err)
		close(ch)
		return "Failure"
	}
	// Stop the goroutine
	ch <- 1
	writeString("Finished "+command, c)
	close(ch)
	return "Success"
}

func writeString(s string, c echo.Context) {
	runningString := s + "\n"
	c.Response().Write([]byte(runningString))
}
