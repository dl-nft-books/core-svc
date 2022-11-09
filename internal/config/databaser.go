package config

import (
	"database/sql"
	"github.com/lib/pq"
	"gitlab.com/distributed_lab/kit/pgdb"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

const databaserConfigKey = "database"

type DBConnection struct {
	DB       *pgdb.DB
	RawDB    *sql.DB
	Listener *pq.Listener
}

type Databaser interface {
	BookDB() *DBConnection
	GeneratorDB() *DBConnection
	TrackerDB() *DBConnection
}

type databaser struct {
	getter kv.Getter
	once   comfig.Once
}

func NewDatabaser(getter kv.Getter) Databaser {
	return &databaser{
		getter: getter,
	}
}

type databaserCfg struct {
	BookService      DatabaserFields `fig:"book_svc"`
	GeneratorService DatabaserFields `fig:"generator_svc"`
	TrackerService   DatabaserFields `fig:"tracker_svc"`
}

type DatabaserFields struct {
	URL                      string        `fig:"url,required"`
	MaxOpenConnections       int           `fig:"max_open_connection"`
	MaxIdleConnections       int           `fig:"max_idle_connections"`
	ListenerMinRetryDuration time.Duration `fig:"listener_min_retry_duration"`
	ListenerMaxRetryDuration time.Duration `fig:"listener_max_retry_duration"`
}

func (d *databaser) readConfig() databaserCfg {
	defaultFields := DatabaserFields{
		MaxOpenConnections:       12,
		MaxIdleConnections:       12,
		ListenerMinRetryDuration: time.Second,
		ListenerMaxRetryDuration: time.Minute,
	}

	cfg := databaserCfg{
		BookService:      defaultFields,
		GeneratorService: defaultFields,
	}

	if err := figure.
		Out(&cfg).
		From(kv.MustGetStringMap(d.getter, databaserConfigKey)).
		Please(); err != nil {
		panic(errors.Wrap(err, "failed to figure out"))
	}

	return cfg
}

func (fields *DatabaserFields) DB() *pgdb.DB {
	db, err := pgdb.Open(pgdb.Opts{
		URL:                fields.URL,
		MaxOpenConnections: fields.MaxOpenConnections,
		MaxIdleConnections: fields.MaxIdleConnections,
	})
	if err != nil {
		panic(errors.Wrap(err, "failed to open database connection"))
	}

	return db
}

//NewListener - returns new listener for notify events
func (fields *DatabaserFields) NewListener() *pq.Listener {
	listener := pq.NewListener(fields.URL,
		fields.ListenerMinRetryDuration,
		fields.ListenerMaxRetryDuration,
		nil)
	return listener
}

func (fields *DatabaserFields) RawDB() *sql.DB {
	return fields.DB().RawDB()
}

func (d *databaser) BookDB() *DBConnection {
	cfg := d.readConfig()
	return &DBConnection{
		DB:       cfg.BookService.DB(),
		RawDB:    cfg.BookService.RawDB(),
		Listener: cfg.BookService.NewListener(),
	}
}

func (d *databaser) GeneratorDB() *DBConnection {
	cfg := d.readConfig()
	return &DBConnection{
		DB:       cfg.GeneratorService.DB(),
		RawDB:    cfg.GeneratorService.RawDB(),
		Listener: cfg.GeneratorService.NewListener(),
	}
}

func (d *databaser) TrackerDB() *DBConnection {
	cfg := d.readConfig()
	return &DBConnection{
		DB:       cfg.TrackerService.DB(),
		RawDB:    cfg.TrackerService.RawDB(),
		Listener: cfg.TrackerService.NewListener(),
	}
}
