# go-mysql-test-container

[![report](https://goreportcard.com/badge/github.com/arikama/go-mysql-test-container)](https://goreportcard.com/report/github.com/arikama/go-mysql-test-container)
[![coverage](https://codecov.io/gh/arikama/go-mysql-test-container/branch/master/graph/badge.svg?token=isZCzDyj1N)](https://codecov.io/gh/arikama/go-mysql-test-container)
[![build](https://github.com/arikama/go-mysql-test-container/actions/workflows/build.yml/badge.svg)](https://github.com/arikama/go-mysql-test-container/actions/workflows/build.yml)

MySQL wrapper for https://github.com/testcontainers/testcontainers-go

## Usage

Execute

```
go get github.com/arikama/go-mysql-test-container
```

Code

```go
package main

import (
	"testing"

	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
)

func Test(t *testing.T) {
	db, _ := mysqltestcontainer.Start("test", "./migration")
	db.Ping()
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
