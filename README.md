# coverit

> Code Coverage for iOS/Mac

[![CI Status](http://img.shields.io/travis/coverit/coverit/master.svg?style=flat)](https://travis-ci.org/coverit/coverit)

## Modules

- [web](web), Frontend
- [cli](ios), iOS SDK
- [api](api), RestAPI Implementation
- [worker](worker), worker

## Contribution

    docker run -it -v $PWD:/go/src/github.com/coverit/coverit coverit/golang-dev  /bin/bash
    cd /go/src/github.com/coverit/coverit

Add dep,

    PWD=~/go/src/github.com/coverit/coverit godep get ./...
