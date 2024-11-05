package postgresql

import (
	"context"
	"fmt"
	"log"
	"postgres_bench/pkg/storage"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	Db *sqlx.DB
}

func UrlPostgres(opt *storage.OptionsM) string {
	return fmt.Sprintf(`postgres://%s:%s@%s:%s/%s`, opt.DbUser, opt.DbPassw, opt.DbHost, opt.DbPort, opt.DbName)
}

func NewDB(opt *storage.OptionsM) (*DB, error) {

	dsn := UrlPostgres(opt)

	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	db.DB.SetMaxOpenConns(opt.SetMaxOpenConns)
	db.DB.SetMaxIdleConns(opt.SetMaxIdleConns)
	db.SetConnMaxLifetime(200 * time.Millisecond)
	db.SetConnMaxIdleTime(20 * time.Second)

	DB := &DB{
		Db: db,
	}

	return DB, nil
}

func (d *DB) Close(ctx context.Context) {
	d.Db.Close()
}
