# Rise JVM Core

[![Go](https://img.shields.io/badge/--00ADD8?logo=go&logoColor=ffffff)](https://golang.org/)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

This is the core of Rise JVM.

Rise JVM is a Java Virtual Machine based on WASM, written in Go.

## Structure

Project structure:

```
.
├── Add.class  // Demo Java Class
├── Add.java   // Demo Java Source Code
├── LICENSE
├── README.md  // You're here :)
├── entity     // Code: Entities, definition of structures
├── go.mod
├── jvm        // Code: Virtual Machine
├── loader     // Code: Class Loader
├── main.go    // Code: Entrance
└── utils      // Code: Utilities, like logger
```

### `loader`

Class loader and related stuffs are here. They load bytes from `class` file. The loader will convert these bytes into Go `struct`.

### `entity`

This directory stores the definition of structures, and its methods. The methods are only related to the `struct`, e.g. `IsPublic` for `Class`.
Other things like deserialization is not included.

### `jvm`

Code of JVM.

## Acknowledgement

[zserge/tojvm](https://github.com/zserge/tojvm). Some snippets in `loader` are from here.

## License

This project is licensed under GPLv3.