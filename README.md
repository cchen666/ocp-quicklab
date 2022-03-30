# OCP-QuickLab

## Usage

### Install

~~~bash

$ curl <server>:<port>/install/<version>

$ curl 10.72.46.240:1323/install/4.8.18

Running tar xvf /root/openshift/4.8.18/openshift-install-tar.gz -C /root/openshift/4.8.18/
README.md
openshift-install
Finished tar xvf /root/openshift/4.8.18/openshift-install-tar.gz -C /root/openshift/4.8.18/
Running /root/openshift/4.8.18/openshift-install create cluster --dir=/root/openshift/4.8.18/
level=info msg=Credentials loaded from the "default" profile in file "/root/.aws/credentials"
level=info msg=Consuming Install Config from target directory

~~~
