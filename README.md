# OpenShift QuickLab for gcg-shift team

## Install

~~~bash

$ curl <IP address>:<port>/install/<version>

$ curl 127.0.0.1:1323/install/4.8.24
<Snip>
 96000K .......... .......... .......... .......... .......... 99% 16.7M 0s
 96050K .......... .......... .......... .......... .......... 99% 4.61M 0s
 96100K .......... .......... .......... .......... .......... 99% 18.0M 0s
 96150K .......... .......... .......... .......... .......... 99% 18.7M 0s
 96200K .......... .......... .......... .......... .......... 99% 20.0M 0s
 96250K .......... .......... .......... .......... .......... 99% 3.33M 0s
 96300K .......... .......... .......... .......... .         100%  129M=14s

2022-01-04 14:23:06 (6.89 MB/s) - ‘openshift-install-mac.tar.gz’ saved [98654083/98654083]

Finished wget https://mirror.openshift.com/pub/openshift-v4/x86_64/clients/ocp/4.8.24/openshift-install-mac.tar.gz
Running tar xvf openshift-install-mac.tar.gz
x README.md
x openshift-install
Finished tar xvf openshift-install-mac.tar.gz
Running mkdir -p /tmp/openshift/4.8.24
Finished mkdir -p /tmp/openshift/4.8.24
Running cp /tmp/install-config.yaml /tmp/openshift/4.8.24
Finished cp /tmp/install-config.yaml /tmp/openshift/4.8.24
Running ./openshift-install create cluster --dir=/tmp/openshift/4.8.24

<Snip>
~~~

## List

~~~bash

$ curl <IP address>:<port>/install/<version>

$ $ curl 127.0.0.1:1323/list

Running grep 'Install complete' -A3 /tmp/openshift/*/.openshift_install.log
time="2022-01-04T13:08:13+08:00" level=info msg="Install complete!"
time="2022-01-04T13:08:13+08:00" level=info msg="To access the cluster as the system:admin user when using 'oc', run 'export KUBECONFIG=/tmp/openshift/4.8.24/auth/kubeconfig'"
time="2022-01-04T13:08:13+08:00" level=info msg="Access the OpenShift web-console here: https://console-openshift-console.apps.mycluster.XXXXXX.com"
time="2022-01-04T13:08:13+08:00" level=info msg="Login to the console with user: \"kubeadmin\", and password: \"XXXXXXXXXX\""
Finished grep 'Install complete' -A3 /tmp/openshift/*/.openshift_install.log
LabList Done

~~~

## Delete

~~~bash

$ curl <IP address>:<port>/install/<version>

$ curl 127.0.0.1:1323/delete/4.8.24

~~~
