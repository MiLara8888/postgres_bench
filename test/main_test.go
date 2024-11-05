package test

import (
	"context"
	"log"
	"postgres_bench/pkg/storage/postgresql"
	"testing"
)


func TestPostgresqlConn(t *testing.T) {

	ctx := context.Background()
	db, err := postgresql.NewDB(&Opt)
	if err != nil {
		log.Println(err)
	}

	defer db.Close(ctx)

}
