package lab

import "github.com/labstack/echo/v4"

type Cluster struct {
	URL      string
	Password string
	Version  string
}

var mirrorURL = "https://mirror.openshift.com/pub/openshift-v4/x86_64/clients/ocp/"
var packageName = "openshift-install-mac.tar.gz"

func LabInstall(version string, c echo.Context) string {
	deployInstaller(version, c)
	install(version, c)
	return "LabInstall Done \n"
}

func LabDelete(version string, c echo.Context) string {
	var deleteCluster = "./openshift-install destroy cluster --dir=/tmp/openshift/" + version
	runBash(deleteCluster, c)
	return "LabDelete Done \n"
}

func LabList(c echo.Context) string {
	var deleteCluster = "grep 'Install complete' -A3 /tmp/openshift/*/.openshift_install.log"
	runBash(deleteCluster, c)
	return "LabList Done \n"
}

func deployInstaller(version string, c echo.Context) {
	var rmPackage = "rm -rf openshift-install*"
	runBash(rmPackage, c)
	var downloadPackage = "wget " + mirrorURL + version + "/" + packageName
	for i := 0; i < 5; i++ {
		if runBash(downloadPackage, c) == "Success" {
			break
		}
	}
	var extractInstaller = "tar xvf openshift-install-mac.tar.gz"
	runBash(extractInstaller, c)
	var mkdirWorkingDir = "mkdir -p " + "/tmp/openshift/" + version
	runBash(mkdirWorkingDir, c)
	var copyInstallYaml = "cp /tmp/install-config.yaml " + "/tmp/openshift/" + version
	runBash(copyInstallYaml, c)
	// var runInstaller = "./openshift-install create cluster --dir=/tmp/openshift/" + version
	// runBash(runInstaller, c)
}

func install(version string, c echo.Context) {
	var runInstaller = "./openshift-install create cluster --dir=/tmp/openshift/" + version
	runBash(runInstaller, c)
}

func TestCLI(c echo.Context) string {
	var commandLine = "ping -c4 8.8.8.8"
	runBash(commandLine, c)
	return "Finished running TestCLI \n"
}
