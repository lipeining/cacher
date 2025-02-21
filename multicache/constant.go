package multicache

const (
	RedisKeyTypeHash    = "hash"
	RedisKeyTypeString  = "string"
	RedisKeyTypeSet     = "set"
	RedisMaxConcurrency = 5  // 最大并发打包数量
	ReloadShuffleSec    = 30 // 随机延迟reload因子，用于各实例reload打散
	ReloadMaxCount      = 5  // 最大重试次数

)

type Config struct {
	ScanRule         string
	MaxConcurrency   int
	ReloadShuffleSec int
	ReloadMaxCount   int
	ReloadTickerSec  int
	LocalCacheClose  bool
}
