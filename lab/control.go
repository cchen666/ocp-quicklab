package lab

import "github.com/labstack/echo/v4"

var mirrorURL = "https://mirror.openshift.com/pub/openshift-v4/x86_64/clients/ocp/"
var packageName = "openshift-install-linux.tar.gz"
var baseDir = "/root/openshift/"
var globalVersion = ""

func LabInstall(version string, c echo.Context) string {
	deployInstaller(version, c)
	install(version, c)
	globalVersion = version
	return "LabInstall Done \n"
}

func LabDelete(version string, c echo.Context) string {
	var workingDir = baseDir + version + "/"
	var deleteCluster = workingDir + "openshift-install destroy cluster --dir=" + workingDir
	runBash(deleteCluster, c)
	globalVersion = ""
	return "LabDelete Done \n"
}

func LabList(c echo.Context) string {
	if globalVersion != "" {
		var listCluster = "grep 'Install complete' -A3 .openshift_install.log"
		runBash(listCluster, c)
		return "LabList Done \n"
	}
	return "Failed to List because we don't have installed cluster"

}

func deployInstaller(version string, c echo.Context) {
	var workingDir = baseDir + version + "/"
	var mkdirWorkingDir = "mkdir -p " + workingDir
	runBash(mkdirWorkingDir, c)
	var copyInstallYaml = "cp /root/install-config.yaml " + workingDir
	runBash(copyInstallYaml, c)
	var rmPackage = "rm -rf " + workingDir + "openshift-install*"
	runBash(rmPackage, c)
	var downloadPackage = "wget -O " + workingDir + "openshift-install-tar.gz " + mirrorURL + version + "/" + packageName
	for i := 0; i < 5; i++ {
		if runBash(downloadPackage, c) == "Success" {
			break
		}
	}
	var extractInstaller = "tar xvf " + workingDir + "openshift-install-tar.gz" + " -C " + workingDir
	runBash(extractInstaller, c)

}

func install(version string, c echo.Context) {
	var workingDir = baseDir + version + "/"
	var runInstaller = workingDir + "openshift-install create cluster --dir=" + workingDir
	runBash(runInstaller, c)
}
