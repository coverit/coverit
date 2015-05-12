# coverit

> Code Coverage for iOS/Mac

[![CI Status](http://img.shields.io/travis/coverit/coverit/master.svg?style=flat)](https://travis-ci.org/coverit/coverit)

## Modules

- [web](web), Frontend
- [cli](ios), iOS SDK
- [api](api), RestAPI Implementation
- [worker](worker), worker

## Contribution

Development environment is provided by Vagrant.

    vagrant up
    cd /vagrant

Add dep,

    PWD=~/go/src/github.com/coverit/coverit godep get ./...
