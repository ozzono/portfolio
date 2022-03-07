package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

const (
	lazyConnect = false
)

type Config struct {
	Host     string
	Port     string
	User     string
	DBNAme   string
	SSLMode  string
	Password string
}

func NewPgxConn(cfg Config) (*pgxpool.Pool, error) {
	ctx := context.Background()
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.DBNAme,
		cfg.SSLMode,
		cfg.Password,
	)

	poolCfg, err := pgxpool.ParseConfig(dataSourceName)
	if err != nil {
		return nil, err
	}

	poolCfg.LazyConnect = lazyConnect

	connPool, err := pgxpool.ConnectConfig(ctx, poolCfg)
	if err != nil {
		return nil, errors.Wrap(err, "pgx.ConnectConfig")
	}

	return connPool, nil
}
