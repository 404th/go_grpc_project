package postgres

import (
	"context"

	"github.com/404th/go_grpc_project/position_dpl_service/config"
	"github.com/404th/go_grpc_project/position_dpl_service/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db *pgxpool.Pool

	// functions below...
}

func NewPostgres(psqlConnString string, cfg config.Config) (storage.StorageI, error) {
	// First set up the pgx connection pool
	cnfg, err := pgxpool.ParseConfig(psqlConnString)
	if err != nil {
		return nil, err
	}

	cnfg.AfterConnect = nil
	cnfg.MaxConns = int32(cfg.PostgresMaxConnections)

	pool, err := pgxpool.ConnectConfig(context.Background(), cnfg)
	if err != nil {
		return nil, err
	}

	// check for Ping
	if err = pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return &Store{
		db: pool,
	}, nil
}

// methods below...
