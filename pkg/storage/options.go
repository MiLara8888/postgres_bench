package storage

type OptionsM struct {
	DbHost  string
	DbPort  string
	DbUser  string
	DbPassw string
	DbName  string
	DbShema string
	//sql запрос
	Sql string
	//время исполнения в миллисекндах
	TimeMS      int
	WorkerCount int

	//устанавливает максимальное количество открытых подключений к базе данных
	//по умолчанию 20
	SetMaxOpenConns int

	// устанавливает максимальное количество соединений в пуле простаивающих соединений
	//по умолчанию 10
	SetMaxIdleConns int
}

type OptionFunc func(*OptionsM)

func NewOptions(options ...OptionFunc) *OptionsM {

	service := &OptionsM{
		SetMaxOpenConns: 20,
		SetMaxIdleConns: 10,
	}
	for _, option := range options {
		option(service)
	}
	return service
}

func WithHost(host string) OptionFunc {
	return func(cs *OptionsM) {
		cs.DbHost = host
	}
}

func WithPort(port string) OptionFunc {
	return func(cs *OptionsM) {
		cs.DbPort = port
	}
}

func WithUser(user string) OptionFunc {
	return func(cs *OptionsM) {
		cs.DbUser = user
	}
}

func WithPassw(passw string) OptionFunc {
	return func(cs *OptionsM) {
		cs.DbPassw = passw
	}
}

func WithName(name string) OptionFunc {
	return func(cs *OptionsM) {
		cs.DbName = name
	}
}

func WithShema(shema string) OptionFunc {
	return func(cs *OptionsM) {
		cs.DbShema = shema
	}
}

func WithSql(sql string) OptionFunc {
	return func(cs *OptionsM) {
		cs.Sql = sql
	}
}

func WithTimeMs(t int) OptionFunc {
	return func(cs *OptionsM) {
		cs.TimeMS = t
	}
}

func WithWorkerCount(w int) OptionFunc {
	return func(cs *OptionsM) {
		cs.WorkerCount = w
	}
}

func WithSetMaxOpenConns(w int) OptionFunc {
	return func(cs *OptionsM) {
		cs.SetMaxOpenConns = w
	}
}

func WithSetMaxIdleConns(w int) OptionFunc {
	return func(cs *OptionsM) {
		cs.SetMaxIdleConns = w
	}
}
