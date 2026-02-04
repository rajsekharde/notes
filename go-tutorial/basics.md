# Go Basics

## Project Creation and first program:

```bash

# Create a new project directory:
mkdir hello-go
cd hello-go

# Enable dependency tracking:
go mod init example.com/hello-go

# Create a new Go file:
touch hello.go
```
### Write code for first program in hello.go:
```bash
package main # Declare a main package

import "fmt" # Import fmt package

func main() { # main function
    fmt.Println("Hello World!")
}
```

### Run the program:
```bash
go run hello.go
```

### Compile and run the program:
```bash
go build hello.go # Creates a compiled binary- hello
./hello
```


Package: Groups functions. Made up of all the files in the same directory

fmt package: Contains functions for formatting text, including printing to console

main function: Executes by default when main package is run. Entrypoint to a program

### Random info:

No line should start with '{'.
Correct way to use braces:
name {
    # statements
}

Go implicitly adds a semicolon at the end of each line
