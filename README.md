# <div align="center">Rise JVM</div>

<div align="center">
<nav>
<a href="">Live Demo</a> | 
<a href="https://github.com/amazingrise/rise-jvm-core">Rise JVM Core</a>
</nav>
</div>

[![Go](https://img.shields.io/badge/--00ADD8?logo=go&logoColor=ffffff)](https://golang.org/)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

Rise JVM is a minimal Java VM based on WASM. No server needed.

## ✨Have a try

<a href="">Live Demo</a>

## About Rise JVM Core

This is only a frontend of Rise JVM Core.

## Compiling

Automatically via GNU make:

```bash
make all
```

For TinyGo users:

```bash
tinygo build -o main.wasm -target wasm ./main.go
cp $(shell tinygo env TINYGOROOT)/targets/wasm_exec.js .
```

For Go users:

```bash
GOARCH=wasm GOOS=js go build -o main.wasm main.go
cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" .
```

We recommend TinyGo instead of Go, which results in a smaller binary size.

## Features

These features are tested and work well:

- Simple Integer Arithmetic
- Calling functions
- System.out.println
- Creating Objects
- Looping

Not working yet:

- Polymorphism
- Standard Input

You can find implemented opcodes [here](https://github.com/AmazingRise/rise-jvm-core/blob/main/jvm/opcode.go).

## Acknowledgement

[Pico.css](https://github.com/picocss/pico)