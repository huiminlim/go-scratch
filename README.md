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
