## Go TCP Client

The approach I took was to go through the following steps
* Parse flags first
* Parse arguments next, and ignore any flags that were described in the problem
* If SSL is true then make sure to create correct config for go connection
* Loop until server sends a solution back, the solution is anything that is
  returned and is only 2 strings

I also handle errors with handleError(). Since go doesn't throw errors, I had
to check for errors after each function.

### Running source code
`> go run main.go [domain] [neuid]`

### Building source code
`> make`

### Running tests
`> go test -v ./test/*`
