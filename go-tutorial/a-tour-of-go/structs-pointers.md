# Structs & Pointers
- pointers
- structs

## Pointers

A pointer holds the memory address of a value. The type *T is a pointer to a T value. Its zero value is nil. The & operator generates a pointer to its operand. The * operator denotes the pointer's underlying value (dereferencing).
```bash
func main() {
    i, j := 42, 2701
    var  p *int // pointer variable declaration

    p = &i // pointer to i
    q := &j // direct assignment

    fmt.Println(*p) // dereference pointer and access value

    *q = 10 // change underlying value of a pointer
}
```

## Structs

Struct: collection of fields. Struct fields are accessed using a dot. Struct fields can be accessed through a struct pointer without explicit dereferencing. 
```bash
type Vertex struct {
    X int
    Y int
}
func main() {
    fmt.Println(Vertex{1, 2})

    v := Vertex{4, 5}
    v.X = 3
    fmt.Println(v.X)

    p := &v
    p.Y = 100
    fmt.Println(v)
}
```

struct literals: a struct literal denotes a newly allocated struct value by listing the values of its fields. Just a subset of fields can be listed using the Name: syntax. Order of named fields is irrelevant.
```bash
type Vertex struct {
    X, Y int
}

var (
    v1 = Vertex{1, 2}
    v2 = Vertex{X: 1} // Y initialised to 0
    v3 = Vertex{} // X: 0, Y: 0
    p = &Vertex{1, 2} // p has type *Vertex
)

func main() {
    fmt.Println(v1, v2, v3, p)
}
```