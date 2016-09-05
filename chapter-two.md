* Note a lot of this is taken from the excellent tour.golang.org
<<<<<<< HEAD
* Some information was taken from http://www.golangbootcamp.com/

## Exported Names

In Golang a packages publicly accessible properties all begin with a capital letter. From within a package namespace you can
refer to private functions and variables but from outside the package, you can only access the things exported from that package.
Think public and private key word in Java.
=======
>>>>>>> e37a23817d406e31c394a8c539f0df8ca16ee6ca

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

<<<<<<< HEAD

There are two main types of assignment in Go. 
1) the ``` var ``` key word and the ```:=``` short assignment.

#### VAR:

The var keyword can be used both outside function scopes and within whereas the short assignment can only be used within the scope of a function.

```go 
package main 

import(
	"fmt"
)

var MyGlobal string = "hello world"


func main (){
	msg := "hello world again"
	fmt.Println(MyGlobal)
	fmt.Println(msg)
}

``` 

You can run this by using the handy ``` go run ``` command

```go

go run main.go 

```

In the above example we set the type, but if an initializer is present, the type can be omitted, the variable will take the type of the initializer (inferred typing).
The below outlines the different ways you can use the var keyword:

```go 
package main

import "fmt"

func main(){
	//as a list of vars
  var(
	  one = 1 //infferred type 
	  two = 2 // infferred type
	  three int // default value of type int 0
  )
  //individual declarations
  var four int // default value of type int 0
  var five = 5 // inferred type 

  // multiple inline initialization is also fine
  var name, location, age = "golang", "is great", 42 

  fmt.Printf("one %d two %d three %d four %d five %d", one,two,three,four,five)

}

```
The above will ouput 

```
one 1 two 2 three 0 four 0 five 5
```

Notice that three and four have a value of 0 that is because the default value for int is 0

#### Short Assignment:

Inside a function, the := short assignment statement can be used.
Here are some example of short assignment:

```go 

func main (){
	msg := "hello world again"
	name,age,fingers := "john",27,10
}
```

A variable can contain any type, including functions:

```go 
package main 

import "fmt"

func main(){
	fn := func (){
		fmt.Println("hello world")
	}
	fn()
}
``` go 
These things allow go to feel quite familar to javascript developers.

=======
There are two main types of assignment in Go. 
1) the ``` var ``` key word and the ```:=``` short assignment.

The var keyword can be used both outside function scopes and within. 
>>>>>>> e37a23817d406e31c394a8c539f0df8ca16ee6ca
