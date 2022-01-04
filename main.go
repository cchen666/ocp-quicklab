package main

import (
	"net/http"
	"ocp-quicklab/lab"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func labInstall(c echo.Context) error {
	version := c.Param("version")
	return c.String(http.StatusOK, lab.LabInstall(version, c))
}

func testCLI(c echo.Context) error {
	return c.String(http.StatusOK, lab.TestCLI(c))
}

func test(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	c.Response().WriteHeader(http.StatusOK)
	b := []byte{1, 2, 3, 4}
	for range b {
		c.Response().Write([]byte("data: "))
		c.Response().Write([]byte("\n"))
		c.Response().Flush()
		time.Sleep(1 * time.Second)
	}
	return nil
}

func main() {
	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/install/:version", labInstall)
	e.GET("/", test)
	e.GET("/testcli", testCLI)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
