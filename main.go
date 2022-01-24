package main

import (
	"net/http"
	"ocp-quicklab/lab"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func labInstall(c echo.Context) error {
	version := c.Param("version")
	return c.String(http.StatusOK, lab.LabInstall(version, c))
}

func labDelete(c echo.Context) error {
	version := c.Param("version")
	return c.String(http.StatusOK, lab.LabDelete(version, c))
}

func labList(c echo.Context) error {
	return c.String(http.StatusOK, lab.LabList(c))
}

func labTest(c echo.Context) error {
	return c.String(http.StatusOK, lab.LabTest(c))
}

func main() {
	// Echo instance
	file, err := os.OpenFile("/var/log/ocp-quicklab.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/install/:version", labInstall)
	e.GET("/delete/:version", labDelete)
	e.GET("/test", labTest)
	e.GET("/list", labList)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
