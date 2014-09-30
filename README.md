chatspry.go
===========

Golang client for Chatspry

```go
package main

import chatspry "github.com/chatspry/chatspry.go"

const baseURL = "foo.bar.baz/api"

func main() {
  client := chatspry.NewV1Client(baseURL)
  client.Login("foo", "bar")
  // Do things with your client here!
}
```
