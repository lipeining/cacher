package rdbcache

import (
	"context"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type clusterCache struct {
	client *redis.ClusterClient
	logger *log.Helper
}

func NewClusterCache(config RedisConfig, logger log.Logger) RedisCache {
	options := &redis.ClusterOptions{
		Addrs:    config.Addrs,
		Password: config.Password,
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

	rdb := redis.NewClusterClient(options)

	return &clusterCache{client: rdb, logger: log.NewHelper(logger)}
}

func (r *clusterCache) Nil() error {
	return redis.Nil
}

func (r *clusterCache) Scan(ctx context.Context, prefix, keyType string) ([]string, error) {
	// 获取所有节点
	nodes, err := r.client.ClusterSlots(ctx).Result()
	if err != nil {
		return nil, err
	}

	var (
		mu sync.Mutex
		wg sync.WaitGroup

		allKeys []string
	)

	// 消重
	materArrMap := make(map[string]struct{})
	for _, node := range nodes {
		if len(node.Nodes) > 0 {
			materArrMap[node.Nodes[0].Addr] = struct{}{}
		}
	}

	// 并发请求多个分片上的keys
	for nodeArr := range materArrMap {
		wg.Add(1)
		addr := nodeArr
		go func() {
			// defer log
			defer wg.Done()
			cli := redis.NewClient(&redis.Options{
				Addr:        addr,
				Password:    r.client.Options().Password,
				IdleTimeout: time.Second,
				ReadTimeout: 3 * time.Second,
				PoolTimeout: time.Second,
			})
			defer cli.Close()

			keys := make([]string, 0)
			cursor := uint64(0)
			for {
				ret, nextCursor, e := cli.ScanType(ctx, cursor, prefix, RedisMaxScanSize, keyType).Result()
				if e != nil {
					err = e
					return
				}
				keys = append(keys, ret...)

				if nextCursor <= 0 {
					break
				}
				cursor = nextCursor
			}

			mu.Lock()
			defer mu.Unlock()
			allKeys = append(allKeys, keys...)
		}()
	}
	wg.Wait()

	if err != nil {
		return nil, err
	}
	return allKeys, nil
}

func (r *clusterCache) HMGet(ctx context.Context, key string, field ...string) ([]interface{}, error) {
	results, err := r.client.HMGet(ctx, key, field...).Result()
	if err != nil {
		return nil, errors.Wrap(err, "redis hmget fail")
	}
	return results, nil
}

func (r *clusterCache) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	results, err := r.client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, errors.Wrap(err, "redis hgetall fail")
	}
	return results, nil
}

func (r *clusterCache) HGetAllByBatch(ctx context.Context, key string) (map[string]string, error) {
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

func (r *clusterCache) Get(ctx context.Context, key string) (string, error) {
	v, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", errors.Wrap(err, "redis get fail")
	}

	return v, nil
}
