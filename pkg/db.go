package pkg

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Database *pgxpool.Pool

func Connect() error {
	dbEnv := []any{}
	dbEnv = append(dbEnv, os.Getenv("DBUSER"))
	dbEnv = append(dbEnv, os.Getenv("DBPASS"))
	dbEnv = append(dbEnv, os.Getenv("DBHOST"))
	dbEnv = append(dbEnv, os.Getenv("DBPORT"))
	dbEnv = append(dbEnv, os.Getenv("DBNAME"))
	dbString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbEnv...)

	var err error
	Database, err = pgxpool.New(context.Background(), dbString)
	if err != nil {
		return err
	}
	return nil
}
