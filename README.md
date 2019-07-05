# zrange
###### 迭代包

## 获得zrange
`go get -u github.com/zlyuancn/zrange`

## 使用zrange

```go
package main

import (
    "fmt"
    "github.com/zlyuancn/zrange"
)

func main() {
    zrange.FRange(5, func(i int) {
        fmt.Println(i)
    })
}
```
