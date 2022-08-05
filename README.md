# Get Started Go

## Tools

- [pkg.go.dev](https://pkg.go.dev/): Go package registry
- [Tutorials](https://go.dev/doc/tutorial/)

## Command

- `go mod init`: initializes and writes a new `go.mod` file in the current directory
- `go mod edit`: edits and formats `go.mod` files, for use primarily by tools and scripts
- `go mod tidy`: ensures that the `go.mod` file matches the source code in the module

## Learning Progress

### Procedural Programming

- [x] Hello World: https://go.dev/doc/tutorial/getting-started
- [ ] Create a Go Module: https://go.dev/doc/tutorial/create-module
  - Create a module
  - Call your code from another module(`import`)
  - Return and handle an error(`errors` module)
  - Return a random greeting(slice: `[]string`, array: `[3]string`)
  - Return greetings for multiple people(`map[<key-type>]<value-type>`)
  - Add a test(unit test)
- [ ] Go predefined data structure

### Object-Oriented Programming

### functional Programming

## Best Practices

- [Effective Go](https://go.dev/doc/effective_go)
