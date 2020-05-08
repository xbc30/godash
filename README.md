## Godash
![image](./gopher.png)

### Example

* With Type Input
```go
package main

import (
  "fmt"

  "github.com/xbc30/godash"
)

func main() {
  input := []int{1,2,3,4,5}
  output := godash.ArrayIntChunk(input, 2)
  fmt.Println(output) // [[1,2] [3,4] [5]]
}

```

* interface{} Input

### List

* Array
  * [Chunk](https://github.com/xbc30/godash/blob/master/array.go#L43)
  * [Diffenece](https://github.com/xbc30/godash/blob/master/array.go#L101)
  * [Fill](https://github.com/xbc30/godash/blob/master/array.go#L147)
  * [Find](https://github.com/xbc30/godash/blob/master/array.go#L165)
  * [FindIndex](https://github.com/xbc30/godash/blob/master/array.go#L212)

### Unit Test & Benchmark Test

* Unit Test
  * Single Func Test

  ```go
    go test -v -run="TestArrayIntChunk" -count 1
  ```

  * File Test

  ```go
    go test -v -run="" -count 1
  ```

* Benchmark Test

```go
go test --bench=. array_test.go array.go -benchmem
```  
