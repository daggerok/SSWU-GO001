# SSWU-GO001 [![ci](https://github.com/daggerok/SSWU-GO001/actions/workflows/ci.yaml/badge.svg)](https://github.com/daggerok/SSWU-GO001/actions/workflows/ci.yaml)

<!--

```bash
mkdir 00-shared-library ; cd $_
go mod init github.com/daggerok/SSWU-GO001/00-shared-library
cd ..

mkdir 01-hello-gopher ; cd $_
go mod init github.com/daggerok/SSWU-GO001/01-hello-gopher
cd ..

mkdir 01-hello-other ; cd $_
go mod init github.com/daggerok/SSWU-GO001/02-hello-other
cd ..
```

-->

## init workspace

```bash
git clone https://github.com/daggerok/SSWU-GO001.git
go work init ; go work use -r . ; go work sync ; go list -m all
```

## test

```bash
go test ./00-shared-library/...
# or brew reinstall jq
go test `go work edit -json | jq -r '.Use[].DiskPath + "/..."'`  
```

## run

```bash
go run ./01-hello-gopher
# Hello, Gopher!
```
