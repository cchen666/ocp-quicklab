package lab

import (
	"os"
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
			time.Sleep(100 * time.Millisecond)
			c.Response().Flush()
		}
	}
}

func runBash(command string, c echo.Context) {
	ch := make(chan int)
	go alwaysFlush(ch, c)
	writeString("Running "+command, c)
	cmd := exec.Command("/bin/bash", "-c", command)
	// Redirect the stdout and stderr to ioWriter of echo
	cmd.Stdout = c.Response()
	cmd.Stderr = c.Response()
	err := cmd.Run()
	if err != nil {
		c.Response().Write([]byte(err.Error()))
		os.Exit(1)
	}
	// Stop the goroutine
	ch <- 1
	writeString("Finished "+command, c)
	close(ch)
}

func writeString(s string, c echo.Context) {
	runningString := s + "\n"
	c.Response().Write([]byte(runningString))
}
