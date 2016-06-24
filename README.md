# Golang lib for botan.io

[Botan](http://botan.io) is a telegram bot analytics system based on [Yandex.Appmetrica](http://appmetrica.yandex.com/)

Example
```go
package main

import (
	"github.com/onrik/botanio"
)

func main() {
	botan := botanio.New("<token>")
	answer, err := botan.Track(123456, "search", botanio.Map{
		"query": "cartoon",
	})
	
	// Track without client creating
	answer, err := botanio.Track("<token>", 123456, "search", botanio.Map{
		"query": "cartoon",
	})
}

```
