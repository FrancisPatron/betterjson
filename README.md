# Better Json

[![Go CI](https://github.com/FrancisPatron/betterjson/actions/workflows/go.yml/badge.svg)](https://github.com/FrancisPatron/betterjson/actions/workflows/go.yml)
[![Release](https://github.com/FrancisPatron/betterjson/actions/workflows/release.yml/badge.svg)](https://github.com/FrancisPatron/betterjson/actions/workflows/release.yml)

`Better Json` is a Go library that enhances JSON parsing by supporting comments within JSON files.

## Features

- **Comment Support**: Easily parse JSON files with single-line and multi-line comments.
- **Simple API**: Uses a familiar API similar to the standard `encoding/json` package.

## Usage

Here's a basic example of how to use `betterjson`:

```go
package main

import (
	"fmt"
	"log"
	"os"
	"github.com/FrancisPatron/betterjson"
)

// ... [Your struct definitions here]

func main() {
	// Read example.json
	data, err := os.ReadFile("example.json")
	if err != nil {
		log.Fatalf("Failed to read example.json: %v", err)
	}

	var example ExampleData
	err = betterjson.Unmarshal(data, &example)
	if err != nil {
		log.Fatalf("Failed to parse example.json: %v", err)
	}

	fmt.Printf("Parsed data: %+v\n", example)
}

```
for a complete example checkout this [example code](https://github.com/FrancisPatron/betterjson/blob/main/example/main.go)

## How It Works
`betterjson` offers a streamlined approach to parsing JSON files with comments. At its core, it employs a single, two-step function. First, the `uncommenter` strips away comments from the JSON content. Instead of simply removing comments, it replaces them with whitespaces, ensuring that the original structure and line count remain intact. This replacement is crucial, as it preserves the accuracy of line and column numbers. Once the comments are replaced, the standard `json.Unmarshal` function takes over, parsing the now-comment-free JSON. This approach ensures that any errors reported by `json.Unmarshal` accurately reflect the correct line and column numbers from the original file, providing precise feedback even in the presence of comments.

## Instalation
To install `betterjson`, use `go get`:   
```sh
 go get github.com/FrancisPatron/betterjson
```