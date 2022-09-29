# go-mysql-test-container

[![pkg](https://pkg.go.dev/badge/github.com/andylongstaffe/go-mysql-test-container.svg)](https://pkg.go.dev/github.com/andylongstaffe/go-mysql-test-container)
[![report](https://goreportcard.com/badge/github.com/andylongstaffe/go-mysql-test-container)](https://goreportcard.com/report/github.com/andylongstaffe/go-mysql-test-container)
[![coverage](https://codecov.io/gh/andylongstaffe/go-mysql-test-container/branch/master/graph/badge.svg?token=isZCzDyj1N)](https://codecov.io/gh/andylongstaffe/go-mysql-test-container)
[![build](https://github.com/andylongstaffe/go-mysql-test-container/actions/workflows/build.yml/badge.svg)](https://github.com/andylongstaffe/go-mysql-test-container/actions/workflows/build.yml)

Golang MySQL testcontainer that makes integration testing MySQL a breeze

## Usage

Execute

```
go get github.com/andylongstaffe/go-mysql-test-container
```

Code

```go
package main

import (
	"testing"

	"github.com/andylongstaffe/go-mysql-test-container/mysqltestcontainer"
)

func Test(t *testing.T) {
	mySql, _ := mysqltestcontainer.Create("test")
	db := mySql.GetDb()
	err := db.Ping()
	if err != nil {
		log.L.Errorln(err.Error())
	}
}
```

## Development

1. Make sure Docker is installed on your local machine

   ```
   docker --version
   ```

2. Make sure you can run Docker locally

   ```
   docker run hello-world
   ```

3. Run the test

   ```
   ./gotest.sh
   ```
