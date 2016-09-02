* Note a lot of this is taken from the excellent tour.golang.org

## Functions

Functions
A function can take zero or more arguments.

In this example, add takes two parameters of type int.

```go 
func add(x int, y int) int {
	return x + y
}
```

Notice that the type comes after the variable name.

When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

```go 
func add(x, y int) int {
	return x + y
}
```

A function can return any number of results to do this it adds braces around the return types:

```go 
func xandy(x,y int)(int,int){
    return x,y
}

```

Lets try this out: 

Create a new dir in your gopath
```
mkdir -p $GOPATH/src/github.com/YOUR_USER/functions
touch $GOPATH/src/github.com/YOUR_USER/functions/main.go
```

Add the following to main.go

```go 

package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
    //notice the assignment here. This is called a short assignment and it inferrs the type. More on this later.
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

```

Just as in javascript, functions have closure and can be passed as values:

```go 

package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

```


## Variables and infferred types

There are two main types of assignment in Go. 
1) the ``` var ``` key word and the ```:=``` short assignment.

The var keyword can be used both outside function scopes and within. 