# go-toolkit

Reusable Go utilities for everyday application code.

This repository is a personal toolkit of small, generic helpers you can import into other projects instead of rewriting the same patterns again.

## Packages

| Package | Description |
|---------|-------------|
| [`colletions`](./colletions) | Slice helpers: `Map`, `Filter`, `Find`, `ForEach` |

## Requirements

- Go 1.18+ (generics)

## Install

```bash
go get github.com/alexandre-mahumane/go-toolkit
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/alexandre-mahumane/go-toolkit/colletions"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}

	doubled := colletions.Map(nums, func(n int) int { return n * 2 })
	evens := colletions.Filter(nums, func(n int) bool { return n%2 == 0 })
	first, ok := colletions.Find(nums, func(n int) bool { return n > 3 })

	colletions.ForEach(doubled, func(n int) {
		fmt.Println(n)
	})

	fmt.Println(evens, first, ok)
}
```

## Run tests

```bash
go test ./...
```

## Idea

Keep a growing set of focused packages (`colletions`, and later things like strings, concurrency, or HTTP helpers) with clear APIs and tests — a toolbox you can pull from when building services and CLIs in Go.
