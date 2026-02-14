# Output Functions in Go

Print() - Default format
fmt.Print("Hello", 1, "\n")

Println() - Default format + space between arguments + \n at end

Printf() - formats arguments based on given formatting verbs
fmt.Printf("i has value: %v and type: %T\n", i, i)

%v - value, %T - type

Integer formatting verbs:
%b - base 2
%d - base 10
%o - base 8
%x - base 16, lower case
%X - base 16, upper case
%#v - Go-syntax format

Float formatting verbs:
%e - scientific notation
%.2f - default width, precision 2
%6.2f - width 6, precision 2