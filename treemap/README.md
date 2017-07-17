# TreeMap

## Example

```go
package main

import (
    "github.com/neoql/container/treemap"
)

func main() {
    // Create a new TreeMap
    tm := treemap.New()

    // Put key and value into the map.
    tm.Put("Hello", 666)
    tm.Put("Hi", 2333)

    // Update value.
    tm.Put("Hello", 888)

    // Remove data
    tm.Remove("Hi")

    // Get size of the Treemap.
    size := tm.Size()

    // Traverse the TreeMap
    iter := tm.EntryIterator()
    for iter.HasNext() {
        e := iter.Next()
        // Get the key of entry.
        k := e.GetKey()
        // Get the value of entry
        v := e.GetValue()
    }
}
```