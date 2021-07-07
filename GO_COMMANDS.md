# Go Commands

## go help

displays all the options on go cli

## go version

displays the version of the go installed

## go env

displays all the go environment variable.

## go get <>

install a package, similar to pip install or npm install

## go install

- creates a binary of the project. Sometime `go mod init` is requried else it may throw error that main program is not found.
- it will also moves everything to $GOPATH/bin

## go build

- this leaves the result in the current working directory, instead of moving it to $GOPATH/bin folder.

## go run

```sh
#to compile and run the code
go run main.go
```

## go mod

### go mod init

initialize a go module. This is recommended if using VCS(like github or gitlab)

## go mod tidy

It cleans up ununsed dependencies or adds missing dependencies.