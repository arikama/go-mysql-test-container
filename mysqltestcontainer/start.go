package mysqltestcontainer

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hooligram/kifu"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	rootPassword = "password"
)

func Start(databaseName string) (*Result, error) {
	kifu.Info("Starting MySQL test container...")
	req := testcontainers.ContainerRequest{
		Image:        "mysql:5.6",
		ExposedPorts: []string{"3306/tcp", "33060/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": rootPassword,
			"MYSQL_DATABASE":      databaseName,
		},
		WaitingFor: wait.ForLog("3306"),
	}
	ctx := context.Background()
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}
	host, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}
	port, err := container.MappedPort(ctx, "3306/tcp")
	if err != nil {
		return nil, err
	}
	p := fmt.Sprint(port.Int())
	db, err := open(host, p, rootPassword, databaseName)
	if err != nil {
		return nil, err
	}
	kifu.Info("MySQL test container started successfully!")
	result := &Result{
		Db:       db,
		Username: "root",
		Password: rootPassword,
		Ip:       host,
		Port:     p,
		Database: databaseName,
	}
	return result, nil
}
