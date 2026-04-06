# Basics

## Packages

Every Go program is made of packages. Programs start running in package main. Package name is same as the element of the import path.

```bash
package main

import (
    "fmt",
    "math/rand"
)

func main() {
    fmt.Println("Number = ", rand.Intn(10))
}
```

## Imports
```bash
Both valid:

import "fmt"
import "math"

import (
    "fmt",
    "math"
)
```

## Exported Names

Names beginning with capital letter are exported. Files importing a package can only access the exported elements of the package. Other elements are inaccessible from outside the package.

