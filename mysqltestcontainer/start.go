package mysqltestcontainer

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hooligram/kifu"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	rootPassword = "password"
)

func Start(databaseName string, migrationDir string) (*sql.DB, error) {
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
	container, _ := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	host, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}
	port, err := container.MappedPort(ctx, "3306/tcp")
	if err != nil {
		return nil, err
	}
	db, err := open(host, fmt.Sprint(port.Int()), rootPassword, databaseName)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	if migrationDir != "" {
		err := migrate(db, migrationDir)
		if err != nil {
			return nil, err
		}
	}
	kifu.Info("MySQL test container started successfully!")
	return db, nil
}
