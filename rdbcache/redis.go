package rdbcache

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type redisCache struct {
	client *redis.Client
	logger *log.Helper
}

func NewRedisCache(config RedisConfig, logger log.Logger) RedisCache {
	options := &redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	}
	if config.DialTimeout > 0 {
		options.DialTimeout = config.DialTimeout
	}
	if config.ReadTimeout > 0 {
		options.ReadTimeout = config.ReadTimeout
	}
	if config.WriteTimeout > 0 {
		options.WriteTimeout = config.WriteTimeout
	}

	if config.MinIdleConns > 0 {
		options.MinIdleConns = config.MinIdleConns
	}
	if config.PollSize > 0 {
		options.PoolSize = config.PollSize
	}

	rdb := redis.NewClient(options)

	return &redisCache{client: rdb, logger: log.NewHelper(logger)}
}

func (r *redisCache) Nil() error {
	return redis.Nil
}

func (r *redisCache) Scan(ctx context.Context, prefix, keyType string) ([]string, error) {
	allKeys := make([]string, 0)
	cursor := uint64(0)
	for {
		ret, nextCursor, err := r.client.ScanType(ctx, cursor, prefix, RedisMaxScanSize, keyType).Result()
		if err != nil {
			return nil, err
		}

		allKeys = append(allKeys, ret...)
		if nextCursor <= 0 {
			break
		}

		cursor = nextCursor
	}

	return allKeys, nil
}

func (r *redisCache) HMGet(ctx context.Context, key string, field ...string) ([]interface{}, error) {
	results, err := r.client.HMGet(ctx, key, field...).Result()
	if err != nil {
		return nil, errors.Wrap(err, "redis hmget fail")
	}
	return results, nil
}

func (r *redisCache) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	results, err := r.client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, errors.Wrap(err, "redis hgetall fail")
	}
	return results, nil
}

func (r *redisCache) HGetAllByBatch(ctx context.Context, key string) (map[string]string, error) {
	fields, err := r.client.HKeys(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	offset, total, fail := 0, len(fields), 0
	results := make(map[string]string, total)
	if total == 0 {
		return results, nil
	}

	for offset < total {
		count := offset + RedisMaxHScanSize
		if count > total {
			count = total
		}

		batches := fields[offset:count]
		ret, err := r.client.HMGet(ctx, key, batches...).Result()
		if err != nil {
			return nil, errors.WithMessage(err, "hmget fields fail")
		}
		if len(ret) != len(batches) {
			return nil, errors.WithMessage(err, "hmget fields result number no match")
		}

		for idx, k := range batches {
			if v, ok := ret[idx].(string); ok {
				results[k] = v
			} else {
				fail++
			}
		}

		offset = count
	}

	if fail > 0 {
		r.logger.Infof("get hash key batch miss some fields, key: %s, total fields: %d, fail: %d", key, total, fail)
	}
	return results, err
}

func (r *redisCache) Get(ctx context.Context, key string) (string, error) {
	v, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", errors.Wrap(err, "redis get fail")
	}

	return v, nil
}
