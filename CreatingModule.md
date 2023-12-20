# Creat Module

### 
```
```

### 
```
mkdir practice
cd practice
go mod init example.com/practice
```

### Module code, ./practice/practice.go
```go
package practice

import (
	"fmt"
)

func Test() {
	fmt.Println("Called practice module")
}
```

### At file ./go.work
```
go 1.18
use ./practice
```

### Commands to import module 
```
mkdir tester
cd tester
go mod init example.com/tester
go mod edit --replace example.com/practice=../practice
go mod tidy
go run tester.go
```
### Code at ./tester/tester.go
```go
package main

import (
	"example.com/practice"
)

func main() {
	practice.Test()
}
```
