package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"time"
)

const driverName = "pgx"

type Config struct {
	Host     string `config:"envVar"`
	Port     string `config:"envVar"`
	Name     string
	User     string `config:"envVar"`
	Password string `config:"envVar"`
	TimeZone string

	MaxOpenConns    *int
	MaxIdleConns    *int
	ConnMaxLifetime *time.Duration
	ConnMaxIdleTime *time.Duration
}

func New(l *slog.Logger, cfg Config) (*sqlx.DB, error) {
	dbString := fmt.Sprintf("DB host: [%s:%s].", cfg.Host, cfg.Port)

	l.Info(dbString + "Подключение...")

	connString := connectionString(cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.TimeZone)

	db, err := sqlx.Open(driverName, connString)
	if err != nil {
		return nil, err
	}

	applyConfig(db, cfg)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	l.Info(dbString + "Подключено!")

	return db, nil
}

func applyConfig(db *sqlx.DB, cfg Config) {
	if cfg.MaxOpenConns != nil {
		db.SetMaxOpenConns(*cfg.MaxOpenConns)
	}

	if cfg.MaxIdleConns != nil {
		db.SetMaxIdleConns(*cfg.MaxIdleConns)
	}

	if cfg.ConnMaxIdleTime != nil {
		db.SetConnMaxIdleTime(*cfg.ConnMaxIdleTime)
	}

	if cfg.ConnMaxLifetime != nil {
		db.SetConnMaxLifetime(*cfg.ConnMaxLifetime)
	}
}

func connectionString(host, port, user, password, name, timeZone string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		host, port, user, password, name, timeZone)
}
