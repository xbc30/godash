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
```go
package main

import (
  "fmt"

  "github.com/xbc30/godash"
)

func main() {
  input := []int{1,2,3,4,5}
  output := godash.ArrayChunk(godash.ArrayIntToInterface(input), 2)
  fmt.Println(output) // [][]interface{}{{1, 2}, {3, 4}, {5}}
}
```

### List

* Array
  * [Chunk](https://github.com/xbc30/godash/blob/master/array.go#L43)
  * [Diffenece](https://github.com/xbc30/godash/blob/master/array.go#L101)
  * [Fill](https://github.com/xbc30/godash/blob/master/array.go#L147)
  * [Find](https://github.com/xbc30/godash/blob/master/array.go#L165)
  * [FindIndex](https://github.com/xbc30/godash/blob/master/array.go#L212)
* String
  * [Capitalize](https://github.com/xbc30/godash/blob/master/string.go#L13)
  * [StartsWith](https://github.com/xbc30/godash/blob/master/string.go#L31)
  * [EndsWith](https://github.com/xbc30/godash/blob/master/string.go#L42)
* [Map](https://github.com/xbc30/godash/blob/master/map.go#L8)

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
