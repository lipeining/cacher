package rdbcache

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

const (
	RedisMaxScanSize  = 500
	RedisMaxHScanSize = 30
)

type RedisConfig struct {
	Driver       string
	Addrs        []string
	Addr         string
	DB           int
	Password     string
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	MinIdleConns int
	PollSize     int
}

type RedisCache interface {
	Nil() error

	Scan(ctx context.Context, prefix, keyType string) ([]string, error)
	HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error)
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	HGetAllByBatch(ctx context.Context, key string) (map[string]string, error)

	Get(ctx context.Context, key string) (string, error)
}

func InitInstance(cf RedisConfig, logger log.Logger) RedisCache {
	var ins RedisCache
	switch cf.Driver {
	case "redis":
		ins = NewRedisCache(cf, logger)
	case "redis_cluster":
		ins = NewClusterCache(cf, logger)
	default:
		panic("unsupported driver: " + cf.Driver)
	}
	return ins
}
