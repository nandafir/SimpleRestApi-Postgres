package main

import (
	"anaconda/config"
	"anaconda/server"
)

func main() {
	config.Init()
	// cfg := config.Get()
	// cache.Init(cfg.RedisHost, cfg.RedisPassword, cfg.RedisDialTimeout, cfg.RedisExpiredDuration)
	// log.Init(cfg.LogDSN, cfg.LogLevel)
	server.Start()
}
