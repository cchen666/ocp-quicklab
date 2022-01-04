package lab

import (
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

var wg sync.WaitGroup

func alwaysFlush(ch chan int, c echo.Context) {
	defer wg.Done()
	for {
		select {
		case <-ch:
			return
		default:
			time.Sleep(100 * time.Millisecond)
			c.Response().Flush()
		}
	}
}

func runBash(command string, c echo.Context) {
	runningString := "Running " + command + "\n"
	// go func() {
	// 	for {
	// 		time.Sleep(500 * time.Millisecond)
	// 		if c.Response() != nil {
	// 			c.Response().Flush()
	// 		}
	// 	}
	// }()
	ch := make(chan int)
	wg.Add(1)
	go alwaysFlush(ch, c)
	c.Response().Write([]byte(runningString))
	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.Stdout = c.Response()
	cmd.Stderr = c.Response()
	err := cmd.Run()
	if err != nil {
		c.Response().Write([]byte(err.Error()))
		c.Response().Flush()
		os.Exit(1)
	}
	finishString := "Finished " + command + "\n"
	c.Response().Write([]byte(finishString))
	// Stop the goroutine
	ch <- 1
	close(ch)
	wg.Wait()
}
