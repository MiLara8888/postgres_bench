package test

import (
	"context"
	"log"
	"postgres_bench/pkg/storage"
	"postgres_bench/pkg/storage/postgresql"
	"runtime"
	"testing"
)

// go test -v


var Opt = *storage.NewOptions(
	storage.WithHost("127.0.0.1"),
	storage.WithPort("5432"),
	storage.WithUser("admin"),
	storage.WithPassw("123456789"),
	storage.WithName("test"),
	storage.WithShema("test"),
	storage.WithSql("select * from users"),
	storage.WithTimeMs(3000),
	storage.WithWorkerCount(runtime.GOMAXPROCS(runtime.NumCPU()-1)),
)


// кастомный бенч
func TestBenchmarkCustomPostgresql(t *testing.T) {

	//общий контекст базы
	//нужен для закрытия
	ctx := context.Background()
	db, err := postgresql.NewDB(&Opt)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close(ctx)

	err = db.Db.Ping()
	if err != nil {
		t.Fatal(err)
		return
	}

	r := db.PostgresT(&Opt)

	//высчитываю количество запросов в секунду
	log.Printf("Запросов в секунду: %.2f\n", r/(float64(Opt.TimeMS)/float64(1000)))
	log.Println("Всего запросов за", (float64(Opt.TimeMS) / float64(1000)), "секунд", r)
	//высчитываю количество запросов в милисекунду
	log.Printf("Запросов в милиисекунду: %.2f\n", float64(r)/(float64(Opt.TimeMS)))

}

// некастомный бенч
func BenchmarkPostgresql(b *testing.B) {

	ctx := context.Background()
	db, err := postgresql.NewDB(&Opt)
	if err != nil {
		b.Fatal(err)
	}

	err = db.Db.Ping()
	if err != nil {
		b.Fatal(err)
	}

	defer db.Close(ctx)

	b.Run("benchmark", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {

			b.StartTimer()

			sqlt := db.Db.Rebind(Opt.Sql)

			_, err := db.Db.ExecContext(ctx, sqlt)

			b.StopTimer()
			if err != nil {
				b.Fatal(err)
			}
		}
	})

}
