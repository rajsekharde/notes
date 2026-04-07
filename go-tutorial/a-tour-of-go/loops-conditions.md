# Loops, Conditions
- for
- if-else
- defer

# Loops

Go has only one looping construct- the for loop. The basic for loop has three components separated by semicolons:
- The init statement. Executed before the first iteration
- The condition expression. Evaluated before each iteration
- The post statement. Executed at the end of every iteration

Init statement: Short variable declaration. Variables declared here are only visible in the local scope of the for statement.

The loop stops iterating once the boolean condition evaluated to false. 

```bash
package main

func main() {
    sum := 0
    for i := 1; i <= 10; i++ {
        sum += i
    }
    fmt.Println(sum)
}
```

The init and post statements are optional.
```bash
func main() {
    sum := 0
    for ; sum < 1000; {
        sum += sum
    }
    fmt.Println(sum)
}

// Semicolons can be dropped to turn it into a while loop
func main() {
    sum := 0
    for sum < 1000 {
        sum += sum
    }
    fmt.Println(sum)
}
```

Infinite loop: Dropping the loop condition the statement loops forever
```bash
func main() {
    for {
        // infinite loop
    }
}
```


## Conditional statement

If statements are used for conditions. Like loops, parenthesis are not mandatory, but braces are required.
```bash
func main() {
    a int = 10;
    if a > 5 {
        fmt.Println("Larger")
    }
}
```

If with a short statement: Like the for statement, the if statement can start with a short statement to execute before the condition. Variables declared in the statement are only in the scope of the if statement.
```bash
func main() {
    if v := math.Pow(5, 3); v < 250 {
        fmt.Println("False")
    }
}
```

If and else: Variables declared inside an if short statement are also available inside any of the else blocks.
```bash
func main() {
    if i := 10; i < 15 {
        fmt.Printf("%d not valid\n", v)
    } else {
        fmt.Printf("%d valid\n", v)
    }
}
```

## switch statememt

switch statement: shorter way of writing if-else statements. It runs the first case whose value is equal to the condition expression. Break statement is provided automatically, and only selected case is run. Short statement can be added at the beginning, like for and if statements. The cases need not be constants and the values involved need not be integers.
```bash
func main() {
    switch i := 10; i {
        case 5:
            fmt.Println("first")
        case 10:
            fmt.Println("second")
        case 11:
            fmt.Println("third")
        default:
            fmt.Println("last")
        }
    }
}
```

switch with no condition: same as switch true. can be a clean way to write long if-then-else chains.
```bash
func main() {
    t := time.Now()
    switch {
        case t.Hour() < 12:
            fmt.Println("Good Morning")
        case t.Hour() < 17:
            fmt.Println("Good Afternoon")
        default:
            fmt.Println("Good Evening")
    }
}
cases are executed from top to bottom.
```


## defer

Delays the execution of a statement to the end of current function. No matter how the current function ends, the defer statement gets executed. multiple defer statements are arranged in a stack (LIFO). Used to schedule something at the end of a function, like file closing, mutex unlocks, etc.

```bash
func main() {
    defer fmt.Println("World")
    fmt.Println("Hello")
}
Output:
Hello
World

func main() {
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")
}
Output:
3
2
1
```