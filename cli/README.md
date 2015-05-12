# coverit-ios

> Coverit iOS SDK

## Usage
Coverit need some parameter to running as below:

 - coverit_repo_name
 - coverit_commit_sha
 - coverit_branch_name
 - coverit_api_token
 > More parameter ...

```
coverit [--help] <command> [<args>...]

Options:
    -h, --help

Basic commands

build
report
```

## Integrate with travis

For integrated in travis you need do some config in `.travis.yml`

The first step, you need install `coverit`. You can specified it at hook `before_install`
```
before_install:
  - brew install coverit
```


And than you need to export some important configuration before using coverit, such as:

```
after_success:
  -export COVERIT_BARNCH_NAME = $TRAVIS_BRANCH
  -export coverit_commit_sha  = $TRAVIS_COMMIT
  -export coverit_repo_name   = "AFNetowrking"
```
> You can also set these parameters into .coverit.yml. But we'll preferred to read the environment variables.

Finally we need create coverage report with `coverit` such as:
```
after_success:
  -coverit build  # upload gcno

  # Or use this command to analysis of the coverage locally
  -coverit report # upload lcov

```


### Create a build

$ coverit build create

### Report code coverage

$ coverit report


## Contribution

$ docker-compose up
