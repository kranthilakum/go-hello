`go.mod` file is the root of dependency management in Go. All the modules which are needed or to be used in the project are maintained in go.mod file.



Import an external Go module:
- `go get <module-name>` e.g. `go get rsc.io/quote`
- `go mod tidy` - adds any missing modules necessary to build the current module's
packages and dependencies

Run your code:
- `go run .`
- `go run <filename>.go`

Get some help:
- `go help` or e.g. `go help mod init`



