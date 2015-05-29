# coverit

> Code Coverage for iOS/Mac

[![wercker status](https://app.wercker.com/status/ffe09e833162a56b84fd78c1d0231b55/s/master "wercker status")](https://app.wercker.com/project/bykey/ffe09e833162a56b84fd78c1d0231b55)

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
