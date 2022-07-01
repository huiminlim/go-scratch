# Go Scratchpad

This repo serves as a scratchpad for code snippets written in Go for learning.

## Getting Started

### Packages

Packages are similar to libraries to be used in code.

To load a package, use the `import` keyword.

```go
// Import a formatting package
import "fmt"

// Batch import multiple packages
import (
    "fmt"
    "temp"
)

// Importing all packages and renaming them
import (
    . "name"                // Import all
    pkg "package\name"      // Import and rename
)
```

### Modules

A `go.mod` file must exist at the root directory of all projects.

It will contain formation about the project, e.g. dependencies, Go version, package info.

For example:

```go
module example.com/practice

go 1.17

require (
    github.com/alexflint/go-arg v1.4.2
)
```

To create a Go module file, issue the `go mod init` command

Info is [here](https://go.dev/doc/modules/gomod-ref).

## Introduction

### Variables

Variables can be created using the `var` keyword.

The type may be specified, else type inference from the value given will be done.

```go
// Create variable with var
var example = 3

// Type annotation used to assign variable of type int
var example int = 3

// Type annotation of uninitialized variable
var example int
example = 3

// Compound creation
var a, b, c = 1, 2, "hello"

// Block creation
var (
    a int = 1
    b int = 2
    c = "sample"
)
```

The create and assign shorthand `:=` can be used also.

```go
// Create and assign shorthand
example := 4
a, b := 1, "sample"

 // Copies by value b to c
c := b
```

The comma ok idiom allows recreation and reassignment of variables in the same scope without error. Convenient for variables often repeated.

```go
x, err := 1, "none"
y, err := 2, "none"
z, err := 3, "none"
```

Constant variables may be created using `const` keyword.

```go
// Convention to use capital case for words
const MaxSpeed = 40;
```

### Functions

The syntax of a function is:

```go
// Use camelCase
func name(param1 type, param2 type) type {
    // Body
}
```

Multiple return values can be supported.

```go
// Multiple return function
func many() (int, int, int) {
    return 1, 2, 3
}
```

### If-Else

The syntax for if-else:

```go
if age >= 21 {
    // Do something
} else if age < 21 && age >= 15 {
    // Do something
} else {
    // Do something
}
```

### Switch

The syntax for switch statements:

```go
// Variable assignment may be done in switch
switch x := check(3) {
    // Multiple cases can share the same action
    case 1, 2, 3:
        // Do something

    // Fallthrough keyword allows bottom cases to be checked
    // and evaluated also
    case 100:
        fallthrough

    // Checks may be done here
    case x > 233:
        // Do something

    default:
        // Do something
}
```

### Loops

The syntax for for-loops:

```go
for i := 0; i < 10; i++{
    // Do something
}
```

Loop counter can also be adjusted outside of the for statement:

```go
for i < 10;{
    // Do something
    i++
}
```

Infinite loops may be created using `for` also.

```go
for {
    // Do something
    if done {
        // Exit loop
        break
    }
    else {
        // Move on to next iteration of loop
        continue
    }
    // Do other things
}
```

## Structures

A structure can be defined with the following syntax:

```go
// Defining a struct named "Sample"
type Sample struct {
    field string
    a, b int
}

// Instantiating a sample
data := Sample{"word", 1, 2}

// Alternative method
data2 := Sample {
    field: "word",
    a: 1,
    b: 2
}

// Using default values for uninitialized
data3 := Sample{}

// Defining only certain fields
data := Sample{a: 5}
```

Inline anonymous structures can be defined also.

```go
// All fields must be created and assigned upon struct definition
sample := struct{
    field string
    a, b int
}{
    "hello",
    1, 2,
}
```
