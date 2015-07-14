chatspry.go
===========

Go client for Chatspry

```go
package main

import chatspry "github.com/chatspry/chatspry.go/v1"

const baseURL = "foo.bar.baz/api"

func main()
  client := chatspry.NewClient(baseURL)
  client.Login("foo", "bar")
  // Do things with your client here!
}
```
