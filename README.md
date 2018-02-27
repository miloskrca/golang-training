## Environment variables

```sh
export PATH=$PATH:$(go env GOPATH)/bin
export GOPATH=$(go env GOPATH)
```

### New program

```sh
mkdir -p $GOPATH/src/github.com/user
```

To compile and run a simple program, first choose a package path (we'll use github.com/user/hello) and create a corresponding package directory inside your workspace:

```sh
$ mkdir $GOPATH/src/github.com/user/hello
```

Next, create a file named hello.go inside that directory, containing the following Go code.

```go
package main

import "fmt"

func main() {
	fmt.Printf("Hello, world.\n")
}
```

Now you can build and install that program with the go tool:

```sh
$ go install github.com/user/hello
```

Note that you can run this command from anywhere on your system. The go tool finds the source code by looking for the github.com/user/hello package inside the workspace specified by GOPATH.

You can also omit the package path if you run go install from the package directory:

```sh
$ cd $GOPATH/src/github.com/user/hello
$ go install
```

This command builds the hello command, producing an executable binary. It then installs that binary to the workspace's bin directory as hello (or, under Windows, hello.exe). In our example, that will be $GOPATH/bin/hello, which is $HOME/go/bin/hello.

The go tool will only print output when an error occurs, so if these commands produce no output they have executed successfully.

You can now run the program by typing its full path at the command line:

```sh
$ $GOPATH/bin/hello
Hello, world.
```

Or, as you have added $GOPATH/bin to your PATH, just type the binary name:

```sh
$ hello
Hello, world.
```

## Structure

```sh
bin/
    hello                          # command executable
    outyet                         # command executable
pkg/
    linux_amd64/
        github.com/golang/example/
            stringutil.a           # package object
src/
    github.com/golang/example/
        .git/                      # Git repository metadata
	hello/
	    hello.go               # command source
	outyet/
	    main.go                # command source
	    main_test.go           # test source
	stringutil/
	    reverse.go             # package source
	    reverse_test.go        # test source
    golang.org/x/image/
        .git/                      # Git repository metadata
	bmp/
	    reader.go              # package source
	    writer.go              # package source
    ... (many more repositories and packages omitted) ...

```

## How do I download remote packages?

To get all dependencies for the current package:

```sh
go get ./...
```

To download a particular remote package:

```sh
go get <package import path>  # see 'go help packages' for details

# for instance,
go get github.com/user/package
```

All downloaded packages end up in `$GOPATH/src`. They are also automatically built and installed into `$GOPATH/pkg`. You can skip installation by passing `-d` flag to `go get`:

```sh
go get -d github.com/user/package
```

## Dep

```sh
$ dep init
$ ls
Gopkg.toml Gopkg.lock vendor/
```

Install a package:

```sh
$ dep ensure -add github.com/pkg/errors
```

Make sure to import it in the code before running `dep ensure` again!

Ensure dependencies are in order:

```sh
$ dep ensure
```

Careful: this command removes packages that are not used anymore and adds packages that have been imported in the meantime, but it does niot update the Gopkg.toml file! To update Gopkg.toml you need to use `-add`.

### Gopkg.toml
Gopkg.toml files contain five basic types of rules. The Gopkg.toml docs explain them in detail, but here's an overview:

* *required*, which are mostly equivalent to import statements in .go files, except that it's OK to list a main package here
* *ignored*, which causes dep to black hole an import path (and any imports it uniquely introduces)
* *[[constraint]]*, stanzas that express version constraints and some other rules on a per-project dependency basis
* *[[override]]*, stanzas identical to [[constraint]] except that only the current project can express them and they supersede [[constraint]] in both the current project and dependencies
* *[prune]*, global and per-project rules that govern what kinds of files should be removed from vendor/

### Links 

* https://golang.org/
* https://golang.github.io/dep/
* https://golangweekly.com/
