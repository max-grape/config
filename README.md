# config

## env

Sets values for your structure fields from environment variables.

### Usage

```go
package main

import (
	"fmt"
	"os"

	"github.com/max-grape/config/env"
)

func main() {
	if err := os.Setenv("FOO_STRING", "bar"); err != nil {
		panic(err)
	}

	type config struct {
		Foo string `env:"FOO_STRING"`
	}

	cfg := config{
		Foo: "defaultValue",
	}

	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	fmt.Println(cfg.Foo) // Prints bar
}
```

### Supported types

- string
- int
- int64
- uint
- uint64
- bool
- duration

## TODO

- Add more parsing types to env package
