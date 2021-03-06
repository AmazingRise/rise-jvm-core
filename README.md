# Rise JVM Core

[![Go](https://img.shields.io/badge/--00ADD8?logo=go&logoColor=ffffff)](https://golang.org/)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

This is the core of Rise JVM.

Rise JVM is a Java Virtual Machine based on WASM, written in Go.

Tested under:
- Ubuntu 20.04
- OpenJDK javac 11.0.13

## β¨Quick Start

Just one line.

```bash
go test
```

## πStructure

Project structure:

```
.
βββ LICENSE
βββ README.md
βββ demo
βββ entity
βββ go.mod
βββ go.sum
βββ jvm
βββ loader
βββ logger
βββ main.go
βββ main_test.go
βββ rt
βββ utils
```

### `demo`

Demo Java classes and their source code.

### `loader`

Class loader and related stuffs are here. They load bytes from `class` file. The loader will convert these bytes into Go `struct`.

### `entity`

This directory stores the definition of structures, and its methods. The methods are only related to the `struct`, e.g. `IsPublic` for `Class`.
Other things like deserialization is not included.

### `jvm`

VM and byte code execution engine.

### `logger`

Global logger. It should be initialized.

## πAcknowledgement

[zserge/tojvm](https://github.com/zserge/tojvm). Some snippets in `loader` are from here. They are noted in comments.

## πLicense

This project is licensed under GPLv3.