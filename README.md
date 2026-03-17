# SSWU-GO001 [![ci](https://github.com/daggerok/SSWU-GO001/actions/workflows/ci.yaml/badge.svg)](https://github.com/daggerok/SSWU-GO001/actions/workflows/ci.yaml)

<!--

```bash
mkdir 00-shared-library ; cd $_
go mod init github.com/daggerok/SSWU-GO001/00-shared-library
cd ..
go work init ; go work use -r . ; go work sync ; go list -m all

mkdir 01-hello-gopher ; cd $_
go mod init github.com/daggerok/SSWU-GO001/01-hello-gopher
cd ..
go work init ; go work use -r . ; go work sync ; go list -m all

mkdir 02-hello-other ; cd $_
go mod init github.com/daggerok/SSWU-GO001/02-hello-other
cd ..
go work init ; go work use -r . ; go work sync ; go list -m all

mkdir 03-fail-fast-and-first-success-goroutines-channels ; cd $_
go mod init github.com/daggerok/SSWU-GO001/03-fail-fast-and-first-success-goroutines-channels
cd ..
go work init ; go work use -r . ; go work sync ; go list -m all

mkdir 04-primitive-types ; cd $_
go mod init github.com/daggerok/SSWU-GO001/04-primitive-types
cd ..
go work init ; go work use -r . ; go work sync ; go list -m all

mkdir 05-functions-methods-structs ; cd $_
go mod init github.com/daggerok/SSWU-GO001/05-functions-methods-structs
cd ..
go work init ; go work use -r . ; go work sync ; go list -m all

mkdir 06-goroutines-and-concurrency-model ; cd $_
go mod init github.com/daggerok/SSWU-GO001/06-goroutines-and-concurrency-model
cd ..
go work init ; go work use -r . ; go work sync ; go list -m all

mkdir 07-concurrency-patterns ; cd $_
go mod init github.com/daggerok/SSWU-GO001/07-concurrency-patterns
cd ..
go work init ; go work use -r . ; go work sync ; go list -m all

mkdir 08-race-conditions-and-mutex ; cd $_
go mod init github.com/daggerok/SSWU-GO001/08-race-conditions-and-mutex
cd ..
go work init ; go work use -r . ; go work sync ; go list -m all

mkdir 08-concurrent-task-management ; cd $_
go mod init github.com/daggerok/SSWU-GO001/08-concurrent-task-management
cd ..
go work init ; go work use -r . ; go work sync ; go list -m all
```

-->

## init workspace

```bash
git clone https://github.com/daggerok/SSWU-GO001.git
go work init ; go work use -r . ; go work sync ; go list -m all
```

## test

```bash
# go test ./00-shared-library/... # or:
go test `go work edit -json | jq -r '.Use[].DiskPath + "/..."'` # requires "brew reinstall jq"

# and with -v flag to see logs output  
go test -v -race `go work edit -json | jq -r '.Use[].DiskPath + "/..."'` # requires "brew reinstall jq"  
```

## run

```bash
go run ./01-hello-gopher
# Hello, G0!
```
