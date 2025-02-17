package canal

import "github.com/go-mysql-org/go-mysql/canal"

type Option func(cfg *canal.Config)

func WithWhere(where string) Option {
	return func(cfg *canal.Config) {
		cfg.Dump.Where = where
	}
}
