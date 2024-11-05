package postgresql

import (
	"context"
	"postgres_bench/pkg/storage"
	"sync"
	"time"
)

// принимает options с настройками запроса и бд
func (db *DB) PostgresT(opt *storage.OptionsM) float64 {
	var (
		//колличество реализованных запросов
		result int
		mu     = &sync.Mutex{}
		r      float64
	)

	//контекст на длинну теста
	ctxMs, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(opt.TimeMS))
	defer cancel()

	for i := 1; i <= opt.WorkerCount; i++ {

		sqlt := db.Db.Rebind(opt.Sql)
		go func(ctx context.Context) {
			for ctx.Err() == nil {
				_, err := db.Db.ExecContext(ctx, sqlt)
				if err != nil {
					return
				}
				//защищаю во избежание race
				mu.Lock()
				result += 1
				mu.Unlock()

			}
		}(ctxMs)
	}

	select {
	case <-ctxMs.Done():
		//возвращаю количество запросов
		r = float64(result)
	}
	return r
}
