package lab

import "github.com/labstack/echo/v4"

type Cluster struct {
	URL      string
	Password string
	Version  string
}

var mirrorURL = "https://mirror.openshift.com/pub/openshift-v4/x86_64/clients/ocp/"
var packageName = "openshift-install-linux.tar.gz"

func LabInstall(version string, c echo.Context) string {
	//cluster := deployInstaller(version)
	deployInstaller(version, c)
	//return cluster.URL + cluster.Password
	return "Done"
}

func deployInstaller(version string, c echo.Context) {
	var rmPackage = "rm -rf openshift-install*"
	runBash(rmPackage, c)
	var downloadPackage = "wget " + mirrorURL + version + "/" + packageName
	runBash(downloadPackage, c)
	var extractInstaller = "tar xvf openshift-install-linux.tar.gz"
	runBash(extractInstaller, c)
	var mkdirWorkingDir = "mkdir -p " + "/tmp/openshift/" + version
	runBash(mkdirWorkingDir, c)
	var copyInstallYaml = "cp /root/install-config.yaml " + "/tmp/openshift/" + version
	runBash(copyInstallYaml, c)
	//var runInstaller = "./openshift-install create cluster --dir=/tmp/openshift/" + version
}

func TestCLI(c echo.Context) {
	var commandLine = "ping -c4 8.8.8.8"
	runBash(commandLine, c)
}
