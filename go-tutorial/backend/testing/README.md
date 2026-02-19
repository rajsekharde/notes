# Testing in Go

Rules:
- File naming: Test file should end in _test.go (e.g., server_test.go)
- Function naming: Test function must start with Test followed by a capitalized word (e.g., TestLogin)
- Signature: The function must accept exactly one argument: (t *testing.T)

Simple Unit Test
```bash
// math.go
func Add(a, b int) int { return a + b }

// math_test.go
func TestAdd(t *testing.T) {
    got := Add(2, 3)
    want := 5
    if got != want {
        t.Errorf("Add(2, 3) = %d; want %d", got, want)
    }
}
```

Running tests:
```bash
// Run all tests in current directory
go test

// Run a specific test function
go test -run TestLogin

// Get a detailed result
go test -v
go test -v -run TestLogin

flag: -count=1 // Bypass Cache

```

Multiple test cases- Use Table-Driven testing
- Define a table of inputs and outputs and loop through them
- On running go test -v, Go executes the main function and iterates through the table. If one row fails, it reports it and moves to the next row immediately


Using t.Parallel():

If you have 100 edge cases and they don't depend on each other, you can make them run incredibly fast. Add t.Parallel() inside the t.Run block:
```bash
for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        t.Parallel() // Runs all subtests at the same time!
        // ... test logic ...
    })
}
```