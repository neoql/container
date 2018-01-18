# Queue

## Example

demo_queue.go
```go
package main

import (
    "fmt"
    "github.com/neoql/container/queue"
)

func main() {
    q := queue.New()

    q.Put("Tom")
    q.Put("Jack")
    q.Put("Peter")
    q.Close()

    for {
        name, flag := q.Pop()
        if !flag {
            break
        }
        fmt.Println(name)
    }
}
```

```
$ go run demo_queue.go
Tom
Jack
Peter
```