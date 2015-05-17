GIT_COMMIT=`git rev-parse --short HEAD`
GIT_BRANCH=`git rev-parse --abbrev-ref HEAD`
GIT_TAG=`git describe --tags --exact-match --match "v*" 2>/dev/null || echo "none"`
GIT_DIRTY=`test -n "$(git status --porcelain)" && echo true || echo false`

all: test

dev:
	gin -g -t api/ -a 3000 -p 5803 run server.go
clean:
	git clean -Xdf -e '!.vagrant' -e '!script/custom-vagrant'

test: test-unit test-integration

test-unit:
	godep go test ./...

test-integration:
	@echo TODO: Integration testing is not implement yet

deps:
	godep restore
	PWD=~/go/src/github.com/coverit/coverit go get -t ./...
	PWD=~/go/src/github.com/coverit/coverit godep save ./...
.PHONY: all clean test test-unit test-integration
