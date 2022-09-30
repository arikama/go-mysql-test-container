package mysqltestcontainer

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	rootUsername = "root"
	rootPassword = "password"
)

// Create creates a container containing MySQL in Docker & returns the connection info along with
// the created database.
// You are responsible for calling GetContainer().Terminate() to tear down the container since the
// reaper for testcontainers has been turned off
func Create(databaseName string) (*MySqlTestContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        "mariadb:10.5",
		ExposedPorts: []string{"3306/tcp", "33060/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": rootPassword,
			"MYSQL_DATABASE":      databaseName,
		},
		WaitingFor: wait.ForLog("3306"),
		SkipReaper: true,
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
	db, err := open(host, p, rootPassword, databaseName)
	if err != nil {
		return nil, err
	}
	result := &MySqlTestContainer{
		db: db,
		dbInfo: &DbInfo{
			Username: rootUsername,
			Password: rootPassword,
			Ip:       host,
			Port:     p,
			DbName:   databaseName,
		},
		container: container,
	}
	return result, nil
}

func CreateWithMigrate(databaseName string, migrationURL string) (*MySqlTestContainer, error) {
	result, err := Create(databaseName)
	if err != nil {
		return nil, errors.Wrap(err, "error creating database")
	}

	driver, err := mysql.WithInstance(result.GetDb(), &mysql.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "error creating database driver")
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationURL,
		"mysql", driver)
	if err != nil {
		return nil, errors.Wrap(err, "error configuring migration")
	}

	err = m.Up()
	if err != nil {
		return nil, errors.Wrap(err, "error running migration")
	}

	return result, nil
}
