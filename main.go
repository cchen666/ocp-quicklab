package main

import (
	"net/http"
	db "ocp-quicklab/db"
	service "ocp-quicklab/service"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func labInstall(c echo.Context) error {
	version := c.Param("version")
	return c.String(http.StatusOK, service.LabInstall(version, c))
}

func labDelete(c echo.Context) error {
	version := c.Param("version")
	return c.String(http.StatusOK, service.LabDelete(version, c))
}

func labList(c echo.Context) error {
	return c.String(http.StatusOK, service.LabList(c))
}

func labTest(c echo.Context) error {
	return c.String(http.StatusOK, service.LabTest(c))
}

func dbTest(c echo.Context) error {
	return c.String(http.StatusOK, db.Test())
}

func main() {
	// Echo instance
	file, err := os.OpenFile(".ocp-quicklab.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
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
	e.GET("/dbtest", dbTest)
	// Start server
	e.Logger.Fatal(e.Start(":1333"))
}
