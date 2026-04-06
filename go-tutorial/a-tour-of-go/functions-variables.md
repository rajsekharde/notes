# Functions & Variables

In Go, a function can take zero or more arguments and can return zero or more results.

```bash
package main

func add(x int, y int) int {
    return x + y
}

// Two or more consecutive parameters with same type can share the type
func mul(x, y, int) int {
    return x * y
}

// A Function can return any number of results
func swap(x, y string) (string, string) {
    return y, x
}

// Named return values
// Using a return statement without any variables - naked return
// Naked return returns named variables declared with return types
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func main() {
    fmt.Println(add(45, 55))
}
```

# Variables
- var statement declares a list of variables
- type comes after variable name
- Scope- can be at package level or function level
- var declaration can include initializers, one per variable
- if initializer is present, type can be omitted. Variable will take the type of initializer
- short assignment statement ':=' can be used instead of var inside functions
- := not valid at package level

```bash
func main() {
    // var with initializers
    var i, j int = 1, 2

    // short assignment
    k := 3
    a, b, c := true, false, "no!"
}
```

Zero values: Variables declared without an explicit value are given their zero value.
```bash
Zero values:
0 for numeric types
false for boolean type
"" for strings

var i int // 0
var f float // 0
var b bool // false
var s string // ""
```

## Basic data types
```bash
bool

string

int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64, uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32, float64

complex64, complex128


int, uint, uintptr - 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems
```

Type conversions: T(v) converts the value v to the type T
```bash
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

or

i := 42
f := float64(i)
u := uint(f)
```

Constants: Declared like variables, but with the const keyword. Can be character, string, boolean, or numeric values. Cannot be declared using the := syntax.
```bash
const Pi = 3.14

// Numeric constants - high-precision values
const (
    Big = 1 << 100 // 1 followed by 100 zeros
    Small = Big >> 99 // 2
)

func main() {
    const s = "hello"
    const t = true
}
```