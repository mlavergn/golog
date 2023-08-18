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
	gh release create v${VERSION} ./${PROJECT}.zip --target master --notes "${VERSION} - ${PROJECT}"
	open "https://github.com/${REPO}/${PROJECT}/releases"

st:
	open -a SourceTree .