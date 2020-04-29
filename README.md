## Godash

### List

* Array
  * chunk
  * diffenece

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
