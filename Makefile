SHELL=/bin/bash
CMD=gchccf
GOBUILD=go build
USER=stvnwrgs
REPOSITORY=request-forwarder
VERSION=$(shell git describe --always --tags | rev |cut -d- -f2-  | rev )

XC_OS_ARCH="darwin/amd64 linux/386 linux/amd64"

build:
	XC_OS_ARCH=${XC_OS_ARCH} CMD=${CMD} ./scripts/build.sh

release-gh:
	USER=${USER} REPOSITORY=${REPOSITORY} VERSION=${VERSION} CMD=${CMD} ./scripts/release.sh

save-version:
	git describe --always --tags
	echo ${VERSION} > git-info

