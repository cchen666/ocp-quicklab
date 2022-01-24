package lab

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

var mirrorURL = "https://mirror.openshift.com/pub/openshift-v4/x86_64/clients/ocp/"
var packageName = "openshift-install-linux.tar.gz"
var baseDir = "/root/openshift/"
var logFile = "/var/log/ocp-quicklab.log"
var deletedString = "Lab Delete Finished"

func LabInstall(version string, c echo.Context) string {
	if isDeleted() {
		deployInstaller(version, c)
		install(version, c)
		return "Lab Install Done \n"
	}
	LabList(c)
	return "Can not install cluster because the cluster is already existed"
}

func LabDelete(version string, c echo.Context) string {
	var workingDir = baseDir + version + "/"
	var deleteCluster = workingDir + "openshift-install destroy cluster --dir=" + workingDir
	runBash(deleteCluster, c)
	log.Infof("%s %s", deletedString, version)
	return "Lab Delete Done \n"
}

func LabList(c echo.Context) string {
	var listCluster = "tail -5 /var/log/ocp-quicklab.log"
	return runBash(listCluster, c)
}

func deployInstaller(version string, c echo.Context) {
	var workingDir = baseDir + version + "/"
	var mkdirWorkingDir = "mkdir -p " + workingDir
	runBash(mkdirWorkingDir, c)
	var copyInstallYaml = "cp /root/install-config.yaml " + workingDir
	runBash(copyInstallYaml, c)
	var rmPackage = "rm -rf " + workingDir + "openshift-install*"
	runBash(rmPackage, c)
	var downloadPackage = "https_proxy=squid.redhat.com:3128 wget -O " + workingDir + "openshift-install-tar.gz " + mirrorURL + version + "/" + packageName
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
	var grepCompleted = "grep 'Install complete' -A3 " + baseDir + version + "/.openshift_install.log >> /var/log/ocp-quicklab.log"
	runBash(grepCompleted, c)
	log.Infof("Lab Installation Finished %s", version)
}

func isDeleted() bool {
	fmt.Println(ReadFile(logFile))
	fmt.Println(strings.Contains(ReadFile(logFile), deletedString))
	return strings.Contains(ReadFile(logFile), deletedString)
}

func LabTest(c echo.Context) string {
	if isDeleted() {
		return "Deleted and we can install"
	}
	return "We can't install"
}

func ReadFile(file_name string) string {
	// Read the last line of a particular file
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatalf("Error is %s", err)
	}
	defer file.Close()
	var lineText string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText = scanner.Text()
	}
	return string(lineText)
}
