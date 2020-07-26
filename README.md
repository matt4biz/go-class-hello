[![Run on Repl.it](https://repl.it/badge/github/matt4biz/go-class-hello)](https://repl.it/github/matt4biz/go-class-hello)

# Go class: Hello example (with unit test)
This example builds on the perennial first program, "Hello, world!"

1. We have a function that given some names, generates the "Hello, ..." string (using "world" if no names are given)

	say.go

1. We have a unit test which checks 0, 1, and 2 or more names as inputs to the function

	say_test.go

1. We use this function together with the command line arguments (except the first, which is the program name) to produce a result

	cmd/hello.go

## Running and testing

```shell
$ go run ./cmd
Hello, world!

$ go run ./cmd class
Hello, class!

$ go run ./cmd class world
Hello, class, world!

$ go test ./...
ok  	hello	0.103s
?   	hello/cmd	[no test files]
```
