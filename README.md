# filc
> filc -> field nil checker

# What is this?
Look for struct fields that are nil.
Used when the argument passed to the constructor etc. is a structure etc.

# Usage

```go
package main

import (
	"log"

	"github.com/go-utils/filc"
)

type (
	Repository1 struct{}
	Repository2 struct{}
	Repository3 struct{}
)

type ServiceRepositories struct {
	Repository1 *Repository1
	Repository2 *Repository2
	Repository3 *Repository3
}

type Service struct {
	repos ServiceRepositories
}

func NewService(repos ServiceRepositories) *Service {
	if nilFields := filc.Look(repos); len(nilFields) > 0 {
		log.Fatalf("%+v", nilFields)
		// Output: ["ServiceRepositories.Repository2", "ServiceRepositories.Repository3"]
	}
	return &Service{repos: repos}
}

func main() {
	repos := ServiceRepositories{
		Repository1: new(Repository1),
	}
	service := NewService(repos)
	...
}
```

## License
- Under the [MIT License](./LICENSE)
- Copyright (C) 2022 go-utils
