# Rise JVM Core

[![Go](https://img.shields.io/badge/--00ADD8?logo=go&logoColor=ffffff)](https://golang.org/)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

This is the core of Rise JVM.

Rise JVM is a Java Virtual Machine based on WASM, written in Go.

Tested under:
- Ubuntu 20.04
- OpenJDK javac 11.0.13

## ✨Quick Start

### Run a specific `.class` file

1. Build it.
```bash
go build .
```
2. Pick a class from `demo` and run it!
```bash
./rise-jvm-core demo/Add
```

**NOTE**: the suffix `.class` should be emitted.

### Run all test cases

Just one line.

```bash
go test
```

## 🎄Structure

Project structure:

```
.
├── LICENSE
├── README.md
├── demo
├── entity
├── go.mod
├── go.sum
├── jvm
├── loader
├── logger
├── main.go
├── main_test.go
├── rt
└── utils
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

## 👏Acknowledgement

[zserge/tojvm](https://github.com/zserge/tojvm). Some snippets in `loader` are from here. They are noted in comments.

## 📜License

This project is licensed under GPLv3.