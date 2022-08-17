# Get Started Go

## Tools

- [pkg.go.dev](https://pkg.go.dev/): Go package registry
- [Tutorials](https://go.dev/doc/tutorial/)

## Command

- `go mod init`: initializes and writes a new `go.mod` file in the current directory
- `go mod edit`: edits and formats `go.mod` files, for use primarily by tools and scripts
- `go mod tidy`: ensures that the `go.mod` file matches the source code in the module
- `go test -v`: executes test functions (whose names begin with `Test`) in test files (whose names end with `_test.go`)
- `go build`: compiles the packages, along with the dependencies but it **doesn't install the build result**.
- `go install`: compiles and installs the packages
  - `go env -w GOBIN=/path/to/your/bin`: changes the install target by setting the `GOBIN` variable using `go env`
- `go work init <dir>`: creates a `go.work` file for the workspace containing the modules in `<dir>` directory

## Learning Progress

### Procedural Programming

- [x] Hello World: https://go.dev/doc/tutorial/getting-started
- [x] Create a Go Module: https://go.dev/doc/tutorial/create-module
  - Create a module
  - Call your code from another module(`import`)
  - Return and handle an error(`errors` module)
  - Return a random greeting(slice: `[]string`, array: `[3]string`)
  - Return greetings for multiple people(`map[<key-type>]<value-type>`)
  - Add a test(unit test)
  - Compile and install the application
- [x] Getting started with multi-module [workspaces](https://go.dev/ref/mod#workspaces): https://go.dev/doc/tutorial/workspaces
- [ ] A Tour of Go: https://go.dev/tour/basics/
  - Basic Type: https://go.dev/tour/basics/11
  - Flow Control: https://go.dev/tour/flowcontrol/12
- [ ] Go predefined data structure
- [ ] Go Web Framework - [Gin](https://gin-gonic.com/docs/): https://go.dev/doc/tutorial/web-service-gin

### Object-Oriented Programming

### functional Programming

## Best Practices

- [Effective Go](https://go.dev/doc/effective_go)
