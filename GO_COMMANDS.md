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

To change simple go program into a go module. This command initializes a go.mod file.

### go get

to download all the modules

#### go get -u

update packages to latest minor version

#### go get -u=patch

update package to latest patch version

#### go get -u github.com/gorilla/mux

download some modules
this will make entry in the go.mod file its similar to package.json file in nodejs
go.sum is similar to package.lock file, Note go.sum keep history of all the modules which are once used, hence its bit different from package.lock as its track of old module, in case they are used in future.
command is similar to npm install
NOTE on how packages and projects are managed
Historically, all Go code was stored in one giant monorepo, because that’s how Google organizes their codebase internally and that took its toll on the design of the language.
Go Modules is somewhat of a departure from this approach. You’re no longer required to keep all your projects under $GOPATH.
However, technically all your downloaded dependencies are still placed under $GOPATH/pkg/mod. If you use Docker containers when developing stuff locally, this may become an issue because dependencies are stored outside of project path (shared with a host filesystem via a bind-mounted volume). By default, they are simply not visible from your IDE.

### go mod tidy

in go.mod file +incompatible suffix is added for all packages that not opted in to Go
Modules yet or violate its versioning guidelines

as we didn't import the package anywhere in our project it, was marked as //indirect, tidy this up using go mod tidy
the below command will either prune the unused modules or remove //indirect comment from go.mod file
the below command is IMPORTANT for every release, as its add any OS dependencies needed for other combinations of OS architecture and build tags.
