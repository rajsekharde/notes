# VIM CLI editor tutorial

Creating a file and opening it using vim:
```bash
vim hello.c # Creates a file- hello.c in current directory and opens it
<ESC> # Set to normal mode. Not necessary as mode is set to Normal by default
:x! # Save the file
```

Quitting Vim CLI:
```bash
:q <Enter> # Quit
:q! # Discard changes and quit
:wq # Save changes and quit
```

## Workflow for creating a C file and running it:
```bash
vim hello.c # Create hello.c and open it in vim

<ESC> # Setsmode to Normal

<i> # Set mode to insert. Can write code now.

# Write code

<ESC> # Set mode to Normal

<:wq> # Set mode to Command Line, save file and exit
```

Modes:
- Normal Mode- Default: for navigation and simple editing. Command: <ESC>
- Insert: For explicitly inserting and modifying text. Command: <i>
- Command line: For operations like saving, quitting etc. Command: <:>

General Commands:
```bash
:set number # Enable line numbers
:2 # Jump to line no.2

```


