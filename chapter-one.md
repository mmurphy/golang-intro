## Installing Go for your environment

You might use your package manager to install Go or install manually from the
official release binaries. Either way should be easy.

Because the `go` tool is able to download code from remote repositories, it is
often useful to have installed clients for the various supported version control
systems. At the time of this writing, they are
[documented](https://golang.org/cmd/go/#hdr-Remote_import_paths) as: Bazaar,
Git, Mercurial and Subversion.

They are not mandatory requirements, you might go a long way with just having
`git` installed, and installing others later as required. In any case, to
install all of them:

```
dnf install bzr git mercurial subversion      # Fedora
yum install bzr git mercurial subversion      # RHEL / CentOS
brew install bazaar git mercurial subversion  # macOS with Homebrew
```

### Installation using a package manager

```
dnf install golang  # Fedora
yum install golang  # RHEL / CentOS
brew install go     # macOS with Homebrew
```

### Installation from official binary releases

1. Download an archive for the desired version and platform from
https://golang.org/dl/.
2. Follow the [installation
instructions](https://golang.org/doc/install). The main steps are outlined
below.

For macOS, execute the package installer from the downloads page.

For Linux, download a `.tar.gz` archive and extract it into `/usr/local`,
creating a Go tree in `/usr/local/go`. For example:

```
tar -C /usr/local -xzf go1.7.linux-amd64.tar.gz
```

Finally, add `/usr/local/go/bin` to the PATH environment variable. You can do
this by adding this line to your `/etc/profile` (for a system-wide installation)
or `$HOME/.profile` (or `~/.bashrc`, `~/.bash_profile`, ...):

```
export PATH=$PATH:/usr/local/go/bin
```


## Understanding the GOPATH environment variable

The GOPATH environment variable lists places to look for Go code. It defines
your workspace. It is likely the only environment variable you'll need to set
when developing Go code.

Official documentation:

- [How to Write Go Code: The GOPATH environment variable](https://golang.org/doc/code.html#GOPATH)
- [Command go: GOPATH environment variable](https://golang.org/cmd/go/#hdr-GOPATH_environment_variable)

Normally the GOPATH is set in your shell profile (one of `~/.bashrc`,
`~/.bash_profile`, etc).

When you install packages and build binaries, the Go tool will look for source
code in `$GOPATH/src/` followed by a package import path in order to resolve
dependencies. More on this shortly.

The GOPATH works like the PATH environment variable, where you can have multiple
locations separated by a `:` (`;` on Windows).

Some people prefer to use a single path workspace, like
`GOPATH=/Users/kelly/work/go`. Others use multiple paths, like
`GOPATH=/home/rodolfo/.go-vendor:/home/rodolfo/Dropbox/go:/home/rodolfo/openshift`.
When using multiple paths, the Go tool will download and install new packages to
the first path in the list, while preserving the same path when building
existing source code.

Make sure to have a line defining your GOPATH in your `~/.bashrc` or equivalent:

```
export $GOPATH=$HOME/go
```

---
**Pro tip**: include also a line to add GOPATH/bin to your PATH, so that you can
easily run built and installed binaries:

```
# Include every ./bin directory from GOPATHs into PATH
export PATH="$PATH:${GOPATH//://bin:}/bin"
```
---


## Setting up and verifying your workspace

Once we have the GOPATH setup, we need to create a directory structure for our
projects that Go understands.

Go's best practice is to have packages / projects as resolvable paths. An
example would be: `github.com/golang/lint/golint`. This is a package on GitHub
such that if you browse to https://github.com/golang/lint you can see the code
there.

The `go get` command expects a resolvable path backed by one of the common VCS
systems (Git, Mercurial, etc). When you run the `go get` command, it will
checkout the repository at `$GOPATH/src/github.com/golang/lint/golint` and you
can then refer to it in an import statement. The import path tells the Go
compiler where to find the source for a certain package.

An `example.go` file could look like:

```go
package main

import "github.com/golang/lint/golint"

// ...
```

So to be good Gophers we will also follow this best practice. Before we finish
let's complete our workspace by creating the following directories:

```
mkdir -p $GOPATH/src  # Where go get will store the source of packages. Our code will also live here under our own namespace.
mkdir -p $GOPATH/bin  # Where executables are installed when you run go install or go get.
mkdir -p $GOPATH/pkg  # Where package objects are stored, e.g., .a files.
```

## Some common Go commands that you should know

1. `go build` (see below): builds / compiles your code into a binary or library
2. `go install` (see below): the same as `go build`, plus installs binaries into
   $GOPATH/bin and libraries to $GOPATH/pkg
3. `go get`: downloads source code and dependencies, then perform a `go install`
4. `go test` (see below): compile and run tests in the given package(s)
5. `go vet`: check code for common coding errors


## Installing some common tools

Use `go get` to download packages and their dependencies, then install them:

```
go get github.com/golang/lint/golint \
       golang.org/x/tools/cmd/goimports
```

Golint looks for common code style mistakes. Golint makes suggestions for many
of the mechanically checkable items listed in [Effective
Go](https://golang.org/doc/effective_go.html) and the [CodeReviewComments wiki
page](https://golang.org/wiki/CodeReviewComments).

`goimports` formats your code correctly (it's a drop-in replacement for `gofmt`)
and also removes any unused imports. It will also try to resolve imports that
haven't been added to the imports definition.

The `go get` command takes the path to the package and expects a version
controlled and accessible URL. It then clones the source code into your
$GOPATH/src directory ready for you to use in import paths as part of your
projects.

```
# recursively list directories in the first path of GOPATH, with limited depth:
find ${GOPATH%%:*}/src/ -type d -maxdepth 2 | less
```


## Create your first program

Create a directory in your workspace:

```
mkdir -p $GOPATH/src/github.com/<your_user>/hello
cd $GOPATH/src/github.com/<your_user>/hello
touch main.go
```

Open your favorite editor, and add the following to `main.go`:

```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```

So what have we done here? The `package` statement indicates the name of the
package we are in and the `main` function is similar to main for other languages.
It is the entrypoint of an executable program.

Functions are defined with the `func` keyword.

We are importing the `fmt` package from the stdlib and then using it to print to
*stdout*.

Now run:

```
go build
```

The `go build` command will compile the source of the package in the current
directory and drop a binary also in the current directory.

```
./hello
```

This should output `hello world`.

Now, try:

```
rm ./hello
go install
hello
```

The `go install` command installs the binary to the $GOPATH/bin directory. If
you have added that directory to your PATH, then you can execute the program
without the `./`.

Finally, open the `main.go` file and change the formatting to be terrible...
Introduce spurious new lines and spaces, then run:

```
go fmt
```

Open `main.go` again and it should be correctly formatted again.

It is highly recommendable to configure your editor to run `gofmt` (or
`goimports`) on save, so that you don't need to bother about it.


## Create the first test for your program

The convention in Go is to put unit tests alongside the code that they test
rather than in a different directory. This has the advantage of giving a very
visible way of seeing what has tests and what doesn't. All test files are named
`*_test.go`; anything named that way will not be added to the final binary, but
will be taken into account by `go test`.

Create a file `main_test.go` and add the following code:

```go
package main

import (
	"testing"
)

// All test function names should follow the TestAbc pattern and must take an
// argument of type *testing.T, conventionally named t.
func TestHelloWorld(t *testing.T) {
	val := HelloWorld()
	if val != "hello world" {
		t.Fail()
	}
}
```

Here we have imported the [testing package](https://golang.org/pkg/testing/) and
defined a test. The test calls a `HelloWorld` function (to be defined) and check
that the value returned matches a given value.

To run this test run the following command:

```
go test
```

To get more verbose output, add `-v`:

```
go test -v
```

You should see the test fail. Now let's change `main.go` to make it pass.
Reopen `main.go` and change it to look like the following:

```go
package main

import (
	"fmt"
)

// This is a public/exported function from the package. Exported functions must
// start with an Uppercase letter. As shown by the signature, the return type of
// this function is a string and it takes no arguments.
func HelloWorld() string {
	return "hello world"
}

func main() {
	fmt.Println(HelloWorld())
}
```

Run the tests again, and you should see success!

```
go test
```

---

**Pro tip**: having exported names in a `main` package is somewhat useless, for
you are not supposed to be importing that package from another one. But it is
fundamental to understand that Go uses the case of the first character in
identifier names instead of `private`/`public` keywords to determine visibility.
That is very nice when you are reading code. For instance, it is easy to tell
that `Println` is a public function defined in the `fmt` package. That package
may have private definitions, such that you can never do something like
`fmt.privatefunc()`.

---

If you are coming from another language, you are certainly missing assertion
methods, what is commonly used in xUnit-style test frameworks. Go's `testing`
package takes a different approach, notably simpler.

You are given a reference `t` to a `*testing.T` value, that allows you to
communicate with the test framework. The basic action is to mark the test as
failed. All the rest of the test code is normal Go code, using the same language
features you'd use elsewhere, like `if` and `for` constructs, function calls,
variable declarations and assignments.

We'll see in future chapters how powerful that idea is, as it allows for testing
more with less boilerplate code, e.g., using table driven tests.


## Reading documentation

It is easy to access documentation for any package, function, or any other
exported value in the standard library, or any other package in your workspace.

You can read the online documentation for the latest release at:
https://golang.org/pkg/

The above includes only the standard library. If you want to read the
documentation of any publicly accessible package (including also the stdlib),
check this out: https://godoc.org

GoDoc.org even helps you identifying popular packages, and consulting the docs
for projects in your own GitHub repository!

Often it is useful to read docs offline. That case is also very well covered.
You can start a local server very similar to golang.org:

```
godoc -http=:6060
```

Now browse to http://localhost:6060/pkg to find docs for every package in your
workspace. Check `godoc -h` for more options.

You can also use the `go doc` (note the space) to print documentation in a
convenient format for command line usage. For example, try out the following
commands:

```
go doc testing
go doc testing T.Fail
```


## IDEs and Editor Plugins for Go

Most main editors have some support for Go. Check the full list of [IDEs and
Plugins for Go](https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins) on
the official Go Wiki.

Though many editors support Go, they do not offer the same level of features and
integrations.

These are things you might look for when considering an editor, from the most
basic to the more advanced use cases. Keep in mind people tend to have different
opinions here, so this is just a list to let you know some of what is possible,
value the ones you care about most:

- Syntax coloring
- Integration with `gofmt` to automatically format source code (best when always
  triggered on save)
- Ability to replace `gofmt` with `goimports` to automatically
  add/remove/organize imports
- Go to definition of identifier under cursor
- Code auto completion
- Integration with code linters (`go vet`, `golint`, `gometalinter`, ...)
- Automatic refactors (rename, extract variable/function, ...)
- Debugger Integration
- Integration with `go test` to run tests
- Integration with `godoc` to provide contextual documentation
- Integration with Go [`guru`](https://docs.google.com/document/d/1_Y9xCEMj5S-7rv2ooHpZNH15JgRT5iM742gJkw5LtmQ/view)
  (formerly known as Go `oracle`) to answer questions about code


## Optional Homework: A Tour of Go

If you want to learn more (we will be going through more next week), now that
you have your environment setup, take a look at [A Tour of
Go](https://tour.golang.org/) that essentially start where we have left off.
