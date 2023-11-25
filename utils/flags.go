package utils

import (
	"github.com/codegangsta/cli"
)

var RedisServer string
var RedisServerFlag = cli.StringFlag{
	EnvVar:      "REDIS_SERVER",
	Name:        "redis-server",
	Usage:       "Redis Server",
	Value:       defaultRedisServer,
	Destination: &RedisServer,
}

var RedisPrefix string
var RedisPrefixFlag = cli.StringFlag{
	EnvVar:      "REDIS_PREFIX",
	Name:        "redis-prefix",
	Usage:       "Redis Prefix",
	Value:       defaultRedisPrefix,
	Destination: &RedisPrefix,
}

var KafkaServersFlag = cli.StringSliceFlag{
	EnvVar: "KAFKA_SERVER",
	Name:   "kafka-host",
	Usage:  "env var holds comma separated list of kafka brokers; use multiple flags on command line",
	Value:  &cli.StringSlice{defaultKafkaServer},
}

var ESIndexLatencyFlag = cli.IntFlag{
	Name:   "es-index-latency",
	Value:  3,
	EnvVar: "ES_INDEX_LATENCY",
}

var ESIndexThresholdFlag = cli.IntFlag{
	Name:   "es-index-threshold",
	Value:  60,
	EnvVar: "ES_INDEX_THRESHOLD",
}
