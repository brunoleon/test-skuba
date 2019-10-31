# test-skuba

This repository holds a few test using Terratest to test rook deployment

## Prerequisites
- terraform >= v0.12 
- terraform-provider-libvirt (for running on libvirt)
- dep
- skuba (using latest available should do)
- kubectl
- libvirt
- SLES15-SP1-JeOS.x86_64-15.1-OpenStack-Cloud-Build35.30.qcow2 ISO available in terraform receipe folder (or update vars accordingly)

## Installation

### skuba

To install skuba using go, run

```
go get github.com/SUSE/skuba/cmd/skuba
```

### dep

To install terratest and the other dependencies used by the test framework,
dep is used. To install dep, either follow the installation instructions
on golang.github.io/dep/docs/installation.html, or run

```
go get -u github.com/golang/dep/cmd/dep
```

Then, to ensure that Terratest and its dependencies are installed, run

```
cd test
dep ensure
```

### terraform-provider-libvirt

The libvirt provider for terraform can be installed for openSUSE/SLES from OBS:

https://build.opensuse.org/package/show/systemsmanagement:terraform/terraform-provider-libvirt

### OS image

openSUSE/kubic based images can be downloaded here:

https://download.opensuse.org/distribution/leap/15.1/jeos/

## Terratest

Terratest is primarily used to test terraform deployments. It is written in Go and has mutiple functions available to interact with various subsytems: https://github.com/gruntwork-io/terratest

In this example, tests have been splitted in three stages corresponding to what we actually need to do in a real deployment:
- `01terraform_test.go`
- `02skuba_test.go`
- `03rook_test.go`

Each of this stage is supposed to do the deployment and test its result.
One of the problem (that could be fixed by refactoring properly) is that after deploying + test terratest will destroy the deployment.
This does not work for us as each stage relies on the previous one. Stage 1 should only be destroyed upon stage 3 execution.


## Executing a test

In the test folder: `go test -v -timeout 90m TEST_FILE.go` (where test file is one of the three above)

```
go test -v -timeout 90m 01terraform_test.go
go test -v -timeout 90m 02skuba_test.go
go test -v -timeout 90m 03rook_test.go
```

