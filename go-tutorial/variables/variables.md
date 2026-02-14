# Variables in Go

Created using 'var' keyword or ':=' sign

var <variable-name> <type> = <value>

<variable-name> := <value>

```bash

var s1 string = "Hello" # String type is defined

var s2 = "World" # Type is inferred

message := "Hello" # Type is inferred

```

All variables are initialised. If no value is assigned, default values of the data
type are assigned.

Default values:
string: ""
int: 0
bool: false


var - can be used inside and outside of functions. Variable declaration and value
assignment can be done separately

:= - can only be done inside functions. Variable declaration and value assignment
cannot be done separately


Multiple variable declaration:
```bash
var a, b, c, d int = 1, 2, 3, 4

var a, b = 6, "Hello"

c, d := 7, "World!"

var (
    a int
    b int = 1
    c string = "Hello"
)
```

## Constants

Used when a variable should have a fixed value that cannot be changed.

const <name> <type> = <value> (Value must be assigned on declaration)

const PI = 3.14

const g float = 9.8