Poseidon is Firmament's (http://www.firmament.io) integration with
Kubernetes.

[![Build Status](https://travis-ci.org/camsas/poseidon.svg)](https://travis-ci.org/camsas/poseidon)

***Note: this repo contains an initial prototype, it may break at any time! :)***

# Getting started

The easiest way to get Poseidon up and running is to use our Docker image:

```
$ docker pull camsas/poseidon:dev
```
Once the image has downloaded, you can start Poseidon as follows:
```
$ docker run camsas/poseidon:dev /usr/bin/poseidon \
    --logtostderr \
    --k8s_apiserver_host=<host> \
    --k8s_apiserver_port=<port> \
    --cs2_binary=/usr/bin/cs2.exe \
    --max_tasks_per_pu=<max pods per node>
```
Note that Poseidon will try to schedule for Kubernetes even if `kube-scheduler`
is running -- to avoid conflicts, shut it down first. If your `kube-scheduler`
is running locally then you can execute `sudo service kube-scheduler stop`,
otherwise stop its pod.

You will also need to ensure that the API server is reachable from the Poseidon
container's network (e.g., using `--net=host` if you're running a local
development cluster).

# Building from source

## System requirements

 * CMake 2.8+
 * Docker 1.7+
 * Kubernetes v1.1+
 * Boost 1.54+
 * libssl 1.0.0+

The build process will install local version of Poseidon's dependencies, which
currently include:

 * the [Microsoft C++ REST SDK](https://github.com/Microsoft/cpprestsdk) v2.7.0
 * the [Firmament scheduler](https://github.com/camsas/firmament) (HEAD)

and their dependencies.

A known-good build environment is Ubuntu 16.04 with gcc 5.1.3.


## Build process

First, generate the build infrastructure:

```
$ mkdir build
$ cd build
build$ cmake ..
```

Then, build Poseidon:

```
build$ make
```

Following, make sure you have a Kubernetes cluster running. To deploy one you can execute:

```
./deploy/build_kubernetes.sh
./deploy/run_kubernetes.sh
```

To start up, run from the root directory:

```
$ build/poseidon --flagfile=deploy/poseidon.cfg \
                 --k8s_apiserver_host=<hostname> \
                 --k8s_apiserver_port=8080
```

Arguments in `deploy/poseidon.cfg` control flow scheduling features
(e.g. to choose a scheduling policy). For more info check
[accepted by Firmament](https://github.com/camsas/firmament/blob/master/README.md).

## Contributing

We always welcome contributions to Poseidon. We use GerritHub for our code
reviews, and you can find the Poseidon review board there:

https://review.gerrithub.io/#/q/project:camsas/poseidon+is:open

In order to do code reviews, you will need an account on GerritHub (you can link
your GitHub account).

The easiest way to submit changes for review is to check out Poseidon from
GerritHub, or to add GerritHub as a remote. Alternatively, you can submit a pull
request on GitHub and we will import it for review on GerritHub.
