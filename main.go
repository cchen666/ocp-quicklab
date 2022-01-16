package main

import (
	"net/http"
	"ocp-quicklab/lab"

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

func main() {
	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/install/:version", labInstall)
	e.GET("/delete/:version", labDelete)
	e.GET("/list", labList)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
