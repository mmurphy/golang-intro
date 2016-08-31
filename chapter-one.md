## Installing Go for your environment

Downloads can be found here: [golang.org downloads](https://golang.org/dl/)

Installation instruction are here [Go installation](https://golang.org/doc/install) I will outline the main ones below.

- Mac OSX there is a package installer under the downloads
- Linux there is a tar.gz under the downloads.

```
  Download the archive and extract it into /usr/local, creating a Go tree in /usr/local/go. For example:
  tar -C /usr/local -xzf go1.7.linux-amd64.tar.gz
```

```
  Add /usr/local/go/bin to the PATH environment variable. You can do this by adding this line to your /etc/profile (for a system-wide installation) or $HOME/.profile:
```  

## Understanding the $GOPATH env var

The $GOPATH env var tells the Go tools where your workspace is. I would normally set this in my bash profile as well.
When you install packages and build binaries, the Go tools will default to the $GOPATH/src followed by a fully resolovable path in order to resolve depenencies. (more on this shortly)

Mine is currently set to /Users/kelly/work/go

```
  export $GOPATH=$HOME/work

```

The $GOPATH env var works like the $PATH var where you can have multiple locations separated by a ``` : ```

## Setting up and verifying your worksapce

Once we have the $GOPATH env var setup, we need to create a directory structure for our projects that Go understands.
Go's best practice is to have packages / projects as resolvable paths. An example would be: ``` github.com/golang/lint/golint ```
This is a package on github if you browe to ``` https://github.com/golang/lint ``` you can see the code there. The go get command expects a resovable path backed by one of the common CVS systems (git,mercurial etc)
When you run the go get command, it will checkout the repo at $GOPATH/src/github.com/golang/lint/golint and you refer to it in your files as an import. The import path tells the Go compiler where to find the source for your projects.
```
  import "github.com/golang/lint/golint"
```

So to be good Gophers we will also follow this best practice. Before we finish lets complete our workspace by creating the following dirs

```
  mkdir -p $GOPATH/src // where go get will store the source of packages. Our code will also live here under our own namespace.
  mkdir -p $GOPATH/bin // where executables are install when you run go install
  mkdir -p $GOPATH/pkg // where package objects are stored things like .a files

```  

## Installing some common tools
 - golint  ``` go get github.com/golang/lint/golint ```  
 - goimports ``` go get  golang.org/x/tools/cmd/goimports ```

golint looks for common code style. Things such as documenting Exported functions.
goimports formats your code correctly and also removes any unused imports. It will also try to resolve imports that haven't been added to the imports definition.

What we just did and how go get works. The go get command takes the path to the package and expects a version controlled and accessible path. It then clones this source code into your
$GOPATH/src directory ready for you to use in import paths as part of your projects.

```
cd $GOPATH/src/github
ls
```

## Some common Go commands that you should know

1) ``` go build ``` (see below) this will build and compile your code into a binary

2) ``` go install ``` (see below) this will do the same as go build except drop it into your $GOPATH/bin dir

3) ``` go vet ``` This will check your code for common coding errors

4) ``` go test ``` (see below) This will run the test files in the given package

5) ``` go get ``` Pulls a dependency into your $GOPATH

## Create your first program
Create the following directory
```
  mkdir -p $GOPATH/src/github.com/<your_user>/hello
  cd $GOPATH/src/github.com/<your_user>/hello
  vi main.go
```
Add the following to main.go

```
package main
import(
    "fmt"
)

func main(){
    fmt.Println("hello world")
}

```
So what have we done here. the package indicates the name of the package we are in and the main function is similar to main for other languages. It is the entrypoint.

functions are defined as func.

We are importing the fmt package from the stdlib and then using it to print to stdout.

```

go build .

```

The go build command will compile the source into a binary and drop it in the current directory.

```
./hello

```

This should output ``` hello world ```

```
rm ./hello

go install .

hello

```
The go install command installs the binary to the $GOPATH/bin dir which is on your path so you can use the program without the ./

Finally open the main.go file and change the formatting to be terrible...

```
go fmt .

```

Open main.go and it will be correctly formatted again.



## Create the first test for your program

The convention in Go is to put your tests alongside the code that it tests rather than in a different directory somewhere. This has the advantage of giving a very visible
way of seeing what has tests and what doesn't. All test files in Go are named _test.go anything named that way will not be added to the final binary.


reopen main.go and change it to look like the following :

```
package main
import(
    "fmt"
)

//This is a public function exposed from the package. A exported functions must start with an Uppercase letter.
// the return type of this function is a string as shown by the definition
func HelloWorld()string{
    return "hello world"
}

func main(){
    fmt.Println(HelloWorld())
}

```

open main_test.go

```
  vi main_test.go
```

Add the following code

```
pacakge main

import(
    "testing"
)

//All test functions accept an argument of a pointer to testing.T
func Test_HelloWorld(t *testing.T){
    val := HelloWorld()
    if (val != "hello world"){
        t.Fail()
    }
}

```

Here we have imported the stdlib [testing package](https://golang.org/pkg/testing/) and defined a test in the standard way that accepts a pointer to a type of t.Testing.
Then we call our HelloWorld function and check that the value returned matches what we expected.

To run this test run the following command

```
go test

```

To get more verbose output add -v

```
go test -v

```

## Viewing docs about stdlib packages or any package

You can google the package and read the online docs at :

```
https://golang.org/pkg/
```

You can also use the godoc command any any package. For example run the following command to see the docs for the testing pkg

```
godoc testing

```


## Editors and integration
Most main editors have some support for Go

- Visual Studio Code Has a good Go plugin [download](https://code.visualstudio.com)
- Intellij Also has a good Go plugin [download](https://www.jetbrains.com/idea/)
- Sublime Text3 Takes a bit of setting up but works quite well [tutorial](https://www.wolfe.id.au/2015/03/05/using-sublime-text-for-go-development/)
- Atom plugin [download](https://atom.io/packages/go-plus)

## Optional Homework

If you wanted to learn more (I will be going through more next week), now that you have your
environment setup, I would recommend looking at [A Tour of Go](https://tour.golang.org/welcome/1) that essentially start where we have left off.
