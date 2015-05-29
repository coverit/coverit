# coverit

> Code Coverage for iOS/Mac

[![CI Status](http://img.shields.io/travis/coverit/coverit/master.svg?style=flat)](https://travis-ci.org/coverit/coverit)

## Modules

- [web](web), Frontend
- [cli](ios), iOS SDK
- [api](api), RestAPI Implementation
- [worker](worker), worker

## Contribution

### Test

Test in docker container,

    docker run -it -v $PWD:/go/src/github.com/coverit/coverit coverit/golang-dev  /bin/bash
    cd /go/src/github.com/coverit/coverit
    make test

or

    docker-compose run dev make test

### Deps management

Add dep,

    PWD=~/go/src/github.com/coverit/coverit godep get ./...
