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
	url, err := botan.Short(123456, "https://google.com")
	
	// Usage without client creating
	botanio.SetToken("<token>")
	answer, err := botanio.Track(123456, "search", botanio.Map{
		"query": "cartoon",
	})
	url, err := botanio.Short(123456, "https://google.com")
}

```
