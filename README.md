# Golang programming language cheatsheet

## Table of contents

- [Basics](#basics)
  - [Compilation](#compilation)

## Basics

### Compilation

Go is a compiled language. The Go toolchain converts a source program and the things it
depends on int o instructions in the native machine language of a computer. These tools are
accessed through a single command called go that has a number of subcommands. The simplest
of these sub commands is run, which compiles the source code from one or more source
files whose names end in .go, links it with librar ies, then runs the resulting executable file.

```bash
go run helloworld.go
```

If the program is more than a one-shot experiment, it’s likely that you would want to compile
it once and save the compiled result for later use. That is done with go build:

```bash
go build helloworld.go
```

This creates an executable binary file called helloworld that can be run anytime without further
processing:

```bash
./helloworld
```