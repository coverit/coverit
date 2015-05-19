# coverit-api

> Coverit REST API service

## Customizations

- `MONGO_URL`, mongodb host. e.g. `localhost`
- `MONGO_DB`, ,mongodb db. e.g. `juliet`
- `PORT`, listening port. e.g. `8000`

## Storage Models

### Build

- `Id`
- `Git`, Git info of current build
  - `Head`, current head info
    - `Id`, commit SHA
    - `AuthorName`
    - `AuthorEmail`
    - `CommiterName`
    - `CommiterEmail`
    - `Message`
  - `Branch`, current branch of build
  - `Tag`, current tag of build
- `GcnoArchive`, zip of all gcon files of current build
- `SourceTree`, source tree
```yaml
file1:
    body: "file body"
file2:
    body: "file body"
dir1:
    body: "dir body"
    /:
        file3:
            body: "file body"
        file4:
            body: "file body"
```
- `CreatedAt`, create timestamp
- `UpdatedAt`, update timestamp


### Report

- `Id`
- `BuildId`, build Id of current test
- `RunAt`
    - `Platform`, test platform of current test, i.e. gitlab, travis
    - `Device`, testing hardware of current test, i.e. iPhone 6+, Simulator
    - `SDK`, testing SDK, i.e. iOS 8.3, Android 5.2
- `GcdaArchive`, zip of all gcda files of current test
- `CoverageTree`, coverage info tree
    - `stats`, file level coverage info
    - `functions`, function level coverage info
    - `lines`, line level coverage info
```yaml
stats:
    - source: "file1"
      lines": 10
      lines_hit: 11
      functions: 2
      functions_hit: 1
    - source: "dir/file2"
      lines": 15
      lines_hit: 10
      functions: 1
      functions_hit: 1
functions:
    - source: "file1"
      name: "foo()"
      line_no: 10
      hit: 1
    - source: "file1"
      name: "bar()"
      line_no: 12
      hit: 0
    - source: "dir/file2"
      name: "go()"
      line_no: 10
      hit: 1
lines:
    - source: "file1"
      line_no: 13
      hit: 10
    - source: "file1"
      line_no: 15
      hit: 1
    - source: "dir/file2"
      line_no: 11
      hit: 1
```
- `CreatedAt`, create timestamp
- `UpdatedAt`, update timestamp

### Source


### Report

- coverages
