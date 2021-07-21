# MODULES

## HOW TO INSTALL MODULES IN GO

```sh
# initialize a go module, in the current directory
go mod init

# download some modules
# this will make entry in the go.mod file its similar to package.json file in nodejs
# go.sum is similar to package.lock file, Note go.sum keep history of all the modules which are once used, hence its bit different from package.json as its track of old module, in case they are used in future.
# command is similar to npm install
## NOTE on how packages and projects are managed
## Historically, all Go code was stored in one giant monorepo, because that’s how Google organizes their codebase internally and that took its toll on the design of the language.
# Go Modules is somewhat of a departure from this approach. You’re no longer required to keep all your projects under $GOPATH.
# However, technically all your downloaded dependencies are still placed under $GOPATH/pkg/mod. If you use Docker containers when developing stuff locally, this may become an issue because dependencies are stored outside of project path (shared with a host filesystem via a bind-mounted volume). By default, they are simply not visible from your IDE.

go get -u github.com/gorilla/mux

# to download all the modules
go get

# in go.mod file +incompatible suffix is added for all packages that not opted in to Go
# Modules yet or violate its versioning guidelines

# as we didn't import the package anywhere in our project it, was marked as //indirect, tidy this up using go mod tidy
# the below command will either prune the unused modules or remove //indirect comment from go.mod file
# the below command is IMPORTANT for every release, as its add any OS dependencies needed for other combinations of OS architecture and build tags.
go mod tidy

# update packages to latest minor version
go get -u

# update package to latest patch version
go get -u=patch
```
