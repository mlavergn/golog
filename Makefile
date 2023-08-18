###############################################
#
# Makefile
#
###############################################

.DEFAULT_GOAL := run

.PHONY: test

run:
	cd demo; go run demo.go

test:
	go test -v -count=1 ./...

lint:
	go vet *.go

format:
	go fmt *.go

#
# Publishing
#

VERSION := 1.0.0
PROJECT := golog
REPO := mlavergn

github:
	open "https://github.com/${REPO}/${PROJECT}"

release:
	zip -r ${PROJECT}.zip LICENSE README.md Makefile *.go
	hub release create -m "${VERSION} - ${PROJECT}" -a ${PROJECT}.zip -t master "v${VERSION}"
	open "https://github.com/${REPO}/${PROJECT}/releases"

st:
	open -a SourceTree .