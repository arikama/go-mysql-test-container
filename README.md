# go-mysql-test-container

![build](https://github.com/arikama/go-mysql-test-container/actions/workflows/build.yml/badge.svg)
![coverage](https://codecov.io/gh/arikama/go-mysql-test-container/branch/master/graph/badge.svg?token=isZCzDyj1N)

Migration runner & MySQL wrapper for https://github.com/testcontainers/testcontainers-go

## Install

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

func TestAbc(t *testing.T) {
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
