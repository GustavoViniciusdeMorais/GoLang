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

### At file ./go.work
```
go 1.18
use ./practice
```

### 
```
mkdir tester
cd tester
go mod init example.com/tester
go mod edit --replace example.com/practice=../practice
go mod tidy
go run tester.go
```

