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

### Array

Arrays store elements of the same kind of data type.

Arrays are fixed size and start with index 0;

An array can be created this way:

```go
// Creating array of int uninitialized
var myArray [3]int

// Creating an array with specified number of default values
myArray := [3]int{7, 8, 9}

// Creating an array with some number of default values
myArray := [...]int{1, 2, 3, 4}

// Creating an array with 3 initialized values and 1 uninitialized
myArray := [4]int{7, 8, 9}
```

### Slices

Slices are "views" of arrays, or subarray.

A slice can be created with an array underlying it.

```go
// Slice when there is no number of elements in square bracket
mySlice := []int{1, 2, 3}

// Accessing element in slice
a := mySlice[0]
```

Slices may be created from arrays.

```go
// Creating a slice from an array
myArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// Creating slice from entire array
slice := myArray[:]

// Creating slice from start index to end
slice2 := myArray[1:]

// Creating slice from element 1 to idx - 1
slice3 := myArray[1:idx-1]
```

Slices may be appended to with additional elements or other slices.

```go
slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// Appending with extra elements
// and reassign to existing slice
slice = append(slice, 1, 2, 3)

// Appending slice with slices
// and reassign to existing slice
part1 := []int{0, 1}
part2 := []int{2, 3}
// 3 dots required to unwrap the previous slice into numbers
slice = append(slice, part1, part2...)
```

Preallocation of slices capacity may be done. For performance to cut computation time.

```go
slice := make([]int, 10)
```

Multidimensional slices can also be created.

```go
board := [][]string{
    []string {"_", "_", "_"},
    {"_", "_", "_"},
    {"_", "_", "_"},
}
board[0][0] = "x"
```

### Ranges

Ranges are largely useful to access each element in a string.

For example:

```go
a := "Hello"
for _, ch := range a {
    fmt.Printf("%q ", ch)
}
```

### Maps

A map can be created for easy key-value pair access.

```go
// Map with string key and integer value
// Uninitialized map
myMap := make(map[string]int)

// Initialize map
myMap2 :=  map[string]int{
    "item 1": 1,
    "item 2": 2,
    "item 3": 3,
}
```

There are some map operations - insert, read, delete, check existence.

```go
// Creating uninitialized map
myMap := make(map[string]int)

// Insert key-pair
myMap["Favourite number"] = 5

// Read value from key
fav := myMap["Favourite number"]
missing := myMap["age"] // returns default

// Delete key
delete(myMap, "Favourite number")

// Check existence
price, found := myMap["price"]
if !found{
    fmt.Println("price not found")
    return
}
```

Iterating between maps can be done with range.

```go
// Initialize map
myMap2 :=  map[string]int{
    "item 1": 1,
    "item 2": 2,
    "item 3": 3,
}

// Iterating maps using range
for key, value := range myMap {
    // print
}
```

### Pointers

Pointers are indicated with a `*`.

To create a pointer from a variable, use the `&`.

```go
// Creating a pointer from variable
value := 10
var valuePtr *int // creating uninitialized pointer
valuePtr = &value

// Create and assign pointer variable
value2 := 10
valuePtr2 := &value2
```

The `*` dereferences the pointer to the value at the memory.

```go
func increment(x *int){
    *x += 1
}

i := 1
increment(&i)
// x is now 2
```

## Idiomatic Go

### Receiver Functions

Receiver functions are functions to operate on structs directly.

There are **Pointer** and **Value** receiver functions.

For example, **Pointer** receiver function:

```go
// struct to use
type Coordinate struct {
    x, y int
}

// Pointer Receiver function to shift x, y by some number
func (coord *Coordinate) shiftBy(x, y int){
    coord.x += x
    coord.y += y
}

// Execute the receiver function
coord := Coordinate{5, 5}
coord.shiftBy(1, 1) // (6, 6)
```

And **Value** receiver function.

```go
// struct to use
type Coordinate struct {
    x, y int
}

// Value Receiver function to return by value another struct
func (c Coordinate) Dist(other Coordinate) Coordinate {
    return Coordinate{c.x - other.x, c.y - other.y}
}

first := Coordinate{5, 5}
second := Coordinate{1, 1}
distance := first.Dist(second) // (4, 4)
```

### Iota

The `iota` can automatically assign values to constants, similar to enumerations.

For example:

```go
// Long form
const (
    l0 = iota   // 0
    l1 = iota   // 1
)

// Short form
const (
    s0 = iota   // 0
    s1          // 1
)
```

Skipping values and starting from another default start may be done also.

```go
// Skipping a value unrequired
const (
    s0 = iota   // 0
    _           // 1 (skipped)
    _           // 2 (skipped)
    s3          // 3
)

// Starting at 3
const (
    i3 = iota + 3   // 3
    i4
    i5
)
```

### Variadics

Variadic functions accept multiple input parameters into the function, denoted by `...`.

For example:

```go
// Variadic functions accept unknown number of input parameters
// here, "nums" are slices
func sum(nums ...int) int {
    sum := 0
    for _, n := range nums {
        sum += n
    }
    return sum
}
```

### Text Formatting

Using `fmt.Printf` allows the text to be formatted according to verbs.

| Verb | Description           |
| ---- | --------------------- |
| `%v` | Default               |
| `%t` | `true` or `false`     |
| `%c` | Character from a rune |
| `%X` | Hex                   |
| `%U` | Unicode format        |
| `%e` | Scientific notation   |

`fmt.Sprintf` returns a formatted string, while `fmt.Fprintf` writes to a file stream.

### `init` Function

With an `init()` function, it will be run before the `main()` function and commonly used to perform initialization.

All packages can have an `init()` function, and all the `init()` functions will run before the `main()` actually runs.

### Testing

To create tests for Go packages, import the `testing` package.

Test file to accompany a Go file is as such: `sample.go` and `sample_test.go`.

Some test functions include: `Fail()`, `Errorf(string)`, `FailNow()`, `Fatalf(string)`, `Logf()`.

Test tables can be used to test more than 1 set of test data to drive the tests.

For example:

```go
func TestAddQueue(t *testing.T){ // Required function signature for test
    q := New(3)
    for i := 0; i < 3; i++ {
        if len(q.items) != i {
            t.Errorf("Incorrect queue element count: %v, want %v", len(q.items), i)
        }
        if !q.Append(i) {
            t.Errorf("Failed to append item %v to queue", i)
        }
    }
}
```

Use command `go test -v ./demo/testing` to run test and show all tests ran.
