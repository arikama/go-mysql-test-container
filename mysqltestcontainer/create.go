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
	rootUsername = "root"
	rootPassword = "password"
)

// Create creates a container containing MySQL in Docker & returns the connection info along with
// the created database.
func Create(databaseName string) (*MySql, error) {
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
	host, _ := container.Host(ctx)
	port, _ := container.MappedPort(ctx, "3306/tcp")
	p := fmt.Sprint(port.Int())
	kifu.Info("Connecting to MySQL inside test container: host=%v, port=%v, database_name=%v", host, p, databaseName)
	db, err := open(host, p, rootPassword, databaseName)
	if err != nil {
		return nil, err
	}
	kifu.Info("MySQL test container started successfully!")
	mySql := &MySql{
		db: db,
		dbInfo: &DbInfo{
			Username: rootUsername,
			Password: rootPassword,
			Ip:       host,
			Port:     p,
			DbName:   databaseName,
		},
	}
	return mySql, nil
}
