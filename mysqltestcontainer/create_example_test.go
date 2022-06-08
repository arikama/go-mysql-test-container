package mysqltestcontainer_test

import (
	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
	"github.com/containerd/containerd/log"
)

func ExampleCreate() {
	mySql, _ := mysqltestcontainer.Create("test")
	db := mySql.GetDb()
	err := db.Ping()
	if err != nil {
		log.L.Errorln(err.Error())
	}
}
